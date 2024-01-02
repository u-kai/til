# CNCF: Platform White Paper

- 良い取り組みだが、間接的なものなので上に認めてもらうい、助けてもらうことが大事

## Why platforms?

- ソフトウェアのアジリティがここ数年で上がってきており、様々な関心ごとも出てきた
- 自動やプロセスの工場によってサービス開発者の責務が大きくなったり、インフラの認知負荷に関して考える時間が増えたし、組織に関連するの価値提供の時間が減った？
- 開発者たちは再度自分らのコアについてよく考えたいし、エンタープライズ全体での被りの作業を減らしたい

- platform に投資するとエンタープライズは

1. プロダクトチームは認知負荷を減らすことができるそして開発、デリバリーの速度を早めることができる
1. 設定や管理を専門家に任せることで、プラットフォーム機能に依存する信頼性とレジリエンスを高めることができる
1. プロダクト開発を促進できたり、プラットフォームツールの使い回しができたり、知見を全体に共有できたりする
1. セキュリティリスクを減らせるし規制や機能的な問題も
1. コスト効果も出せる(うまくやればだと思う)

- 一般的な仕事を減らすだけでなく、最適化していく必要がある

## What is a platform

- user experience が一番大事
- プラットフォームは提供される実装全体で一貫性を提供し、組織の要件を満たす最も薄い合理的なレイヤー
- 一番シンプルなプラットフォームは wiki page
- クラウドネイティブなアプローチはサービスのロジックとは別で必要な多くのコンポーネントがあり、それをプラットフォームとして出して、アプリケーションやシステムに簡単に統合できる様にすると良い
  - 認証、オブザーバビリティ、データベース、タスクランナーなどなど

## Platform maturity

## Attributes of platforms

プラットフォームとは何か、そして組織がプラットフォームを構築したい理由を定義した後、プラットフォームの成功に影響を及ぼすいくつかの主要な属性を特定しましょう。

プロダクトとしてのプラットフォーム。プラットフォームは、ユーザーの要件を満たすために存在し、その要件に基づいて設計および進化する必要があります。他のソフトウェアと同様に、プラットフォームは製品チーム全体で最も一般的なユースケースをサポートするための必要な機能を提供し、単一のチームだけが使用する特定の機能よりも優先します。

ユーザーエクスペリエンス。プラットフォームは、一貫したインターフェースを通じて機能を提供し、ユーザーエクスペリエンスに焦点を当てるべきです。プラットフォームは、GUI、API、コマンドラインツール、IDE、ポータルなど、ユーザーが利用する方法に合わせるよう努めるべきです。例えば、プラットフォームは通常、アプリケーションのデプロイ能力を提供します。開発者は IDE を介してそのような機能を利用するかもしれませんし、テスターはコマンドラインツールを使用するかもしれませんが、プロダクトオーナーは GUI ベースの Web ポータルを使用するかもしれません。

ドキュメントとオンボーディング。ドキュメントは成功したソフトウェア製品の重要な要素です。プラットフォームの提供物を利用するためには、ユーザーはドキュメントとサンプルが必要です。プラットフォームは、ユーザーのニーズに応える適切なドキュメントを提供する必要があります。また、新しいプロジェクトのオンボーディングを迅速かつ簡単に行うためのツールも提供するべきです。例えば、プラットフォームは、Kubernetes 上の Web アプリケーションのビルド、スキャン、テスト、デプロイ、監視に関する再利用可能なサプライチェーンワークフローを提供することができます。このようなワークフローは、初期のプロジェクトテンプレートとドキュメントと一緒に提供されることがよくあります。

セルフサービス。プラットフォームはセルフサービスであるべきです。ユーザーは能力を自律的かつ自動的に要求し、受け取ることができるようになる必要があります。この特性により、プラットフォームチームは複数の製品チームを活性化し、必要に応じてスケールすることができます。プラットフォームの機能は、上記で説明したインターフェースを介してオンデマンドで最小限の手動介入で利用できるべきです。例えば、ユーザーはコマンドラインツールを実行したり、ウェブポータルのフォームを記入することで、データベースを要求し、そのロケーターと資格情報を受け取ることができるべきです。

