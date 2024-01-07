Standing at an inflection point is a magical experience. It’s where we look at what has come before, learn from it, and choose a new path. It’s a point where we have a choice to turn to a new direction, with an eye on a different destination. This chapter introduces the destination and the outcomes to expect when choosing data mesh at your organization’s inflection point. Data mesh assumes the environmental conditions I introduced in the previous chapter as a default state. By default, data mesh assumes the ubiquitous nature of data. Data can be of any origin; it can come from any system within an organization, or outside, and across boundaries of organizational trust. Any underlying platform can serve it on one cloud hosting service or another. Data mesh assumes the diversity of data use cases and their unique modes of access to data. The data use cases range from historical data analysis and reporting to training machine learning models and data-intensive applications. And lastly, data mesh assumes complexity of the business landscape—continuous growth, change, and diversity—as a natural state of being. Data mesh learns from the past solutions and addresses their shortcomings. It reduces points of centralization that act as coordination bottlenecks. It finds a new way of decomposing the data architecture without slowing the organization down with synchronizations. It removes the gap between where the data originates and where it gets used and removes the accidental complexities—aka pipelines—that happen in between the two planes of data. Data mesh departs from data myths such as a single source of truth, or one tightly controlled canonical data model. Ultimately, data mesh’s goal is to enable organizations to get value from data at scale, using data to not only improve and optimize their business but also reshape it. Data mesh outcomes can be summarized as (Figure 7-1): Respond gracefully to change: a business’s essential complexity, volatility, and uncertainty Sustain agility in the face of growth Increase the ratio of value from data to the investment Figure 7-1. Data mesh outcomes for organizations In this chapter, I describe each of these outcomes and how data mesh principles accomplish them. Respond Gracefully to Change in a Complex Business Businesses are complex systems, composed of many domains that each have their own accountability structure and goals; they each change at a different pace. The behavior of the business as a whole is the result of an intricate network of relationships between its domains and functions and their interactions and dependencies. The volatility and rapid change of the markets and regulations within which the businesses operate compound their complexity. How can businesses manage the impact of such complexity on their data? How can organizations keep going through change while continuing to get value from their data? How can businesses avoid the increased cost of managing the change of their data landscape? How can they provide truthful and trustworthy data without disruption, in the face of continuous change? This comes down to embracing change in a complex organization. Let’s look at a few ways data mesh achieves embracing change despite the increased complexity of the business. Align Business, Tech, and Now Analytical Data One way to manage complexity is to break it down into independently managed parts. Businesses do so by creating domains. For example, Daff breaks down its business domains according to relatively independent outcomes and functions—managing podcasts, managing artists, player applications, playlists, payments, marketing, etc. This allows each domain to move fast without tight synchronization dependencies on other parts of the business. Just as a business divides its work through business domains, technology can, and should, align itself to these business divisions. Modern digital businesses orient their technology staff around their business units, allowing each business unit to be supported, enabled, and shaped by dedicated digital products and services, built and maintained by a long-standing dedicated technology team. The recent movement toward microservices is largely about performing this kind of decomposition. Business units control and manage their operational applications and data, supported by their partnering technology team. The first principle of data mesh carries out the same decomposition for analytical data, resulting in the domain ownership of data.1 Each business unit takes on the responsibility for analytic data ownership and management. This is because the people closest to the data are best able to understand what analytical data exists and how it should best be interpreted. Domain ownership results in a distributed data architecture, where the data artifacts—datasets, code, metadata, and data policies—are maintained by their corresponding domains. Figure 7-2 shows the concept of business, technology, and data alignment applied to Daff. Each domain has a business function and goal, enabled and shaped by a set of technology solutions—applications and services—and empowered by data and analytics. Domains have dependencies through explicitly defined data and service contracts. Figure 7-2. Aligning business, tech, and data to manage complexity Close the Gap Between Analytical and Operational Data To make good decisions in the moment, analytical data must reflect business truthfulness. It must be as close as possible to the facts and reality of the business at the moment the decision is made. This is hard to achieve with two separate data planes—analytical and operational—that are far from each other and connected through fragile data pipelines and intermediary data teams. Data pipelines must dissolve and give way to a new way of providing the analytical data and capabilities as close to the source as possible. How can changes in the business, such as adding new features to a product, introducing new services, or modifying a business process be reflected in near real time in the analytical data? Data mesh suggests closing the gap and feedback loop between the two planes, through data shared as a product and oriented around the domains. Data mesh connects the two planes under a new structure—a network of peer-to-peer connected data products and applications, a mesh that exchanges analytical data. The data mesh principle of data as a product introduces a new accountability for each domain to share their analytical data as a product, with the goal of delighting the experience of data users by streamlining their experience in discovering, understanding, trusting, and ultimately using quality data. The data as a product principle is designed to address data quality and the age-old siloed data problem and unhappy data users.2 Figure 7-3 shows data mesh’s approach to integrating operational and analytical planes with tighter and faster feedback loops. The concept of centralized pipelines across the two planes is eliminated. Here, the planes are divided by business domains. The integration between data products, the analytical data plane, and their corresponding domain’s operational plane services are rather simple and unintelligent and a matter of simple movement of data. Data products will embed and abstract the intelligence and code required to transform the operational data into its analytical form. By embedding machine intelligent decisions and actions into modern systems through embedding analytics in digital experiences, the boundaries between the analytical and operational planes are dissolving. Data mesh continues to respect the fundamental technical differences between operational data and analytical data, while it closes the gap and tightly integrates the two as demonstrated in this section. Figure 7-3. Closing the gap between operational and analytical data Localize Data Changes to Business Domains Data mesh allows for data models to change continuously without fatal impact to downstream data users or slowing down access to data. It does so by removing the shared global canonical data model and hence removing the need for synchronizing changes. Data mesh localizes change to domains and gives them autonomy to model their data based on their most intimate understanding of the business without the need for central coordination of a single shared canonical model. Data mesh imposes well-defined and guaranteed data contracts to share domain data. Domain data contracts support older revisions until they gracefully migrate their data users to the new revision. This liberates domains to change their data models continuously. Reduce Accidental Complexity of Pipelines and Copying Data As Fred Brooks laid out in his widely popular paper, “No Silver Bullet—Essence and Accident in Software Engineering”, there are two types of complexity when building software systems. First, we have the essential complexity that is essential and inherent to the problem space. This is business and domain complexity. Second, there is accidental complexity: the complexity that we—engineers, architects, and designers—create in our solutions. Accidental complexity can and should be reduced. The world of analytical solutions is full of opportunities to remove accidental complexities. Let’s talk about a few of the accidental complexities that data mesh reduces. Today, we keep copying data around because we need the data for yet another mode of access, or yet another model of computation. We copy data from operational systems to a landing zone and then to the data lake and after that to feature stores for data scientists. We copy the data again from the lake into lakeshore marts for data analyst access and then into the downstream dashboard or reporting databases for the last mile. We build complex and brittle pipelines to do the copying. The copying journey continues from one technology stack to another and from one cloud vendor to another. Today, to run analytical workloads you need to decide up front which cloud provider copies all of your data in its lake or warehouse before you can get value from it. Data mesh addresses this problem by creating a new architectural unit that encapsulates a domain-oriented data semantic while also providing multiple modes of access to the data suitable for different use cases and users. This architectural unit is called the data product quantum (data quantum for short). A data quantum has an explicit set of contracts and guarantees for each of its native access modes—SQL, files, events, etc. It can be accessed anywhere across the internet, in case it chooses to provide data to external data users. It provides access control and policy enforcement on each of its interfaces at the time of access. A data quantum encapsulates the code that transforms and maintains its data. Data pipelines break down and become internal implementations of data quantum logic. A data quantum shares data without the need for intermediary pipelines. Removing complex, brittle, and labyrinth pipelines reduces the opportunity for failure in case of an upstream data change. Sustain Agility in the Face of Growth Today, the success of businesses is predicated on their multifaceted growth—new acquisitions, new service lines, new products, geolocation expansions, and so on. All this leads to new sources of data to manage and new data-driven use cases to build. Many organizations slow down or plateau in the speed of delivering value from their data, onboarding new data, or serving the use cases as they grow. Data mesh’s approach to sustaining agility in the face of growth can be summarized in a few techniques that aim to reduce organization-wide bottlenecks, coordination, and synchronization. Agility relies on business domains’ ability to achieve outcomes autonomously with minimal dependencies. Remove Centralized and Monolithic Bottlenecks A centralized data team managing a monolithic data lake or warehouse limits agility, particularly as the number of sources to onboard or number of use cases to serve grow. Data mesh looks carefully for centralized bottlenecks, particularly where they are the focal point of a multiparty synchronization, from both the architecture and human communication perspective. Architecturally, these bottlenecks include data lakes and data warehouses. Data mesh proposes an alternative, a peer-to-peer approach in data collaboration when serving and consuming data. The architecture enables consumers to directly discover and use data from the source data products. For example, an ML training function or a report can directly access independent data products, without the intervention of a centralized architectural component such as a lake or warehouse and without the need for an intermediary data (pipeline) team. Figure 7-4 demonstrates the conceptual shift. Each data product provides versioned interfaces that allow peer-to-peer consumption of data. The data from multiple data products can be composed and aggregated into new higher-order data products. Figure 7-4. Data mesh removes centralized architecture bottlenecks Reduce Coordination of Data Pipelines In recent decades, technologies that have exceeded their operational scale have one thing in common: they have minimized coordination and synchronization. Asynchronous I/O has scaled the throughput of networked applications over blocking I/O. Reactive applications have resulted in faster parallel processing of messages. MapReduce functional programming has distributed large-volume data processing across many servers. Choreographed event-driven microservices have scaled business workflows. Despite the relentless effort to remove coordination and synchronization in core technologies, we have, for the most part, neglected organizational and architectural coordination. As a result, no matter how fast our computer systems run, achieving outcomes have fallen behind coordinating activities of teams and humans. Data mesh reduces architectural and human coordination. Existing architectures build on the technical decomposition of components—i.e., pipeline tasks such as ingestion, processing, serving, etc. This style of architectural decomposition results in heavy coordination between these functions each time a new data source or a new use case is delivered. Data mesh moves away from technical partitioning of data management to domain-oriented partitioning. Domain-oriented data products develop and evolve independently of other data products. This domain-oriented decomposition reduces the need for coordination to achieve an outcome. For the most part, the domain-oriented data product team can take care of the new data sources for their new use cases. In cases where a new use case requires access to a new data product outside of the domain, the consumer can make progress by utilizing the standard contracts of the new data product, mocks, stubs, or synthetic data3 interfaces, until the data product becomes available. This is the beauty of contracts, as they ease the coordination between consumer and provider during development. Figure 7-5 shows the shift in reducing pipeline coordination. Figure 7-5. Reduce architectural coordination of pipelines Reduce Coordination of Data Governance Another major coordination bottleneck is the central function of data governance. Today, data governance coordination is necessary to permit access to data, approve the quality of data, and validate the conformance of data changes with the organization’s policies. The central and heavily manual processes of data governance inhibit agility in data sharing. Data mesh reduces governance coordination friction through two functions: Automating and embedding policies as code in each data product Delegating central responsibilities of governance to individual domain data product owners These changes are implemented by data mesh’s federated and computational data governance model.4 Operationally, the governance team is composed of the individual domain data product owners—the long-term product owners responsible for domain data sharing. Architecturally, the governance function embeds policy execution into every data product in a computational and automated fashion. This vastly improves the function of governance today, which is one of the main synchronization points for discovering data, approving data, and making sure it follows the necessary policies. As you can imagine, the autonomy of the domains can have undesirable consequences if not checked: isolation of domains, incompatibility and disconnection of one domain’s data product from others, and a fragmented experience when consuming multiple domains’ data. Data mesh governance heavily relies on the automation of governance concerns for a consistent, connected, and trustworthy experience using the domains’ data products. Figure 7-6 shows the replacement of manual and central governance functions with automated delivery of data products with policies embedded as code. Figure 7-6. Reduce synchronization of data governance Enable Autonomy The correlation between team autonomy and team performance has been the subject of team management studies. Empirical studies show that teams’ freedom in decision making to fulfill their mission can lead to better team performance. On the other hand, too much autonomy can result in inconsistencies, duplicated efforts, and team isolation. Data mesh attempts to strike a balance between team autonomy and interteam interoperability and collaboration. It gives domain teams autonomy to take control of their local decision making, for example, choosing the best data model for their data products, while it uses computational governance policies to impose a consistent experience across all data products, for example, standardizing the data modeling language that all domains utilize. Data mesh gives domain teams autonomy to build and maintain their data products, while it places a domain-agnostic data platform in place for teams to do so in a consistent and cost-effective way. The principle of a self-serve data platform essentially makes it feasible for domain teams to manage the life cycle of their data products with autonomy and utilize the skillsets of their generalist developer to do so.5 The self-serve data infrastructure allows data product developers to build, deploy, monitor, and maintain their data products. It allows data consumers to discover, learn, access, and use the data products. The self-serve infrastructure makes it possible for the mesh of data products to be joined, correlated, and used as a whole, while maintaining the independence of the domain teams. Increase the Ratio of Value from Data to Investment Industry reports, such as the NewVantage Partners report I shared in the previous chapter, and my personal experience, point to the fact that we are getting little value from data compared to our investments in data management. If we compare the value we get from our data teams and data solutions, compared to other technical investments such as app development infrastructure, it’s evident that we are facing headwinds when it comes to data. Data mesh looks at ways to improve the ratio of value over effort in analytical data management: the creation of a new archetype of data platform that abstracts today’s technical complexity through open data interfaces that enable sharing data across organizational trust boundaries or physical locations and through applying product thinking to remove friction from the experience of data users. Abstract Technical Complexity with a Data Platform Today’s landscape of data management technology is undoubtedly too complex. The litmus test for technical complexity is the ever-growing need for data engineers and data experts. We don’t seem to ever have enough of them. Another litmus test is the low value to effort ratio of data pipeline projects. Much effort is spent with little value returned—i.e., getting frictionless access to quality datasets. Data mesh looks critically at the existing technology landscape and reimagines the technology solutions as a data-product-developer (or user)-centric platform. It intends to remove the need for data specialists and enable generalist experts to develop data products. Additionally, data mesh defines a set of open and standard interfaces for different affordances that all data products share—discovering, requesting access, querying, serving data, securing data, etc.—to enable a more collaborative ecosystem of technologies. This is to reduce the cost of integration across vendors.6 Embed Product Thinking Everywhere Data mesh introduces a few shifts to get us laser focused on value, as perceived by the data users. It shifts our thinking from data as an asset to data as a product. It shifts how we measure success from data volume to data user happiness. Data is not the only component of a data mesh ecosystem that is treated as a product. The self-serve data platform itself is also a product. In this case, it serves the data product developers and data product users. Data mesh shifts the measure of success of the platform from the number of capabilities to the impact of its capabilities on improving the experience of data product development and reducing the lead time to deliver, or discover and use, a data product. Product thinking leads to reduced effort and cost, hidden in the everyday experience of data product users and developers. Go Beyond the Boundaries Improvement to a business function almost always requires insights beyond the unit’s boundary. It needs data from many different business domains. Similarly, the data-driven value that an organization generates in serving its customers, employees, and partners requires data beyond what it generates and controls. Consider Daff. In order to provide a better experience to the listeners with auto-play music, it not only requires data from listeners’ playlists, but also their network of friends and their social and environmental influences and behaviors. It requires data from many corners of Daff and beyond—including news, weather, social platforms, etc. Multidomain and multiorg access to data is an assumption built into data mesh. Data mesh’s data quantum concept can provide access to data no matter where the data physically resides. A data quantum provides a set of interfaces that essentially allow anyone with the proper access control to discover and use the data product independent of its physical location. The identification schema, access control, and other policy enforcement assumes using open protocols that are enabled over the internet. Data mesh architecture delivers more value by connecting data beyond organizational boundaries. Recap After reading this chapter you might assume that data mesh is a silver bullet. Quite the contrary. Data mesh is an important piece of the solution. It enables us to truly democratize access to data. However, to close the loop of deriving value from data, there is much more to be done beyond sharing data. We need to continuously deliver repeatable and production-quality analytical and ML-based solutions. But, to bootstrap, we need data sharing at scale, and that is what data mesh focuses on. The data mesh goals listed in this chapter invite us to reimagine data, specifically how to design solutions to manage it, how to govern it, and how to structure our teams. In this chapter, I linked data mesh goals to their enablers. That was a lot to cover, so allow me to summarize it in Table 7-1. Table 7-1. Summary of after the inflection point with data mesh Data mesh goal What to do How to do it Manage changes to data gracefully in a complex, volatile, and uncertain business environment Align business, tech, and data Create cross-functional business, tech, and data teams each responsible for long-term ownership of their data
Principle of domain data ownership Close the gap between the operational and analytical data planes Remove organization-wide pipelines and the two-plane data architecture
Integrate applications and data products more closely through dumb pipes
Principle of data as a product Localize data changes to business domains Localize maintenance and ownership of data products in their specific domains
Create clear contracts between domain-oriented data products to reduce impact of change
Principle of data as a product Manage changes to data gracefully in a complex, volatile, and uncertain business environment (continued) Reduce the accidental complexity of pipelines and copying of data Breakdown pipelines, move the necessary transformation logic into the corresponding data products, and abstract them as an internal implementation
Principle of data as a product
Data product quantum architectural component Sustain agility in the face of growth Remove centralized architectural bottlenecks Remove centralized data warehouses and data lakes
Enable peer-to-peer data sharing of data products through their data interfaces
Principle of domain ownership
Principle of data as a product Reduce the coordination of data pipelines Move from a top-level functional decomposition of pipeline architecture to a domain-oriented decomposition of architecture
Introduce explicit data contracts between domain-oriented data products.
Principle of domain ownership
Principle of data as a product Reduce coordination of data governance Delegate governance responsibilities to autonomous domains and their data product owners
Automate governance policies as code embedded and verified by each data product quantum
Principle of federated computational governance Enable team autonomy Give domain teams autonomy in moving fast independently.
Balance team autonomy with computational standards to create interoperability and a globally consistent experience of the mesh.
Provide domain-agnostic infrastructure capabilities in a self-serve manner to give domain teams autonomy.
Principle of federated computational governance
Principle of the self-serve data platform Increase value from data over cost Abstract complexity with a data platform Create a data-developer-centric and a data-user-centric infrastructure to remove friction and hidden costs in data development and use journeys
Define open and standard interfaces for data products to reduce vendor integration complexity
Principle of data as a product
Principle of the self-serve data platform Embed product thinking everywhere Focus and measure success based on data user and developer happiness
Treat both data and the data platform as a product
Principle of the self-serve data platform
Principle of data as a product Go beyond the boundaries of an organization Share data across physical and logical boundaries of platforms and organizations with standard and internet-based data sharing contracts across data products
Principle of data as a product
Principle of the self-serve data platform In the next chapter, I will give an overview of what has happened before the inflection point: why the data management approach that got us here won’t take us to the future.

