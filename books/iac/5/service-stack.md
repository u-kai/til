Pattern: Service Stack A service stack manages the infrastructure for each deployable application component in a separate infrastructure stack (see Figure 5-4).
Motivation Service stacks align the boundaries of infrastructure to the software that runs on it. This alignment limits the blast radius for a change to one service, which simplifies the process for scheduling changes. Service teams can own the infrastructure that relates to their software.
Applicability Service stacks can work well with microservice application architectures.3 They also help organizations with autonomous teams to ensure each team owns its infrastructure.4
Consequences If you have multiple applications, each with an infrastructure stack, there could be an unnecessary duplication of code. For example, each stack may include code that specifies how to provision an application server. Duplication can encourage inconsistency, such as using different operating system versions, or different network configurations. You can mitigate this by using modules to share code (as in Chapter 16). Implementation Each application or service has a separate infrastructure code project. When creating a new application, a team might copy code from another application’s infrastructure. Or the team could use a reference project, with boilerplate code for creating new stacks. In some cases, each stack may be complete, not sharing any infrastructure with other application stacks. In other cases, teams may create stacks with infrastructure that supports multiple application stacks. You can learn more about different patterns for this in Chapter 17. Related patterns The service stack pattern falls between an application group stack (“Pattern: Application Group Stack”), which has multiple applications in a single stack, and a micro stack, which breaks the infrastructure for a single application across multiple stacks.
パターン：サービススタック
サービススタックは、デプロイ可能な各アプリケーションコンポーネントのインフラストラクチャを別々のインフラストラクチャスタックで管理します（図5-4参照）。

動機
サービススタックは、インフラストラクチャの境界をそれが実行されるソフトウェアに合わせることができます。この整合性により、変更の爆発範囲を1つのサービスに限定することができ、変更のスケジュール管理を簡素化することができます。サービスチームは、自分たちのソフトウェアに関連するインフラストラクチャを所有することができます。

適用性
サービススタックは、マイクロサービスアプリケーションアーキテクチャとよく組み合わせて使用できます。また、各チームが独自のインフラストラクチャを所有することを確認するために、自律的なチームを持つ組織にも役立ちます。

結果
複数のアプリケーションがあり、各アプリケーションに対してインフラストラクチャスタックがある場合、コードの不必要な重複が発生する可能性があります。例えば、各スタックにはアプリケーションサーバーをプロビジョニングする方法を指定するコードが含まれる場合があります。重複は、異なるオペレーティングシステムバージョンやネットワーク構成の使用など、一貫性の欠如を引き起こす可能性があります。これをモジュールを使用してコードを共有することで緩和することができます（第16章参照）。

実装
各アプリケーションまたはサービスには、別々のインフラストラクチャコードプロジェクトがあります。新しいアプリケーションを作成する場合、チームは別のアプリケーションのインフラストラクチャからコードをコピーする場合があります。または、チームは新しいスタックを作成するためのボイラープレートコードを持つ参照プロジェクトを使用する場合もあります。一部の場合、各スタックは完全であり、他のアプリケーションスタックとインフラストラクチャを共有しません。他の場合では、チームは複数のアプリケーションスタックをサポートするインフラストラクチャを持つスタックを作成する場合があります。これに関するさまざまなパターンについては、第17章で詳しく学ぶことができます。

関連するパターン
サービススタックパターンは、「パターン：アプリケーショングループスタック」（複数のアプリケーションが1つのスタックにある）とマイクロスタック（単一のアプリケーションのインフラストラクチャを複数のスタックに分割する）の中間に位置します。