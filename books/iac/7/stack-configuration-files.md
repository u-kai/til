Pattern: Stack Configuration Files Stack configuration files manage parameter values for each instance in a separate file, which you manage in version control with your stack code, such as in Example 7-9. Example 7-9. Example project with a parameter file for each environment ├── src/
│   ├── cluster.infra
│   ├── host_servers.infra
│   └── networking.infra
├── environments/
│ ├── test.properties
│ ├── staging.properties
│ └── production.properties
└── test/
Motivation Creating configuration files for a stack’s instances is straightforward and easy to understand. Because the file is committed to the source code repository, it is easy to:
See what values are used for any given environment (“What is the maximum cluster size for production?”) Trace the history for debugging (“When did the maximum cluster size change?”) Audit changes (“Who changed the maximum cluster size?”) Stack configuration files enforce the separation of configuration from the stack code. Applicability Stack configuration files are appropriate when the number of environments doesn’t change often. They require you to add a file to your project to add a new stack instance. They also require (and help enforce) consistent logic in how different instances are created and updated, since the configuration files can’t include logic. Consequences When you want to create a new stack instance, you need to add a new configuration file to the stack project. Doing this prevents you from automatically creating new environments on the fly. In “Pattern: Ephemeral Test Stack”, I describe an approach for managing test environments that relies on creating environments automatically. You could work around this by creating a configuration file for an ephemeral environment on demand. Parameter files can add friction for changing the configuration of downstream environments in a change delivery pipeline of the kind described in “Infrastructure Delivery Pipelines”. Every change to the stack project code must progress through each stage of the pipeline before being applied to production. It can take a while for this to complete and doesn’t add any value when the configuration change is only applied to production. Defining parameter values can be a source of considerable complexity in provisioning scripts. I’ll talk about this more in “Using Scripts to Wrap Infrastructure Tools”, but as a teaser, consider that teams often want to define default values for stack projects, and for environments, and then need logic to combine these into values for a given instance of a given stack in a different environment. Inheritance models for parameter values can get messy and confusing. Configuration files in source control should not include secrets. So for secrets, you either need to select an additional pattern from this chapter to handle secrets or implement a separate secrets configuration file outside of source control.
Implementation You define stack parameter values in a separate file for each environment, as shown earlier in Example 7-9. The contents of a parameter file could look like this:
env = staging
cluster_minimum = 2
cluster_maximum = 3 Pass the path to the relevant parameter file when running the stack command: stack up --source ./src --config ./environments/staging.properties If the system is composed of multiple stacks, then it can get messy to manage configuration across all of the environments. There are two common ways of arranging parameter files in these cases. One is to put configuration files for all of the environments with the code for each stack: ├── cluster_stack/
│ ├── src/
│ │   ├── cluster.infra
│ │   ├── host_servers.infra
│ │   └── networking.infra
│ └── environments/
│ ├── test.properties
│ ├── staging.properties
│ └── production.properties
│
└── appserver_stack/
├── src/
│   ├── server.infra
│   └── networking.infra
└── environments/
├── test.properties
├── staging.properties
└── production.properties The other is to centralize the configuration for all of the stacks in one place:
├── cluster_stack/
│ ├── cluster.infra
│ ├── host_servers.infra
│ └── networking.infra
│
├── appserver_stack/
│ ├── server.infra
│ └── networking.infra
│
└── environments/
├── test/
│ ├── cluster.properties
│ └── appserver.properties
├── staging/
│ ├── cluster.properties
│ └── appserver.properties
└── production/
├── cluster.properties
└── appserver.properties
Each approach can become messy and confusing in its own way. When you need to make a change to all of the things in an environment, making changes to configuration files across dozens of stack projects is painful. When you need to change the configuration for a single stack across the various environments it’s in, trawling through a tree full of configuration for dozens of other stacks is also not fun. If you want to use configuration files to provide secrets, rather than using a separate pattern for secrets, then you should manage those files outside of the project code checked into source control. For local development environments, you can require users to create the file in a set location manually. Pass the file location to the stack command like this:
stack up --source ./src \
 --config ./environments/staging.properties \
 --config ../.secrets/staging.properties
In this example, you provide two --config arguments to the stack tool, and it reads parameter values from both. You have a directory named .secrets outside the project folder, so it is not in source control.
It can be trickier to do this when running the stack tool automatically, from a compute instance like a CD pipeline agent. You could provision similar secrets property files onto these compute instances, but that can expose secrets to other processes that run on the same agent. You also need to provide the secrets to the process that builds the compute instance for the agent, so you still have a bootstrapping problem. Related patterns Putting configuration values into files simplifies the provisioning scripts described in “Pattern: Scripted Parameters”. You can avoid some of the limitations of environment configuration files by using the stack parameter registry pattern instead (see “Pattern: Stack Parameter Registry”). Doing this moves parameter values out of the stack project code and into a central location, which allows you to use different workflows for code and configuration.
パターン：スタック構成ファイル
スタック構成ファイルは、各インスタンスのパラメータ値を別々のファイルで管理します。これらのファイルはスタックのコードとともにバージョン管理され、例えばExample 7-9のようになります。

Example 7-9. 各環境のパラメータファイルを持つ例のプロジェクト
├── src/
│   ├── cluster.infra
│   ├── host_servers.infra
│   └── networking.infra
├── environments/
│   ├── test.properties
│   ├── staging.properties
│   └── production.properties
└── test/

