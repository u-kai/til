use std::collections::HashMap;
use std::rc::Rc;
use std::sync::Mutex;

use rand::thread_rng;
use rand::Rng;
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;

// When the `wee_alloc` feature is enabled, this uses `wee_alloc` as the global
// allocator.
//
// If you don't want to use `wee_alloc`, you can safely delete this.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// This is like the `main` function, except for JavaScript.
#[wasm_bindgen(start)]
pub fn main_js() -> Result<(), JsValue> {
    // This provides better error messages in debug mode.
    // It's disabled in release mode so it doesn't bloat up the file size.
    console_error_panic_hook::set_once();

    let window = web_sys::window().unwrap();
    let document = window.document().unwrap();
    let canvas = document
        .get_element_by_id("canvas")
        .unwrap()
        .dyn_into::<web_sys::HtmlCanvasElement>()
        .unwrap();

    let context = canvas
        .get_context("2d")
        .unwrap()
        .unwrap()
        .dyn_into::<web_sys::CanvasRenderingContext2d>()
        .unwrap();
    wasm_bindgen_futures::spawn_local(async move {
        let json = fetch_json("images/rhb.json")
            .await
            .expect("Could not fetch JSON");
        let sheet: Sheet = serde_wasm_bindgen::from_value(json).expect("Could not parse JSON");
        let (tx, rx) = futures::channel::oneshot::channel::<Result<(), JsValue>>();
        let success_tx = Rc::new(Mutex::new(Some(tx)));
        let error_tx = Rc::clone(&success_tx);
        let image = web_sys::HtmlImageElement::new().unwrap();
        let callback = Closure::once(move || {
            if let Some(success_tx) = success_tx.lock().ok().and_then(|mut opt| opt.take()) {
                success_tx.send(Ok(())).unwrap();
            };
        });
        let error_callback = Closure::once(move |err: JsValue| {
            if let Some(error_tx) = error_tx.lock().ok().and_then(|mut opt| opt.take()) {
                web_sys::console::log_1(err.as_ref());
                error_tx.send(Err(err)).unwrap();
            };
        });
        image.set_onload(Some(callback.as_ref().unchecked_ref()));
        image.set_onerror(Some(error_callback.as_ref().unchecked_ref()));
        image.set_src("images/rhb.png");
        rx.await.unwrap();
        let sprite = sheet
            .frames
            .get("Run (1).png")
            .expect("Could not find sprite");
        context.draw_image_with_html_image_element_and_sw_and_sh_and_dx_and_dy_and_dw_and_dh(
            &image,
            sprite.frame.x as f64,
            sprite.frame.y as f64,
            sprite.frame.w as f64,
            sprite.frame.h as f64,
            300.0,
            300.0,
            sprite.frame.w as f64,
            sprite.frame.h as f64,
        );
        let (tx, rx) = futures::channel::oneshot::channel::<Result<(), JsValue>>();
        let success_tx = Rc::new(Mutex::new(Some(tx)));
        let error_tx = Rc::clone(&success_tx);
        let image = web_sys::HtmlImageElement::new().unwrap();
        let callback = Closure::once(move || {
            if let Some(success_tx) = success_tx.lock().ok().and_then(|mut opt| opt.take()) {
                success_tx.send(Ok(())).unwrap();
            };
        });
        let error_callback = Closure::once(move |err: JsValue| {
            if let Some(error_tx) = error_tx.lock().ok().and_then(|mut opt| opt.take()) {
                web_sys::console::log_1(err.as_ref());
                error_tx.send(Err(err)).unwrap();
            };
        });
        image.set_onload(Some(callback.as_ref().unchecked_ref()));
        image.set_onerror(Some(error_callback.as_ref().unchecked_ref()));
        image.set_src("images/Idle (1).png");
        rx.await.unwrap();
        context.draw_image_with_html_image_element(&image, 0.0, 0.0);
    });
    Ok(())
}

#[derive(serde::Deserialize)]
struct Sheet {
    frames: HashMap<String, Cell>,
}

#[derive(serde::Deserialize)]
struct Rect {
    x: u32,
    y: u32,
    w: u32,
    h: u32,
}
#[derive(serde::Deserialize)]
struct Cell {
    frame: Rect,
}
async fn fetch_json(json_path: &str) -> Result<JsValue, JsValue> {
    let window = web_sys::window().unwrap();
    let resp_value = wasm_bindgen_futures::JsFuture::from(window.fetch_with_str(json_path)).await?;
    let resp: web_sys::Response = resp_value.dyn_into()?;
    wasm_bindgen_futures::JsFuture::from(resp.json()?).await
}
