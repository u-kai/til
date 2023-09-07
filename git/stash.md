# stash で変更を退避させる

```shell
git stash -u
```

- -u は--include-untracked の略.新規作成ファイル(追跡対象に含まれていないファイル)も退避することができる

```
git stash list
```

- 一覧を見ることができる

```
git stash apply stash@{0}
```

- stash@{0}の作業を元に戻すことができる

```
git stash drop stash@{0}
```

- stash@{0}の作業を削除できる

```
git stash pop stash@{0}
```

- 作業を戻しつつ削除する
