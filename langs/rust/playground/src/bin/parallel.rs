use std::{
    fmt::Debug,
    io,
    net::TcpListener,
    thread::{self, sleep},
    time::Duration,
};

macro_rules! calc_time {
    ($fn:ident,$($args:ident),*) => {
        let timer = std::time::Instant::now();
        $fn($($args),*);
        println!("time = {:?}", timer.elapsed());
    };
}
fn main() {
    let v = vec![None, Some(1), Some(2), None, None];
    let mut exist_users = v.into_iter().filter_map(|s| s);
    let num = exist_users.by_ref().count() as i32;
    let ave: i32 = exist_users.fold(0, |acc, user| acc + user) / num;
    fn parallel_print(len: usize) {
        let mut handler = Vec::new();
        for _ in 0..len {
            let fork = thread::spawn(|| {
                sleep(Duration::from_secs(1));
            });
            handler.push(fork);
        }
        for h in handler {
            h.join().unwrap();
        }
    }
    fn print(len: usize) {
        for _ in 0..len {
            sleep(Duration::from_secs(1))
        }
    }
    let d = 10;
    calc_time!(parallel_print, d);
    calc_time!(print, d);
}
