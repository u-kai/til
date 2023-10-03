# Domain Modeling with Types

- ドメインエキスパートの人もレビューできるようなドキュメント性のあるプログラムコードにしないといけない
- UML のようなものを必要とせずにソースコードに直書きで伝わるのが最終ゴールで良いとのこと

## Seeing Patterns in a Domain Model

## Modeling Simple Values

- ドメインエキスパートは int とか string とかは思わず，OrderIds とか ProductCode と思う
  - これは確かにで，プログラムに慣れてしまうと，int,string という選択肢が最初に出てしまうが，プログラムを全く知らない人に，例えば人の名前ってどんなものですか？って聞いて，文字列，と答える人はいなそうだし，それで OK(人の名前は人の名前)

## Avoiding Performance Issues with Simple Types

- なんか型をつけることによるオーバヘッドの話
- F#固有の話な気がする
  - Rust では new type はゼロコストのはず
- コンパイル時に型情報をつけなければいいのでは？と思うが結構むずいんかな？

## Modeling Complex Data

- Undefined 型で未定義部分のコードを先に書くことをお勧めしている？
  - そしてできてないところはもう少しエキスパートの人と話し合いながら決定する感じっぽい？

## Modeling Workflows with Functions

- これまでは型に対して名刺のユビキタス言語をつけてきたが，これからは動詞(workflow)の関数型をモデリングする

```fss
type ValidateOrder = UnvalidatedOrder -> ValidatedOrder
```

## Value Objects

- F#はデフォルトで全て比較可能なフィールドであればすぐに比較してくれる
  - 逆に rust はなんで derive しないといけないのかな？
    - 色々理由はありそうだけど
    - 結局関数が増えることになるので，メモリ領域は増えるよね？って話かな？

## A Question of Identity:Entities

- Entities は Id があったり life cycle があるもので，モデリングする際は Value Objects か Entities かどうかはコンテキストに依存する
- Value Objects

## Identifiers for Entities

- Entities は変わっていく値を属性として多く持つことになるが，ID だけは変わらないように設計しなくてはいけず，なかなかむずいよねっていうはなし

## Adding Identifiers to Data Definitions

- ID を型に付加する時にどうすればいいかの話
- outside
  - ID を複数のケースとは関連させず，一つの概念に関連させるやり方

```fs
type UnpaidInvoiceInfo = ...
type PaidInvoiceInfo = ...

type InvoiceInfo =
| UnpaidInvoiceInfo
| PaidInvoiceInfo

type InvoiceId = ...

type Invoice = {
    InvoiceId,
    InvoiceInfo
}
```

- これだと Unpaid か，Paid を判定することが難しい

- inside
  - ID を複数のケースとそれぞれ関連させるやり方

```fs
type InvoiceId = ...
type UnpaidInvoice = {
    InvoiceId
    // and other Info for Unpaid case
}
type PaidInvoice = {
    InvoiceId
    // and other Info for Paid case
}

type Invoice =
| Unpaid
| Paid

```

- これによって patter match を使った ID へのアクセスなどが簡単になる

## Implementing Equality for Entities

- ID が一緒であれば等価なので，eq と hash のメソッドを override する必要がある
- ただしこれは後にバグをうむ可能性があるため，NoEquality を付与して明示的に普通の等価計算はできないようにすることには利点がある
  - rust ではデフォルトで比較できないので OK??

## Aggregates

- なんでも一つの Aggregates に入れたくなる時もあるがしっかり考えること

  - Aggregates そのものではなく，ID ではだめか？など
  - お互いの関連具合や独立具合を考えること
  - 設計を考えるとまた DB の保存にも的するようになる
    - でもこれって，DB のことは考えないでモデリングしようねとバッティングしてないか？
    - 副次的な効果として DB のこともよくなるであればわかるが，そういうことだと信じたい

- Aggregate の重要なまとめでわかったが，一貫性の境界っていう意味がわかってきた
- Aggregate に変更ライフサイクルや関連度合いが強いものを詰めておいて，Aggregate に変更のお願いをする
- それによって一貫性が保たれる
- DB のトランザクションとかも Aggregate 一体であれば分散せずに良いよねってことね
- 確かによく考えられているかも
- アトミックなもの
- Entity とかのあつまりでモデリングの際に重要.特に関連の薄いものを一緒にしないこと
- root の ID を ID として持っておくべし
