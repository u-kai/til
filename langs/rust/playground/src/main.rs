use std::io::stdin;

fn main() {
    let etos = [
        "子", "丑", "寅", "卯", "辰", "🐍", "午", "未", "申", "酉", "戌", "亥",
    ];
    let mut s = String::new();
    stdin().read_line(&mut s).unwrap();
    s.remove(s.len() - 1);
    while s.parse::<usize>().is_err() || s.parse::<usize>().unwrap() < 1900 {
        println!("1900 以上を入れてください");
        s.clear();
        std::io::stdin().read_line(&mut s).unwrap();
        s.remove(s.len() - 1);
    }
    let mod_ = (s.parse::<usize>().unwrap() - 1900) % 12;
    println!("{}", etos[mod_]);
}
