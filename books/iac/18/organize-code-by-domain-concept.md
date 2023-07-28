Organize Code by Domain Concept Code within a single project can include multiple pieces. The application infrastructure project in the ShopSpinner example defines a server cluster and a database instance, and networking structures and security policies for each. Many teams define networking structures and security policies in their own files, as shown in Example 18-2. Example 18-2. Source files organized by technology └── src/
       ├── cluster.infra
       ├── database.infra
       ├── load_balancer.infra
       ├── routing.infra
       ├── firewall_rules.infra
       └── policies.infra The firewall_rules.infra file includes firewall rules for the virtual machines created in cluster.infra as well as rules for the database instance defined in database.infra. Organizing code this way focuses on the functional elements over how they’re used. It’s often easier to understand, write, change, and maintain the code for related elements when they’re in the same file. Imagine a file with thirty different firewall rules for access to eight different services, versus a file that defines one service, and the three firewall rules related to it.
This concept follows the design principle of designing around domain concepts rather than technical ones (see “Design components around domain concepts, not technical ones”). Organizing Configuration Value Files Chapter 7 described the configuration files pattern for managing parameter values for different instances of a stack project (see “Pattern: Stack Configuration Files”). The description suggested two different ways to organize per-environment configuration files across multiple projects. One is to store them within the relevant project: ├── application_infra_stack/
   │   ├── src/
   │   └── environments/
   │       ├── test.properties
   │       ├── staging.properties
   │       └── production.properties
   │
   └── shared_network_stack/
       ├── src/
       └── environments/
           ├── test.properties
           ├── staging.properties
           └── production.properties The other is to have a separate project with the configuration for all of the stacks, organized by environment: ├── application_infra_stack/
   │   └── src/
   │
   ├── shared_network_stack/
   │   └── src/
│
   └── configuration/
       ├── test/
       │   ├── application_infra.properties
       │   └── shared_network.properties
       ├── staging/
       │   ├── application_infra.properties
       │   └── shared_network.properties
       └── production/
           ├── application_infra.properties
           └── shared_network.properties Storing configuration values with the code for a project mixes generalized, reusable code (assuming it’s a reusable stack, per “Pattern: Reusable Stack”) with details of specific instances. Ideally, changing the configuration for an environment shouldn’t require modifying the stack project. On the other hand, it’s arguably easier to trace and understand configuration values when they’re close to the projects they relate to, rather than mingled in a monolithic configuration project. Team ownership and alignment is a factor, as usual. Separating infrastructure code and its configuration can discourage taking ownership and responsibility across both.

ドメイン概念によってコードを整理する 単一のプロジェクト内のコードは複数の部分で構成されることがあります。ShopSpinnerの例でのアプリケーションインフラストラクチャプロジェクトでは、サーバクラスタとデータベースインスタンス、およびネットワーキング構造とセキュリティポリシーが定義されています。多くのチームは、ネットワーキング構造とセキュリティポリシーを独自のファイルに定義します（例18-2を参照）。 例18-2.技術別に整理されたソースファイル └──src/
       ├── clusetr.infra
       ├── database.infra
       ├── load_balancer.infra
       ├── routing.infra
       ├── firewall_rules.infra
       └── policies.infra firewall_rules.infraファイルには、clusetr.infraで作成された仮想マシンおよびdatabase.infraで定義されたデータベースインスタンスに対するファイアウォールルールが含まれています。このようにコードを整理することは、使用方法よりも機能要素に焦点を当てています。関連する要素が同じファイルにある場合、関連する要素のコードを理解しやすく、書きやすく、変更しやすく、維持しやすい場合があります。8つの異なるサービスへのアクセスのための30の異なるファイアウォールルールを含むファイルと、1つのサービスを定義し、それに関連する3つのファイアウォールルールを定義するファイルを想像してみてください。 このコンセプトは、技術的なものではなくドメインの概念を基にデザインする原則に従っています（「ドメインの概念に基づいて設計コンポーネントを作成する」を参照）。 構成値ファイルの整理 第7章では、スタックプロジェクトの異なるインスタンスのパラメータ値を管理するための構成ファイルパターンについて説明しました（「パターン：スタック構成ファイル」を参照）。説明では、複数のプロジェクト間で環境ごとの構成ファイルを2つの異なる方法で整理することが提案されています。1つは関連するプロジェクト内にそれらを格納する方法です： ├──application_infra_stack/
   │   ├── src/
   │   └──environments/
   │       ├── test.properties
   │       ├── staging.properties
   │       └── production.properties
   │
   └──shared_network_stack/
       ├── src/
       └──environments/
           ├── test.properties
           ├── staging.properties
           └── production.properties もう1つは、環境別に整理されたすべてのスタックの構成を持つ別のプロジェクトを持つ方法です： ├──application_infra_stack/
   │   └── src/
   │
   ├──shared_network_stack/
   │   └── src/
│
   └──configuration/
       ├── test/
       │   ├── application_infra.properties
       │   └── shared_network.properties
       ├── staging/
       │   ├...