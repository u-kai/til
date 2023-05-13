- Deployment は複数の ReplicaSet を管理することでローリングアップデートやロールバックなどを実現するリソース
- Deployment が ReplicaSet を管理し，ReplicaSet が Pod を管理する 3 層の親子関係になっている
- 仕組み

  1. 新しい ReplicaSet を作成
  1. 新しい ReplicaSet 上のレプリカ数を徐々に増やす
  1. 古い ReplicaSet 上のレプリカ数を徐々に減らす
  1. 2，3 繰り返す
  1. 古い ReplicaSet はレプリカ数 0 で保持する

- Deployment は k8s では最も推奨されているコンテナの起動方法とされている
  - もし Pod 単体でデプロイした場合は Pod に障害が発生した際に自動で Pod が再作成されることはないし，ReplicaSet でデプロイした場合にはローリングアップデートなどの機能は利用できない

## Deployment のアップデート(ReplicaSet が作成される)条件

- レプリカ数の変更では Deployment は ReplicaSet を作成しない
- 作成される Pod の内容変更が必要
- 上記は，spec.template に変更があると，作成される Pod の設定が変わるため，ReplicaSet を新規で作成して，ローリングアップデートが行われる
- 実際にマニフェストを k8s に登録した後の ReplicaSet の定義を見てみると，spec.template 以下の構造体のハッシュ値を計算し，それを利用したラベル付きで管理を行う
