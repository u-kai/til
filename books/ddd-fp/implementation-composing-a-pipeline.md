# Implementation Composing a Pipeline

- 小さい関数の集まりで一つの workflow を書けるとドキュメント性も良いし，副作用もなく，テストもしやすい
- ただし，Result 型や Async 型などそのままでは次の関数の input として利用できないケースがあり，何かはしないといけない
- 関数同士は 2 つの理由から何もしないと合成できない
  - いくつかの関数は依存という外部パラメータを持っており，それは pipeline の一部に書いてはないが，実装上必要なもの
    - これには DI で解決するらしい
  - Result 型などのは直接 input として利用しないことが望ましいので，Result 型などを output にする関数はどうにかしなければいけない
  - adapter function がそれかもしれない

## Working with Simple Types

## Using Function Types to Guide the Implementation

## Implementation the Validation Step

## Implementation the Rest of the Step

## Composing the Pipeline Steps Together

- 異なる input とかを合わせる技として monad がある??

## Injecting Dependencies

- IoC のような不明確な依存は fp はいやだ
- 代わりに明示的なパラメータでやって依存を明確にすべき
- これには Reader Monad と Free Monad を使うと良い
- ここでは簡単な説明にとどめている
- oop ではトップレベルの関数をコンポジションルートとよんでおり，ここでもそれを使っている
- composition root fn はできる限りアプリケーションのエントリーポイントに閉じているべき
  - main 関数で workflow(composition root)に必要な全ての依存を作成や設定して突っ込んでいる感じ

## Testing Dependencies

- 依存関係を上のように解決するようにしてしまえば，特定の mock ライブリー無しに fake を提供できる
- fp はステートレスなのでテストしやすい
- 全ての依存関係は明示的になっているので，どのように動くのか理解しやすい
- 全ての副作用はパラメータとしてカプセル化されているのでテストをコントロールしやすい
