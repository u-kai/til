## Create Events

- ECS

  - CreateService
  - CreateTask
  - CreateTaskSet

- RDS
  - CreateDBInstance
  - CreateDBCluster
  - CreateDBInstanceReadReplica
  - CreateDBProxy

# Public Subnet について

- Public Subnet を監視しようと思ったら以下の Event をとる必要がある

  - CreateRouteTable

    - Main かつ IGW と接続していて，Subnet に RouteTable が紐づいてない場合は RouteTable と明示的な紐付けがない全ての Subnet が PublicSubnet になる
    - でも Create した瞬間に Route って紐づくっけ？
      - 紐づかないのであれば不要

  - CreateRoute

    - Create された Route が igw かどうか & 左の Route が紐づいている RouteTable と Subnet が紐づいているかどうか
    - RouteTable への Route 追加がコンソール上の動き

  - AssociationRouteTable
    - あたらしく紐づけられた RouteTable に IGW があればダメ
    - おそらく，VPC の Main テーブルを使っていた Subnet とかが新しく Public な RouteTable をつけられるとダメ
  - ReplaceRouteTableAssociation
    - RouteTable の関連付けを変更したとき
    - 新しい AssociationId が Event で拾える
    - おそらく新しい RouteTableId も拾える
  - ReplaceRoute
    - コンソールでは CreateRoute が実行されたが，API ドキュメントとして存在するのでおそらく必要

## Public Subnet に Instance があるかどうか

- 一番安直なのが Instance の状態変更を Config などで常に監視すること
  - ただし，関係ない操作についても発行することや，Config がややこしいこと，Terminate した時もそのイベントを Config が取得してしまう
  - Config は高額で様々な機能が面倒
- できれば EventBridge で補足したいが，EC2 以外の Instance についても記述するのがめんどくさい(勉強にはなるかも)

- 大まかな考え方

  - 構築された時に何らかの形で紐づいている Subnet が Public な Subnet かどうかを調べる
  - 構築された後に Subnet を移動できないのであれば Create した時だけ考えれば良い
    - Subnet 移動不可
      - Lambda
      - EC2
      - ECS
        - CreateService
        - CreateTask
        - CreateTaskSet
        - RunTask
        - StartTask
  - RouteTable を変更するなどして Public Subnet の条件に変更された時に，その Subnet 内の Instance 全てが PublicInstance になる
    - これをどうやって Instance の方を調べるか(Public Subnet かどうかは Event から簡単に補足できるはず)
    - すべての Instance の Describe を行って，Subnet と紐づいているのかを見る
      - これって Too Many Request エラーとか出ないよな？(サービス同士の依存度は少ないんよな？AWS は)

- RDS

  - CreateDBSubnetGroup
  - CreateDBInstance
  - CreateDBCluster

- Redis

  - CreateCluster
    - SubnetGroupName が Request で必要で，それ以上の情報は Fetch する必要あり
  - CreateSubnetGroup

    - Subnet 作成の時点で破壊してあげる方が早いけど，どうするか？
    - 正直 CreateCluster の時にデフォルトで作成されてしまうのであればいらない気がする