動機
スタックのインスタンスのための設定ファイルを作成するのは簡単で理解しやすいです。ソースコードリポジトリにファイルがコミットされているため、以下のことが簡単に行えます。
- 任意の環境で使用される値を確認することができる（「プロダクションの最大クラスタサイズは何ですか？」）
- デバッグのための履歴を追跡することができる（「最大クラスタサイズはいつ変更されましたか？」）
- 変更の監査を行うことができる（「最大クラスタサイズを変更したのは誰ですか？」）

スタック構成ファイルは、構成とスタックのコードの分離を強制します。

適用範囲
スタック構成ファイルは、環境の数が頻繁に変更されない場合に適しています。新しいスタックインスタンスを追加するためには、プロジェクトにファイルを追加する必要があります。これにより、異なるインスタンスの作成と更新の一貫したロジックが必要になります。

結果
新しいスタックインスタンスを作成する場合、スタックプロジェクトに新しい設定ファイルを追加する必要があります。これにより、環境を自動的に作成することができなくなります。「パターン：一時的なテストスタック」では、環境を自動的に作成する方法について説明しています。必要に応じて、一時的な環境のための設定ファイルを作成することでこれを回避することができます。

パラメータファイルは、変更配信パイプラインの下流環境の構成を変更する際に摩擦を引き起こす可能性があります。プロジェクトのコードの変更は、本番環境に適用される前にパイプラインの各ステージを進む必要があります。これには時間がかかる場合があり、構成の変更が本番環境にのみ適用される場合には価値がありません。

プロビジョニングスクリプトでパラメータ値を定義することは、かなり複雑さの源泉になります。これについては「スクリプトを使用してインフラツールをラップする」で詳しく説明しますが、チームはスタックプロジェクトと環境のデフォルト値を定義したいと考えることがよくあり、これらを異なる環境のインスタンスの値に組み合わせるためのロジックが必要です。パラメータ値の継承モデルは、ごちゃごちゃしたり混乱したりすることがあります。

ソースコントロールの構成ファイルには、シークレット情報を含めるべきではありません。シークレット情報の場合は、この章から別のパターンを選択してシークレット情報を扱うか、ソースコントロール以外に別の秘密の設定ファイルを実装する必要があります。

実装
各環境ごとにスタックのパラメータ値を別々のファイルで定義します。前述のExample 7-9のような形式です。パラメータファイルの内容は、次のようになります。

env = staging
cluster_minimum = 2
cluster_maximum = 3

スタックコマンドを実行する際に、関連するパラメータファイルのパスを渡します。例えば、以下のようになります。

stack up --source ./src --config ./environments/staging.properties

システムが複数のスタックから構成されている場合、すべての環境での構成を管理するのは複雑になる可能性があります。その場合、以下の2つの一般的な方法があります。

1つは、各スタックのコードと一緒にすべての環境の構成ファイルを配置する方法です。

├── cluster_stack/
│   ├── src/
│   │   ├── cluster.infra
│   │   ├── host_servers.infra
│   │   └── networking.infra
│   └── environments/
│       ├── test.properties
│       ├── staging.properties
│       └── production.properties
│
└── appserver_stack/
    ├── src/
    │   ├── server.infra
    │   └── networking.infra
    └── environments/
        ├── test.properties
        ├── staging.properties
        └── production.properties

もう1つは、すべてのスタックの構成を一箇所に集約する方法です。

├── cluster_stack/
│   ├── cluster.infra
│   ├── host_servers.infra
│   └── networking.infra
│
├── appserver_stack/
│   ├── server.infra
│   └── networking.infra
│
└── environments/
    ├── test/
    │   ├── cluster.properties
    │   └── appserver.properties
    ├── staging/
    │   ├── cluster.properties
    │   └── appserver.properties
    └── production/
        ├── cluster.properties
        └── appserver.properties

どちらのアプローチもそれぞれの方法で複雑さや混乱を引き起こす可能性があります。一つの環境のすべての設定を変更する必要がある場合、数十ものスタックプロジェクトの構成ファイルを変更することは手間です。さらに、複数の環境にまたがる単一のスタックの構成を変更する場合、数十の他のスタックの構成が含まれるツリーを探索する必要があります。

シークレット情報を提供するために構成ファイルを使用したい場合、別のシークレットのパターンではなく、プロジェクトコードがソースコントロールにチェックインされた別の場所でこれらのファイルを管理する必要があります。ローカルな開発環境では、ユーザーに手動でファイルを指定された場所に作成するよう要求することができます。ファイルの場所を以下のようにしてスタックコマンドに渡します。

stack up --source ./src \
 --config ./environments/staging.properties \
 --config ../.secrets/staging.properties

この例では、2つの--config引数をstackツールに提供し、両方のパラメータ値を読み取ります。プロジェクトフォルダの外側に.secretsというディレクトリがあり、ソースコントロールされていないため、その場所にファイルを格納します。

このような処理を自動的に行う場合、CDパイプラインエージェントのようなコンピュートインスタンスからstackツールを実行するのは少し難しくなります。同様のシークレットプロパティファイルをこれらのコンピュートインスタンスにプロビジョニングすることもできますが、それにより同じエージェントで実行される他のプロセスにシークレットがさらされる可能性があります。エージェントのためのコンピュートインスタンスを構築するプロセスにシークレットを提供する必要もありますので、まだ起動の問題があります。

関連するパターン
ファイルに構成値を配置することは、「パターン：スクリプトパラメータ」で説明されるプロビジョニングスクリプトを簡素化することができます。環境の構成ファイルの制約をいくつか回避するために、スタックパラメータレジストリパターンを使用することもできます（「パターン：スタックパラメータレジストリ」を参照）。これにより、パラメータ値がスタックプロジェクトのコードから分離され、異なるワークフローをコードと構成に使用できるようになります。
