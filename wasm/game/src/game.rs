use std::collections::HashMap;

use anyhow::Result;
use async_trait::async_trait;
use web_sys::HtmlImageElement;

use crate::{
    browser::fetch_json,
    engine::{load_image, Game, KeyState, Rect, Renderer},
};
use serde::Deserialize;

use self::red_hat_boy_states::{RedHatBoyContext, RedHatBoyState};

struct RedHatBoy {
    state_machine: RedHatBoyStateMachine,
    sprite_sheet: Sheet,
    image: HtmlImageElement,
}
impl RedHatBoy {
    fn new(sheet: Sheet, image: HtmlImageElement) -> Self {
        Self {
            state_machine: RedHatBoyStateMachine::Idle(RedHatBoyState::new()),
            sprite_sheet: sheet,
            image,
        }
    }
    fn run_right(&mut self) {
        self.state_machine = self.state_machine.transition(Event::Run)
    }

    fn update(&mut self) {
        self.state_machine = self.state_machine.update();
    }

    fn draw(&self, renderer: &Renderer) {
        let frame_name = format!(
            "{} ({}).png",
            self.state_machine.frame_name(),
            (self.state_machine.context().frame / 3) + 1
        );
        let sprite = self
            .sprite_sheet
            .frames
            .get(&frame_name)
            .expect("Could not find frame");
        renderer.draw_image(
            &self.image,
            &Rect {
                x: sprite.frame.x.into(),
                y: sprite.frame.y.into(),
                width: sprite.frame.w.into(),
                height: sprite.frame.h.into(),
            },
            &Rect {
                x: self.state_machine.context().position.x.into(),
                y: self.state_machine.context().position.y.into(),
                width: sprite.frame.w.into(),
                height: sprite.frame.h.into(),
            },
        );
    }
}

#[derive(Debug, Clone, Copy)]
struct Idle;
#[derive(Debug, Clone, Copy)]
struct Running;

#[derive(Debug, Clone, Copy)]
enum RedHatBoyStateMachine {
    Idle(RedHatBoyState<Idle>),
    Running(RedHatBoyState<Running>),
}
impl RedHatBoyStateMachine {
    fn transition(self, event: Event) -> Self {
        match (self, event) {
            (RedHatBoyStateMachine::Idle(state), Event::Run) => state.run().into(),
            _ => self,
        }
    }
    fn update(self) -> Self {
        match self {
            RedHatBoyStateMachine::Idle(mut state) => {
                state.update();
                RedHatBoyStateMachine::Idle(state)
            }
            RedHatBoyStateMachine::Running(mut state) => {
                state.update();
                RedHatBoyStateMachine::Running(state)
            }
        }
    }

    fn frame_name(&self) -> &str {
        match self {
            RedHatBoyStateMachine::Idle(state) => state.frame_name(),
            RedHatBoyStateMachine::Running(state) => state.frame_name(),
        }
    }
    fn context(&self) -> &RedHatBoyContext {
        match self {
            RedHatBoyStateMachine::Idle(state) => state.context(),
            RedHatBoyStateMachine::Running(state) => state.context(),
        }
    }
}
impl From<RedHatBoyState<Running>> for RedHatBoyStateMachine {
    fn from(state: RedHatBoyState<Running>) -> Self {
        RedHatBoyStateMachine::Running(state)
    }
}
pub enum Event {
    Run,
}

mod red_hat_boy_states {
    use crate::game::Point;

    use super::{Idle, Running};

