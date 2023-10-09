# Working with Errors

## Working with Domain Errors

- 全てのエラーについて型づけすることは難しいからまずはエラーのクラス分けをするべき
- 以下の 3 つのエラーに分けることができる

  - Domain Errors
    - ビジネスプロセスで予想することができ，ドメインの設計にて考慮する必要があるもの
  - Panics
    - 道の状態に陥ったときでで例えば，out of memory のような対処不能のものや 0 割などのプログラマーの見落とし
  - Infrastructure Errors
    - ネットワークのタイムアウトや認証失敗など

- これを見分けるならドメインエキスパートに聞くのがよい

  - 例では loadbalancer の前でミスったらどうする？っていう問いかけをしており，ドメインエキスパートは？？？となっているため，これはドメイン Error ではない

- ドメインエラーは他のものと同様にドメインの一部

  - できるだけ型システムで表現すること

- panics は main 関数などの一番上の層でハンドルすべし
- Infrastructure error は上記のどちらでも対処可能

  - これはアーキテクチャに依存する

- ドメインのエラーと同じように扱いすぎると開発者が多くのことを考える必要があるので大変
- だた，確かに特定の場合においてはドメインエキスパートにエスカレーションする必要がある
  - 例えばリモートアドレス検証サービスが利用できなかったら，どのようにビジネスプロセスを変えるべきか？なにを顧客に伝えるべきか？
  - このような問いは開発者だけでは解決できない，ドメインエキスパートやシステムのおー何考えてもらうこと

## Modeling Domain Errors in Types

```f#
type PlaceOrderError =
| ValidationError of string
| ProductOutOfStock of ProductCode
| RemoteServiceError of RemoteServiceError

```

- RemoteServiceError は infra の関心ごとだが，ドメイン的に大切なことなので入っている

  - 技術的レイヤー感がバラバラな気がするが，ドメイン的に関心があることは全て突っ込む感じか

- choice 型を使うのは変更に対して難しいが，コンパイラーが変更を検知してくれて安全だよねってある

  - 逆にこれって interface とかだとコンパイラーが何も言わないから柔軟っていうけど，それってほんまに大丈夫なんか？とは思う
    - switch だと検査し忘れるし，共通の振る舞いを実行するようにしている側は大丈夫だけど，この型を作成する側はケースとして忘れる可能性あるよな
    - だから，choice 型を使ってパターンマッチするところはできるだけ少なくして隠蔽することが大事だけど，choice 型自体は interface を使う場合よりも良い気がする

- choice 型にしたから edge case についてもドメインエキスパートと喋れるよね

## Chaining Result-Generating Functions

- two-track 関数は bind か flatMap としてよく fp で使われる
- 入力はスイッチ関数で，出力は新しい two-track 関数で，input と output が two-track なもの
- もし two-track input が成功したら，スイッチ関数を pass する
- もし失敗したら，スイッチ関数をバイパスして失敗を返す

```f#
let bind switchFn =

    fun twoTrackInput ->
        match twoTrackInput with
            |Ok success -> switchFn success
            |Error failure -> Error failure

```

- mapError で error の方変換ができて，簡単に合成可能になる

## Using bind and map in Our Pipeline

- error の型を pipeline で一致してあげるのは結構やるみたい
  - rust で map_err のために同じようなことをしていて，少し冗長か？と思っていたが問題なさそう
  - ただし，おそらく domain error に対してのみだと思われる
- error が発生するものは bind でくっつけて，発生しないものは map でくっつける？
  - one-track なので
  - ただ，map するところまで来ているということは error 型ではない気がするので，わざわざ Result.map する必要がある？
  - ただ，Result は関数全体にかかっているから失敗し得ない関数に対しても input は Result を map したものでありたいということ？
  - rust で bind ってあるんやろか？try_into(シンタックスシュガーは?)がこれに当たる？

## Adapting Other Kinds of Functions to the Two Track Model

## Monads and More

- monad はモナディック関数を連続的に繋げるようにできるパターン
- ここでいうモナディック関数というのは普通の値を引数にとって，いくつかの強化された値を返すもの
- Result 型で包む行為は強化に値する
- monad はシンプルな以下の 3 つの要素における用語

  - データ構造
    - ここでいうと Result 型
  - いくつかの関連している関数

    - return と bind を持っているべき
      - return(pure として知られている) は monadic 型の通常の返り値を返す関数?
      - Result であれば Ok コンストラクタ
    - bind(flatMap としても知られている)は他の monadic 関数と連鎖しやすいようにする

  - 関数がどのように動作しなければならないかに関するいくつかの規則

## Composing in Parallel with Applicative

- Applicative はモナドに似ているが monadic 関数をチェインしやすい
- Applicative は monadic を並行に合成させてくれる
