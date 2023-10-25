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
