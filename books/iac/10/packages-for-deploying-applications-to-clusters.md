Packages for Deploying Applications to Clusters Modern applications often involve multiple processes and components deployed across complex infrastructure. A runtime environment needs to know how to run these various pieces: What are the minimum and maximum number of instances to run? How should the runtime know when to add or remove instances? How does the runtime know whether an instance is healthy or needs to be restarted? What storage needs to be provisioned and attached to each instance? What are the connectivity and security requirements? Different runtime platforms provide different functionality, and many have their own packaging and configuration format. Often, these platforms use a deployment manifest that refers to the actual deployment artifacts (for example, a container image), rather than an archive file that includes all of the deployable pieces. Examples of cluster application deployment manifests include: Helm charts for Kubernetes clusters Weave Cloud Kubernetes deployment manifests AWS ECS Services, which you can define as code using your favorite stack management tool Azure App Service Plans CNAB Cloud Native Application Bundle Different deployment manifests and packages work at different levels. Some are focused on a single deployable unit, so you need a manifest for each application. Some define a collection of deployable services. Depending on the tool, each of the services within the collection may have a separate manifest, with a higher-level manifest defining common elements and integration parameters. A manifest for the ShopSpinner web server deployment might look like the pseudocode shown in Example 10-1. Example 10-1. Example of an application cluster deployment manifest service:
name: webservers
organization: ShopSpinner
version: 1.0.3
application:
name: nginx
container_image:
repository: containers.shopspinner.xyz
path: /images/nginx
tag: 1.0.3
instance:
count_min: 3
count_max: 10
health_port: 443
health_url: /alive
connectivity:
inbound:
id: https_inbound
port: 443
allow_from: $PUBLIC_INTERNET
ssl_cert: $SHOPSPINNER_PUBLIC_SSL_CERT
outbound:
port: 443
allow_to: [ $APPSERVER_APPLICATIONS.https_inbound ] This example specifies where and how to find the container image (the container_image block), how many instances to run, and how to check their health. It also defines inbound and outbound connection rules.
クラスターにアプリケーションをデプロイするためのパッケージ
現代のアプリケーションでは、複数のプロセスやコンポーネントが複雑なインフラストラクチャに展開されることがよくあります。ランタイム環境は、これらのさまざまな要素を実行する方法を知る必要があります。最小および最大のインスタンス数は何ですか？いつインスタンスを追加または削除するかをランタイムはどのように知るのでしょうか？インスタンスが健康であるか再起動が必要か、ランタイムはどのように知るのでしょうか？各インスタンスにはどのストレージが必要で取り付けられる必要があるのでしょうか？連携およびセキュリティの要件は何ですか？異なるランタイムプラットフォームは異なる機能を提供し、多くのプラットフォームには独自のパッケージ化および設定形式があります。これらのプラットフォームでは、展開アーティファクト（たとえばコンテナイメージ）を含むアーカイブファイルの代わりに、デプロイメントマニフェストが使用されることがよくあります。クラスタアプリケーションの展開マニフェストの例は次のとおりです。

- KubernetesクラスタのためのHelmチャート
- Weave Cloud Kubernetes展開マニフェスト
- AWS ECSサービス（お気に入りのスタック管理ツールを使用してコードとして定義できます）
- Azure App Serviceプラン
- CNABクラウドネイティブアプリケーションバンドル

さまざまな展開マニフェストとパッケージは、さまざまなレベルで機能します。単一の展開可能なユニットに焦点を当てたものもあれば、展開可能なサービスのコレクションを定義したものもあります。ツールによっては、コレクション内の各サービスには別々のマニフェストがあり、共通の要素や統合パラメータを定義する上位レベルのマニフェストもあります。ShopSpinnerウェブサーバーの展開に対するマニフェストの疑似コードの例は、次のようになります。

service:
  name: webservers
  organization: ShopSpinner
  version: 1.0.3

application:
  name: nginx
  container_image:
    repository: containers.shopspinner.xyz
    path: /images/nginx
    tag: 1.0.3

instance:
  count_min: 3
  count_max: 10
  health_port: 443
  health_url: /alive

connectivity:
  inbound:
    id: https_inbound
    port: 443
    allow_from: $PUBLIC_INTERNET
    ssl_cert: $SHOPSPINNER_PUBLIC_SSL_CERT
  outbound:
    port: 443
    allow_to: [$APPSERVER_APPLICATIONS.https_inbound]

この例では、コンテナイメージの位置や方法（container_imageブロック）、実行するインスタンス数、およびその健康状態を確認する方法が指定されています。さらに、インバウンドおよびアウトバウンドの接続ルールも定義されています。