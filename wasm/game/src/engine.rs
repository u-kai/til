use std::cell::RefCell;
use std::{rc::Rc, sync::Mutex};

use anyhow::{anyhow, Result};
use futures::channel::oneshot::channel;
use wasm_bindgen::JsCast;
use wasm_bindgen::{closure::WasmClosure, prelude::Closure};
use web_sys::{CanvasRenderingContext2d, HtmlImageElement};

use crate::browser::{self, context, window};

pub trait Game {
    fn update(&mut self);
    fn draw(&self, context: &CanvasRenderingContext2d);
}
pub struct GameLoop;
type SharedLoopClosure = Rc<RefCell<Option<LoopClosuer>>>;
impl GameLoop {
    pub async fn start(mut game: impl Game + 'static) -> Result<()> {
        let f: SharedLoopClosure = Rc::new(RefCell::new(None));
        let g = Rc::clone(&f);

        *g.borrow_mut() = Some(create_raf_closure(move |time| {
            game.update();
            game.draw(&context().unwrap());
            request_animation_frame(&g.borrow().as_ref().unwrap()).unwrap();
        }));
        request_animation_frame(
            g.borrow()
                .as_ref()
                .ok_or_else(|| anyhow!("GameLoop: Loop is None"))?,
        )?;
        Ok(())
    }
}

pub type LoopClosuer = Closure<dyn FnMut(f64)>;

pub fn create_raf_closure(f: impl FnMut(f64) + 'static) -> LoopClosuer {
    closure_wrap(Box::new(f) as Box<dyn FnMut(f64)>)
}
pub fn closure_wrap<T: WasmClosure + ?Sized>(data: Box<T>) -> Closure<T> {
    Closure::wrap(data)
}

pub fn request_animation_frame(callback: &LoopClosuer) -> Result<i32> {
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
            tx.send(Ok(()));
        }
    });
    let error_callback = browser::closure_once(move || {
        if let Some(tx) = error_tx.lock().ok().and_then(|mut opt| opt.take()) {
            tx.send(Err(anyhow!("Could not load image")));
        }
    });
    image.set_onload(Some(success_callback.as_ref().unchecked_ref()));
    image.set_onerror(Some(error_callback.as_ref().unchecked_ref()));
    image.set_src(source);
    rx.await??;
    Ok(image)
}
