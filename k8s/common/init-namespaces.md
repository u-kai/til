- 初期状態では下記 4 つの Namespace が作成される
- kube-system
  - k8s クラスタのコンポーネントやアドオンがデプロイされる Namespace
- kube-public
  - 全ユーザが利用できる ConfigMap などを配置する Namespace
- kube-node-lease
  - ノードのハートビート情報が保存されている Namespace
- default
  - デフォルトの Namespace