    const FLOOR: i16 = 475;
    const IDLE_FRAMES: u8 = 29;
    const RUNNING_FRAMES: u8 = 23;
    const IDLE_FRAME_NAME: &str = "Idle";
    const RUN_FRAME_NAME: &str = "Run";
    const RUNNING_SPEED: i16 = 3;
    #[derive(Debug, Clone, Copy)]
    pub struct RedHatBoyState<S> {
        pub context: RedHatBoyContext,
        _state: S,
    }
    #[derive(Debug, Clone, Copy)]
    pub struct RedHatBoyContext {
        pub frame: u8,
        pub position: Point,
        pub velocity: Point,
    }
    impl RedHatBoyContext {
        pub fn update(mut self, frame_count: u8) -> Self {
            if self.frame < frame_count {
                self.frame += 1;
            } else {
                self.frame = 0;
            }
            self.position.x += self.velocity.x;
            self.position.y += self.velocity.y;
            self
        }
        fn reset_frame(mut self) -> Self {
            self.frame = 0;
            self
        }
        pub fn run_right(mut self) -> Self {
            self.velocity.x += RUNNING_SPEED;
            self
        }
    }
    impl<S> RedHatBoyState<S> {
        pub fn context(&self) -> &RedHatBoyContext {
            &self.context
        }
    }
    impl RedHatBoyState<Running> {
        pub fn frame_name(&self) -> &str {
            RUN_FRAME_NAME
        }
        pub fn update(&mut self) {
            self.context = self.context.update(RUNNING_FRAMES);
        }
    }
    impl RedHatBoyState<Idle> {
        pub fn new() -> Self {
            RedHatBoyState {
                context: RedHatBoyContext {
                    frame: 0,
                    position: Point { x: 0, y: FLOOR },
                    velocity: Point { x: 0, y: 0 },
                },
                _state: Idle,
            }
        }
        pub fn frame_name(&self) -> &str {
            IDLE_FRAME_NAME
        }
        pub fn run(self) -> RedHatBoyState<Running> {
            RedHatBoyState {
                context: self.context.reset_frame().run_right(),
                _state: Running,
            }
        }
        pub fn update(&mut self) {
            self.context = self.context.update(IDLE_FRAMES);
        }
    }
}

pub struct WalkTheDog {
    rhb: Option<RedHatBoy>,
}

#[derive(serde::Deserialize, Debug, Clone, Copy)]
pub struct Point {
    pub x: i16,
    pub y: i16,
}

#[derive(serde::Deserialize, Debug, Clone)]
struct Sheet {
    frames: HashMap<String, Cell>,
}

#[derive(serde::Deserialize, Debug, Clone)]
struct Cell {
    frame: SheetRect,
}
impl WalkTheDog {
    pub fn new() -> Self {
        Self { rhb: None }
    }
}

#[async_trait(?Send)]
impl Game for WalkTheDog {
    async fn initialize(&self) -> Result<Box<dyn Game>> {
        let sheet: Option<Sheet> =
            serde_wasm_bindgen::from_value(fetch_json("images/rhb.json").await?)
                .expect("Could not deserialize json");
        let image = Some(load_image("images/rhb.png").await?);

        Ok(Box::new(WalkTheDog {
            rhb: Some(RedHatBoy::new(
                sheet.clone().ok_or_else(|| anyhow::anyhow!("No sheet"))?,
                image.clone().ok_or_else(|| anyhow::anyhow!("No image"))?,
            )),
        }))
    }

    fn update(&mut self, key_state: &KeyState) {
        let mut velocity = Point { x: 0, y: 0 };
        if key_state.is_pressed("ArrowDown") {
            velocity.y += 3;
        }
        if key_state.is_pressed("ArrowUp") {
            velocity.y -= 3;
        }
        if key_state.is_pressed("ArrowLeft") {
            velocity.x -= 3;
        }
        if key_state.is_pressed("ArrowRight") {
            velocity.x += 3;
            self.rhb.as_mut().unwrap().run_right();
        }
        self.rhb.as_mut().unwrap().update();
    }
    fn draw(&self, renderer: &Renderer) {
        renderer.clear(&Rect {
            x: 0.0,
            y: 0.0,
            width: 600.0,
            height: 600.0,
        });
        self.rhb.as_ref().unwrap().draw(renderer);
    }
}

#[derive(Deserialize, Debug, Clone)]
struct SheetRect {
    x: i16,
    y: i16,
    w: i16,
    h: i16,
}
