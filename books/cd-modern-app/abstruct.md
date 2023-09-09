---
marp: true
---

# ソフトウェア工学について

---

# 目次

- 目的
- IT 技術の重要性について
- ソフトウェア開発の難しさ
- どうすれば良いか？

---

# 目的

- IT 技術の重要性について理解する
- ソフトウェア開発がなぜ難しいのか，その理由を理解する
- ソフトウェア開発をどうすれば良いのか，その方法や理由を理解する

---

# IT 技術の重要性について

- 現代社会において IT 技術は非常に重要
- 様々な課題が IT 技術によって解決されていっており，解決されている領域も日々どんどん広がっている
- そのため IT 技術を迅速かつ継続的にデプロイ/デリバリーしている企業が優位に立っている
  - これは IT 企業に限らず，あらゆる企業において当てはまっている事実として分析されている

---

# Accelerate 本抜粋

- 数多くの分析を行い,どのような組織文化やプラクティスなどが企業を優位にしているのかをまとめた有名な本

| 2017               | ハイパフォーマ | ミディアムパフォーマ | ローパフォーマ    |
| ------------------ | -------------- | -------------------- | ----------------- |
| デプロイの頻度     | 1 日複数回     | 週 1 回から月 1 回   | 週一回から月一回  |
| 変更のリードタイム | １時間未満     | 1 週間から 1 ヶ月    | 1 週間から 1 ヶ月 |
| MTTR               | 1 時間未満     | 1 日未満             | 1 日から 1 週間   |
| 変更失敗率         | 0-15%          | 0-15%                | 31-45%            |

### 上の表は品質と速度にトレードオフな関係がないことを示している

---

# 品質と速度

- ソフトウェア開発の文脈ではこの 2 つはどちらか一方を選択しなければいけないものではない

  - 前ページの表が物語っている ，

## 逆に言えばどちらか一方を犠牲にすると，両方とも犠牲になる

## 本当か？

---

# 考察:品質を重視して速度を軽視するケース

ここでは以下のように定義

- 品質を重視とはプログラムのバグを無くしたり，コードを綺麗にしたり，テストを充実させたりすることを過度に行うこと
- 速度を軽視とはサービスのリリース頻度が少ないこと

デメリット

- プログラムを入念に作り込んだが，利用者のニーズとマッチしなかった.
  - プログラムから作成されたサービスの品質は利用者が決めるものでもあり，利用者が満足しなければそれは品質が悪いと言える
- リリースすることで初めてわかるバグが出てきた.
- 先に同じようなサービスを他社にデプロイされた.

### 品質は高速なフィードバックによっても高めることができ,高速なフィードバックには実際にデプロイすることが必要

---

# 考察:速度を重視して品質を軽視するケース

ここでは以下のように定義

- 速度を重視とはとにかく動くプログラムを作成してとにかくデプロイする
- 品質を軽視とは汚いプログラムでも構わず，テストも準備しない.ひとまず動けばいい

デメリット

- 品質の悪いプログラムはどこを修正すれば良いかわからない状態になりやすく，さらに壊れやすい.
- 速度を重視して機能を追加したくても，どこに機能を追加すれば良いか，機能を追加してもプログラムが壊れないかは，品質が高くないとわからない.
- テストが十分でないので，変更によってバグが生じないか，ちゃんとデプロイできるかに自信がなくなる.

### 品質が高くないことには，安全/確実なプログラム作成は難しい.速度を求めるだけでは速度は手に入らない.

---

# プログラムの品質？テスト？

- プログラムの品質とは何か？
- なぜプログラムの品質が低いと機能の追加や修正が難しいのか?
- テストが十分でないと自信が持てなくなるのか？

---

# プログラムの品質の前に，品質とは...

- よく使う言葉ではあるが,抽象的で難しい話
- 誰かにとっての価値!という人もいる

---

# ソフトウェアにおける品質

- 以下は JIS に記載されているソフトウェア製品に対する品質特性

