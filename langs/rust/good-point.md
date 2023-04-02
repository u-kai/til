---
marp: true
---

# 最先端？言語 Rust の紹介

---

# 注意事項

- 本発表は自組織とは全く関係がありません.
- Rust にバイアスがかかった発表内容かもしれません.ご了承ください.

---

# 目次

- 本日の狙い
- Rust とは？
- 私の推しポイント
- おすすめの Rust コンテンツについて

---

# 本日の狙い

- 一人でも多くの方に Rust に興味持ってもらう!
  - よくわからんけど，なんかすごそう！試してみよう！と思ってもらえるように資料を作成しました！
- あわよくば弊社 Rust の輪を広げる!
  - Rust で書きたいんです...🙇

---

# Rust とは

## よく言われていることは．．．

- 2015 年に ver1.0 がリリースされた比較的若い言語
- stack overflow の愛されている言語ランキング 7 年連続 1 位！
- 実行スピードは C,C++ 言語並みにも関わらずメモリ安全！
- 安全性と実行スピードから, 他言語からの置き換えとして名だたる企業が注目！
  - Amazon
  - Google
  - Microsoft ...

---

# C 言語並みのスピードなんてほとんどのケースで不要では？🤔

---

# スピードだけじゃないんです！！👍

# 仕様や標準パッケージなどもヤバいんです！！😍

# そして何より．．．

---

# 書いてて楽しい！！！！！

## これが私の 1 番の推しポイントです!

---

# 1.堅牢でスマートな型システム

- Rust は静的型付け言語の中でも強い型整合性でいろんな種類の型がある
- 型整合性さえクリアすればコードが動くと自信を持てる！💪
- 型が多くを語ってくれるため，可読性が高いコードができる!👍

```rust
fn bin_to_hex(bin:&str)->Result<String,Error> {
  // 名前から察するに2進数を16進数に変更するのかな？
  // 引数は文字列なのね
  // 失敗するかもしれないのね！確かにbinに01以外のものがあれば失敗するよね
}
```

- 型推論がエグいから，定義以外で型を書くことはほとんどない！🥰

```rust
let data:String = String::from("data")
// わざわざ型を書かなくてもOK
let data = String::from("data")
```

---

# でも型の整合性が強いと柔軟性がなくならない？？🤔

```rust
fn my_print(data:&str){
  println!("{}",data);
}
// OK
my_print("Hello world!");

// NG
my_print(100);

// NG
// Rustにはたくさんの文字列を表現する型があります．．．
my_print(String::from("Hello world!"));
```

---

# trait やジェネリクスによる柔軟性

```rust
// traitを定義
trait Printable:Debug {}

// ジェネリクスを使って関数の引数を定義
// このPはPrintableを実装していれば何でもOK!
fn my_print_use_generic<P:Printable>(data:P){
  println!("{:#?}",data);
}

// traitのすごいところは型の定義よりも後に実装を行うこと!
// 語弊を恐れずに言うなら，どんな型でも定義を変更せずに後から振る舞いを変えられると言うこと
impl Printable for usize {}
impl Printable for str{}
impl Printable for String{}
impl Printable for MyStruct{}
// Third partyが定義した型にも独自の振る舞いを追加できる(逆もまた然り)
// interfaceではできない(はず．．．)
impl Printable for LibraryStruct{}

// All OK
my_print_use_generic("Hello world!");
my_print_use_generic(String::from("Hello world!"));
my_print_use_generic(100 as usize);
my_print_use_generic(MyStruct::new());
my_print_use_generic(LibraryStruct::new());
// NG
// -100は型推論されるとi32型
// i32型にPrintableは実装されていないのでNG
my_print_use_generic(-100)
```

---

# 2.test が言語仕様に組み込まれている

- 多くの言語はテストライブラリのインストールが必要
- Rust だと下のようにコードを書いて cargo test で OK

```rust

fn add(left:i32,right:i32) -> i32 {
    left + right
}

#[cfg(test)]
#[test]
fn test_add(){
    assert_eq!(add(3,4),7);
}
```

- しかもテストを実装コードと同じファイルに書ける!
  -> 書きやすい＆コードの挙動をすぐに確認できる!

---

