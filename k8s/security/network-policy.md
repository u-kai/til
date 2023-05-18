# NetworkPolicy

- k8s クラスタ内で Pod 同士が通信する際のトラフィックルールを規定するもの
- NetworkPolicy を利用しない場合，クラスタ内のすべての Pod 同士で通信を行うことが可能
- 一方で NetworkPolicy を利用できれば，Namespace ごとにトラフィックを転送しないようにしたり，基本的にすべての Pod 間の通信を断絶して特定の Pod 間だけ通信を許可するホワイトリスト方式にしたりすることが可能

- NetworkPolicy はすべての k8s 環境で利用できるわけではなく，Calico などの NetworkPolicy をサポートする CNI プラグインを利用してクラスタが構築されている必要がある
- GKE の場合は作成時に明示的に有効化することで Calico がデプロイされ，NetworkPolicy を利用できるようになる

- Ingress と Egress では，Policy の指定として podSelector,namespaceSelector,ipBlock がある
- podSelector は特定の Pod からの通信に対して制御を行うポリシー
- ipBlock はクラスタ外のネットワークに関する制御を行う場合には利用するイメージ
