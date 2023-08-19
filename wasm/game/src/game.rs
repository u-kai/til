use std::collections::HashMap;

use anyhow::Result;
use async_trait::async_trait;
use web_sys::HtmlImageElement;

use crate::{
    browser::fetch_json,
    engine::{load_image, Game, Rect, Renderer},
};
use serde::Deserialize;

pub struct WalkTheDog {
    image: Option<HtmlImageElement>,
    sheet: Option<Sheet>,
    frame: u8,
}
#[derive(serde::Deserialize, Debug)]
struct Sheet {
    frames: HashMap<String, Cell>,
}

#[derive(serde::Deserialize, Debug)]
struct Cell {
    frame: SheetRect,
}
impl WalkTheDog {
    pub fn new() -> Self {
        Self {
            image: None,
            sheet: None,
            frame: 0,
        }
    }
}

#[async_trait(?Send)]
impl Game for WalkTheDog {
    async fn initialize(&self) -> Result<Box<dyn Game>> {
        let sheet: Sheet = serde_wasm_bindgen::from_value(fetch_json("images/rhb.json").await?)
            .expect("Could not deserialize json");
        let image = load_image("images/rhb.png").await?;

        Ok(Box::new(WalkTheDog {
            image: Some(image),
            sheet: Some(sheet),
            frame: self.frame,
        }))
    }

    fn update(&mut self) {
        if self.frame < 23 {
            self.frame += 1;
        } else {
            self.frame = 0;
        }
    }
    fn draw(&self, renderer: &Renderer) {
        let current_sprite = (self.frame / 3) + 1;
        let frame_name = format!("Run ({}).png", current_sprite);
        let sprite = self
            .sheet
            .as_ref()
            .and_then(|sheet| sheet.frames.get(&frame_name))
            .expect("Could not find frame");
        renderer.clear(&Rect {
            x: 0.0,
            y: 0.0,
            width: 600.0,
            height: 600.0,
        });
        self.image.as_ref().map(|image| {
            renderer.draw_image(
                &image,
                &Rect {
                    x: sprite.frame.x.into(),
                    y: sprite.frame.y.into(),
                    width: sprite.frame.w.into(),
                    height: sprite.frame.h.into(),
                },
                &Rect {
                    x: 300.0,
                    y: 300.0,
                    width: sprite.frame.w.into(),
                    height: sprite.frame.h.into(),
                },
            );
        });
    }
}

#[derive(Deserialize, Debug)]
struct SheetRect {
    x: i16,
    y: i16,
    w: i16,
    h: i16,
}
