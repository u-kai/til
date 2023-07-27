Pattern: Resource Matching A consumer stack uses resource matching to discover a dependency by looking for infrastructure resources that match names, tags, or other identifying characteristics. For example, a provider stack could name VLANs by the types of resources that belong in the VLAN and the environment of the VLAN (see Figure 17-1)
In this example, vlan-appserver-staging is intended for application servers in the staging environment. The application-infrastructure-stack code finds this resource by matching the naming pattern: virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
  vlan: "vlan-appserver-${ENVIRONMENT_NAME}" A value for ENVIRONMENT_NAME is passed to the stack management tool when applying the code, as described in Chapter 7. Motivation Resource matching is straightforward to implement with most stack management tools and languages. The pattern mostly eliminates hardcoded dependencies, reducing coupling. Resource matching also avoids coupling on tools. The provider infrastructure and consumer stack can be implemented using different tools. Applicability Use resource matching for discovering dependencies when the teams managing provider and consumer code both have a clear understanding of which resources should be used as dependencies. Consider switching to an alternative pattern if you experience issues with breaking dependencies between teams. Resource matching is useful in larger organizations, or across organizations, where different teams may use different tooling to manage their infrastructure, but still need to integrate at the infrastructure level. Even where everyone currently uses a single tool, resource matching reduces lock-in to that tool, creating the option to use new tools for different parts of the system. Consequences As soon as a consumer stack implements resource matching to discover a resource from another stack,
the matching pattern becomes a contract. If someone changes the naming pattern of the VLAN in the shared networking stack, the consumer’s dependency breaks. So a consumer team should only discover dependencies by matching resources in ways that the provider team explicitly supports. Provider teams should clearly communicate what resource matching patterns they support, and be sure to maintain the integrity of those patterns as a contract. Implementation There are several ways to discover infrastructure resources by matching. The most straightforward method is to use variables in the name of the resource, as shown in the example code from earlier: virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
  vlan: "vlan-appserver-${ENVIRONMENT_NAME}" The string vlan-appserver-${ENVIRONMENT_NAME} will match the relevant VLAN for the environment. Most stack languages have features to match other attributes than the resource name. Terraform has data sources and AWS CDK’s supports resource importing. In this example (using pseudocode) the provider assigns tags to its VLANs: vlans:

-   appserver_vlan
    address_range: 10.2.0.0/8
    tags:
    network_tier: "application_servers"
    environment: ${ENVIRONMENT_NAME} The consumer code discovers the VLAN it needs using those tags:
    external_resource:
    id: appserver_vlan
    match:
    tag: name == "network_tier" && value == "application_servers"
    tag: name == "environment" && value == ${ENVIRONMENT_NAME}

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: external_resource.appserver_vlan Related patterns The resource matching pattern is similar to the stack data lookup pattern. The main difference is that resource matching doesn’t depend on the implementation of the same stack tool across provider and consumer stacks.
パターン：リソースのマッチング
コンシューマースタックは、名前、タグ、またはその他の識別特性に一致するインフラストラクチャリソースを探すことによって依存関係を見つけるために、リソースのマッチングを使用します。たとえば、プロバイダースタックでは、VLANをVLANに所属するリソースの種類とVLANの環境に基づいて命名することができます（図17-1を参照）。
この例では、vlan-appserver-stagingはステージング環境のアプリケーションサーバー向けです。アプリケーションインフラストラクチャスタックのコードは、命名パターンに一致することでこのリソースを見つけます：virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: "vlan-appserver-${ENVIRONMENT_NAME}"
ENVIRONMENT_NAMEの値は、コードを適用する際にスタック管理ツールに渡されます（第7章で説明）。
モチベーション
リソースのマッチングは、ほとんどのスタック管理ツールや言語で簡単に実装できます。このパターンは、ハードコードされた依存関係をほとんど排除し、カップリングを減らすことができます。リソースのマッチングはまた、ツールにも依存しません。プロバイダーのインフラストラクチャとコンシューマースタックは、異なるツールを使用して実装することができます。  
適用範囲
プロバイダーとコンシューマーコードの両方のチームが、依存関係として使用するリソースが明確にわかる場合に、依存関係を見つけるためにリソースのマッチングを使用します。チーム間の依存関係の破損に問題がある場合は、代替パターンに切り替えることを検討してください。リソースのマッチングは、異なるチームがそれぞれのインフラストラクチャを管理するために異なるツールを使用する必要がある、または組織を横断する大規模な組織で有用です。現在はすべての人が単一のツールを使用している場合でも、リソースのマッチングによってそのツールに依存することを減らし、システムのさまざまな部分に新しいツールを使用するオプションを作成できます。
結果
コンシューマースタックが他のスタックからリソースを発見するためにリソースのマッチングを実装すると、一致するパターンは契約となります。共有ネットワーキングスタックのVLANの命名パターンを変更すると、コンシューマーの依存関係が失われます。したがって、コンシューマーチームは、プロバイダーチームが明示的にサポートする方法でリソースをマッチングしてのみ依存関係を発見するべきです。プロバイダーチームは、どのリソースのマッチングパターンをサポートするかを明確に伝えるべきであり、それらのパターンを契約として維持する必要があります。

実装
一致させるためのインフラストラクチャリソースを発見する方法はいくつかあります。最も直接的な方法は、リソースの名前に変数を使用することです。前述の例コードのように：virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: "vlan-appserver-${ENVIRONMENT_NAME}"
文字列vlan-appserver-${ENVIRONMENT_NAME}は、環境に関連する適切なVLANに一致します。ほとんどのスタック言語には、リソース名以外の属性と一致する機能があります。Terraformにはデータソースがあり、AWS CDKではリソースのインポートがサポートされています。この例（疑似コードを使用）では、プロバイダーはVLANにタグを割り当てます：vlans：
-   appserver_vlan
    address_range: 10.2.0.0/8
    tags:
    network_tier: "application_servers"
    environment: ${ENVIRONMENT_NAME}
コンシューマーコードは、これらのタグを使用して必要なVLANを見つけます：
external_resource:
    id: appserver_vlan
    match:
    tag: name == "network_tier" && value == "application_servers"
    tag: name == "environment" && value == ${ENVIRONMENT_NAME}

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: external_resource.appserver_vlan
関連パターン
リソースのマッチングパターンは、スタックデータルックアップパターンに類似しています。主な違いは、リソースのマッチングが、プロバイダースタックとコンシューマースタックの間で同じスタックツールの実装に依存しないことです。