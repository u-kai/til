---
marp: true
---

# Copilot に関するレポート

---

# 目次

- 本レポートの目的
- Copilot とは
- Copilot の使い方
- Copilot の評価
  - 実用性
  - ChatGPT との違い
  - 危険性や課題
- まとめ

---

# 本レポートの目的

GitHub Copilot とは何か，どのくらい有用なのか，企業での課題は何かなどを共有する

---

# Copilot とは(基本情報)

- 正式名称は GitHub Copilot
- 2021/7 月からプレビューとして公開され，2023/2/15 から Copilot Business が GA
- 現在の利用形態は個人利用と，企業利用の 2 種類
- 学習データは GitHub 上のパブリックコード

  - 数億行!(Amazing!)

> GitHub Copilot は AI ペア プログラマーです。 GitHub Copilot を使うと、エディター内で行全体または関数全体の候補を得ることができます。
> https://docs.github.com/ja/copilot/quickstart

## コードを書くときに AI がコードを提案(サジェスト)してくれる!!

---

# Copilot 前提条件

- Copilot は IDE(JetBrains や VsCode,NeoVim 等)が開発者とのインターフェイス
  - API は明示的に公開されていない
- Copilot に対応している IDE に Copilot の機能をインストールして利用する

---

# Copilot 実際の使い方 1

コードを少し記述するとその文脈でサジェストしてくれる

![height:400](suggest.png)

## ちゃんと最適なアルゴリズム!!

---

# Copilot 実際の使い方 2

## Q & A できちゃう！

![](q-and-stackoverflow-a.png)
ちゃんと正常な URL(1) でした(なんで？？)

![](q-and-tdd-a.png)
だそうです

---

# Copilot 実際の使い方 3

- ctl+Enter で 10 候補ほどサジェストしてくれる!

![](many-suggest.png)

---

# Copilot 実際の使い方 3

- Accept Solution ボタンを押すと，選択したサジェストをコードに反映してくれる

![](after-suggest.png)

楽ちん楽ちん！

---

# Copilot の使い方その他

- GitHub Copilot labs などがあり，コードレビューやコメント作成などの実験的ソリューションが検討されている
- 凄まじいソリューションが登場するのも時間の問題かも...

---

# 有用性について(GitHub 社調べ)

![](copilot-numbers.png)
※ https://github.com/features/copilot/

---

# 利用形態について

### 1. Copilot for Individuals(10$ / month \* user)

### 2. Copilot Business(19$ / month \* user)

- 弊社は GitHub Enterprise を利用しているため Business での利用が可能
- Business の利点

  - Individual の機能全てを保有
  - Enterprise 全体への統一した policy 設定が可能
  - 自己証明書を利用する Proxy 設定が可能(Zscaler!!!)
  - Business の利用者は Enterprise 外の個人開発にも Copilot を利用可能

- 弊社環境では統一した policy 設定と Proxy の設定が可能な Business がベスト

---

# Copilot の動作原理

- OpenAI の Codex というモデルがバックエンドで推論，サジェストをしてくれる
- Copilot は VSCode などの IDE をインターフェイスとして動作する

![width:700](copilot-arch.png)

---

#　 Copilot の危険性

---

# MS の主張

## Copilot がサジェストしたコードはあなたのコードです※

※ 責任持てよという意味

---

# 参考

(1)https://stackoverflow.com/questions/5811151/why-do-we-check-up-to-the-square-root-of-a-number-to-determine-if-the-number-is

- https://github.com/features/copilot/
