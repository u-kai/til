# Introducing DDD

- EventStorming による関係者全員参加の Event 探しが重要らしい
- business events にフォーカスをまずは当てるべき

- 全ての詳細をカバーすることは event storming でできないので，他側面の要求を集めることも大事

  - business モデルを共有すること
  - 全てのチームの認識
  - 要求のギャップを見つけること
  - チーム間のコネクション

    - おそらくどこかのイベント(output)がどこかのシステムの input になるってこと
    - 上記がわかるようになるということ

  - レポート要求の認識
    - reporting はいつでもドメインの一部？
      - 過去に何があったのかを理解する必要がある
    - これも重要な一つで event storming の一つで語られるべき

- Domain Event は何かシステムに起こったことを表現するデータ.
  - 常に過去形で記述される
  - 他の追加のアクティビティのトリガーになり得る
- event の前と後に何があるのかを深掘りすると，さらに event が出てくる
- この時には，実装が簡単なものや優先順位が高いものなど様々あるが，この段階では，問題領域を知ることが重要なので，実装のことはその後に考えるべき
- 何か Event が始まるのに対するものを commands と DDD ではいう
- Commands は常に命令的に記述される(わたしのために行って)
- command による命令が成功した場合，それに続いて対応するドメインイベントが生成されるはず
  - ワークフローが X を実現した場合，対応するドメインイベントは X が実現したとなる
- これが pipeline しょい r
- もし なんの Command によってもトリガーされない場合は，それは時間によるスケジューリングされたとりがーになる

## Partitioning the Domain into Subdomains

- 小さい問題領域に区切ること
- domain とは domai expert が

## Creating a Solution Using Bounded Contexts

- problem space と solution space の間の区別を作成する必要がある
  - そしてそれらを違うものとして扱う必要がある
- problem space は現実世界のよう
- solution space は domain model すなわち，process として定義されたもの
- 現実世界は絡まっているが，ソフトウェアの世界ではなるべくでカップリングしたい
- 現実世界のコンテキストがシステムの世界のコンテキスト 1 対 1 にならないこともある

## Getting the Contexts Right

- 正しく Context を切るのは芸術でいくつかのガイドラインがある
- Listen to the domain experts
- Pay attention existing team and department boundaries

  - 今のチーム編成などが domain や subdomain を見つけるのに大きな手掛かりになることがある
  - 逆に多くのコラボレーションをしていればそれは同じ domain であるべき

- Don't forget the bounded par of a bounded context
  - よくわからん
- Design for autonomy

  - 独立して成長可能にせよ

- Design for friction-free business workflows

- 最後に，静的なモデルはないので継続して成長させましょう

## Creating Context Maps

- 上流と下流がわかる程度の詳細で ContextMap を書くと良い

## Focusing on the Most Important Bounded Contexts

- 一般的には最重要なドメインがある(ドメインには優先順位がある)
  - このようなドメインをコアドメインと呼ぶ
  - コアドメインによってお金を稼いでいる

## Creating Ubiquitous Language

- ユビキタス言語は全てのところで利用すべき
- ユビキタス言語もどんどん進化すべき．新しい概念は見つかることもある
- ユビキタス言語は方言で，コンテキストごとに異なる
