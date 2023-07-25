Infrastructure Languages for Stack Components Chapter 4 describes different types of infrastructure code languages. The two main styles of language for defining stacks are declarative (see “Declarative Infrastructure Languages”) and imperative (see “Programmable, Imperative Infrastructure Languages”). That chapter mentions that these different types of languages are suitable for different types of code (see “Declarative Versus Imperative Languages for Infrastructure”). These differences often collide with stack components, when people write them using the wrong type of language. Using the wrong type of language usually leads to a mixture of both declarative and imperative languages, which, as explained previously, is a Bad Thing (see “Separate Declarative and Imperative Code”). The decision of which language to use tends to be driven by which infrastructure stack management tool you’re using, and the languages it supports.1 The patterns defined later in this chapter should encourage you to think about what you’re trying to achieve with a particular stack and its components. To use this to consider the type of language, and potentially the type of stack tool, that you should use, consider two classes of stack component based on language type. Reuse Declarative Code with Modules Most stack management tools with declarative languages let you write shared components using the same language. CloudFormation has nested stacks and Terraform has modules. You can pass parameters to these modules, and the languages have at least some programmability (such as the HCL expressions sublanguage for Terraform). But the languages are fundamentally declarative, so nearly all complex logic written in them is barbaric. So declarative code modules work best for defining infrastructure components that don’t vary very much. A declarative module works well for a facade module (see “Pattern: Facade Module”), which wraps and simplifies a resource provided by the infrastructure platform. These modules get nasty when you use them for more complex cases, creating spaghetti modules (see “Antipattern: Spaghetti Module”). As mentioned in “Challenge: Tests for Declarative Code Often Have Low Value”, testing a declarative module should be fairly simple. The results of applying a declarative module don’t vary very much, so you don’t need comprehensive test coverage. This doesn’t mean you shouldn’t write tests for these modules. When a module combines multiple declarations to create a more complex entity, you should test that it fulfills its requirement. Dynamically Create Stack Elements with Libraries Some stack management tools, like Pulumi and the AWS CDK, use general-purpose, imperative languages. You can use these languages to write reusable libraries that you can call from your stack project code. A library can include more complex logic that dynamically provisions infrastructure resources according to how it’s used. For example, the ShopSpinner team’s infrastructure includes different application server infrastructure stacks. Each of these stacks provisions an application server and networking structures for that application. Some of the applications are public facing, and others are internally facing. In either case, the infrastructure stack needs to assign an IP address and DNS name to the server and create a networking route from the relevant gateway. The IP address and DNS name will be different for a public-facing application versus an internally facing one. And a public-facing application needs a firewall rule to allow the connection. The checkout_service stack hosts a public-facing application:
application_networking = new ApplicationServerNetwork(PUBLIC_FACING, "checkout")

