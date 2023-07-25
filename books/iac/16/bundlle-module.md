Pattern: Bundle Module A bundle module declares a collection of related infrastructure resources with a simplified interface. The stack code uses the module to define what it needs to provision: use module: application_server
service_name: checkout_service
application_name: checkout_application
application_version: 1.23
min_cluster: 1
max_cluster: 3
ram_required: 4GB The module code declares multiple infrastructure resources, usually centered on a core resource. In Example 16-5, the resource is a server cluster, but also includes a load balancer and DNS entry. Example 16-5. Module code for an application server declare module: application_server

server_cluster:
id: "${service_name}-cluster"
    min_size: ${min_cluster}
    max_size: ${max_cluster}
    each_server_node:
      source_image: base_linux
      memory: ${ram_required}
      provision:
        tool: servermaker
        role: appserver
        parameters:
          app_package: "${checkout_application}-${application_version}.war"
app_repository: "repository.shopspinner.xyz"

load_balancer:
protocol: https
target:
type: server_cluster
target_id: "${service_name}-cluster"

dns_entry:
id: "${service_name}-hostname"
    record_type: "A"
hostname: "${service_name}.shopspinner.xyz"
ip_address: {$load_balancer.ip_address} Motivation A bundle module is useful to define a cohesive collection of infrastructure resources. It avoids verbose, redundant code. These modules are useful to capture knowledge about the various elements needed and how to wire them together for a common purpose. Applicability A bundle module is suitable when you’re working with a declarative stack language, and when the resources involved don’t vary in different use cases. If you find that you need the module to create different resources or configure them differently depending on the usage, then you should either create separate modules, or else switch to an imperative language and create an infrastructure domain entity (see “Pattern: Infrastructure Domain Entity”). Consequences A bundle module may provision more resources than you need in some situations. Users of the module should understand what it provisions, and avoid using the module if it’s overkill for their use case. Implementation Define the module declaratively, including infrastructure elements that are closely related to the declared purpose. Related patterns A facade module (“Pattern: Facade Module”) wraps a single infrastructure resource, while a bundle module includes multiple resources, although both are declarative in nature. An infrastructure domain entity (“Pattern: Infrastructure Domain Entity”) is similar to a bundle module, but dynamically generates
generates infrastructure resources. A spaghetti module is a bundle module that wishes it was a domain entity but descends into madness thanks to the limitations of its declarative language.
パターン：バンドルモジュール
バンドルモジュールは、簡略化されたインターフェースで関連するインフラストラクチャリソースのコレクションを宣言します。スタックコードは、必要なものをプロビジョニングするためにモジュールを使用します：use module: application_server
service_name: checkout_service
application_name: checkout_application
application_version: 1.23
min_cluster: 1
max_cluster: 3
ram_required: 4GB
モジュールコードは、通常、コアリソースを中心とした複数のインフラストラクチャリソースを宣言します。例16-5では、リソースはサーバクラスタですが、ロードバランサとDNSエントリも含まれます。例16-5. アプリケーションサーバのためのモジュールコードを宣言します。

server_cluster:
id: "${service_name}-cluster"
    min_size: ${min_cluster}
    max_size: ${max_cluster}
    each_server_node:
      source_image: base_linux
      memory: ${ram_required}
      provision:
        tool: servermaker
        role: appserver
        parameters:
          app_package: "${checkout_application}-${application_version}.war"
app_repository: "repository.shopspinner.xyz"

load_balancer:
protocol: https
target:
type: server_cluster
target_id: "${service_name}-cluster"

dns_entry:
id: "${service_name}-hostname"
    record_type: "A"
hostname: "${service_name}.shopspinner.xyz"
ip_address: {$load_balancer.ip_address}
動機
バンドルモジュールは、一貫したインフラストラクチャリソースのコレクションを定義するために役立ちます。冗長なコードを避けることができます。これらのモジュールは、必要な要素やそれらを共通の目的のためにどのように組み合わせるかについての知識を把握するために役立ちます。
適用範囲
バンドルモジュールは、宣言的なスタック言語で作業している場合や、異なるユースケースで変化するリソースに関与していない場合に適しています。ユースケースに応じてモジュールを使用して異なるリソースを作成したり、それらを異なる方法で構成する必要がある場合は、別々のモジュールを作成するか、または手続き型言語に切り替えてインフラストラクチャドメインエンティティを作成する必要があります（「パターン：インフラストラクチャドメインエンティティ」を参照）。
結果
バンドルモジュールは、一部の状況で必要以上のリソースをプロビジョニングする場合があります。モジュールの使用者は、それが自分のユースケースには過剰である場合は使用しないようにし、モジュールがプロビジョニングする内容を理解する必要があります。
実装
目的に密接に関連するインフラストラクチャ要素を含めて、モジュールを宣言的に定義します。
関連パターン
ファサードモジュール（「パターン：ファサードモジュール」）は、単一のインフラストラクチャリソースをラップしますが、バンドルモジュールは複数のリソースを含みます。両者とも宣言的な性質です。インフラストラクチャドメインエンティティ（「パターン：インフラストラクチャドメインエンティティ」）は、バンドルモジュールに似ていますが、動的にインフラストラクチャリソースを生成します。スパゲッティモジュールは、ドメインエンティティであることを望んでいるバンドルモジュールですが、宣言的な言語の制限により狂気に陥ってしまいます。