転換点に立つことは、魔法のような経験です
そこでは、これまでの経験を振り返り、そこから学び、新しい道を選択します
それは、異なる目的地を目指して新しい方向に進む選択肢があるポイントです
この章では、データメッシュを組織の転換点で選択した場合の目的地と予想される結果を紹介します
データメッシュでは、前章で紹介した環境条件をデフォルトの状態として仮定しています
デフォルトでは、データメッシュはデータの普遍的な性質を前提としています
データはどのような起源であっても構いません
組織内外、組織の信頼の境界を越えてどのシステムからでもデータが得られます
アンダーライイングプラットフォームは、いずれかのクラウドホスティングサービスでそれを提供します
データメッシュでは、データの使用事例の多様性とデータへのユニークなアクセス方法を前提としています
データの使用事例は、歴史的なデータ分析やレポート作成から、機械学習モデルのトレーニングやデータ集約型アプリケーションまでさまざまです
そして、データメッシュは、ビジネスの景観の複雑さ、持続的な成長、変化、多様性を自然な状態として前提としています
データメッシュは、過去の解決策から学び、それらの欠点に対処します
それは、調整のボトルネックとなる中央集権ポイントを削減します
同期化によって組織の進行を遅延させることなく、データアーキテクチャを新しい方法で分解します
データの起源と使用先の間に生じる偶発的な複雑さ、すなわちパイプラインを除去します
データメッシュは、真実の単一の情報源や厳格に制御された典型的なデータモデルなどといったデータの神話から逸脱します
最終的に、データメッシュの目標は、組織がスケールに応じたデータの価値を得ることであり、データを利用してビジネスを改善・最適化するだけでなく、再構築することです
データメッシュの結果は以下のようにまとめることができます（図7-1）： 変化にうまく対応する：ビジネスの本質的な複雑さ、不安定さ、不確実性 成長に対するアジリティの維持 データからの価値の投資への比率の向上 図7-1
組織におけるデータメッシュの結果 この章では、これらの結果それぞれとデータメッシュの原則がそれらを達成する方法について説明します
複雑なビジネスにおいて変化にうまく対応する ビジネスは複雑なシステムであり、それぞれが独自の責任構造と目標を持つ多くのドメインで構成されています
それぞれのドメインは異なるペースで変化します
ビジネス全体の振る舞いは、そのドメインと機能の複雑なネットワークの関係、相互作用、依存関係の結果です
ビジネスが運営される市場と規制の流動性と急速な変化は、その複雑さを複雑にします
企業はどのようにしてそのような複雑さの影響をデータに対して管理することができるのでしょうか？組織はどのようにして変化を続けながらデータから価値を得続けることができるのでしょうか？企業はデータランドスケープの変化を管理するためのコスト増加を避けることができるのでしょうか？継続的な変化に直面して、データに支障をきたさず真実かつ信頼性のあるデータを提供することはどのように可能になるのでしょうか？これは、複雑な組織において変化を受け入れることにつながります
ビジネス、テクノロジー、そして現在の分析データを調整する ポイントをうまく管理する方法の1つは、それを独立して管理される部分に分割することです
ビジネスは、ドメインを作成することでそれを実現しています
例えば、Daffは、ポッドキャストの管理、アーティストの管理、プレイヤーアプリケーション、プレイリスト、支払い、マーケティングなど、比較的独立した成果物や機能に基づいてビジネスドメインを分割しています
これにより、各ドメインは他のビジネスの部分との緊密な同期の依存関係なしに迅速に進むことができます
ビジネスがビジネスドメインを通じて作業を分割するように、テクノロジーもまた、これらのビジネス分野に合わせて調整することができます
現代のデジタルビジネスは、各ビジネスユニットをサポートし、担当のデジタル製品やサービスによって形成される専門のテクノロジーチームによってサポートされるように、テクノロジースタッフをビジネスユニットの周りに配置しています
最近のマイクロサービスへの移行は、このような分解を行うためのものです
ビジネスユニットは、オペレーションアプリケーションとデータを制御し管理し、パートナーとなるテクノロジーチームによってサポートされています
データメッシュの第1の原則は、分析データに対して同様の分解を行い、データのドメイン所有権を実施することです
各ビジネスユニットは、分析データの所有権と管理の責任を負います
これは、データに最も近い人々が分析データの存在や最適な解釈方法を最もよく理解できるためです
ドメイン所有権により、データアーティファクト（データセット、コード、メタデータ、データポリシー）は、それに対応するドメインによって維持される分散型のデータアーキテクチャになります
図7-2は、ビジネス、テクノロジー、データの整合性をDaffに適用した概念を示しています
各ドメインにはビジネス機能と目標があり、技術ソリューション（アプリケーションやサービス）によってサポートされ、形成され、データと分析によって強化されます
ドメイン間には、明示的に定義されたデータとサービスの契約を通じて依存関係があります
図7-2
業務、テクノロジー、データを整合させて複雑さを管理するためのアプローチを示しています
分析データとオペレーションデータの差を埋める 分析データが正確であるためには、ビジネスの真実にできるだけ近づける必要があります
起こりうる変更を適切に分析データに反映させるにはどうすればよいのでしょうか？データメッシュは、分析データとオペレーションデータの間のギャップとフィードバックループを狭めるために、ドメインを中心にデータ共有するというアプローチを提案しています
データメッシュは、分析データを交換するメッシュである、ピアツーピア接続されたデータ製品とアプリケーションのネットワークとして、オペレーションデータと分析データの統合を行います
データ製品の原則であるデータを製品として共有することにより、各ドメインは分析データの品質と利用者の体験を向上させることを目指します
データ製品の原則は、データの品質とデータの孤立化の問題、そして不満なデータ利用者に対処するために設計されています
図7-3は、データメッシュがオペレーションデータと分析データを緊密で迅速なフィードバックループで統合するアプローチを示しています
2つのデータプレーン間の中央パイプラインのコンセプトは排除されます
ここでは、プレーンはビジネスドメインによって分割されます
データ製品、分析データプレーン、およびそれに対応するドメインのオペレーションプレーンサービスの統合は非常に単純で知能的であり、データの単純な移動によるものです
データ製品は、オペレーションデータを分析データに変換するために必要な知能とコードを埋め込むことがあります
デジタル体験に分析を埋め込むことで、分析とオペレーションのプレーンの境界が解消されます
データメッシュは、オペレーションデータと分析データの基本的な技術的な違いを尊重しながら、このセクションで示されているように、ギャップを埋め、その2つを緊密に統合します
オペレーションデータと分析データのギャップを埋める ビジネスドメインにデータ変更をローカライズする データメッシュは、データモデルが連続的に変更されても、ダウンストリームのデータユーザーに致命的な影響を与えず、データへのアクセスを遅らせることなく可能にします
これは、共有されたグローバルな通称データモデルを取り除き、変更を同期する必要をなくすことによって行われます
データメッシュは変更をドメインにローカライズし、ビジネスの最も深い理解に基づいてデータモデルを設計する自治を与えます
これにより、単一の共有された通称モデルの中央調整の必要なしに、ドメインはデータモデルを連オペレーションデータと分析データの差を埋めること
データの変更をビジネスドメインにローカライズするためのデータメッシュは、下流のデータユーザーに致命的な影響を与えることなく、データへのアクセスを遅くすることなく、データモデルの連続的な変更を可能にします
これは、共有されたグローバルな正準データモデルを取り除き、変更の同期を不要にすることによって行われます
データメッシュは、変更をドメインにローカライズし、ビジネスに関する最も緻密な理解に基づいてデータをモデル化する自治権を与えます
これにより、単一の共有正準モデルの中央調整の必要もなくなります
データメッシュは、ドメインデータを共有するために明確で保証されたデータ契約を課します
ドメインデータ契約は、新しいリビジョンにデータユーザーを円滑に移行するまで、古いリビジョンをサポートします
これにより、ドメインはデータモデルを連続的に変更することができます

