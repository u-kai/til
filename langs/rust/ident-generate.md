# Rust のマクロで識別子を生成する方法

-   paste クレートを使うと以下のように関数名などの識別子を自動生成できる

```
macro_rules! call_with_n {
    ( $num:expr ) => {
        paste::item! {
            fn [<call_with_ $num>]() {
                f($num)
            }
        }
    };
}
```
- [\<default_name $ident\>]の形式でできるらしい
- 参考
    - https://qiita.com/termoshtt/items/6e008905d3f5f85ce707
