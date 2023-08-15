use anyhow::{anyhow, Result};
use futures::Future;
use wasm_bindgen::{prelude::Closure, JsCast, JsValue};
use wasm_bindgen_futures::JsFuture;
use web_sys::{
    CanvasRenderingContext2d, Document, HtmlCanvasElement, HtmlImageElement, Response, Window,
};

macro_rules! log {
    ($($t:tt)*) => {
        web_sys::console::log_1(&format!($($t)*).into());
    };
}
pub fn window() -> Result<Window> {
    web_sys::window().ok_or_else(|| anyhow!("no window"))
}

pub fn document() -> Result<Document> {
    window()?.document().ok_or_else(|| anyhow!("no document"))
}

pub fn canvas() -> Result<HtmlCanvasElement> {
    document()?
        .get_element_by_id("canvas")
        .ok_or_else(|| anyhow!("no canvas"))?
        .dyn_into()
        .map_err(|element| anyhow!("element is not a canvas {:#?}", element))
}

pub fn context() -> Result<CanvasRenderingContext2d> {
    canvas()?
        .get_context("2d")
        .map_err(|js_value| anyhow!("could not get context {:#?}", js_value))?
        .ok_or_else(|| anyhow!("could not get context"))?
        .dyn_into()
        .map_err(|element| anyhow!("element is not a canvas {:#?}", element))
}
pub fn spawn_local<F>(future: F)
where
    F: Future<Output = ()> + 'static,
{
    wasm_bindgen_futures::spawn_local(future);
}
pub async fn fetch_with_str(resource: &str) -> Result<JsValue> {
    JsFuture::from(window()?.fetch_with_str(resource))
        .await
        .map_err(|err| anyhow!("fetch error: {:#?}", err))
}
pub async fn fetch_json(resource: &str) -> Result<JsValue> {
    let resp_value = fetch_with_str(resource).await?;
    let resp: Response = resp_value
        .dyn_into()
        .map_err(|err| anyhow!("fetch error: {:#?}", err))?;
    JsFuture::from(
        resp.json()
            .map_err(|err| anyhow!("fetch error: {:#?}", err))?,
    )
    .await
    .map_err(|err| anyhow!("fetch error: {:#?}", err))
}

pub fn new_image() -> Result<HtmlImageElement> {
    HtmlImageElement::new().map_err(|err| anyhow!("new image error: {:#?}", err))
}
pub fn closure_once<F, A, R>(fn_once: F) -> Closure<F::FnMut>
where
    F: 'static + wasm_bindgen::closure::WasmClosureFnOnce<A, R>,
{
    Closure::once(fn_once)
}
