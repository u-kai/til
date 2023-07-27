Dependency Injection These patterns describe strategies for a consumer to discover resources managed by a provider stack. Most stack management tools support directly using these patterns in your stack definition code. But there’s an argument for separating the code that defines a stack’s resources from code that discovers resources to integrate with.
Consider this snippet from the earlier implementation example for the dependency matching pattern (see “Pattern: Resource Matching”): external_resource:
id: appserver_vlan
match:
tag: name == "network_tier" && value == "application_servers"
tag: name == "environment" && value == ${ENVIRONMENT_NAME}

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: external_resource.appserver_vlan The essential part of this code is the declaration of the virtual machine. Everything else in the snippet is peripheral, implementation details for assembling configuration values for the virtual machine. Issues with mixing dependency and definition code Combining dependency discovery with stack definition code adds cognitive overhead to reading or working with the code. Although it doesn’t stop you from getting things done, the subtle friction adds up. You can remove the cognitive overhead by splitting the code into separate files in your stack project. But another issue with including discovery code in a stack definition project is that it couples the stack to the dependency mechanism. The type and depth of coupling to a dependency mechanism and the kind of impact that coupling has will vary for different mechanisms and how you implement them. You should avoid or minimize coupling with the provider stack, and with services like a configuration registry. Coupling dependency management and definition can make it hard to create and test instances of the stack. Many approaches to testing use practices like ephemeral instances (“Pattern: Ephemeral Test
Stack”) or test doubles (“Using Test Fixtures to Handle Dependencies”) to enable fast and frequent testing. This can be challenging if setting up dependencies involves too much work or time. Hardcoding specific assumptions about dependency discovery into stack code can make it less reusable. For example, if you create a core application server infrastructure stack for other teams to use, different teams may want to use different methods to configure and manage their dependencies. Some teams may even want to swap out different provider stacks. For example, they might use different networking stacks for public-facing and internally facing applications. Decoupling dependencies from their discovery Dependency injection (DI) is a technique where a component receives its dependencies, rather than discovering them itself. An infrastructure stack project would declare the resources it depends on as parameters, the same as instance configuration parameters described in Chapter 7. A script or other tool that orchestrates the stack management tool (“Using Scripts to Wrap Infrastructure Tools”) would be responsible for discovering dependencies and passing them to the stack. Consider how the application-infrastructure-stack example used to illustrate the dependency discovery patterns earlier in this chapter would look with DI: parameters:

-   ENVIRONMENT_NAME
-   VLAN

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: ${VLAN} The code declares two parameters that must be set when applying the code to an instance. The ENVIRONMENT_NAME parameter is a simple stack parameter, used for naming the application server virtual machine. The VLAN parameter is the identifier of the VLAN to assign the virtual machine to.
To manage an instance of this stack, you need to discover and provide a value for the VLAN parameter. Your orchestration script can do this, using any of the patterns described in this chapter. It might set the parameter based on tags, by finding the provider stack project’s outputs using the stack management tool, or it could look up the values in a registry. An example script using stack data lookup (see “Pattern: Stack Data Lookup”) might use the stack tool to retrieve the VLAN ID from the provider stack instance, as in Example 17-1, and then pass that value to the stack command for the consumer stack instance: #!/usr/bin/env bash

ENVIRONMENT_NAME=$1

VLAN_ID=$(
  stack value \
    --stack_instance shared_network-${ENVIRONMENT_NAME} \
 --export_name appserver_vlan_id
)

stack apply \
 --stack_instance application_infrastructure-${ENVIRONMENT_NAME} \
  --parameter application_server_vlan=${VLAN_ID} The first command extracts the appserver_vlan_id value from the provider stack instance named shared_network-${ENVIRONMENT_NAME}, and then passes it as a parameter to the consumer stack application_infrastructure-${ENVIRONMENT_NAME}. The benefit of this approach is that the stack definition code is simpler and can be used in different contexts. When you work on changes to the stack code on your laptop, you can pass whatever VLAN value you like. You can apply your code using a local API mock (see “Testing with a Mock API”), or to a personal
