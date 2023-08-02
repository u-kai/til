The People A reliable automated IT system is like Soylent Green — its secret ingredient is people.2 While human hands shouldn’t be needed to get a code change through to production systems, other than perhaps reviewing test results and clicking a few buttons, people are needed to continuously build, fix, adapt, and improve the system. There are a handful of roles involved with most infrastructure systems, automated or otherwise. These roles don’t often map one to one to individuals — some people play more than one role, and some people share roles with others:
Users Who directly uses the infrastructure? In many organizations, application teams do. These teams may develop the applications, or they may configure and manage third-party applications. Governance specialists Many people set policies for the environment across various domains, including security, legal compliance, architecture, performance, cost control, and correctness. Designers People who design the infrastructure. In some organizations, these people are architects, perhaps divided into different domains, like networking or storage. Toolmakers People who provide services, tools, and components that other teams use to build or run environments. Examples include a monitoring team or developers who create reusable infrastructure code libraries. Builders People who build and change infrastructure. They could do this manually through consoles or other interfaces, by running scripts, or by running tools that apply infrastructure code. Testers People who validate infrastructure. This role isn’t limited to QAs (quality analysts). It includes people who test or review systems for a governance domain like security or performance. Support People who make sure the system continues to run correctly and fix it when it doesn’t. Figure 20-1 shows a classic structure, with a dedicated team for each part of the workflow for changing a system. Many roles may be divided across different infrastructure domains, such as networking, storage, or servers. They are also potentially split across governance domains like security, compliance, architecture,
architecture, and performance. Many larger organizations create baroque organizational structures of micro-specialties.3
However, it’s also common for a person or team to work across these roles. For example, an infosec (information security) team might set standards, provide scanning tools, and conduct security audits. A bit later in this chapter we’ll look at ways to reshuffle responsibilities (see “Reshuffling Responsibilities”).

信頼性のある自動化された IT システムは、ソイレント・グリーンのようなもので、その秘密の要素は人々です。コードの変更をプロダクションシステムに反映させるためには、おそらくテスト結果のレビューと数つのボタンのクリック以外に人々は必要ありませんが、システムを継続的に構築し、修正し、適応し、改善するためには人々が必要です。自動化されたシステムやそれ以外のインフラストラクチャシステムには、いくつかの役割が関与しています。これらの役割は個人と一対一で対応することはめったにありません-1 つの役割を複数の人々が担当することもあり、また、一部の人々が他の人々と役割を共有することもあります。

ユーザー
どのような人々がインフラストラクチャを直接使用しますか？多くの組織では、アプリケーションチームが使用します。これらのチームはアプリケーションを開発するか、サードパーティのアプリケーションを構成および管理するかのいずれかです。

ガバナンスの専門家
セキュリティ、法的準拠、アーキテクチャ、パフォーマンス、コスト管理、正確さなど、さまざまなドメインにおける環境のポリシーを設定する多くの人々。

デザイナー
インフラストラクチャを設計する人々。一部の組織では、これらの人々はネットワーキングやストレージなどの異なるドメインに分けられているかもしれません。

ツール作成者
他のチームが環境の構築や実行に使用するサービス、ツール、コンポーネントを提供する人々。モニタリングチームや再利用可能なインフラストラクチャのコードライブラリを作成する開発者などが含まれます。

ビルダー
インフラストラクチャを構築し変更する人々。これは、コンソールや他のインターフェースを介して手動で行うか、スクリプトを実行するか、インフラストラクチャコードを適用するツールを実行することによって行うことができます。

テスター
インフラストラクチャを検証する人々。この役割は QA（品質分析者）に限定されるわけではありません。セキュリティやパフォーマンスなどのガバナンスドメインのシステムをテストまたはレビューする人々も含まれます。

サポート
システムが正常に動作し続け、動作しないときに修正する人々。

図 20-1 は、システムの変更のワークフローごとに専任チームが存在するクラシックな構造を示しています。多くの役割は、ネットワーキング、ストレージ、サーバーなどのインフラストラクチャドメインで分けられる場合もあります。これらの役割は、セキュリティ、コンプライアンス、アーキテクチャ、パフォーマンスなどのガバナンスドメインでも分割される可能性があります。多くの大規模な組織では、マイクロスペシャリストの複雑な組織構造が作られることもあります。

ただし、個人やチームがこれらの役割を横断して働くことも一般的です。例えば、情報セキュリティ（infosec）チームは基準を設定し、スキャンツールを提供し、セキュリティ監査を実施するかもしれません。この章の後半では、責任の再配置方法について見ていきます（「責任の再配置」を参照）。

- 自動化にもそれを支える人がいるってことね

  - まさにこれ
  - でも忙しいところほど，人がいなくて忙しいから自動化もできないみたいなジレンマになっている気がしている

- めっちゃ役割があるな
  - こんなに役割や人が必要ってなるとどうしても人が足りなくなる
  - でも別に被ってても良さそうだけどな
  - って思ってたらやっぱりそうなのね
  - 責任の再配置か