パイプラインとデータのコピーの偶発的な複雑さを削減する
Fred Brooksが彼の広く知られた論文「No Silver Bullet—Essence and Accident in Software Engineering」で述べたように、ソフトウェアシステムを構築する際には2種類の複雑さがあります
まず、問題空間において本質的で固有の本質的な複雑さがあります
これはビジネスとドメインの複雑さです
次に、偶発的な複雑さがあります
これは私たち、エンジニア、アーキテクト、デザイナーが私たちの解決策で作り出す複雑さです
偶発的な複雑さを削減することは可能ですし、すべきです
分析ソリューションの世界は、偶発的な複雑さを削減するための機会でいっぱいです
データメッシュが削減するいくつかの偶発的な複雑さについて話しましょう
今日、私たちはデータをコピーし続けています
なぜなら、データをさらに別のアクセスモードや計算モデルに必要とするためです
オペレーションシステムからランディングゾーン、データレイク、それからデータサイエンティストのためのフィーチャーストアにデータをコピーします
また、データをコピーして、データアナリストがアクセスできるようにするためのマートにデータを再度コピーします
そして最後に、下流のダッシュボードやレポートデータベースにコピーします
コピーには複雑で脆弱なパイプラインが必要です
コピーの旅は、1つのテクノロジースタックから別のテクノロジースタック、または1つのクラウドベンダーから別のベンダーに移り続けます
今日、分析ワークロードを実行するためには、価値を得る前にどのクラウドプロバイダがすべてのデータをそのデータレイクやデータウェアハウスにコピーするかを事前に決定する必要があります
データメッシュは、ドメイン指向のデータセマンティクスをカプセル化し、異なるユースケースやユーザーに適したデータの複数のアクセスモードも提供する新しいアーキテクチャユニットを作成することで、この問題に対処します
このアーキテクチャユニットは、データ製品の量子（データクォンタムと呼ばれる）と呼ばれます
データクォンタムは、各ネイティブアクセスモード（SQL、ファイル、イベントなど）に対して明示的な契約と保証を持っています
外部のデータユーザーにデータを提供する場合、インターネット上のどこからでもアクセスできます
データクォンタムは、アクセス時に各インターフェイスでアクセス制御とポリシーの適用を提供します
データクォンタムは、データを変換して維持するコードをカプセル化します
データパイプラインは分解し、内部のデータクォンタムロジックの実装となります
データクォンタムは中間パイプラインの必要なくデータを共有します
複雑で脆弱なパイプラインを削除することは、上流データの変更があった場合の障害の可能性を減らす助けとなります

