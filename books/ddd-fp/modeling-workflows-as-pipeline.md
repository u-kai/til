# Modeling Workflows as Pipelines

- 徹底してドメインエキスパートがわかりやすいように書きましょうって感じやな

  - システム開発の文脈でドメインエキスパートがいない場合とかってないのか？
    - 例えば全く新しいものを作る時とか(ただ，これも何かの課題を解決するものと捉えたら課題を抱えているドメインエキスパートはいそう)
    - ただし，今回のものを見ると，ドメインエキスパートの現在の要件通りに型づけをして，最終的に出来上がるものは現状の業務を電子化したものになっている
    - それだけで本当にいいのか？
    - アジャイル的に価値の方向性を見つけながらやる手法とマッチしているのかどうか？
    - どっちかというと，ドメインエキスパートのやることが決まっていてそれを少しずつ型づけしているので，ウォーターフォールに近く見えてしまう

- pipeline をビジネスプロセスに見立てる
- このスタイルは transformation-oriented programming とも呼ばれる
- workflow の input は常にドメインオブジェクトであるべき
  - DTO からは deserialized しておくべき
- command によって workflow は走る

## Commands as Input

- workflow の入り口の command にはその workflow で利用する全ての情報が必要
- またいつ，誰がこのコマンドを作成したかを追跡できる必要があるし，その他の情報もログに残せるようにするべきなので，それらのデータも全て command に含める

## Sharing Common Structures Using Generics

- Generics を使って共通部分をうまく省略可能な Command<T>を作成する感じ
  - TimeStamp はわかるけど，それ以外の設計はしっかりしないと難しそう
  - 逆にいうと Workflow のクリティカルな部分が Command の共通部分が依存することはよくなさそう
  - log とか結構どの workflow でも共通で行われているものであれば Command へ依存しても問題ないのでは？

## State Machines

- 状態自体がどきゅめんとになる
- エッジケースを抑えることができる

## Modeling Each Step in the Workflow with Types

- 外部へのリクエストは unit ではなく，独自の Result でリクエストの失敗や成功を表せると良い
- 外部へのリクエストには何を使えば良いのかは，後でよく最初は interface というか，型の定義をしてあげれば良い
- workflow 全体を型として定義している感じ

```
type AcknowledgeOrder =
CreateOrderAcknowledgmentLetter // dependency
-> SendOrderAcknowledgment // dependency
-> PricedOrder // input
-> OrderAcknowledgmentSent option // output
```

## Documenting Effects

- ドキュメントによって依存関係やダブルチェックなどの効果が出る

## Composing the Workflow from the Steps

- input,output 以外の依存性を型やドキュメントとしてどう扱うのか？
- 依存性を明示するかどうかに答えはないが以下のようなガイドラインがある
  - パブリックな API である関数は呼び出されるので依存性を隠すと良い
    - 呼び出す側はその関数が何かに依存していることなどどうでも良い
    - なので input,output だけで良いはず
  - インターナルな関数は依存性を明示的にするべし
    - workflow の internal な step は明示的に依存関係を記述するのがよい
    - しかも internal なので依存先の変更が起きても外部には問題がない

## Wrapping Up

- 型情報は増えるけどそれはドキュメントとして成り立つのでとても良いこと
- 今までのやり方はウォーターフォールに見えるが，実際は多くのフィードバックを継続的にもらうこと
- 型情報もどんどんドメインエキスパートに聞いて，変えていくこと

## 所感

- この型情報自体が一種のプロトタイプのようにも見える
- これをどんどん提示していってやっていく感じかな？
- type driven development とあったがまさにその通りな気がする
- TDD はこんな感じに Object が振る舞ってその結果何かがこのようなことが起こってほしいな〜から書くので，OOP 向けだったりする？
- FP では type driven development で型で多くを表現していく形にすれば同じようなことはできるかも？
  - ただ TDD がやりにくいということではないと思うが，type driven development で TDD が達成したかった，独立していて綺麗なコードはかけるのではないか？
  - tdd と違って最後にテストは残らないが，まずはテストから書く，というよりも全体の workflow をちゃんと見てから型を書いていくので，TDD よりも，抽象的なところから始められる印象
  - どちらも補完し合えればよいが，どうなんだろ
