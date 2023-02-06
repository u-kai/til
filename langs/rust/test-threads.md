# テストを並列実行しない

```
cargo test -- --test-threads=1
```

-   ignored などを入れたい場合はまず--test-threads を入れなければ動かなかった

```
cargo test -- --test-threads=1 --ignored
```