ユーザーの認知負荷の軽減。プラットフォームの主要な目標は、製品チームの認知負荷を軽減することです。プラットフォームは実装の詳細をカプセル化し、アーキテクチャから生じる複雑さを隠すべきです。例えば、プラットフォームは一部のサービスをクラウドプロバイダに委任するかもしれませんが、ユーザーはこのような詳細に触れる必要はありません。同時に、プラットフォームは必要に応じてユーザーが特定のサービスを設定および監視できるようにするべきです。ユーザーはプラットフォームが提供するサービスの運用に責任を持つべきではありません。例えば、ユーザーは頻繁にデータベースを必要とするかもしれませんが、データベースサーバーの管理は行わなくてもよいはずです。

オプションで組み合わせ可能。プラットフォームは製品開発を効率化するために意図されているため、障害になってはいけません。プラットフォームは組み合わせ可能であり、製品チームが提供する要素の一部だけを利用できるようにするべきです。また、必要に応じて、製品チームがプラットフォームの提供物とは別に自分たち自身の機能を提供および管理できるようにするべきです。例えば、プラットフォームがグラフデータベースを提供していないが製品に必要な場合、製品チームが自分たちでグラフデータベースをプロビジョニングおよび運用できるようにすることができるべきです。

デフォルトでセキュア。プラットフォームはデフォルトで安全であり、組織が定義したルールや基準に基づいてコンプライアンスと検証を保証する機能を提供するべきです。

## Attributes of platform teams

最も重要なことは、プラットフォームチームがプラットフォームユーザーの要件について学び、プラットフォームが提供する機能とインターフェースを情報提供し、継続的に改善することです。ユーザーの要件を学ぶ方法には、ユーザーインタビュー、対話型のハッカソン、課題トラッカーと調査、そして観測ツールを通じた利用の直接観察などがあります。たとえば、プラットフォームチームは、ユーザーが機能リクエストを提出するためのフォームを公開したり、今後の機能を共有するためのロードマップミーティングを主催したり、ユーザーの利用パターンを確認して優先順位を設定することができます。

ユーザーからのフィードバックと熟考された設計は、製品の提供の一面です。もう一方の面は、アウトバウンドマーケティングと支持です。もしプラットフォームが本当にユーザーの要件に基づいて構築されていれば、それらのユーザーは提供される機能を使うことに興奮するでしょう。プラットフォームチームがユーザーの採用を促進するいくつかの方法には、広報発表、魅力的なデモ、定期的なフィードバックとコミュニケーションセッションなどの内部マーケティング活動があります。重要なのは、ユーザーがどこにいても迎え、プラットフォームとの関わりで利益を得るための旅に連れて行くことです。

プラットフォームチームは、コンピューティング、ネットワーク、ストレージ、その他のサービスを必ずしも実行する必要はありません。実際、内部のプラットフォームはできるだけ外部の提供サービスと機能に依存すべきです。プラットフォームチームは、管理されたプロバイダーや内部インフラチームから利用できない場合にのみ、独自の機能を構築し、維持するべきです。代わりに、プラットフォームチームが最も責任を負うべきなのは、サービスと機能のインターフェース（つまり GUI、CLI、および API）とユーザーエクスペリエンスです。

たとえば、プラットフォーム内の Web ページには、アプリケーションのアイデンティティを提供するためのボタンが記載されていることもあり、その機能の実装はクラウドホストされたアイデンティティサービスを介して行われる場合もあります。内部のプラットフォームチームは Web ページと API を管理するかもしれませんが、実際のサービスの実装は管理しません。プラットフォームチームは、通常は他の場所で必要な機能が利用できない場合に、独自の機能を作成し、維持することを検討すべきです。

## Challenges with platforms

プラットフォームの課題
プラットフォームは多くの価値を約束していますが、次に挙げるような課題もあり、実装者はこれに留意する必要があります。

プラットフォームチームは、自分たちのプラットフォームを商品のように扱い、ユーザーと一緒に開発する必要があります。
プラットフォームチームは、優先順位と初期のパートナーアプリケーションチームを注意深く選ぶ必要があります。
プラットフォームチームは、企業のリーダーシップの支持を得て、価値ストリームに対するインパクトを示す必要があります。
もっとも重要なのは、プラットフォームを顧客向けの製品として扱い、その成功が直接ユーザーや製品の成功に依存することを認識することです。そのためには、プラットフォームチームがアプリチームや他のユーザーと協力し、プラットフォームの機能やユーザーエクスペリエンスを優先し、計画し、実装し、繰り返す必要があります。フィードバックなしで機能やエクスペリエンスをリリースしたり、トップダウンの命令に頼ることで採用を達成する場合、ユーザーからは抵抗や憤りを受け、約束された価値の多くを見逃すことになります。これを軽減するために、プラットフォームチームはプロダクトマネージャーを初めから含め、ロードマップを共有し、フィードバックを集め、プラットフォームユーザーのニーズを理解し、代表する必要があります。