成長に伴った機敏性の維持
現在、ビジネスの成功は多様な成長に基づいています
新しい買収、新しいサービスライン、新しい製品、地理的拡大などです
これにより、管理するための新しいデータソースが生まれ、新しいデータ駆動型のユースケースを構築する必要があります
多くの組織は、データからの価値を提供するスピード、新しいデータのオンボーディング、または成長するにつれてユースケースを提供するスピードを遅くするか、停滞してしまいます
データメッシュは、機敏性を持続するためのアプローチをいくつかのテクニックにまとめたものであり、組織全体のボトルネック、調整、同期を減らすことを目指しています
機敏性は、ビジネスドメインが最小限の依存関係で目標を自律的に達成できる能力に依存しています
中央集権化および一元化のボトルネックの解消
中央集権的なデータチームが一つの統一されたデータレイクまたはデータウェアハウスを管理することは、特に取り込むソースの数や対応するユースケースの数が増えるにつれて、アジリティを制限します
データメッシュは、アーキテクチャと人間のコミュニケーションの両面から、特に複数の当事者の同期の焦点となる場合に、中央集権的なボトルネックを注意深く探し出します
アーキテクチャ的には、これらのボトルネックにはデータレイクとデータウェアハウスが含まれます
データメッシュは、データの提供と利用時にピアツーピアのアプローチを提案しています
このアーキテクチャでは、消費者は直接データソースからデータを発見して使用することができます
たとえば、MLトレーニング機能やレポートは、レイクやウェアハウスなどの中央集権的なアーキテクチャコンポーネントや中間データ（パイプライン）チームの介入なしで、独立したデータ製品に直接アクセスすることができます
図7-4は、概念のシフトを示しています
各データ製品は、ピアツーピアのデータの消費を可能にするバージョン化されたインターフェースを提供します
複数のデータ製品からのデータは、新しい上位データ製品に組み立てられて集約することができます
図7-4
データメッシュは中央集権的なアーキテクチャのボトルネックを解消します


