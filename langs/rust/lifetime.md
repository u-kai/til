# lifetime について

## Box dny TRAIT だとデフォルトのライフタイムが'static

## クロージャ内で状態を持った値を取り扱うのはむずい

-   FnMut または Fn のクロージャーないにおいて，状態を保持するデータを持つことは結構むずい
    -   できるのかわからん
-   なのでクロージャーではなく,共通の trait を実装した struct で代替するのがいいかも

```rust
let v = vec!["1","2"];
let containe = |number_str:&str|->bool {
   v.containes(&number_str)
};

for _ in 0..10 {
    // error
    // containeはFnOnce方になってしまう
    // 保持している状態によって関数の振る舞いを変更したければstructが無難かもしれん
    println!("{}",containe);
}
```
