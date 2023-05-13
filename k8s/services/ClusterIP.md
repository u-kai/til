# ClusterIP

- type :ClusterIP を作成すると，K8s クラスタ内からのみ疎通性がある Internal Network に作り出される仮想 IP が割り当てられる
- ClusterIP 宛の通信は各ノード上で実行しているシステムコンポーネントの kube-proxy が Pod 向けに転送を行う
- ClusterIP は K8gs Cluster 外からトラフィックを受ける必要がない箇所などで，クラスタ内のロードバランサとして利用する
- デフォルトでは K8s API に接続するための kubernetes Service が作成されており，ClusterIP が発行されている

## ClusterIP 仮想 IP の静的な指定

- 例えば，アプリケーションからデータベースサーバを指定する場合，基本的に K8s Service で登録されているクラスタ内 DNS レコードを用いてホスト指定をするのが望ましい
- ただ，一方，IP アドレスで指定しなければいけない場合には ClusterIP を静的に指定することも可能
- ClusterIP を自動付でなく，手動で指定する場合には spec.clusterIP を指定する