データパイプラインの調整の軽減
運用スケールを超えた技術が共中央集権的で巨大なボトルネックを解消する データメッシュでは、特にソースの数やユースケースの数が増えるにつれて、敏捷性が制限されるような、一元管理されたデータレイクやデータウェアハウスの運営を中心に据えたボトルネックを慎重に検討します
アーキテクチャや人間のコミュニケーションの観点から見て、このようなボトルネックはデータレイクやデータウェアハウスなどを含んでいます
データメッシュは、データの提供と利用時にピアツーピアのアプローチを提案しています
このアーキテクチャでは、消費者はソースのデータ製品から直接データを発見して使用することができます
例えば、MLトレーニング機能やレポートは、レイクやウェアハウスといった中央集権的なアーキテクチャ要素や中間のデータ（パイプライン）チームの介入なしで、独立したデータ製品に直接アクセスすることができます
図7-4は、この概念的な変化を示しています
各データ製品は、ピアツーピアでデータを消費するためのバージョン管理されたインターフェースを提供します
複数のデータ製品からのデータは、新しい上位のデータ製品に組み合わせて集約することができます
 図7-4. データメッシュは中央集権的なアーキテクチャのボトルネックを除去する データパイプラインの調整を減らす 近年、オペレーションのスケールが超えられた技術に共通していることは1つあります