virtual_machine:
name: appserver-checkout
vlan: $(application_networking.address_block)
ip_address: $(application_networking.private_ip_address) The stack code creates an ApplicationServerNetwork object from the application_networking library, which provisions or references the necessary infrastructure elements: class ApplicationServerNetwork {

def vlan;
def public_ip_address;
def private_ip_address;
def gateway;
def dns_hostname;

public ApplicationServerNetwork(application_access_type, hostname) {
if (application_access_type == PUBLIC_FACING) {
vlan = get_public_vlan()
public_ip_address = allocate_public_ip()
dns_hostname = PublicDNS.set_host_record(
"${hostname}.shopspinners.xyz",
this.public_ip_address
)
} else {
// Similar stuff but for a private VLAN
}

    private_ip_address = allocate_ip_from(this.vlan)
    gateway = get_gateway(this.vlan)
    create_route(gateway, this.private_ip_addressif (application_access_type == PUBLIC_FACING) {
      create_firewall_rule(ALLOW, '0.0.0.0', this.private_ip_address, 443)
    }

}
} This pseudocode assigns the server to a public VLAN that already exists and sets its private IP address from the VLAN’s address range. It also sets a public DNS entry for the server, which in our example will be checkout.shopspinners.xyz. The library finds the gateway based on the VLAN that was used, so this would be different for an internally facing application.
「スタックコンポーネントのためのインフラストラクチャ言語」第4章では、異なるタイプのインフラストラクチャコード言語が説明されています。スタックを定義するための主要な言語スタイルは、宣言的なもの（「宣言的インフラストラクチャ言語」を参照）と命令的なもの（「プログラマブルな命令的インフラストラクチャ言語」を参照）の2つです。この章では、これらの異なるタイプの言語が異なる種類のコードに適していることが述べられています（「インフラストラクチャのための宣言的対命令的言語」を参照）。これらの違いは、スタックコンポーネントでしばしば衝突し、人々が間違ったタイプの言語でそれらを記述すると起こります。間違ったタイプの言語を使用することは通常、宣言的言語と命令的言語の両方が混在した結果をもたらし、それは先述のように悪影響を及ぼします（「宣言的と命令的のコードを分離する」を参照）。どの言語を使用するかの決定は、使用しているインフラストラクチャスタック管理ツールと、そのツールがサポートしている言語によって決まる傾向があります。この章の後半で定義されるパターンは、特定のスタックとそのコンポーネントで何を達成しようとしているかを考えることを促すものです。特定のスタックコンポーネントのタイプを考慮するために、言語タイプに基づいて2つのクラスのスタックコンポーネントを考えてみましょう。モジュールを使用した宣言的コードの再利用宣言的な言語を持つほとんどのスタック管理ツールでは、同じ言語を使用して共有コンポーネントを記述することができます。CloudFormationにはネストされたスタック、Terraformにはモジュールがあります。これらのモジュールにパラメータを渡すことができ、言語には少なくともいくつかのプログラム可能性（TerraformのHCL式のサブ言語など）があります。しかし、これらの言語は基本的に宣言的なため、それらで書かれたほぼすべての複雑なロジックは原始的です。したがって、ほとんど変化しないインフラストラクチャコンポーネントを定義するためには、宣言的なコードモジュールが最適です。宣言的なモジュールはファサードモジュール（「パターン：ファサードモジュール」を参照）にもよく適しており、インフラストラクチャプラットフォームによって提供されるリソースをラップして単純化します。これらのモジュールは、より複雑なケースで使用するとやや面倒です（「アンチパターン：スパゲッティモジュール」を参照）。宣言的なモジュールのテストは比較的シンプルであるべきです（「チャレンジ：宣言的コードのテストは低い価値を持つことが多い」で述べられています）。宣言的モジュールの適用結果はあまり変わらないため、包括的なテストカバレッジは必要ありません。しかし、これはこれらのモジュールのテストを書かないべきではないことを意味するものではありません。モジュールが複数の宣言を組み合わせてより複雑なエンティティを作成する場合は、その要件を満たしているかをテストする必要があります。ライブラリを使用してスタック要素を動적に作成する一部のスタック管理ツール（PulumiやAWS CDKなど）は、汎用の命令的な言語を使用しています。これらの言語を使用して、スタックプロジェクトコードから呼び出すことができる再利用可能なライブラリを作成することができます。ライブラリには、使用方法に応じて動的にインフラストラクチャリソースをプロビジョニングするより複雑なロジックを含めることができます。例えば、ShopSpinnerのチームのインフラストラクチャには、異なるアプリケーションサーバのインフラストラクチャスタックが含まれています。これらのスタックのそれぞれは、そのアプリケーションに対してアプリケーションサーバとネットワーキング構造をプロビジョニングします。いくつかのアプリケーションは公開向けであり、他は内部向けです。どちらの場合でも、インフラストラクチャスタックはサーバにIPアドレスとDNS名を割り当て、関連するゲートウェイからのネットワーキングルートを作成する必要があります。公開向けアプリケーションは接続を許可するためにファイアウォールルールが必要です。checkout_serviceスタックは公開向けアプリケーションをホストします：application_networking = new ApplicationServerNetwork(PUBLIC_FACING, "checkout")

virtual_machine:
name: appserver-checkout
vlan: $(application_networking.address_block)
ip_address: $(application_networking.private_ip_address) スタックコードは、アプリケーション_networkingライブラリからApplicationServerNetworkオブジェクトを作成し、必要なインフラストラクチャ要素をプロビジョニングまたは参照します：class ApplicationServerNetwork {

def vlan;
def public_ip_address;
def private_ip_address;
def gateway;
def dns_hostname;

public ApplicationServerNetwork(application_access_type, hostname) {
if (application_access_type == PUBLIC_FACING) {
vlan = get_public_vlan()
public_ip_address = allocate_public_ip()
dns_hostname = PublicDNS.set_host_record(
"${hostname}.shopspinners.xyz",
this.public_ip_address
)
} else {
// その他のプライベートVLANの場合は同様の処理

    private_ip_address = allocate_ip_from(this.vlan)
    gateway = get_gateway(this.vlan)
    create_route(gateway, this.private_ip_addressif (application_access_type == PUBLIC_FACING) {
      create_firewall_rule(ALLOW, '0.0.0.0', this.private_ip_address, 443)
    }

}
} この疑似コードでは、サーバを既存のパブリックVLANに割り当て、そのVLANのアドレス範囲からプライベートIPアドレスを設定しています。また、サーバのパブリックDNSエントリを設定します。この例では、checkout.shopspinners.xyzになります。ライブラリは使用されたVLANに基づいてゲートウェイを見つけるため、内部向けアプリケーションでは異なるゲートウェイになります。