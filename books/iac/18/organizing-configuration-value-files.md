Chapter 7 described the configuration files pattern for managing parameter values for different instances of a stack project (see “Pattern: Stack Configuration Files”). The description suggested two different ways to organize per-environment configuration files across multiple projects. One is to store them within the relevant project: ├── application_infra_stack/
│ ├── src/
│ └── environments/
│ ├── test.properties
│ ├── staging.properties
│ └── production.properties
│
└── shared_network_stack/
├── src/
└── environments/
├── test.properties
├── staging.properties
└── production.properties The other is to have a separate project with the configuration for all of the stacks, organized by environment: ├── application_infra_stack/
│ └── src/
│
├── shared_network_stack/
│ └── src/
│
└── configuration/
├── test/
│ ├── application_infra.properties
│ └── shared_network.properties
├── staging/
│ ├── application_infra.properties
│ └── shared_network.properties
└── production/
├── application_infra.properties
└── shared_network.properties Storing configuration values with the code for a project mixes generalized, reusable code (assuming it’s a reusable stack, per “Pattern: Reusable Stack”) with details of specific instances. Ideally, changing the configuration for an environment shouldn’t require modifying the stack project. On the other hand, it’s arguably easier to trace and understand configuration values when they’re close to the projects they relate to, rather than mingled in a monolithic configuration project. Team ownership and alignment is a factor, as usual. Separating infrastructure code and its configuration can discourage taking ownership and responsibility across both.
第 7 章では、スタックプロジェクトの異なるインスタンスごとにパラメーター値を管理するための設定ファイルのパターンについて説明しました（「パターン：スタック設定ファイル」を参照）。説明では、複数のプロジェクトにわたる環境ごとの設定ファイルを整理するための 2 つの異なる方法が提案されています。一つは、関連するプロジェクト内にそれらを保存する方法です。 ├── application_infra_stack/
│ ├── src/
│ └── environments/
│ ├── test.properties
│ ├── staging.properties
│ └── production.properties
│
└── shared_network_stack/
├── src/
└── environments/
├── test.properties
├── staging.properties
└── production.properties もう一つは、すべてのスタックの設定を環境別に整理する別のプロジェクトを持つ方法です。 ├── application_infra_stack/
│ └── src/
│
├── shared_network_stack/
│ └── src/
│
└── configuration/
├── test/
│ ├── application_infra.properties
│ └── shared_network.properties
├── staging/
│ ├── application_infra.properties
│ └── shared_network.properties
└── production/
├── application_infra.properties
└── shared_network.properties プロジェクトのコードと一緒に設定値を保存すると、一般化され再利用可能なコード（「パターン：再利用可能なスタック」を仮定）と特定のインスタンスの詳細が混ざってしまいます。理想的には、環境の設定を変更するためには、スタックプロジェクトを修正する必要はありません。一方で、設定値が関連するプロジェクトに近い場所にあると、トレースしやすく理解しやすいと言えるでしょう。一つの巨大な設定プロジェクトに混在させるよりもです。チームの所有権とアライメントも重要な要素です。インフラストラクチャのコードとその設定を分離することは、両方についての所有権と責任を減らす可能性があります。

- 考えまとめ

  - 理想的には、環境の設定を変更するためには、スタックプロジェクトを修正する必要はありません。
  - 一方で、設定値が関連するプロジェクトに近い場所にあると、トレースしやすく理解しやすいと言えるでしょう。
    - 一つの巨大な設定プロジェクトに混在させるよりもです。
  - チームの所有権とアライメントも重要な要素です。インフラストラクチャのコードとその設定を分離することは、両方についての所有権と責任を減らす可能性があります。
