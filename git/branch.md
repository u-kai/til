# Git のブランチについて

- git にはブランチというものがある
- ブランチの実態はコミットある一つのコミットを指すポインタ

```bash
git log

commit 14ec7b80acb29b6979d7249272304109c0991399 (HEAD -> main, origin/main) # mainとorigin/mainがブランチ名で，この例ではどちらも14ec7b80acb29b6979d7249272304109c0991399を指している
Author: Kai
Date:   Fri Sep 8 06:44:26 2023 +0900
```

- 実際は .git/regs/heads/ブランチ名のファイルに記述がある
  - リモートリポジトリに関する記述は .git/regs/remotes/origin(remote のエイリアス)/ブランチ名のファイルに記述がある

```bash
cat .git/regs/heads/main
```

## ブランチは何の役に立つのか？

- ブランチは実はいくらでも作成可能

```bash
git branch ブランチ名
```
