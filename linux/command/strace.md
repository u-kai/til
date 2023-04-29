- 使い方

```
strace -o output.log ./exe
```

- 説明
  - strace の出力は 1 行が１つのシステムコール発行に対応している

strace -T -o hello.log ./hello
これでどのくらいの時間がかかったのか計測することができる
