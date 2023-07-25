Stack Topologies for Application Clusters An application cluster is comprised of different moving parts. One set of parts is the applications and services that manage the cluster. Some of these cluster management services include: Scheduler, which decides how many instances of each application to run, and where to run them Monitoring, to detect issues with application instances so that they can be restarted or moved if necessary
Configuration registry, for storing information needed to manage the cluster and configure applications Service discovery, enabling applications and services to find the current location of application instances Management API and UI, enabling tools and users to interact with the cluster Many cluster deployments run management services on dedicated servers, separate from the services that host application instances. These services should probably also run clustered, for resilience. The other main parts of an application cluster are the application hosting nodes. These nodes are a pool of servers where the scheduler runs application instances. It’s common to set these up as a server cluster (see “Compute Resources”) to manage the number and location of server instances. A service mesh (see “Service Mesh”) may run sidecar processes on the host nodes, alongside applica​tion instances. An example application cluster (Figure 14-2) includes servers to run cluster management services, a server cluster to run application instances, and network address blocks.
Networking structures for an application cluster may be flat. The cluster services assign network addresses to application instances. It should also handle network security, including encryption and connection management, often using a service mesh for this. You use infrastructure stacks to provision this infrastructure. Chapter 5 explained that the size of a stack and the scope of its contents have implications for the speed and risk of changes.
“Patterns and Antipatterns for Structuring Stacks” in particular lists patterns for arranging resources across stacks. The following examples show how these patterns apply to cluster infrastructure. アプリケーションクラスタのスタックトポロジー
アプリケーションクラスタは、異なる動作部分で構成されています。その一部は、クラスタを管理するアプリケーションとサービスです。これらのクラスタ管理サービスには、次のものが含まれます。

- スケジューラ：各アプリケーションのインスタンスの実行数と実行場所を決定します。
- モニタリング：アプリケーションインスタンスに問題がある場合に検出し、必要に応じて再起動または移動します。
- 構成レジストリ：クラスタを管理し、アプリケーションを構成するために必要な情報を保存します。
- サービスディスカバリ：アプリケーションとサービスがアプリケーションインスタンスの現在の場所を見つけるための機能です。
- マネージメント API および UI：ツールとユーザーがクラスタと対話するための機能です。

多くのクラスタ配置では、アプリケーションインスタンスをホストするサービスとは別に、専用のサーバー上で管理サービスを実行します。これらのサービスも耐障害性を考慮してクラスタ化して実行する必要があります。

アプリケーションクラスタのもう一つの主要な部分は、アプリケーションをホストするノードです。これらのノードは、スケジューラがアプリケーションインスタンスを実行するサーバーのプールです。通常、これらはサーバークラスタとして設定され、サーバーインスタンスの数と位置を管理します。サービスメッシュは、ホストノード上でアプリケーションインスタンスと並行してサイドカープロセスを実行する場合もあります。

アプリケーションクラスタのネットワーキング構造はフラットになる場合があります。クラスタサービスがアプリケーションインスタンスにネットワークアドレスを割り当てます。これにはネットワークセキュリティ（暗号化や接続管理など）も含まれます。インフラストラクチャスタックを使用してこのインフラストラクチャをプロビジョニングします。

特に「スタックの構造化に関するパターンとアンチパターン」では、スタック間のリソースの配置パターンについて説明しています。以下の例は、これらのパターンがクラスタインフラストラクチャにどのように適用されるかを示しています。

パッケージ化されたクラスタソリューションを使用する場合、インフラストラクチャスタックはクラスタ管理ソフトウェアのデプロイと構成に必要なインフラストラクチャをプロビジョニングします。インストーラを実行するのは別のステップで行うべきです。これにより、アプリケーションクラスタとは別にインフラストラクチャスタックをテストできます。クラスタの構成をコードで定義できる場合は、コードを適切に管理し、テストとパイプラインを使用して簡単かつ安全に更新や変更を行うことができます。サーバーの構成コードを使用してクラスタ管理システムをサーバーに展開することも有用かもしれません。

以上は、クラスタ管理アプリケーションのサーバーを作成するために「Example 14-1」のものに追加するスニペットの例です。

このコードは、クラスタ管理サーバーをクラスタネットワークに適用するための設定を行っています。