それは、協調と同期を最小限に抑えています
非同期I/Oにより、ネットワークアプリケーションのスループットがブロッキングI/Oを超えてスケーリングされました
リアクティブアプリケーションにより、メッセージの並列処理が高速化されました
MapReduce関数型プログラミングにより、大容量データ処理が多数のサーバーに分散されました
コレオグラフ型のイベント駆動型マイクロサービスは、ビジネスワークフローのスケーリングを実現しました
コア技術での協調と同期の除去に関する不断の努力にもかかわらず、組織とアーキテクチャの協調はほとんど軽視されてきました
その結果、コンピュータシステムの動作がどれほど速くても、チームと人の活動の調整の遅れが成果の達成を阻んでいます
データメッシュは、アーキテクチャと人の協調を減らします
既存のアーキテクチャは、パイプラインのタスク（摂取、処理、提供など）の技術的分解に基づいて構築されています
このようなスタイルのアーキテクチャ分解では、新しいデータソースや新しいユースケースの提供のたびに、これらの機能間での密な調整が必要となります
データメッシュは、データ管理の技術的な分割からドメイン指向の分割へと移行します
ドメイン指向のデータ製品は、他のデータ製品とは独立して開発および進化します
このドメイン指向の分解により、目標の達成には調整が必要ありません
大部分の場合、ドメイン指向のデータ製品チームは、新しいデータソースに対して新しいユースケースを扱うことができます
ドメイン外の新しいデータ製品へのアクセスが新しいユースケースに必要な場合、消費者は新しいデータ製品の標準契約、モック、スタブ、またはシンセティックデータのインターフェースを利用して進捗を得ることができます
これが契約の美しさです
開発中における消費者とプロバイダーの間の調整を容易にするためです
図7-5は、パイプラインの調整の軽減を示しています
 図7-5.パイプラインのアーキテクチャ調整を減らす データガバナンスの調整を減らす もう1つの重要な調整のボトルネックは、データガバナンスの中央機能です
現在、データガバナンスの調整はデータへのアクセスの許可、データの品質の承認、データの変更が組織のポリシーに適合しているかの検証に必要です
データガバナンスの中央化と手作業のプロセスは、データ共有の敏捷性を妨げます
データメッシュは、2つの機能によってガバナンス調整の摩擦を軽減します
1つは、ポリシーをコードとして自動化し、各データ製品に埋め込むことです
もう1つは、ガバナンスの中心的な責任を個々のドメインデータ製品のオーナーに委任することです
これらの変更は、データメッシュの連邦と計算データガバナンスモデルによって実装されます
操作的には、ガバナンスチームは個々のドメインデータ製品のオーナーで構成されます
ドメインデータの共有に責任を持つ長期的な製品オーナーです
アーキテクチャ的には、ガバナンス機能はポリシーの実行を計算化および自動化された形式で各データ製品に組み込みます
これにより、データの発見、データの承認、および必要なポリシーに従うことを確認するための主要な同期ポイントの1つである現在のガバナンス機能が大幅に改善されます
ご想像の通り、ドメインの自律性がチェックされない場合、望ましくない結果が生じる可能性があります
ドメインの孤立、ドメインのデータ製品と他のデータ製品の非互換性および切断、および複数のドメインのデータを消費する場合の断片化された体験です
データメッシュのガバナンスは、ドメインのデータ製品を使用して一貫性のある、接続された、信頼性のあるエクスペリエンスを提供するために、ガバナンスへの自動化の依存度が非常に高いです
図7-6は、ポリシーをコードとして埋め込んだデータ製品の自動配信によって、手動および中央のガバナンス機能が置き換えられる様子を示しています
図7-6
データガバナンスの同期を削減し、自律性を可能にする

 チームの自律性とチームのパフォーマンスの相関は、チームマネジメントの研究の主題となっています
