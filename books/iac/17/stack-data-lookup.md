Pattern: Stack Data Lookup Also known as: remote statefile lookup, stack reference lookup, or stack resource lookup. Stack data lookup finds provider resources using data structures maintained by the tool that manages the provider stack (Figure 17-2).
Many stack management tools maintain data structures for each stack instance, which include values exported by the stack code. Examples include Terraform and Pulumi remote state files. Motivation Stack management tool vendors make it easy to use their stack data lookup features to integrate different projects. Most implementations of data sharing across stacks require the provider stack to explicitly declare which resources to publish for use by other stacks. Doing this discourages consumer stacks from creating dependencies on resources without the provider’s knowledge. Applicability Using stack data lookup functionality to discover dependencies across stacks works when all of the infrastructure in a system is managed using the same tool. Consequences Stack data lookup tends to lock you into a single stack management tool. It’s possible to use the pattern with different tools, as described in the implementation of this pattern. But this adds complexity to your implementation. This pattern sometimes breaks across different versions of the same stack tool. An upgrade to the tool may involve changing the stack data structures. This can cause problems when upgrading a provider stack to the newer version of the tool. Until you upgrade the consumer stack to the new version of the tool as well, it may not be possible for the older version of the tool to extract the resource values from the upgraded provider stack. This stops you from rolling out stack tool upgrades incrementally across stacks, potentially forcing a disruptive coordinated upgrade across your estate.
Implementation The implementation of stack data lookup uses functionality from your stack management tool and its definition language. Terraform stores output values in a remote state file. Pulumi also stores resource details in a state file that can be referenced in a consumer stack using a StackReference. CloudFormation can export and import stack output values across stacks, which AWS CDK can also access.2 The provider usually explicitly declares the resources it provides to consumers: stack:
name: shared_network_stack
environment: ${ENVIRONMENT_NAME}

vlans:

-   appserver_vlan
    address_range: 10.2.0.0/8

export:

-   appserver_vlan_id: appserver_vlan.id The consumer declares a reference to the provider stack and uses this to refer to the VLAN identifier exported by that stack: external_stack:
    name: shared_network_stack
    environment: ${ENVIRONMENT_NAME}

virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: external_stack.shared_network_stack.appserver_vlan.id
This example embeds the reference to the external stack in the stack code. Another option is to use dependency injection (“Dependency Injection”) so that your stack code is less coupled to the dependency discovery. Your orchestration script looks up the output value from the provider stack and passes it to your stack code as a parameter. Although stack data lookups are tied to the tool that manages the provider stack, you can usually extract these values in a script, so you can use them with other tools, as shown in Example 17-1. Example 17-1. Using a stack tool to discover a resource ID from a stack instance’s data structure #!/usr/bin/env bash

VLAN_ID=$(
stack value \
 --stack_instance shared_network-staging \
 --export_name appserver_vlan_id
) This code runs the fictional stack command, specifying the stack instance (shared_network-staging) to look in, and the exported variable to read and print (appserver_vlan_id). The shell command stores the command’s output, which is the ID of the VLAN, in a shell variable named VLAN_ID. The script can then use this variable in different ways. Related patterns The main alternative patterns are resource matching (“Pattern: Resource Matching”) and registry lookup.
パターン: スタックデータの検索
以下、リモートステートファイルの参照、スタックリファレンスの検索、またはスタックリソースの検索とも呼ばれるパターンです。スタックデータの検索は、プロバイダースタックを管理するツールによって維持されるデータ構造を使用してプロバイダーリソースを見つけます（図17-2）。
多くのスタック管理ツールは、スタックの各インスタンスごとにデータ構造を維持しており、その中にはスタックコードによってエクスポートされた値も含まれます。例として、TerraformやPulumiのリモートステートファイルがあります。

動機
スタック管理ツールのベンダーは、異なるプロジェクトを統合するためにスタックデータの検索機能を簡単に利用できるようにします。スタック間でのデータ共有のほとんどの実装では、他のスタックで使用するために明示的にリソースを公開するようにプロバイダースタックを宣言する必要があります。これにより、プロバイダーの知識を持たないままリソースに依存関係を作成するコンシューマースタックの作成を防止することができます。

適用範囲
スタックデータの検索機能を使用してスタック間の依存関係を発見する場合、システム内のすべてのインフラストラクチャを同じツールで管理している場合に機能します。

結果
スタックデータの検索は、通常、単一のスタック管理ツールに固定される傾向があります。このパターンは、このパターンの実装で説明されているように、さまざまなツールで使用することも可能ですが、その場合、実装が複雑になります。このパターンは、同じスタックツールの異なるバージョン間で場合によっては動作しないことがあります。ツールのアップグレードでは、スタックデータ構造の変更が必要になる場合があります。これにより、プロバイダースタックをツールの新しいバージョンにアップグレードする際に問題が発生する可能性があります。古いバージョンのツールでは、アップグレードされたプロバイダースタックからリソース値を抽出することができないため、コンシューマースタックもツールの新しいバージョンにアップグレードするまで、スタックツールのアップグレードを段階的に展開することができなくなります。それにより、ドメイン全体での中断を伴う調整されたアップグレードが強制される可能性があります。

実装
スタックデータの検索の実装では、スタック管理ツールとその定義言語の機能を使用します。Terraformは、出力値をリモートステートファイルに格納します。Pulumiも、ステートファイルにリソースの詳細を保存し、StackReferenceを使用してコンシューマースタックで参照できます。CloudFormationはスタック間でスタック出力値をエクスポート・インポートすることができ、AWS CDKもこれにアクセスできます。通常、プロバイダーは明示的にコンシューマーに提供するリソースを宣言します。

スタックコード内に外部スタックへの参照を埋め込む例を示します。別のオプションとして、依存性注入（"依存性注入"）を使用して、スタックコードを依存性の発見との結びつきを弱めることもできます。オーケストレーションスクリプトは、プロバイダースタックから出力値を検索し、それをスタックコードにパラメータとして渡します。スタックデータの検索は、プロバイダースタックを管理するツールに結び付けられていますが、スクリプトでこれらの値を抽出することができるため、他のツールでも使用することができます（例17-1参照）。

関連するパターン
主な代替パターンはリソースマッチング（"パターン: リソースマッチング"）とレジストリの検索です。