- EKS

  - CreateCluster
    - PrivateSubnet にすべしとドキュメントにはあったかつ，後から Node 用の Subnet を追加？できるとあったので，ここで補足しても良い気がするが，eksctl などのツールがデフォルトで PublicSubnet を指定していたらどうしようか
    - そもそも Cluster って何？Node や Namespace が構築されるものだと思っているけど，この Cluster が存在している Subnet に含まれていない LB とかって後から追加できるのか？
    - ControlPlane は利用者アカウントには存在しないと思えば良い？
    - Subnet の説明に，クラスターとの通信を容易にするためコントロールプレーンが ENI を配置できる Subnet とあるので，これが Public である必要なさそうだし，ELB とかを配置するのはまた違う話な気がする
      - https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/network_reqs.html#:~:text=%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%BF%E3%83%BC%E3%81%AE%E4%BD%9C%E6%88%90%E6%99%82%E3%81%AB%E6%8C%87%E5%AE%9A%E3%81%99%E3%82%8B%E3%81%AE%E3%81%A8%E5%90%8C%E3%81%98%E3%82%B5%E3%83%96%E3%83%8D%E3%83%83%E3%83%88%E3%81%AB%E3%80%81%E3%83%8E%E3%83%BC%E3%83%89%E3%81%A8%20Kubernetes%20%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9%E3%82%92%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%99%E3%80%82%E3%81%9F%E3%81%A0%E3%81%97%E3%80%81%E3%81%93%E3%82%8C%E3%81%AF%E5%BF%85%E9%A0%88%E3%81%A7%E3%81%AF%E3%81%82%E3%82%8A%E3%81%BE%E3%81%9B%E3%82%93%E3%80%82%E3%81%93%E3%82%8C%E3%81%AF%E3%80%81%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%BF%E3%83%BC%E3%81%AE%E4%BD%9C%E6%88%90%E6%99%82%E3%81%AB%E6%8C%87%E5%AE%9A%E3%81%97%E3%81%AA%E3%81%8B%E3%81%A3%E3%81%9F%E3%82%B5%E3%83%96%E3%83%8D%E3%83%83%E3%83%88%E3%81%AB%E3%82%82%E3%80%81%E3%83%8E%E3%83%BC%E3%83%89%E3%81%A8%20Kubernetes%20%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9%E3%82%92%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%A7%E3%81%8D%E3%82%8B%E3%81%9F%E3%82%81%E3%81%A7%E3%81%99%E3%80%82
      - ノードと K8s リソースは別サブネットにおけるとある
        - EKS 自体が，IAMRole の権限内で ELB とか作るのであれば，確かに Subnet はクラスター内になくても良さそう
      - と思ったけど，ELB を置くには PublicSubnet が必要ってあるから，これが Cluster 作成時の話なのか，その後で追加できるのかは知っておきたい
      - もし無理であればどうしようか．
      - ちなみに EKS から ec2 に RunInstances が走っているっぽいので EKS に対してはノータッチでもいいのかもしれない
        - と思ったかもしれないけど，NodeGroup を破壊しないことには Node が復活する可能性ないか？AutoScaling ってことはその可能性はありそう
      - クラスター作成時に指定した Subnet に VPC とクラスターが通信するようの ENI を 2〜4 つ作成するみたい
        - これを使って kubectl などの機能を利用できるっぽい
    - でもコントロールプレーンが ELB を作成したり，コントロールするとしたら，通信できる必要はあるのでは？(同じ VPC にいれば問題はなさそうなので多分 OK)
    - LB をサブネットにデプロイする場合はサブネットにタグが必要みたい
  - UpdateClusterConfig
  - CreateFargateProfile
  - CreateNodegroup
    - RunInstances に責務を任せれば良い
  - 疑問
    - ELB などの Service はどのように作成される？
    - AWS 以外の Ingress とかを使いたくなっても大丈夫なのか？
    - これが普通に Manifest 経由かつ，コントロールプレーンや Node の Subnet を変更せずに行えるのであれば嬉しい

- all list

- Amazon EC2 (Elastic Compute Cloud): 仮想サーバー
- Amazon RDS (Relational Database Service): リレーショナルデータベースサービス
- Amazon Redshift: データウェアハウスサービス
- Amazon Elasticache: インメモリデータストアおよびキャッシュサービス
- Amazon Neptune: グラフデータベースサービス
- Amazon DocumentDB: MongoDB と互換性のあるドキュメントデータベース
- Amazon ECS (Elastic Container Service): コンテナオーケストレーションサービス
- Amazon EKS (Elastic Kubernetes Service): Kubernetes オーケストレーションサービス
- AWS Fargate: ECS および EKS のサーバーレスコンピューティングエンジン
- Amazon EMR (Elastic MapReduce): ビッグデータ処理サービス
- Amazon MSK (Managed Streaming for Apache Kafka): Kafka クラスターサービス
- AWS App Runner: コンテナベースのアプリケーションデプロイメントサービス
- AWS Lambda: イベントドリブンのサーバーレスコンピューティングサービス（VPC 内のリソースへのアクセスが可能）
- Amazon Elastic File System (EFS): マネージドファイルストレージ
- AWS Batch: バッチコンピューティングワークロードの実行
- AWS Outposts: オンプレミス環境での AWS サービスの実行
- AWS Snow Family: エッジロケーションやオンプレミスでのデータ処理および移行
