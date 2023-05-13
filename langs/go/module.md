# モジュール

- Go ではサードパーティのライブラリ管理にの単位はモジュールで Go Modules と呼ばれている
- Go Modules を利用するには go mod コマンドを利用する
- Go には npm や PyPI のような他の言語で見られる中央集権型のモジュール管理サービスではなく，Git リポジトリを指定してモジュールをダウンロードする
- go mod tidy コマンドで利用しなくなったサードパーティモジュールを削除できる
- ある程度コードが動くようになったら，バージョン情報をタグとして付与してリポジトリに push する

```
git tag v1.0.0
git push origin main --tag
```

- リポジトリに push したモジュールは go get でタグとして付与したバージョンを指定することで，取得できる