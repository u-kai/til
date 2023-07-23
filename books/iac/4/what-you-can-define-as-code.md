What You Can Define as Code Every infrastructure tool has a different name for its source code — for example, playbooks, cookbooks, manifests, and templates. I refer to these in a general sense as infrastructure code, or sometimes as an infrastructure definition. Infrastructure code specifies both the infrastructure elements you want and how you want them configured. You run an infrastructure tool to apply your code to an instance of your infrastructure. The tool either creates new infrastructure, or it modifies existing infrastructure to match what you’ve defined in your code. Some of the things you should define as code include:

- An infrastructure stack, which is a collection of elements provisioned from an infrastructure cloud platform. See Chapter 3 for more about infrastructure platforms, and Chapter 5 for an introduction to the infrastructure stack concept.
- Elements of a server’s configuration, such as packages, files, user accounts, and services (Chapter 11).
- A server role is a collection of server elements that are applied together to a single server instance (“Server Roles”).
- A server image definition generates an image for building multiple server instances (“Tools for Building Server Images”). An application package defines how to build a deployable application artifact, including containers (Chapter 10).
- Configuration and scripts for delivery services, which include pipelines and deployment (“Delivery Pipeline Software and Services”).
- Configuration for operations services, such as monitoring checks.
- Validation rules, which include both automated tests and compliance rules (Chapter 8).
コードと定義できるものは何ですか
インフラストラクチャツールには、さまざまな名前があります。例えば、プレイブック、クックブック、マニフェスト、テンプレートなどです。私は一般的な意味でこれらをインフラストラクチャコード、またはインフラストラクチャ定義と呼んでいます。インフラストラクチャコードは、欲しいインフラストラクチャ要素とその設定方法を指定します。インフラストラクチャツールを実行して、コードをインフラストラクチャのインスタンスに適用します。ツールは新しいインフラストラクチャを作成するか、既存のインフラストラクチャをコードで定義したものに合わせて変更します。コードで定義するべきものには、以下のものがあります：

- インフラストラクチャクラウドプラットフォームからプロビジョニングされた要素のコレクションであるインフラストラクチャスタック。インフラストラクチャプラットフォームについては、第3章で詳しく説明します。また、インフラストラクチャスタックの概念については、第5章で紹介します。
- サーバーの構成要素、例えばパッケージ、ファイル、ユーザーアカウント、サービスなど（第11章）。
- サーバーロールは、単一のサーバーインスタンスに一緒に適用されるサーバー要素のまとまりです（「サーバーロール」）。
- サーバーイメージの定義は、複数のサーバーインスタンスの構築用イメージを生成します（「サーバーイメージの構築ツール」）。アプリケーションパッケージは、デプロイ可能なアプリケーションアーティファクトの構築方法を定義します。これには、コンテナも含まれます（第10章）。
- パイプラインやデプロイメントなどの配信サービスの構成とスクリプト（「配信パイプラインソフトウェアとサービス」）。
- モニタリングチェックなどのオペレーションサービスの構成。
- 自動テストとコンプライアンスルールを含む検証ルール（第8章）。