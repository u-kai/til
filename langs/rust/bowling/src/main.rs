fn main() {
    println!("Hello, world!");
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct HitPin(u8);
impl HitPin {
    fn new(pin: u8) -> Self {
        Self(pin)
    }
    fn add(self, other: Self) -> Self {
        Self(self.0 + other.0)
    }
    fn add10(mut self) -> Self {
        self.0 += 10;
        self
    }
    fn addDouble(mut self) -> Self {
        self.0 += 20;
        self
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Score(u8);
impl From<HitPin> for Score {
    fn from(hit_pin: HitPin) -> Self {
        Self(hit_pin.0)
    }
}

#[derive(Debug, PartialEq, Eq, Clone)]
enum Frame {
    NotYet(NotYet),
    FirstThrow(FirstThrow),
    SecondThrow(SecondThrow),
    Spare(Spare),
    Strike(Strike),
}
impl Frame {
    fn new() -> Self {
        Self::NotYet(NotYet::new())
    }
}

#[derive(Debug, PartialEq, Eq, Clone)]
struct NotYet;
impl NotYet {
    fn new() -> Self {
        Self
    }
    fn throw(self, hit_pin: HitPin) -> Frame {
        if hit_pin.0 == 10 {
            Frame::Strike(Strike(hit_pin))
        } else {
            Frame::FirstThrow(FirstThrow(hit_pin))
        }
    }
}
#[derive(Debug, PartialEq, Eq, Clone)]
struct FirstThrow(HitPin);
#[derive(Debug, PartialEq, Eq, Clone)]
struct SecondThrow(HitPin, HitPin);
#[derive(Debug, PartialEq, Eq, Clone)]
struct Spare(HitPin, HitPin);
#[derive(Debug, PartialEq, Eq, Clone)]
struct Strike(HitPin);

#[derive(Debug, PartialEq, Eq, Clone)]
enum LastFrame {
    NotYet,
    FirstThrow(HitPin),
    SecondThrow(HitPin, HitPin),
    Spare(HitPin, HitPin),
    Strike(HitPin, HitPin),
    Bonus(HitPin, HitPin, HitPin),
}

#[derive(Debug, PartialEq, Eq, Clone)]
struct Frames<const T: usize> {
    frames: [Frame; T],
    last_frame: LastFrame,
}

impl Frames<9> {
    fn new() -> Self {
        Self {
            frames: [
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
                Frame::new(),
            ],
            last_frame: LastFrame::NotYet,
        }
    }
}
