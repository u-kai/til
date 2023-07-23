Pattern: Application Group Stack An application group stack includes the infrastructure for multiple related applications or services.services. The infrastructure for all of these applications is provisioned and managed as a group, as shown in Fig. For example, ShopSpinner’s product application stack includes separate services for browsing products, searching for products, and managing a shopping basket. The servers and other infrastructure for all of these are combined in a single stack instance.
Motivation Defining the infrastructure for multiple related services together can make it easier to manage the application as a single unit. Applicability This pattern can work well when a single team owns the infrastructure and deployment of all of the pieces of the application. An application group stack can align the boundaries of the stack to the team’s responsibilities. Multiservice stacks are sometimes useful as an incremental step from a monolithic stack to service stacks.
Consequences Grouping the infrastructure for multiple applications together also combines the time, risk, and pace of changes. The team needs to manage the risk to the entire stack for every change, even if only one part is changing. This pattern is inefficient if some parts of the stack change more frequently than others. The time to provision, change, and test a stack is based on the entire stack. So again, if it’s common to change only one part of a stack at a time, having it grouped adds unnecessary overhead and risk. Implementation To create an application group stack, you define an infrastructure project that builds all of the infrastructure for a set of services. You can provision and destroy all of the pieces of the application with a single command. Related patterns This pattern risks growing into a monolithic stack (see “Antipattern: Monolithic Stack”). In the other direction, breaking each service in an application group stack into a separate stack creates a service stack.
パターン：アプリケーショングループスタック
アプリケーショングループスタックは、複数の関連するアプリケーションやサービスのためのインフラストラクチャを含みます。これらのアプリケーションのインフラストラクチャは、図に示すように、グループとしてプロビジョニングおよび管理されます。例えば、「ShopSpinner」の製品アプリケーションスタックには、製品の閲覧、製品の検索、ショッピングバスケットの管理など、別々のサービスが含まれています。これらのすべてのサーバーおよび他のインフラストラクチャは、1つのスタックインスタンスに組み合わされます。

動機
関連する複数のサービスのインフラストラクチャをまとめて定義することで、アプリケーションを1つの単位として管理しやすくなる場合があります。

適用性
このパターンは、1つのチームがアプリケーションのすべての部分のインフラストラクチャと展開を所有する場合に効果的です。アプリケーショングループスタックは、スタックの境界をチームの責任に合わせることができます。マルチサービススタックは、モノリシックスタックからサービススタックへの段階的な移行時にも有用です。

結果
複数のアプリケーションのインフラストラクチャをグループ化することで、変更の時間、リスク、およびペースも結合されます。チームは、1つの部分のみが変更されている場合でも、常にスタック全体のリスクを管理する必要があります。一部のスタックの変更が他よりも頻繁に行われる場合、このパターンは効率的ではありません。スタックのプロビジョニング、変更、およびテストにかかる時間は、スタック全体に基づいています。したがって、スタックの一部を一度に変更することが一般的な場合、それをグループ化することは不必要なオーバーヘッドとリスクをもたらします。

実装
アプリケーショングループスタックを作成するには、特定のサービスのためのすべてのインフラストラクチャをビルドするインフラストラクチャプロジェクトを定義します。単一のコマンドでアプリケーションのすべての部分をプロビジョニングおよび破棄できます。

関連するパターン
このパターンは、モノリシックスタックに成長するリスクがある（「アンチパターン：モノリシックスタック」を参照）一方で、アプリケーショングループスタックの各サービスを個別のスタックに分割することで、サービススタックを作成することもできます。