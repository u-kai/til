Pattern: Wrapper Stack A wrapper stack uses an infrastructure stack project for each instance as a wrapper to import a stack code component (see Chapter 16). Each wrapper project defines the parameter values for one instance of the stack. It then imports a component shared by all of the stack instances (see Figure 7-2).
Motivation A wrapper stack leverages the stack tool’s module functionality or library support to reuse shared code across stack instances. You can use the tool’s module versioning, dependency management, and artifact repository functionality to implement a change delivery pipeline (see “Infrastructure Delivery Pipelines”). As of this writing, most infrastructure stack tools don’t
have a project packaging format that you can use to implement pipelines for stack code. So you need to create a custom stack packaging process yourself. You can work around this by using a wrapper stack, and versioning and promoting your stack code as a module. With wrapper stacks, you can write the logic for provisioning and configuring stacks in the same language that you use to define your infrastructure, rather than using a separate language as you would with a provisioning script (see “Pattern: Scripted Parameters”). Consequences Components add an extra layer of complexity between your stack and the code contained in the component. You now have two levels: the stack project, which contains the wrapper projects, and the component that contains the code for the stack. Because you have a separate code project for each stack instance, people may be tempted to add custom logic for each instance. Custom instance code makes your codebase inconsistent and hard to maintain. Because you define parameter values in wrapper projects managed in source control, you can’t use this pattern to manage secrets. So you need to add another pattern from this chapter to provide secrets to stacks. Implementation Each stack instance has a separate infrastructure stack project. For example, you would have a separate Terraform project for each environment. You can implement this like a copy-paste environment (see “Antipattern: Copy-Paste Environments”), with each environment in a separate repository. Alternatively, each environment project could be a folder in a single repository: my_stack/
├── test/
│ └── stack.infra
├── staging/
│ └── stack.infra
└── production/
└── stack.infra
Define the infrastructure code for the stack as a module, according to your tool’s implementation. You could put the module code in the same repository with your wrapper stacks. However,
this would prevent you from leveraging module versioning functionality. That is, you wouldn’t be able to use different versions of the infrastructure code in different environments, which is crucial for progressively testing your code. The following example is a wrapper stack that imports a module called container_cluster_module, specifying the version of the module, and the configuration parameters to pass to it: module:
name: container_cluster_module
version: 1.23
parameters:
env: test
cluster_minimum: 1
cluster_maximum: 1 The wrapper stack code for the staging and production environments is similar, other than the parameter values, and perhaps the module version they use. The project structure for the module could look like this: ├── container_cluster_module/
│ ├── cluster.infra
│ └── networking.infra
└── test/ When you make a change to the module code, you test and upload it to a module repository. How the repository works depends on your particular infrastructure stack tool. You can then update your test stack instance to import the new module version and apply it to the test environment. Terragrunt is a stack orchestration tool that implements the wrapper stack pattern. Related patterns A wrapper stack is similar to the scripted parameters pattern. The main differences are that it uses your stack tool’s language rather than a separate scripting language and that the infrastructure code is in a separate component.
パターン: Wrapper Stack
ラッパースタックは、各インスタンスに対してラッパーとしてインフラストラクチャスタックプロジェクトを使用し、スタックコードコンポーネントをインポートします（第16章を参照）。各ラッパープロジェクトは、スタックの1つのインスタンスのパラメータ値を定義します。それから、スタックのすべてのインスタンスで共有されるコンポーネントをインポートします（図7-2を参照）。

動機
ラッパースタックは、スタックツールのモジュール機能やライブラリサポートを活用して、スタックのインスタンス間で共有されるコードを再利用します。ツールのモジュールのバージョン管理、依存関係管理、およびアーティファクトリポジトリの機能を使用して、変更デリバリーパイプラインを実装することができます（「インフラストラクチャデリバリーパイプライン」を参照）。執筆時点では、ほとんどのインフラストラクチャスタックツールには、スタックコードのパイプラインを実装するために使用できるプロジェクトパッケージング形式がありません。そのため、カスタムのスタックパッケージングプロセスを作成する必要があります。ラッパースタックを使用することで、モジュールとしてスタックコードをバージョニングしてプロモーションすることができます。ラッパースタックでは、スタックのプロビジョニングと設定のロジックを、別の言語ではなく、インフラストラクチャの定義に使用する言語で記述することができます（「パターン: スクリプト化されたパラメータ」を参照）。

結果
コンポーネントは、スタックとコンポーネント内のコードとの間に余分な複雑さのレイヤーを追加します。スタックプロジェクト内にラッパープロジェクトが含まれ、その中にスタックのコードが含まれています。各スタックインスタンスごとに別個のコードプロジェクトがあるため、それぞれのインスタンスにカスタムロジックを追加することが誘発されるかもしれません。カスタムインスタンスコードは、コードベースを一貫性がなくして保守が難しくします。ラッパープロジェクト内でパラメータ値を定義してソースコントロールで管理するため、このパターンを使用してシークレットを管理することはできません。そのため、スタックへのシークレットの提供をするために、この章の別のパターンを追加する必要があります。

実装
各スタックインスタンスには、個別のインフラストラクチャスタックプロジェクトがあります。例えば、各環境には個別のTerraformプロジェクトが必要です。これは、個々のリポジトリ内の環境ごとにコピー＆ペーストする方法で実装することができます（「アンチパターン: コピー＆ペースト環境」を参照）。または、各環境プロジェクトを単一のリポジトリのフォルダとして構成することもできます：my_stack/
├── test/
│ └── stack.infra
├── staging/
│ └── stack.infra
└── production/
└── stack.infra
スタックのインフラストラクチャコードを、ツールの実装に応じてモジュールとして定義します。モジュールコードをラッパースタックと同じリポジトリに置くこともできますが、これによりモジュールのバージョン管理機能を活用することはできません。つまり、異なるバージョンのインフラストラクチャコードを異なる環境で使用することができず、コードを段階的にテストするためには重要な要素です。以下の例は、container_cluster_moduleというモジュールをインポートし、そのバージョンと渡す設定パラメータを指定するラッパースタックです。

module:
name: container_cluster_module
version: 1.23
parameters:
env: test
cluster_minimum: 1
cluster_maximum: 1

ステージングおよび製品環境用のラッパースタックコードは、パラメータ値や使用するモジュールバージョンを除いて、類似しています。モジュールのプロジェクト構造は次のようになります。

├── container_cluster_module/
│ ├── cluster.infra
│ └── networking.infra
└── test/

モジュールコードに変更を加えた場合、テストしてモジュールリポジトリにアップロードします。リポジトリの動作方法は、特定のインフラストラクチャスタックツールによって異なります。その後、テストスタックインスタンスを更新して新しいモジュールバージョンをインポートし、テスト環境に適用します。Terragruntは、ラッパースタックパターンを実装したスタックオーケストレーションツールです。

関連パターン
ラッパースタックは、スクリプト化されたパラメータのパターンに類似しています。主な違いは、別個のスクリプト言語ではなく、スタックツールの言語を使用していること、およびインフラストラクチャコードが別のコンポーネントにあることです。