# 3. コンパイラや標準機能が優しい

- 標準機能のほとんどに doctest が記入されているので迷うことがない!
- コンパイラが丁寧に教えてくれる！

  - コンパイラの指示に従っていればあらかた上手くいく！
  - ペアプロしているみたい

---

# 4.macro が面白い！

- 名前の後ろに!がついているもの

```rust
println!("{}",data);
assert_eq!(1,1);
let concat:String = format!("{}/{}"month,day);
let groups:Vec<&str> = vec!["tic","it","dict"];
```

- 関数と似ているけど違うもの
  - 関数のように振る舞うものも多いが実態は違う
- 関数は実行時にコードが実行されるが macro はコンパイル時にコードとして展開される
- コードが生成されるイメージ!

---

# macro の定義例

```rust
// define of macro
macro_rules! fn_time {
    ($fn:ident,$($args:ident),*) => {
        let timer = std::time::Instant::now();
        $fn($($args),*);
        println!("time = {:?}", timer.elapsed());
    };
}

fn heavy_calc(sec:usize) {
    sleep(Duration::from_secs(sec));
}
// use of macro
fn_time!(heavy_calc,1);
// time = 1.00002s
```

---

# macro が展開された時のイメージ

```rust
// before
fn_time!(heavy_calc,1);
```

```rust
// after
let timer = std::time::Instant::now();
heavy_calc(1);
println!("time = {:?}", timer.elapsed());
```

- 関数では到達困難な表現力！
- ただし，可読性は低くなりがちなので，本番プロジェクトではやりすぎ注意！
- 遊ぶ分には最高のおもちゃ！！

---

# おすすめの Rust コンテンツ

- rust the book
  - 無料の Web コンテンツ
  - 基本的なことが丁寧に解説されている
  - https://doc.rust-jp.rs/book-ja/
- プログラミング Rust 第 2 版
  - 詳しくわかりやすく説明されている
  - 本格的に始めるならおすすめです
- rust playground
  - rust の web 実行環境
  - 公開されているライブラリも使うことができる!
  - https://play.rust-lang.org/?version=stable&mode=debug&edition=2021

---

# 終わりに

- 本日はありがとうございました！
- もしご興味あれば私までご連絡ください！
  - Rust の凄さはこれだけでは決してありません！笑
- みなさまが Rust に興味を持っていただけ，挑戦していただけましたら幸いです!

---

# ありがとうございました！

---

# 補足：本番プロジェクトで採用する旨み(予想)

- 型システムによって型安全で品質の高いシステムが作りやすい
- テストが実装の近くに書けることや高度な型システムなどのおかげで，可読性や保守性が高いコードになりやすい
- 宣言的にかける API が多く，うまくかければ可読性の高いものにできる

```python
user_data = get_users()
age_sum = 0
user_num = 0
for user in user_data:
  if user.age is None:
    continue
  age_sum += user.age
  user_num += 1
```

```rust
let user_data:Vec<User> = get_users();
let users_holding_age = user_data.iter().filter_map(|user|user.age);
let count =  users_holding_age.by_ref().count();
let ave = users_holding_age.fold(0,|acc,user|acc+user.age) / count;
```

- 実行時間が早いことや，メモリ消費量を抑えられることから計算リソースの節約が可能

---

# 補足：本番プロジェクトで採用する辛み(予想)

- 学習コストが高い
- クレート(ライブラリ)の充実度が発展途上
  - AWS SDK はまだ Developer Preview
  - 他 サービスの SDK についても基本 Rust 製のものは少ない
- 多機能ゆえにコードの統制が難しいかも
  - for 文を使うのかイテレーターのメソッドを積極的に使うのか？
  - macro を使ってコード量を減らすのか愚直に書いて可読性を上げるのか

---

# 補足：伝えるか悩んだ項目

- if 式
- if let ,let else
- イテレーター
- パターンマッチ
- 手続きマクロ
- 所有権
- ライフタイム
- Rust を学ぶと嬉しい副次効果

---

# Rust を学んでよかったこと

- メモリの割り当てやポインタなど深い知識を身につけることができた
- 設計力が上がった気がする
  - 強力な型システム
  - trait 境界
  - 所有権
