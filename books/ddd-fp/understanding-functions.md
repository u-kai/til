# Understanding Functions

- oop では interface や DI をつかうが，fp では関数のパラメータ化を使う
- oop では継承やデコレーターパターンを使って DRY を達成するが，fp では再利用性のある関数やそれらを統合していく
- oop は全てがもの，fp は function はもの
- カリー化は一つのパラメータをとって一つの値を返すもの
- 二つのパラメータを撮る関数をカリーかすると P1 -> P2 -> T のようになる
- これによって，柔軟に多くの関数を作成することができる

- 手続き型でもできなくはないが，圧倒的に関数型でやる方が楽
  - 手続き型でプログラムを書くときも関数型の思想を取り入れてカリー化をすることってあるんかな？
  - それなら関数型使えよとはなりそう

```go
func SayGreeting(greeting string) func(string) {
    return func (name string) {
        fmt.Printf("%s %s",greeting,name)
    }
}


sayGoodbye := SayGreeting("Goodbye")
sayGoodBye("kai")
```

## Total Functions

- 型情報で嘘をつかないことは大切

## Composition

- 型を変換して新しい型を作る際に Composition を使ってみる
- こうすることで，細部の変換関数や変換される途中の型を気にすることがなくなる
- 小さい関数を組み合わせて完全な workflow を作成する
- パイプ演算子を使って関数の Composition を達成することができる
- workflow を並行に Compose することで一つの Application ができる