instance on your infrastructure platform (see “Personal Infrastructure Instances”). The VLANs you provide in these situations may be very simple. In more production-like environments, the VLAN may be part of a more comprehensive network stack. This ability to swap out different provider implementations makes it easier to implement progressive testing (see “Progressive Testing”), where earlier pipeline stages run quickly and test the consumer component in isolation, and later stages test a more comprehensively integrated system. Origins of Dependency Injection DI originated in the world of object-oriented (OO) software design in the early 2000s. Proponents of XP found that unit testing and TDD were much easier with software written for DI. Java frameworks like PicoContainer and the Spring framework pioneered DI. Martin Fowler’s 2004 article “Inversion of Control Containers and the Dependency Injection Pattern” explains the rationale for the approach for OO software design.
依存性注入
これらのパターンは、コンシューマがプロバイダスタックで管理されるリソースを発見するための戦略を記述しています。ほとんどのスタック管理ツールは、スタックの定義コードで直接これらのパターンを使用できます。ただし、スタックのリソースを発見するコードと統合するコードを分離することについても議論があります。
依存性マッチングパターンの以前の実装例（「パターン：リソースマッチング」を参照）からのスニペットを考えてみましょう。 external_resource:
id: appserver_vlan
match:
tag: name == "network_tier" && value == "application_servers"
tag: name == "environment" && value == ${ENVIRONMENT_NAME}

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: external_resource.appserver_vlan このコードの本質的な部分は、仮想マシンの宣言です。その他のスニペットに含まれるものは、仮想マシンの構成値を組み立てるための実装の詳細です。依存性と定義コードの混在に関する問題
依存性の発見とスタックの定義コードを組み合わせることで、コードの読み取りや処理に認知的な負荷がかかります。これにより、作業を妨げることはありませんが、微妙な摩擦が積み重なります。スタックプロジェクトのコードを別々のファイルに分割することで、認知的な負荷を除去することができます。ただし、スタック定義プロジェクトに発見コードを含める別の問題は、スタックを依存性メカニズムと結合することです。依存性メカニズムへの結合のタイプと深さ、およびその結合がもたらす影響の種類は、異なるメカニズムとその実装方法によって異なります。プロバイダスタックや構成レジストリのようなサービスとの結合を回避または最小限に抑えるべきです。依存性管理と定義を結合することにより、スタックのインスタンスを作成およびテストすることが困難になることがあります。多くのテスト手法では、エフェメラルインスタンス（「パターン：エフェメラルテストスタック」）やテストダブル（「依存性を処理するためのテストフィクスチャの使用」）のような方法を使用して、高速で頻繁なテストを可能にします。これが調整に時間や労力がかかる場合、依存関係の設定が困難になることがあります。スタックコードに依存性の発見に関する具体的な前提条件をハードコードすると、再利用性が低下する可能性があります。たとえば、他のチームが使用するための基本的なアプリケーションサーバーインフラストラクチャスタックを作成した場合、異なるチームは異なる方法で依存関係を構成および管理することを希望する場合があります。一部のチームはさらに、異なるプロバイダスタックを交換したいと考えるかもしれません。たとえば、公開向けと内部向けのアプリケーションに異なるネットワーキングスタックを使用するかもしれません。依存関係を発見から切り離す
依存性インジェクション（DI）は、コンポーネントがその依存関係を自分自身で発見するのではなく、依存関係を受け取る技術です。インフラストラクチャスタックプロジェクトでは、インスタンス構成パラメータと同じように、依存するリソースをパラメータとして宣言します（第7章のインスタンス構成パラメータの説明を参照）。スタック管理ツール（「インフラストラクチャツールをラップするためのスクリプトの使用」）をオーケストレーションするスクリプトや他のツールは、依存関係を発見し、それをスタックに渡す責任を持ちます。この章で以前に説明した依存関係の発見パターンを使用して、依存関係の発見を行うと、アプリケーションインフラストラクチャスタックの例は次のようになります。 parameters:

-   ENVIRONMENT_NAME
-   VLAN

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: ${VLAN} このコードでは、インスタンスにコードを適用する際に設定する必要のある2つのパラメータが宣言されています。ENVIRONMENT_NAMEパラメータは、アプリケーションサーバー仮想マシンに名前を付けるために使用される単純なスタックパラメータです。VLANパラメータは、仮想マシンに割り当てるVLANの識別子です。
このスタックのインスタンスを管理するためには、VLANパラメータの値を発見して提供する必要があります。これは、この章で説明されているパターンのいずれかを使用して、オーケストレーションスクリプトが行うことができます。タグに基づいてパラメータを設定する、スタック管理ツールを使用してプロバイダスタックプロジェクトの出力を検索する、またはレジストリで値を検索することができます。スタックデータを検索するパターン（「パターン：スタックデータの検索」を参照）を使用した例のスクリプトでは、プロバイダスタックインスタンスからVLAN IDを取得し（例17-1を参照）、その値をコンシューマスタックインスタンスに渡します。 #!/usr/bin/env bash

ENVIRONMENT_NAME=$1

VLAN_ID=$(
  stack value \
    --stack_instance shared_network-${ENVIRONMENT_NAME} \
 --export_name appserver_vlan_id
)

stack apply \
 --stack_instance application_infrastructure-${ENVIRONMENT_NAME} \
  --parameter application_server_vlan=${VLAN_ID} 最初のコマンドは、プロバイダスタックインスタンスであるshared_network-${ENVIRONMENT_NAME}からappserver_vlan_idの値を抽出し、それをコンシューマスタックapplication_infrastructure-${ENVIRONMENT_NAME}のパラメータに渡します。このアプローチの利点は、スタック定義コードがシンプルになり、さまざまなコンテキストで使用できることです。スタックコードをラップするスクリプトを作業中にラップトップで実行する場合、お好みのVLAN値を渡すことができます。ローカルAPIモック（「モックAPIを使用したテスト」を参照）やインフラストラクチャプラットフォームの個人向けインスタンス（「個人のインフラストラクチャインスタンス」を参照）にコードを適用することができます。これらの状況で提供するVLANは非常に単純なものかもしれません。より本番に近い環境では、VLANはより包括的なネットワークスタックの一部かもしれません。これにより、異なるプロバイダ実装を交換できるため、前進的なテスト（「プログレッシブテスト」を参照）を実装しやすくなります。前のパイプラインステージは短時間で実行され、コンシューマコンポーネントの単体テストが実施され、後のステージではより包括的に統合されたシステムがテストされます。依存性注入の起源
DIは、2000年代初頭のオブジェクト指向（OO）ソフトウェア設計の世界で生まれました。XPの支持者は、DI向きのソフトウェアではユニットテストとTDDがはるかに容易になることを発見しました。PicoContainerやSpringフレームワークなどのJavaフレームワークがDIを先駆けとして導入しました。Martin Fowlerの2004年の記事「Inversion of Control Containers and the Dependency Injection Pattern」では、OOソフトウェア設計のアプローチの理論を説明しています。