プラットフォームを採用する際、最初に有効化する適切な機能やエクスペリエンスを選ぶことは重要です。パイプライン、データベース、オブザーバビリティなどのように、頻繁に必要で差別化されていない機能は、始める良い出発点となるかもしれません。また、プラットフォームチームは、限られた数の熱心で熟練したアプリチームに最初に焦点を当てることも選ぶことができます。このようなチームからの詳細なフィードバックは、最初のプラットフォーム体験を向上させます。また、これらのチームのメンバーは、後の採用者に対してプラットフォームを推進し、宣伝する役割も果たします。

最後に、大企業においては、プラットフォームチームへのリーダーシップの支持を早急に得ることが重要です。多くの企業リーダーは、IT インフラを一次的な価値ストリームとは無関係な費用と見なし、IT プラットフォームに割り当てられるコストやリソースを制約しようとする場合があり、それにより実装が悪化し、約束された成果が得られず、憤りが生じます。これを緩和するために、プラットフォームチームは直接的なインパクトと製品および価値ストリームチームとの関係を示す必要があります（前の 2 段落を参照）。プラットフォームチームを製品チームの戦略的パートナーとして位置付け、顧客に価値を提供することにおいてチームチームをサポートする必要があります。

## Enabling platform teams

## How to measure the success of platforms

- とにかくプラットフォームを製品として扱え
- そしてのその製品のパフォーマンスをしっかり測ろう
- 社内ユーザーからフィードバックやアクティビティをしっかり測ることが大事
- 以下のメトリクスをしっかり取ることが重要

### User satisfaction and productivity

- ユーザーの満足度と、活動の向上を測るのが良い
  - どのくらいのユーザーが利用しているのか

### Organizational efficiency

##

Capability Description Example CNCF/CDF Projects
Web portals for provisioning and observing capabilities Publish documentation, service catalogs, and project templates. Publish telemetry about systems and capabilities. Backstage, Skooner, Ortelius
APIs for automatically provisioning capabilities Structured formats for automatically creating, updating, deleting and observing capabilities. Kubernetes, Crossplane, Operator Framework, Helm, KubeVela
Golden path templates and docs Templated compositions of well-integrated code and capabilities for rapid project development. ArtifactHub
Automation for building and testing products Automate build and test of digital products and services. Tekton, Jenkins, Buildpacks, ko, Carvel
Automation for delivering and verifying services Automate and observe delivery of services. Argo, Flux, Keptn, Flagger, OpenFeature
Development environments Enable research and development of applications and systems. Devfile, Nocalhost, Telepresence, DevSpace
Application observability Instrument applications, gather and analyze telemetry and publish info to stakeholders. OpenTelemetry, Jaeger, Prometheus, Thanos, Fluentd, Grafana, OpenCost
Infrastructure services Run application code, connect application components and persist data for applications Kubernetes, Kubevirt, Knative, WasmEdge, KEDA
CNI, Istio, Cilium, Envoy, Linkerd, CoreDNS
Rook, Longhorn, Etcd
Data services Persist structured data for applications TiKV, Vitess, SchemaHero
Messaging and event services Enable applications to communicate with each other asynchronously Strimzi, NATS, gRPC, Knative, Dapr
Identity and secret services Ensure workloads have locators and secrets to use resources and capabilities. Enable services to identify themselves to other services Dex, External Secrets, SPIFFE/SPIRE, Teller, cert-manager
Security services Observe runtime behavior and report/remediate anomalies. Verify builds and artifacts don't contain vulnerabilities. Constrain activities on the platform per enterprise requirements; notify and/or remediate aberrations Falco, In-toto, KubeArmor, OPA, Kyverno, Cloud Custodian
Artifact storage Store, publish and secure built artifacts for use in production. Cache and analyze third-party artifacts. Store source code. ArtifactHub, Harbor, Distribution, Porter