| 品質特性     | 説明                                         | 品質副特性                                     |
| ------------ | -------------------------------------------- | ---------------------------------------------- |
| 機能適合性   | 機能がニーズを満たす度合                     | 機能完全性，機能正確性，機能適切性             |
| 性能効率性   | リソース効率や性能の度合い                   | 時間効率性，資源効率性，容量満足性             |
| 互換性       | 他のシステムと情報の共有，交換ができる度合い | 共存性，相互運用性                             |
| 使用性       | 利用者がシステムを満足に利用できる度合い     | アクセシビリティ，UI 快美性,習得性             |
| 信頼性       | 必要な時に機能実行できる度合い               | 成熟性，可用性，回復性                         |
| セキュリティ | 不正利用から保護する度合い                   | 機密性，否認防止性，責任追従性，真正性         |
| 保守性       | システムを修正する有効性や効率の度合い       | モジュール性，再利用性，解析性，修正性，試験性 |
| 移植性       | 他の実行環境に移植できる度合い               | 適応性，設置性，置換性                         |

---

# 狩野モデルの 5 つの品質特性

![](./Kano_model_rev.webp)

---

# 品質には外部と内部の品質がある

- 外部品質

  - 機能適合性,使用性

- 内部品質

  - 性能効率性,互換性,信頼性,セキュリティ,保守性,移植性

- 外部品質は魅力的品質になりうるもので利用者が実際に感じる品質

### 外部品質は外部からのフィードバックを元に継続的に,実験的に高める必要がある

- 内部品質は当たり前品質であったり，利用者からは見えない品質

### 今回の議題であるプログラムの品質は内部品質にマッピングする

---

# プログラム品質について

- おさらいすると，ソフトウェアの内部品質については以下

  - 性能効率性,互換性,信頼性,セキュリティ,保守性,移植性

- 上記はプログラムコードによって達成されることが多い
- どれも大事ではあるが，根底にあるのが**保守性**と考える

---

# なぜ保守性?

- 保守って開発が終わったらやるものではないの？
- 全てが終わってからの開発とは少し関係のない話？
- と思われるかも(保守性という単語だけ見たら僕もそう思う)
- ただし表の説明をみると，"シテムを修正する有効性や効率の度合い"とある
- これを言い換えると，システムに機能追加やバグ修正がどのくらいやりやすいのか？ということ
- そして前のページでも述べたが，内部品質の多くはプログラムコードによって作成される
- つまり，そのプログラムコードに対してどのくらい機能追加，修正がしやすいのかは相当大事
- そのため保守性が根幹にあり，保守性によって他の内部品質が向上すると考えている

---

# じゃあどうやって保守性を高める？

- プログラムコードを機能追加，修正が容易なように記載する!
- 保守性の品質服特性としても書いてあることが良いヒント

モジュール性，再利用性，解析性，修正性，試験性

- また別の本では保守性と同じような意味として，複雑性管理を挙げており,そのためのプラクティスとして以下が記載されている

モジュラー性（modularity）,凝集度（一体性、cohesion）,関心の分離（separation of concerns）,抽象化（abstraction）,疎結合（loose coupling）

これらは全て正しく重要で保守性を高めるために必要な項目で，それぞれの特性同士は強く関連している

---

# 例:保守性が低いコード

- 同ディレクトリにある sample-bad-code.py は何をやっているのか？

```python
def main():
    print("Start")
    flag = True
    player1 = True
    player2 = False
    filed = []
    ...

if __name__ == "__main__":
    main()

```

- 一つの main 関数があり，その中で全ての処理を行っている

---

# 何が悪いのか

- 何がしたいのか直感的にわからない(抽象化がされていない)
- 一連の流れが記述されているだけで構造がない(モジュラー性がない，再利用性がない,関心の分離がされていない,密結合,凝集度がない)
- プログラムを実行する以外の方法で正常性を担保できない(試験性がない)
- 機能追加や修正をするときにどこを治して良いかわからない(修正性がない)
- 記述に重複があり，修正をする際に漏れが生じる可能性がある

### それぞれの特性は関連性が強いため，悪いコードはほとんど全ての関連性に当てはまっている

---

# 本当にそうか？

- 本当にそうなのか確かめるために，以下のようなことが簡単にできるか考えてみよう

  - マスの数を変えてください
  - 出力する文字を変えられるようにしてください
  - Player の名前を任意のものにできるようにしてください
  - ゲームをコンテニューできるようにしてください
  - バグがないか確かめてください

- 間違いなく，もれなくできるか？
- ストレスなく，すぐにできるか？

---

# 相当難しいはず

- できる人でもストレスはかかるはず

---

# 本番のシステムはこんなに単純じゃない

- 今回のサンプルコードはあくまでサンプル
- 価値のあるシステムはこの 100~10000 倍以上複雑
- その複雑なものに対して継続的に機能追加や修正をしなければいけない
- そのために重要なものが保守性の品質ということになる

---

# どうすればいいのか?