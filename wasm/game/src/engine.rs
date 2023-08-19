use std::cell::RefCell;
use std::{rc::Rc, sync::Mutex};

use anyhow::{anyhow, Result};
use async_trait::async_trait;
use futures::channel::oneshot::channel;
use serde::Deserialize;
use wasm_bindgen::JsCast;
use wasm_bindgen::{closure::WasmClosure, prelude::Closure};
use web_sys::{CanvasRenderingContext2d, HtmlImageElement};

use crate::browser::{self, context, window};

#[derive(Debug)]
pub struct Renderer {
    context: CanvasRenderingContext2d,
}
impl Renderer {
    pub fn clear(&self, rect: &Rect) {
        self.context.clear_rect(
            rect.x.into(),
            rect.y.into(),
            rect.width.into(),
            rect.height.into(),
        );
    }
    pub fn draw_image(&self, image: &HtmlImageElement, frame: &Rect, destination: &Rect) {
        self.context
            .draw_image_with_html_image_element_and_sw_and_sh_and_dx_and_dy_and_dw_and_dh(
                image,
                frame.x.into(),
                frame.y.into(),
                frame.width.into(),
                frame.height.into(),
                destination.x.into(),
                destination.y.into(),
                destination.width.into(),
                destination.height.into(),
            )
            .expect("draw_image failed");
    }
}

#[derive(Deserialize)]
pub struct Rect {
    pub x: f32,
    pub y: f32,
    pub width: f32,
    pub height: f32,
}

#[async_trait(?Send)]
pub trait Game {
    async fn initialize(&self) -> Result<Box<dyn Game>>;
    fn update(&mut self);
    fn draw(&self, renderer: &Renderer);
}
pub struct GameLoop {
    last_frame: f64,
    accumulated_delta: f32,
}
type SharedLoopClosure = Rc<RefCell<Option<LoopClosure>>>;
pub type LoopClosure = Closure<dyn FnMut(f64)>;

const FRAME_SIZE: f32 = 1.0 / 60.0 * 1000.0;

impl GameLoop {
    pub async fn start(mut game: impl Game + 'static) -> Result<()> {
        let mut game = game.initialize().await?;
        let mut game_loop = GameLoop {
            last_frame: browser::now()?,
            accumulated_delta: 0.0,
        };
        let f: SharedLoopClosure = Rc::new(RefCell::new(None));
        let g = Rc::clone(&f);
        let renderer = Renderer {
            context: context()?,
        };

        *g.borrow_mut() = Some(create_raf_closure(move |time| {
            game_loop.accumulated_delta += (time - game_loop.last_frame) as f32;
            while game_loop.accumulated_delta > FRAME_SIZE {
                game.update();
                game_loop.accumulated_delta -= FRAME_SIZE;
            }
            game_loop.last_frame = time;
            game.draw(&renderer);
            request_animation_frame(f.borrow().as_ref().unwrap()).unwrap();
        }));
        request_animation_frame(g.borrow().as_ref().unwrap())?;
        Ok(())
    }
}

pub fn create_raf_closure(f: impl FnMut(f64) + 'static) -> LoopClosure {
    closure_wrap(Box::new(f) as Box<dyn FnMut(f64)>)
}
pub fn closure_wrap<T: WasmClosure + ?Sized>(data: Box<T>) -> Closure<T> {
    Closure::wrap(data)
}

pub fn request_animation_frame(callback: &LoopClosure) -> Result<i32> {
    window()?
        .request_animation_frame(callback.as_ref().unchecked_ref())
        .map_err(|err| anyhow!("request_animation_frame error: {:#?}", err))
}

pub async fn load_image(source: &str) -> Result<HtmlImageElement> {
    let image = browser::new_image()?;
    let (tx, rx) = channel::<Result<()>>();
    let success_tx = Rc::new(Mutex::new(Some(tx)));
    let error_tx = Rc::clone(&success_tx);
    let success_callback = browser::closure_once(move || {
        if let Some(tx) = success_tx.lock().ok().and_then(|mut opt| opt.take()) {
            tx.send(Ok(())).unwrap();
        }
    });
    let error_callback = browser::closure_once(move || {
        if let Some(tx) = error_tx.lock().ok().and_then(|mut opt| opt.take()) {
            tx.send(Err(anyhow!("Could not load image"))).unwrap();
        }
    });
    image.set_onload(Some(success_callback.as_ref().unchecked_ref()));
    image.set_onerror(Some(error_callback.as_ref().unchecked_ref()));
    image.set_src(source);
    rx.await??;
    Ok(image)
}