実証的な研究では、チームがミッションを達成するために意思決定の自由を持つことが、より良いチームのパフォーマンスにつながることが示されています
一方、自律性が強すぎると、一貫性の欠如、重複した努力、およびチームの孤立につながる可能性があります
データメッシュは、チームの自律性とチーム間の相互運用性と協力のバランスを取ろうとします
ドメインチームには、例えばデータ製品に最適なデータモデルを選択するなど、ローカルの意思決定を制御する自律性を与えますが、全てのデータ製品に一貫したエクスペリエンスを提供するために、例えば全ドメインが利用するデータモデリング言語を標準化するなど、計算ガバナンスポリシーを使用します
データメッシュは、ドメインチームがデータ製品を構築・維持する自律性を与える一方、ドメインに依存しないデータプラットフォームを提供し、チームが一貫性とコスト効率の良い方法でそれを行うことを可能にします
セルフサーブデータプラットフォームの原則は、ドメインチームがデータ製品のライフサイクルを自律的に管理し、ジェネラリスト開発者のスキルセットを活用することが可能になります
セルフサーブデータインフラストラクチャを使用すると、データ製品開発者はデータ製品の構築、デプロイ、監視、および保守を行うことができます
データの利用者は、データ製品を発見し、学習し、アクセスし、利用することができます
セルフサーブインフラストラクチャにより、データ製品のメッシュを結合、相関、統合して利用することが可能になりますが、ドメインチームの独立性は維持されます
データの投資対効果を向上させる データに対する投資に比べて、データ管理に対する投資からはあまり価値を得られていないという事実を指摘する、ニューバンテージパートナーズの報告書などの業界報告書や、私自身の経験があります
データチームやデータソリューションから得られる価値を、アプリ開発インフラなどの他の技術的投資と比較すると、データに関しては頭風が吹いていることは明らかです
データメッシュは、分析データ管理において、価値と努力の比率を改善する方法を探ります
組織の信頼境界や物理的な場所を越えてデータを共有することを可能にするオープンなデータインターフェースを通じて、現在の技術の複雑さを抽象化し、データユーザーのエクスペリエンスから摩擦を取り除くためにプロダクト思考を適用することで、今日の技術の複雑さを抽象化し、データプラットフォームの新しい原型を作り出します
技術的な複雑さの試金石は、データエンジニアとデータ専門家の必要性がますます高まることです
私たちはそれらを十分に持っているようには思えません
もう1つの試金石は、データパイプラインプロジェクトの付加価値と努力の比率の低さです
多くの努力が投入されても、ほとんど価値が返されないため、品質の高いデータセットへの摩擦のないアクセスを得ることができません
データメッシュは、既存の技術の風景を批判的に見て、技術ソリューションをデータ製品開発者（またはユーザー）中心のプラットフォームとして再構築します
これにより、データ専門家の必要性がなくなり、一般的な専門家がデータ製品を開発できるようになります
さらに、データメッシュは、すべてのデータ製品が共有するさまざまな機能のためのオープンで標準的なインターフェースのセットを定義し、より協力的な技術生態系を実現するための、データの発見、アクセスの要求、データのクエリ、データの提供、データの保護などを可能にします
これにより、ベンダー間の統合コストを削減します
製品思考をどこにも組み込むデータメッシュは、データユーザーが価値に焦点を合わせるためにいくつかの変化をもたらします
データを資産として考えるのではなく、データを製品として考えます
成功の測定基準をデータの量からデータユーザーの満足度に変えます
データは、データメッシュエコシステムの一部として製品として扱われるだけでなく、セルフサーブデータプラットフォーム自体も製品です
この場合、データ製品の開発者とデータ製品のユーザーに役立つものです
データメッシュは、プラットフォームの成功の測定基準を能力の数から、データ製品の開発の経験の改善およびリードタイムの短縮、またはデータ製品の発見と使用に重点を置いた能力の影響に変えます
製品思考は、データ製品のユーザーや開発者の日常の経験に隠れて努力とコストを削減します
業務機能の改善はほとんど常に単位の境界を超えた洞察を必要とします
多くの異なる業務ドメインのデータが必要です
同様に、顧客、従業員、パートナーへのサービスで組織が生成し制御するデータを超えたデータ駆動の価値は、さまざまな角度からのデータを必要とします
たとえば、Daffは、自動再生音楽でリスナーにより良い体験を提供するために、リスナーのプレイリストだけでなく、友人のネットワークやソーシャルおよび環境的な影響や行動に関するデータも必要とします
ニュース、天気、ソーシャルプラットフォームなど、Daffのさまざまな角度からのデータが必要です
データメッシュでは、データメッシュ内に関係なくデータにアクセスできるようにデータ量子のコンセプトを提供します
データ量子は、物理的な場所に依存せずにデータ製品に適切なアクセス制御を持つ人々がデータを発見し利用することを可能にする一連のインターフェースを提供します
識別スキーマ、アクセス制御、その他のポリシーの施行は、インターネット上で有効化されるオープンプロトコルの使用を前提としています
データメッシュアーキテクチャは、組織の境界を超えたデータの接続により、さらなる価値を提供します
この章を読んだ後、データメッシュが万能の解決策であると思われるかもしれません
まったくその逆です
データメッシュは解決策の重要な一部です
データへのアクセスを真に民主化することができます
ただし、データから価値を生み出すためには、データを共有するだけでは不十分です
継続的に繰り返し可能な、本格的な分析および機械学習ベースのソリューションを提供する必要があります
ただし、起動するためには、データの大規模な共有が必要であり、それがデータメッシュが焦点を当てていることです
この章で挙げられているデータメッシュの目標は、データを再構築する方法、それを管理する方法、チームを構築する方法を再考するように私たちに招待します
この章では、データメッシュの目標とそれを実現する要素を紐付けました
これをまとめるために、表7-1をご覧ください
表7-1
データメッシュの転換点後の要約
データメッシュの目標
- 複雑で不安定で不確実なビジネス環境でのデータの適切な管理
- ビジネス、テクノロジー、データの調整
- ビジネス、テクノロジー、データの各チームを作成し、それぞれが自身のデータの長期的な所有権を持つ

