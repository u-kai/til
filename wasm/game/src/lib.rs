#[macro_use]
mod browser;
use std::collections::HashMap;
use std::rc::Rc;
use std::sync::Mutex;

mod engine;
use browser::document;
use browser::window;
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
    console_error_panic_hook::set_once();
    let context = browser::context().expect("Could not get context");
    browser::spawn_local(async move {
        let sheet: Sheet = serde_wasm_bindgen::from_value(
            browser::fetch_json("rhb.json")
                .await
                .expect("Could not fetch json"),
        )
        .unwrap();
        let image = engine::load_image("rhb.png")
            .await
            .expect("Could not load image");
    });
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
