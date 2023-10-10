# Persistence

## Pushing Persistence to the Edges

- これまでの章では I/O を考えずに workflow を作っていたため I/O のことを考えるには 2 つの workflow が必要

  - ビジネスロジックが入っている中心的なもの
  - I/O のコードが入っている edge 的なもの

- 最初は DB にすごいアクセスするものでもドメインロジックのみに摘出することは可能でその例が書かれている
- 永続化の単体テストは取るに足らないので不要

  - 統合テストに力を注ぐべき

- I/O で純粋な関数を挟むべきで，純粋な関数は独立して純粋であるべき
- もし，I/O の次に純粋な関数を置きたくなったら，さらに I/O の関数を後においてサンドイッチできるようにするべき
- Repository パターンはいいパターンだが FP では出てこない
- I/O は全て edge に押しやられるので必要ない??

- 一つの interface に数十個のメソッドを入れる必要がなく，それぞれの I/O に関する関数を必要なだけ定義すればよくなるのでメンテナンス性が上がる

## Command Query Separation

- fp では insert した後の新しい DataStore を作るような関数定義をしている

```
type InsertData = DataStoreState -> Data -> NewDataStoreState
```

- fp での CQS は
  - データを返す関数は副作用がないべき
  - 副作用がある関数はデータを返さないべき，unit を返すべき

## Command-Query Responsibility Segregation

- read と write の結果を同じような 型で表現したくなるが，それはよした方が良い
- はじめに，Query の結果は非正規化されていることが多いが，Command はそうではない

  - そして，write の時には生成 ID やバージョンなどの情報は使用されないが Query の時はある

- 一つの型で多くの目的を試すよりもそれぞれに特有の型を定義した方が良い

- 2 つ目の理由に独立して write も read も成長すべきなので，結合度を減らすべき
- 最後に一度に複数のデータを取得するクエリでパフォーマンスを出すために
- 基本的にはモデリングの際は write と read は別のモデリングが必要
- module ごとに分けた方が良いぐらいに独立すべき
- 本では ReadModel module と WriteModel module に分けていた
- ただし実態を分けるまでじゃないよね．もし分けるなら結構大変で，結界整合性とかでいくのもあり(ただしそれはドメインに依存する)
- 物理的な DB は同じだけど View だけ異なるようにして read 専用にするとかできる?
  - やはり全く DB のことをわかってないかも

## Event Sourcing

- CQRS は Event Sourcing と関連性がある

## Bounded Contexts Must Own Their Data Storage

- マイクロサービスを推奨しなくてもこれは推奨なのか？
- 最終目的としてはやはり独立して成長可能にすることと，結合度を少なくすること

## Working with Relational Database

- fp はデータと振る舞いを一緒にしないので RDB と相性が良い
- しかしまだいくつか解決しなければいけない問題はある

## Mapping Nested Types to Tables

- 内部型が DDD entity で ID を持つ場合は他のテーブルに分ける
- 内部型が DDD value で ID を持たないのであれば，同じテーブルにそのまま書く
