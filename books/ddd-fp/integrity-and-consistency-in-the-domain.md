# Integrity and Consistency in the Domain

- FP には smart constructor っていうアプローチがあるっぽい

```fs
type UnitQuantity = private UnitQuantity of int

```

- 上のようにすることで constructor を private 化できるっぽい
- これによって同じモジュール内のコード以外は UnitQuantity を作成できなくすることができる
- 絶対一つ以上の要素が入っている必要があるなど，できるだけ型として定義してあげることで，ドキュメント性が向上する

## The Integrity of Simple Values

## Units of Measure

## Enforcing Invariants with the Type System

## Capturing Business Rules in the Type System

- flag による型はいくつもの問題がある
- 対処方法としてはドメインをよく観察すること
- 型システムにビジネスルールを載せることができれば，不正値がコードに入ることはなく，そのための単体テストを書く必要性もなくなる
- User が Email か Postal かそれともどっちもかの時，どのような型宣言をすれば良いのか？
- 答えは全てのケースを型として定義して，最後に enum として定義すること
- とことん型定義をすること
- そうすることで test なしで，うまくいくケースはうまくいくと信じることができる(invalid なデータや未検証な型を正常な型と分ける形で定義しているため)

## Consistency

- 一貫性は設計の重荷になるため，できるだけ避けるべし

- 様々な Aggregate での一貫性が書かれていておもろい
- 一つの Aggregate につき一つのトランザクションが好ましい
- もし難しいかつ，コンテキスト境界外であれば Event などを飛ばして，結果整合性を目指すのもあり
- もしおなじコンテキスト境界で 2 つの Aggregate に一つのトランザクションが必要になった時，Aggregate の観察が足りないかもしれない
- 関数型の関数を使い回すことは OOP のそれよりも影響が少ない？
  - なぜなら関数は特定のオブジェクトに紐づいていなく，グローバルな状態にも依存していないから
  - ここはよくわからん
