Create Boundaries That Support Scaling A common strategy for scaling systems is to create additional instances of some of its components. You may add instances in periods of higher demand, and you might also consider deploying instances in different geographical regions. Most people are aware that most cloud platforms can automatically scale server clusters (see “Compute Resources”) up and down as load changes. A key benefit of FaaS serverless (see “Infrastructure for FaaS Serverless”) is that it only executes instances of code when needed.
However, other elements of your infrastructure, such as databases, message queues, and storage devices, can become bottlenecks when compute scales up. And different parts of your software system may become bottlenecks, even aside from the infrastructure. For example, the ShopSpinner team can deploy multiple instances of the product browsing service stack to cope with higher load, because most user traffic hits that part of the system at peak times. The team keeps a single instance of its frontend traffic routing stack, and a single instance of the database stack that the application server instances connect to (see Figure 15-7).
Other parts of the system, such as order checkout and customer profile management services, probably don’t need to scale together with the product browsing service. Splitting those services into different stacks helps the team to scale them more quickly. It reduces the waste that replicating everything would create.
Prefer Vertical Groupings Over Horizontal Groupings Traditionally, many architects organized systems functionally. Networking stuff lives together, database stuff lives together, and operating system stuff lives together. This is often a result of organization design, as predicted by Conway’s Law — when teams are organized around these technical, functional specialties, they will divide the infrastructure based on what they manage. The pitfall with this approach is that a service provided to users cuts across many functions. This is often shown as a vertical service crossing horizontal functional layers, as shown in Figure 15-8. Organizing system elements into cross-cutting, functional infrastructure stacks has two drawbacks. One is that a change to the infrastructure for one service may involve making changes to multiple stacks. These changes need to be carefully orchestrated to ensure a dependency isn’t introduced in a consumer stack before it appears in the provider stack (see Providers and Consumers).
Spreading ownership of infrastructure for a single service across multiple functional teams adds considerable communication overhead and process for changes to any one service. The second drawback of functional stacks emerges when they are shared across services. In Figure 15-9, a server team manages a single stack instance with servers for multiple services. When the team changes the server for one of the services, there’s a risk of breaking other services, because a stack boundary represents the blast radius for a change, as described in “Antipattern: Monolithic Stack”.
スケーリングをサポートする境界を作成する 多くのシステムのスケーリング戦略の一つは、そのコンポーネントのいくつかに追加のインスタンスを作成することです。需要の上昇期にはインスタンスを追加することができ、異なる地理的地域にも展開することが考えられます。ほとんどの人は、クラウドプラットフォームのほとんどが、負荷の変化に応じてサーバクラスタ（「コンピュートリソース」を参照）を自動的にスケールアップおよびスケールダウンできることを知っています。FaaSサーバーレス（「FaaSサーバーレスのインフラストラクチャ」を参照）の主な利点は、コードのインスタンスを必要な時にのみ実行することです。

しかし、データベース、メッセージキュー、ストレージデバイスなど、インフラストラクチャの規模が拡大すると、その他のインフラストラクチャ要素もボトルネックになる可能性があります。インフラストラクチャ以外にも、ソフトウェアシステムの異なる部分がボトルネックになる可能性があります。たとえば、ShopSpinnerチームは、システムのピーク時にほとんどのユーザートラフィックがその部分に集中するため、製品閲覧サービススタックの複数のインスタンスを展開することで高い負荷に対応することができます。チームはフロントエンドトラフィックルーティングスタックと、アプリケーションサーバーインスタンスが接続するデータベーススタックの単一のインスタンスを維持しています（図15-7を参照）。

注文チェックアウトや顧客プロファイル管理サービスなど、他のシステムの一部は、おそらく製品閲覧サービスと一緒にスケールする必要はありません。これらのサービスを異なるスタックに分割することで、チームはより迅速にスケーリングすることができます。すべてを複製することによる無駄を減らすことができます。

水平グループよりも垂直グループを優先する 伝統的に、多くのアーキテクトはシステムを機能ごとに組織化してきました。ネットワーキングの要素は一緒になり、データベースの要素は一緒になり、オペレーティングシステムの要素は一緒になります。これは、Conwayの法則によって予想される組織設計の結果であり、これらの技術的な機能の専門分野を中心にチームが組織されるため、インフラストラクチャは彼らが管理するものに基づいて分割されることがしばしばあります。このアプローチの落とし穴は、ユーザーに提供されるサービスが多くの機能にまたがっていることです。これは、図15-8に示されるように、水平機能レイヤーを横切る垂直サービスとして示されることがしばしばあります。システム要素を横断的な機能インフラストラクチャスタックに組織することには2つの欠点があります。1つは、1つのサービスのインフラストラクチャを変更すると、複数のスタックに変更を加える可能性があることです。これらの変更は注意深く調整する必要があり、提供元スタックに依存関係が現れる前に、消費者スタックで依存関係が導入されないようにする必要があります（プロバイダと消費者を参照）。

単一のサービスのインフラストラクチャの所有権を複数の機能チームで分散すると、1つのサービスへの変更に対するコミュニケーションコストとプロセスが大幅に増加することになります。機能スタックの2番目の欠点は、それらが複数のサービス間で共有される場合に発生します。図15-9に示されるように、サーバーチームは複数のサービスに対してサーバを管理する単一のスタックインスタンスを管理します。チームが1つのサービスのサーバを変更すると、他のサービスが壊れるリスクがあります。なぜなら、スタックの境界は変更の爆発半径を表しているからです（「アンチパターン：モノリシックスタック」を参照）。