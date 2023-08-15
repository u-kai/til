use std::{rc::Rc, sync::Mutex};

use anyhow::{anyhow, Result};
use futures::channel::oneshot::channel;
use wasm_bindgen::JsCast;
use web_sys::HtmlImageElement;

use crate::browser;

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
