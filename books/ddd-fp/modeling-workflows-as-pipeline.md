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
