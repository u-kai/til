# pyinstaller について

*   下記コマンドで実行ファイルを作成できる

```shell
pyinstaller FILE_NAME **onefile
```

*   ただし上記 FILE_NAME のファイル内の import 文は相対パスでは無理で，絶対パスでなければいけない
    *   なぜかというと，pyinstaller の仕様として exe 化を指定したファイルはパッケージに属さないものとして扱われ，どこからの相対パスなのかが理解できないためにエラーとなる模様