- test を書きたい欲が強くなった & test の大切さを再認識できた

---

# Rust を学んでよくなかったこと

- Rust しか書きたく無くなった 😢

---

# 補足：今後やりたいこと

- Rust で大規模開発
- wasm

---

# 型があることの恩恵

```python
def add(i,j):
    return i + j
```

```rust
fn add(i:i32,j:i32) -> i32 {
    i + j
}
```

---

# 堅牢な型システム Option 型

```rust
let data = [1,2,3];
println!("{}",data[3]);
// error: this operation will panic at runtime
// index out of bounds: the length is 3 but the index is 3
```

```rust
let data = [1,2,3];
let maybe_d3:Option<&usize> = data.get(3);

match maybe_d3 {
    Some(d) => println!("{}",d),
    None    => println!("data len is {}",data.len()),
}
// data len is 3

sum = maybe_d + 3;
//error[E0369]: cannot add `{integer}` to `Option<&{integer}>`
```

- 型レベルで存在しない可能性がわかる
- Option 型にはたくさんの便利メソッドがある
- 型によって守られているので，コンパイラに任せてストレス無く開発できる

---

# 堅牢な型システム Result

```python
with open("test.txt","r") as f:
    data = f.read()
# FileNotFoundError: [Errno 2] No such file or directory: 'test.txt'
```

```rust
let data = read_to_string("test.txt");
match &data {
    Ok(d)     => println!("{:#?}",d.split("\n").next()),
    Err(e)    => //something error case
}
lines = data.split("\n");
//error[E0599]: no method named `split` found for enum `Result` in the current scope
```

- 型レベルで失敗する可能性があるのかがわかる!
- Result 型にはたくさんの便利メソッドがある

---

```rust
// Resultは成功しているかどうかの確認を強いる->エラーチェックを忘れない
fn read_to_string(file:String)->Result<String,Error>;
let data = read_to_string("test.txt");

let lines = data.split("\n");
//error[E0599]: no method named `split` found for enum `Result` in the current scope

let lines = match &data {
    Ok(d)     => d.split("\n").next().unwrap(),
    Err(e)    => ""
}
//ok!
```

---

# 型がないと前提知識が必要なことがある

```python
# 誰かが書いたコード(または半年前の自分が書いたコード)
def get_child_from_class(class_,name):
    return class_.get(name)
```

- 関数名から察するに child を返すみたいだけど，child ってどんな値？
- name は str でいいのかな？class\_に name の child がいなかったらどうなるんや？
- そもそも class\_ってどんな型?

---

# 型は多くを語ってくれる

```rust
// 誰かが書いたコード(または半年前の自分が書いたコード)
fn get_child_from_class(class:Class,name:String)->Option<Child>{
    class.get(name)
}
```

- Class 型から文字列を入れると Child 型が帰ってくるんだな
- しかもその Child は存在しない場合があるから，その際のことも考えよう！

---

# 便利な機能

- if let/let else

```rust
if let Some(name) = group.get("kai") {
  // process use name
}
let Some(name) = group.get("kai") else {
  return
}
// process use name
```

---

# 型の紹介: mut 型

- 値が変更するか否かに対しても型がつけられています!

```rust
fn add_str(data:&mut String,added:&str){
  data.push_str(&added);
}

let mut h = String::from("Hello")

// hが書き変わっていることが明白
add_str(&mut h,"world")

println!("{} ",h);
// Hello world
```

- データが知らないうちに変更されている...っていうのが型レベルで明確になる!

---

# 型の紹介: Result,Option 型

- 失敗しうるものや，存在しない可能性があるものに対する型があります
- チェックをコンパイラが強制してくれるので，安全なコードが書ける!

```rust
let data = [1,2,3];
println!("{}",data[3]);
// error: ...
// index out of bounds: the length is 3 but the index is 3

let maybe_d3:Option<&usize> = data.get(3);

//error[E0369]: cannot add `{integer}` to `Option<&{integer}>`
let sum = maybe_d + 3;

// Option型やResult型などの展開を援助してくれるmatch(本当はもっとすごいです！)
let sum = match maybe_d3 {
    Some(d) => d + 3,
    None    => 0
}
```
