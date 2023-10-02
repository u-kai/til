# A Functional Architecture

- 歩くスケルトンで早いフィードバックをもらって認識齟齬をなくすこと
- C4 アプローチ
  - システムコンテキストは全体のシステムの一番上位
  - システムコンテキストにはいくつかのコンテナがあり，website や web service や database などの
  - そのコンテナにはいくつかのコンポーネントがある，これは塊のコード
  - そしてそのコンポーネントにはいくつかのクラスがある，関数型だと modules，そしてそこには最下位のメソッドや関数がある

## Bounded Contexts as Autonomous Software Components

- まずは疎結合につくろう
- デプロイに関するアーキテクチャ(microservice or monolisic)は後で良い
- 最初はモノリスに作って必要な時にマイクロサービスが良い
- Queue は非同期のコンテキスト移動の最適
  - これはデプロイの形がなんであれそう

## Contracts Between Bounded Contexts

- Shared Kernel
- Consumer Drive Contract
- Conformist
- うえの 3 つの考え方がある
- ACL(Anti Corruption Layer)は DTO の変換だけではなく，コンテキストの変換も行う
  - 文章では言語の変換とあったが，おそらく上と同義であると思う
  - これは 3rd Party や外部の Untrust なものに対して腐敗防止として利用できる

## Code Structure Within a Bounded Context

- 技術で層を区切ると，様々な Workflow に影響が出るが，Workflow ごとに切ると影響が出なくなる
- ただしそれでも十分ではない
- workflow を水平のパイプと層に伸ばす?
- オニオンアーキテクチャ
  - ヘキサゴナルアーキテクチャ
  - くりーんあーきてくちゃなど
