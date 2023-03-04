use std::sync::{
    mpsc::{channel, Receiver, Sender},
    Arc, Mutex, RwLock,
};

type PlayerId = u32;

const GAME_SIZE: usize = 8;
type WaitingList = Vec<PlayerId>;
struct FernEmpireApp {
    waiting_list: Mutex<WaitingList>,
}
impl FernEmpireApp {
    fn join_waiting_list(&self, player: PlayerId) {
        let mut guard = self.waiting_list.lock().unwrap();
        guard.push(player);

        if guard.len() == GAME_SIZE {
            let players = guard.split_off(0);
            self.start_game(players)
        }
    }
    fn start_game(&self, players: Vec<PlayerId>) {
        println!("{:#?}", players)
    }
}

struct SharedReceiver<T>(Arc<Mutex<Receiver<T>>>);
impl<T> Iterator for SharedReceiver<T> {
    type Item = T;
    fn next(&mut self) -> Option<Self::Item> {
        let guard = self.0.lock().unwrap();
        guard.recv().ok()
    }
}
fn shared_channel<T>() -> (Sender<T>, SharedReceiver<T>) {
    let (s, r) = channel();
    (s, SharedReceiver(Arc::new(Mutex::new(r))))
}

fn main() {
    let app = Arc::new(FernEmpireApp {
        waiting_list: Mutex::new(Vec::new()),
    });
}
