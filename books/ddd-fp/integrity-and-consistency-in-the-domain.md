# Integrity and Consistency in the Domain

- FP には smart constructor っていうアプローチがあるっぽい

```fs
type UnitQuantity = private UnitQuantity of int

```

- 上のようにすることで constructor を private 化できるっぽい
- これによって同じモジュール内のコード以外は UnitQuantity を作成できなくすることができる
- 絶対一つ以上の要素が入っている必要があるなど，できるだけ型として定義してあげることで，ドキュメント性が向上する

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
