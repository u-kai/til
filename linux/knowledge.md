- 正確には Linux という言葉はカーネルのみを指す
- コンピュータシステムの階層

  - ユーザープログラム
  - OS 外プログラム
  - OS ライブラリ
  - カーネル
  - ハードウェア

- マシンの電源を入れると最初にカーネルが起動する
- カーネルは CPU のカーネルモードとユーザモードの２つをつかって、プロセスからデバイスに直接アクセスできないようにしている
- カーネルのみがカーネルモードで実行をすることができる

  - ユーザモードに飲み特定の命令を実行できないような制約をかけることができる
  - プロセスはカーネルを介して間接的にデバイスにアクセスをする

- デバイス制御に加えて、システム内のすべてのプロセスが共有するリソースを一元管理してシステム上で動作するプロセスに配分するカーネルモードで動作するプログラムがカーネル

- システムコールとはプロセスがカーネルに処理を依頼する方法
  - 新規プロセスの生成やハードウェアの操作など、カーネルの助けが必要になる
  - システムコールの例は以下
    - プロセス生成、削除
    - メモリ確保、解放
    - 通信処理
    - ファイルシステムの操作
    - デバイス操作
- システムコールは CPU の特殊な命令を実行することによって実現してい る

  - プロセスはユーザモードで実行しているが、カーネルに処理を依頼するためにシステムコールを発行すると、CPU において例外というイベントが発生する
  - これをきっかけとして CPU のユーザモードからカーネルモードに遷移し、依頼内容に応じたカーネルの処理が動き始める
  - カーネル内のシステムコール処理が終了すれば、再びユーザモードに戻ってプロセスの動作を継続する

- プロセスからシステムコールを介さずに直接 CPU のモードを変更する方法はない
- strace の使い方

```
strace -o output.log ./exe
```