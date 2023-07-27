Pattern: Integration Registry Lookup Also known as: integration registry. A consumer stack can use integration registry lookup to discover a resource published by a provider stack (Figure 17-3). Both stacks refer to a registry, using a known location to store and read values.
Many stack tools support storing and retrieving values from different types of registries within definition code. The shared-networking-stack code sets the value: vlans:
  - appserver_vlan
    address_range: 10.2.0.0/8

registry:
  host: registry.shopspinner.xyz
  set:
    /${ENVIRONMENT_NAME}/shared-networking/appserver_vlan: appserver_vlan.id The application-infrastructure-stack code then retrieves and uses the value: registry:
  id: stack_registry
  host: registry.shopspinner.xyz
  values:
    appserver_vlan_id: /${ENVIRONMENT_NAME}/shared-networking/appserver_vlan

virtual_machine:
  name: "appserver-${ENVIRONMENT_NAME}"
  vlan: stack_registry.appserver_vlan_id Motivation Using a configuration registry decouples the stack management tools for different infrastructure stacks. Different teams can use different tools as long as they agree to use the same configuration registry service and naming conventions for storing values in it. This decoupling also makes it easier to upgrade and change tools incrementally, one stack at a time.
Using a configuration registry makes the integration points between stacks explicit. A consumer stack can only use values explicitly published by a provider stack, so the provider team can freely change how they implement their resources. Applicability The integration registry lookup pattern is useful for larger organizations, where different teams may use different technologies. It’s also useful for organizations concerned about lock-in to a tool. If your system already uses a configuration registry (“Configuration Registry”), for instance, to provide configuration values to stack instances following the stack parameter registry pattern (“Pattern: Stack Parameter Registry”), it can make sense to use the same registry for integrating stacks. Consequences The configuration registry becomes a critical service when you adopt the integration registry lookup pattern. You may not be able to provision or recover resources when the registry is unavailable. Implementation The configuration registry is a prerequisite for using this pattern. See “Configuration Registry” in Chapter 7 for a discussion of registries. Some infrastructure tool vendors provide registry servers, as mentioned in “Infrastructure automation tool registries”. With any registry product, be sure it’s well-supported by the tools you use, and those you might consider using in the future. It’s essential to establish a clear convention for naming parameters, especially when using the registry to integrate infrastructure across multiple teams. Many organizations use a hierarchical namespace, similar to a directory or folder structure, even when the registry product implements a simple key/value mechanism. The structure typically includes components for architectural units (such as services, applications, or products), environments, geography, or teams. For example, ShopSpinner could use a hierarchical path based on the geographical region:
/infrastructure/
 ├── au/
 │   ├── shared-networking/
 │   │   └── appserver_vlan=
 │   └── application-infrastructure/
 │       └── appserver_ip_address=
 └── eu/
     ├── shared-networking/
     │   └── appserver_vlan=
     └── application-infrastructure/
         └── appserver_ip_address= The IP address of the application server for the European region, in this example, is found at the location /infrastructure/eu/application-infrastructure/appserver​_ip_address. Related patterns Like this pattern, the stack data lookup pattern (see “Pattern: Stack Data Lookup”) stores and retrieves values from a registry. That pattern uses data structures specific to the stack management tool, whereas this pattern uses a general-purpose registry implementation. The parameter registry pattern (see “Pattern: Stack Parameter Registry”) is essentially the same as this pattern, in that a stack pulls values from a registry to use in a given stack instance. The only difference is that with this pattern, the value comes from another stack, and is used explicitly to integrate infrastructure resources between stacks.

パターン：統合レジストリの検索
別名：統合レジストリ。コンシューマースタックは、統合レジストリの検索を使用して、プロバイダースタックによって公開されたリソースを発見することができます（図17-3）。両方のスタックは、値を保存および読み取るために既知の場所を使用してレジストリを参照しています。
多くのスタックツールは、定義コード内でさまざまなタイプのレジストリから値を保存および取得することをサポートしています。共有ネットワーキングスタックのコードは以下の値を設定します：vlans：
   - appserver_vlan
   address_range: 10.2.0.0/8

レジストリ：
　　host: registry.shopspinner.xyz
　　set:
　　　 /${ENVIRONMENT_NAME}/shared-networking/appserver_vlan: appserver_vlan.id

次に、アプリケーションインフラストラクチャスタックのコードは以下の値を取得して使用します：レジストリ：
   id: stack_registry
   host: registry.shopspinner.xyz
   values:
      appserver_vlan_id: /${ENVIRONMENT_NAME}/shared-networking/appserver_vlan

virtual_machine:
   name: "appserver-${ENVIRONMENT_NAME}"
   vlan: stack_registry.appserver_vlan_id

動機
構成レジストリの使用により、異なるインフラストラクチャスタックのスタック管理ツールが分離されます。異なるチームは、同じ構成レジストリサービスと値の格納方法の規則に同意する限り、さまざまなツールを使用することができます。この分離により、ツールのアップグレードや変更を一つのスタックずつ段階的に行うことが容易になります。

構成レジストリの使用により、スタック間の統合ポイントが明示的になります。コンシューマースタックは、プロバイダースタックによって明示的に公開された値のみを使用できるため、プロバイダーチームはリソースの実装方法を自由に変更することができます。

適用範囲
統合レジストリの検索パターンは、異なる技術を使用する可能性のある大規模な組織にとって有用です。特定のツールに依存することを懸念する組織にとっても便利です。たとえば、既に構成レジストリ（「構成レジストリ」）を使用している場合は、スタックパラメータレジストリパターン（「パターン：スタックパラメータレジストリ」）に従ってスタックインスタンスに構成値を提供するために同じレジストリを使用することが合理的です。

結果
統合レジストリの検索パターンを採用すると、構成レジストリは重要なサービスとなります。レジストリが利用できない場合、リソースをプロビジョニングまたは回復することができなくなる可能性があります。

実装
このパターンを使用するには、構成レジストリが前提条件です。レジストリに関するディスカッションについては、「レジストリ」を参照してください（第7章）。一部のインフラストラクチャツールベンダーは、レジストリサーバーを提供しています（「インフラストラクチャ自動化ツールのレジストリ」で言及）。レジストリ製品を使用する場合は、使用しているツールおよび将来使用を検討しているツールによって十分にサポートされていることを確認してください。また、複数のチーム間でインフラストラクチャを統合する場合には、パラメータの命名に明確な規則を確立することが重要です。レジストリ製品がシンプルなキー/値メカニズムを実装している場合でも、多くの組織はディレクトリまたはフォルダ構造に類似した階層的な名前空間を使用します。この構造には、アーキテクチャユニット（サービス、アプリケーション、製品など）、環境、地理、チームなどのコンポーネントが通常含まれます。

関連するパターン
このパターンと同様に、スタックデータの検索パターン（「パターン：スタックデータの検索」）は、レジストリから値を保存および取得します。このパターンでは、スタック管理ツール固有のデータ構造を使用します。一方、このパターンでは汎用的なレジストリの実装が使用されます。パラメータレジストリパターン（「パターン：スタックパラメータレジストリ」）は、スタックがレジストリから値を取得して特定のスタックインスタンスで使用するという点で、基本的にはこのパターンと同じです。唯一の違いは、このパターンでは値が別のスタックから取得され、スタック間のインフラストラクチャリソースを明示的に統合するために使用されることです。