領域データ所有の原則
- 運用と解析のデータプレーンのギャップを縮める
- 組織全体のパイプラインと二つのプレーンデータアーキテクチャを廃止する
- アプリケーションとデータプロダクトを「ダムパイプ」を通じてより密接に統合する

データプロダクトとしての原則
- ビジネスドメインにデータ変更を局所化する
- データプロダクトのメンテナンスと所有権を特定のドメインに局所化する
- ドメイン指向のデータプロダクト間の明確な契約を作成し、変更の影響を減らす

データプロダクトとしての原則(継続)
- パイプラインとデータのコピーの偶発的な複雑さを減らす
- パイプラインを分解し、必要な変換ロジックを対応するデータプロダクトに移し、内部実装として抽象化する

データプロダクトの量子アーキテクチャの要素
- 成長に直面してアジリティを維持する
- 中央集権的なアーキテクチャのボトルネックをなくす
- 中央集権的なデータウェアハウスとデータレイクを廃止する
- データプロダクトを介したピア・ツー・ピアのデータ共有をデータインタフェースを通じて可能にする

ドメイン所有の原則
- データパイプラインの調整を減らす
- パイプラインアーキテクチャの上位機能分解からドメイン指向のアーキテクチャ分解へ移行する
- ドメイン指向のデータプロダクト間に明示的なデータ契約を導入する

データプロダクトとしての原則
- データガバナンスの調整を減らす
- ガバナンスの責任を自治ドメインとそのデータプロダクト所有者に委任する
- ガバナンスポリシーをコードとして自動化し、各データプロダクトの量子に埋め込み検証する

連邦コンピュータガバナンスの原則
- チームの自律性を可能にする
- ドメインチームに独立して迅速に進む自律性を与える
- チームの自律性を計算基準とのバランスを取りながら確立し、メッシュのグローバルな一貫した体験と相互運用性を創造する
- ドメインに中立なインフラストラクチャ機能をセルフサーブ方式でドメインチームに提供する

セルフサーブデータプラットフォームの原則
- コストよりもデータからの価値を高める
- データプラットフォームによる複雑さを抽象化する
- データ開発と利用の旅路で摩擦や隠れたコストをなくすために、データデベロッパーセントリックとデータユーザーセントリックのインフラストラクチャを作成する
- データプロダクトのベンダー統合の複雑さを減らすために、オープンで標準的なインタフェースを定義する

セルフサーブデータプラットフォームの原則
- プロダクト思考をあらゆるところに取データメッシュの転換点以降のまとめ
データメッシュの目標
- 複雑で不安定、不確実なビジネス環境でデータの変更を適切に管理する
- ビジネス、テクノロジー、データを調整する
- ビジネス、テクノロジー、データの各チームを作成し、それぞれが長期的な所有権を持つデータに責任を持つ

ドメインデータ所有の原則
- 運用と分析のデータプレーンのギャップを埋める
- 組織全体のパイプラインと二層のデータアーキテクチャを削除する
- ダムパイプを通じてアプリケーションとデータ製品をより緊密に統合する

データ製品としての原則
- ビジネス領域によってデータ変更を局所化する
- データ製品のメンテナンスと所有権をそれぞれの具体的な領域に局所化する
- 領域指向のデータ製品間の明確な契約を作成し、変更の影響を減らす

データ製品としての原則 (続き)
- 複雑で不安定、不確実なビジネス環境でデータの変更を適切に管理する
- パイプラインとデータのコピーの偶発的な複雑さを減らす
- パイプラインを分解し、必要な変換ロジックを対応するデータ製品に移動し、内部実装として抽象化する

データ製品の量子アーキテクチャコンポーネント
- 成長に対応するためにアジリティを維持する
- 中央集権的なアーキテクチャのボトルネックを削除する
- 中央集権的なデータウェアハウスとデータレイクを削除する
- データ製品を介したピアツーピアのデータ共有をデータインターフェースで可能にする

領域所有の原則

データ製品としての原則
- データパイプラインの調整を減らす
- 上位レベルのパイプラインアーキテクチャから領域指向のアーキテクチャへ移行する
- 領域指向のデータ製品間に明示的なデータ契約を導入する

領域所有の原則

データ製品としての原則
- データガバナンスの調整を減らす
- 自律ドメインとそのデータ製品所有者にガバナンスの責任を委任する
- 各データ製品量子によって埋め込まれ、検証されたガバナンスポリシーをコードとして自動化する

連邦計算ガバナンスの原則
- チームの自律性を可能にする
- ドメインチームに個別に迅速に動作する自律性を与える
- チームの自律性と計算基準をバランスさせ、メッシュの全体的な一貫した体験と相互運用性を作り出す
- ドメインに関係しないインフラストラクチャの機能をドメインチームに自律的に提供する

自己サーブデータプラットフォームの原則
- コスト対効果によるデータの価値を高める
- データプラットフォームによって複雑さを抽象化する
- データ開発と使用の過程で摩擦と隠れたコストを除去するために、データ開発者とデータユーザーを中心としたインフラストラクチャを作成する
- データ製品のベンダー統合の複雑さを減らすために、データ製品のためのオープンで標準的なインタフェースを定義する

自己サーブデータプラットフォームの原則
- プロダクト思考をすべてに組み込む
- データユーザーと開発者の幸福度に基づいて重点を置き、成功を測定する
- データとデータプラットフォームの両方を製品として扱う

データ製品としての原則
- 組織の枠を超える
- データ製品間の標準的なインターネットベースのデータ共有契約により、プラットフォームと組織の物理的および論理的な境界を超えてデータを共有する

自己サーブデータプラットフォームの原則

次の章では、転換点に至るまでの経緯について概説し、ここで行われたデータ管理アプローチが将来には適さない理由について説明します