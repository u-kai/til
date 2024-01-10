level component of data mesh architecture. Many questions are yet to be answered in the gap between this high-level architecture and the actual implementation. What exact technologies to buy or build? How to evaluate these technologies? How to integrate them? This chapter walks you through a framework that you can apply to answer these questions, in the context of your organization and available technology stack. It provides examples of what constitutes a platform to help you kickstart your creative journey. It shows you the road to follow with all its twists and turns to discover what is truthful for your organization. By now you have some familiarity with the concept of the data mesh platform and have a language to describe it. Chapter 4 introduced the platform as: A collection of interoperable domain-agnostic services, tools, and APIs that enable cross-functional domain teams to produce and consume data products with lowered cognitive load and with autonomy. Chapter 9 proposed a multiplane platform to distinguish between different classes of platform services based on their scope of operation without imposing a strict layering. The three planes of the platform include: Data infrastructure (utility) plane Atomic services to provision and manage physical resources such as storage, pipeline orchestration, compute, etc. Data product experience plane Higher-level abstraction services that operate directly with a data product and enable data product producers and consumers to create, access, and secure a data product, among other operations that run on a data product. Mesh experience plane Services that operate on a mesh of interconnected data products such as searching for data products and observing the data lineage across them. All of these planes can be directly accessed by the consumers of the platform: data product developers, consumers, owners, the governance function, etc. Figure 10-1 illustrates the platform planes and their users’ personas. Figure 10-1. Multiplane self-serve platform and platform users The mesh experience plane optimizes the experience of people who need to operate, govern, and query the mesh as a whole. For example, members of the governance team and data product owners interact with the services in this plane to evaluate the current state of policies, monitor the general operational health of the mesh, and search for existing data products. It is also used by data product consumers and providers in scenarios where they need to work with a collection of data products such as searching, browsing, or debugging through lineage and logs. Data product consumers and providers mainly interact with the data product experience plane to discover, learn, understand, consume, build, and maintain data products. The data product experience plane is optimized for the delivery and consumption of data products as the unit of value exchanged between the consumers and providers. To a lesser degree they may use services in the infrastructure plane, often in cases where the data product experience plane doesn’t yet provide the services they need. For example, a particular data product may want to expose its output port via a graph query language, but the data product experience plan does not support this as a packaged output port implementation, yet. In this case, the data product developers may use the infrastructure services directly to provision a graph query engine and then hook it up to their output port implementation. Ideally their graph query output port implementation gets harvested later and built back into the standard output ports supported by the platform. Data mesh planes do not impose a strict layering in accessing their services. The data product experience plane provides a set of operations on a data product and manages the complexity of provisioning its underlying infrastructure. But it does not intend to strictly stop platform users from accessing the infrastructure plane. It’s just going to make it so easy to work with data products that domain teams would not have any incentive to directly work with the infrastructure services. While the data product experience plane optimizes for the user experience, the data infrastructure utility plane makes sure that the resources’ performance and utilizations are optimized. The utility plane optimizes to get the best out of what the underlying infrastructure providers offer around storage, compute, orchestration, etc. It adapts the data product experience plane to the underlying physical hosting environment. It is organized around the underlying resources and its users, the majority of which are the platform engineers who are building the other planes. Many of the data infrastructure plane services are shared with the operational systems. In this chapter I will take you further along the design of your data mesh platform capabilities with a user-experience-centric approach. I will go through a few key journeys of data product developers and consumers and discover what interfaces the platform exposes to streamline the journeys. Design a Platform Driven by User Journeys The ultimate purpose of the platform is to serve the cross-functional domain teams so they can deliver or consume data products. Hence, the best way to begin the design of the platform is to understand the main journeys of your platform users and evaluate how you can make it easy for them to complete their journeys successfully using the platform. There are a few main high-level personas in the data mesh ecosystem: Data product developers This persona covers the role of a data product developer, considering a wide spectrum of skill sets—from generalist developers with general programming skills to specialist data engineers who are well-versed in the existing analytical data processing technologies. Data product consumers This persona covers multiple roles that have one thing in common: they need access and use data to do their job. Their job might be: Training or inferring ML models, as data scientists Developing reports and dashboards, as data analysts Developing new data products that consume data from existing ones, as data product developers Building data-driven services and applications in the operational domain, as an application developer This is quite a wide range, and platform interfaces and services can vary depending on the exact role. Data product owners Data product owners are responsible for delivering and evangelizing successful data products for their specific domains. The success of the data products depends on their adoption and value delivery, as well as conformance to the wider policies of the mesh and their interoperability with other data products. They use the platform to keep the pulse of their data products. Data governance members Since data governance is founded on a federated structure, there is no one specific data governance role. Nonetheless, the members of the governance team have a collective set of responsibilities to assure the optimal and secure operation of the mesh as a whole. There are various roles within the governance group such as subject matter experts on security and legal issues, as well as data product owners with specific accountabilities. The platform facilitates their collective journey as well as the individual roles. Data platform product owner The data platform planes and their services are products whose users play all the other roles mentioned earlier. The data platform product owner is responsible for delivering the platform services as a product, with the best user experience. Based on an understanding of the platform users’ needs and constraints, the data platform product owner prioritizes the services that the platform offers. Data platform developer The platform developers build and operate the data platform as well as use it. Data platform developers who work on the data product experience plane services are users of the utility plane services. Hence, their skillset and journey are important to the design of the utility plane services. In this chapter I demonstrate how to approach the platform’s design using two example roles: a data product developer as well as a data scientist as a data product consumer. Data Product Developer Journey One of the key journeys of a data product developer is to create and operate a data product. This is an end-to-end and long-term responsibility that embraces the continuous delivery principle of “everyone is responsible.” Building and operating a data product is not someone else’s problem. Figure 10-2 shows the high-level phases of the data product creation journey. These stages are iterative with continuous feedback loops. For simplicity of demonstration, I have drawn the journey as a linear succession of stages. The journey begins with the activities around the inception of a data product such as ideation and exploration of sources. It continues to the actual build and delivery of the data product, followed by monitoring it, continuously evolving it, and optionally retiring it. A data product development journey interacts with other journeys. In the case of a source-aligned data product that consumes data from an operational system, the data product developers work closely with the source application developers. They co-design and implement how the application will share its operational data as an input to the data product. Note that these people belong to the same domain team. For example, development of the play events data product requires close collaboration with the player application that generates the raw data. Figure 10-2. High-level example of the data product development journey Similarly, if the data product is built for a known analytical application, in the case of a consumer-aligned data product, there is a collaboration that happens early between the data product developers and the analytical application developers. Data product developers need to understand the goals of the analytical application and its needs in terms of data fields and the guarantees that the data product must support. For example at Daff, an ML-powered playlist recommendation application requires data from the music profiles data product. The music profiles data product developers need to work closely with the playlist recommendation application team to align on the expected classifications of the profile. Let’s look at the high-level stages of data product development and how the platform interfaces are designed to support them (Figure 10-3). Note The notation and phrasing of the interfaces are not representative of a final implementation. The imperative language is chosen to focus on the functionality, while in reality the declarative forms of APIs are more desirable. For example, the /deploy interface is very likely implemented as /deployments, the declarative resources for the deployed data products. The interfaces are exemplary and demonstrate the semantics of the operation and not the actual syntax or mechanism of the interface, e.g., CMD, UI, API, etc. Figure 10-3. High-level example of the data product development journey using the platform Incept, Explore, Bootstrap, and Source Data products often come to exist because we discover one or many potential analytical use cases that need their data. While data products, particularly source-aligned ones, have many diverse use cases, some beyond our imagination today, we still need to start the inception—the creation—of them from a real use case that demonstrates the value of the data product in a real-world context. Consumer-aligned data products directly serve one or multiple specific use cases. The inception of a consumer-aligned data product always involves direct collaboration with its consumers. During the inception of a data product, the developer is in an exploratory phase looking for potential sources for the data product. The source can be discovered as an upstream data product, an external system (to the organization), or an operational system (in the organization). During the exploration, the developer may study the upstream source’s discovery information, to evaluate the source’s suitability, data guarantees, documentation, data profiling, existing users, etc. Once potential sources are nominated, the developer can quickly bootstrap a hello world data product. The platform provisions the necessary infrastructure. It enables the data product developer to connect to the sources and consume data—either in the form of synthesized data, obfuscated, or the actual data. At this point the developer has all the necessary resources to start experimenting with the source data and the data product’s transformation logic. These steps constitute the exploratory steps of incepting a data product. Rapid discovery, access to source output data, quick scaffolding of a data product, and provisioning of infrastructure are all the necessary pieces of the inception and bootstrapping phase. For example, the playlist team plans to generate new playlists for holiday seasons across different countries. They need data about regional holidays and their sentiment, as well as labeled music profiles that associate different tracks with their relevant holiday. They start by searching for existing data products on the mesh with similar information. They find a few sources for music profiling and then connect to those data products and start playing with the data to assess their completeness and relevance. They quickly create a simple holiday playlist data product to see what playlists they can generate from the simplest implementation of music holiday profiling. Table 10-1 provides a summary of platform APIs used during this phase. Table 10-1. Example interfaces provided by the platform planes in support of a data product inception Development phase Platform Plane Platform interface Platform interface description Incept | Explore Mesh experience /search Search the mesh of existing data products to find suitable sources.
The search can be based on various parameters such as source operational systems, domains, and types of data. Mesh experience /knowledge-graph Browse the mesh of related data products’ semantic models. Traverse their semantic relationship to identify the desired sources of data.
Example: the playlist team can browse the semantic relationship between the holiday-related music profilings they have found. Mesh
experience /lineage Traverse the lineage of input-output data between different data products on the mesh to identify the desired sources based on the origin of the data and the transformations the data has gone through.
Example: the playlist team can look at how the existing music holiday profiling is created and trace back what its sources are and what transformation has led to the classifications. This is to evaluate the suitability of the existing music holiday profiling. Bootstrap | Source Data product experience /{dp}/discover Once a source is identified, access all the data product discoverability information such as documentation, data model, available output ports, etc. Data product experience /{dp}/observe Once a source is identified, access the data product’s guarantees and metrics—real-time—such as how often the data is being released, the last release date, data quality metrics, etc.
Example: Once the playlist team finds a promising holiday music profiling, they need to evaluate how often this data is refreshed and what its completeness compared to all the music available on the platform is. This assessment gives them more trust in the data. Data product experience /init To get started with experimenting with the source data, this API bootstraps a barebones data product with enough of the infrastructure available to connect to the sources, access their data, run transformation on the data, and serve the output in a single access mode for verification.
The scaffolding of a data product allocates the continuous delivery pipeline, early environments, provisioning of data processing clusters, and accounts to run and access infrastructure resources for the data product. This marks the beginning of the data product life cycle. Mesh experience /register During the initialization of a new data product, it’s automatically registered with the mesh. It’s given a unique global identifier and address and makes the data product visible to the mesh and the governance process. Data product experience /connect Once the source is discovered, the data product gets access to the source by connecting to it. This step validates the access control policies governing the source.
This might trigger a request for permission to get access to the data product. Bootstrap | Source (continued) Data product experience /{dp}/{output}/
query
/{dp}/{output}/
subscribe Read data from a particular output port of the source data product, either in a pull-based querying model or subscribing to changes. Build, Test, Deploy, and Run Data product developers have an end-to-end responsibility of building, testing, deploying, and operating their data products. The stage I have simply called “build, test, deploy, and run” is a continuous and iterative series of activities that data product developers perform to deliver all the necessary components of a successful data product. This data product is autonomously discoverable, understandable, trustworthy, addressable, interoperable and composable, secure, natively accessible, and valuable on its own (see “Baseline Usability Attributes of a Data Product”). Table 10-2 lists the high-level interfaces to deliver a data product. Table 10-2. Example interfaces provided by the platform planes in support of data product development Development phase Platform plane Platform interface Platform interface description Build Data product experience /build Compile, validate, and compose all the components of a data product into a deployable artifact that can be used to run the data product in various deployment environments.
See Table 10-3 for the data infrastructure plane interfaces used during this phase. Test Data product experience /test Test various aspects of a data product. Test the functionality of data transformation, the integrity of the output data, expectations of the data versioning and upgrade process, expectations of the data profile (its expected statistical characteristics), test bias, etc.
Testing capabilities are offered in different deployment environments. Deploy Data product experience /deploy Deploy the built revision of a data product to an environment, including the developer’s local machine, a development sandbox on the target hosting environment (e.g., a specific cloud provider), and pre-production or production environments. Run Data product experience /start
/stop Run or stop running the data product instance in a particular environment. Build/Test/Deploy/Run Data product experience /local-
policies One of the main components of a data product is the policies that govern its data and function. Policies include encryption, access, retention, locality, privacy, etc.
The platform facilitates configuration and authoring of these policies locally, during the development of a data product, their validation during test, and their execution during access to data. Build/Test/Deploy/Run Mesh experience /global-
policies In many organizations policies that govern a data product are authored by the global (federated) governance team.
Hence, the platform enables authoring of the global policies and application of these policies by all data products. Let’s look more closely at what is involved at the data infrastructure plane. Building a data product has many parts to it. A data product developer must configure or code all of these parts to implement an autonomous data product. The platform is a key enabler in lowering the cost and effort required to code or configure each of these components. Figure 10-4 maps the build and execution of different components of a data product to the lower-level platform interfaces, infrastructure plane APIs. A data product developer, in the majority of the cases, simply interacts with the data product experience plane: building, deploying, and testing a single data quantum. The data product experience plane delegates the implementation of the data product components to the data infrastructure plane. The developers must understand, code, and configure different aspects and components of a data product given the technologies that the data infrastructure plane offers. For example, the choice of how to code the transformation aspect of a data product depends on what computation engines the data infrastructure plane supports.1 Nevertheless, data product developers focus on coding the transformation, using the technology made available to them, and leave it to the data product experience plane to manage the build, deployment, and execution of transformation along with all the other components. In this process, the data product experience plane delegates the details of transformation execution to the data infrastructure plane. Figure 10-4. Example of data infrastructure plane interfaces to support data product delivery Table 10-3 lists some of the data infrastructure plane capabilities utilized during the development of a data product. Table 10-3. Example of data infrastructure plane interfaces that support data product experience plane APIs Platform interface Platform interface description /input-ports Provides different mechanisms for consuming data according to the design of data products, e.g., event streaming input, remote database queries, file read, etc. It triggers executing the transformation when data is available.
The input port mechanisms keep track of consumption progress as new data becomes available. /output-ports Provides different mechanisms for serving data, according to the modes of access data product offers, e.g., streams, files, tables, etc. /transformations Provides different execution mechanisms for running all the computation necessary to transform data. /containers Managing the life cycle of data products and all its required infrastructure resources, as an autonomous unit. /controls Provides a wide range of technologies to enable configuration and execution of diverse and evolving policies such as access control, encryption, privacy, retention, locality, as code. /storage Input ports, output ports, and transformation all require access to permanent and temporary storage for data and their metadata (SLOs, schemas, etc.) The platform must supply access to different types of storage (blob, relational, graph, time series, etc.) and all the operational capabilities such as failover, recovery, backup, and restore. /models The mechanisms to describe, share, and link semantic and syntax schema models of the data. /identities The mechanism to identify and address a data product uniquely on the mesh. Maintain, Evolve, and Retire Maintenance and evolution of a data product involves continuous change to the data product: improving its transformation logic, evolving its data model, supporting new modes of access to its data, and enrichment of its policies. The changes to data products result in continuous iterations of build, test, deploy, and run (“Build, Test, Deploy, Run”), while maintaining an uninterrupted processing and sharing of data. The platform must reduce the overhead of how a data product is maintained to what is to be maintained. For example, if the result of a new build is a change to the semantic and schema, the data product developer simply focuses on changes to the data model, and the transformation of data based on the new model. The complexity of how different versions of schemas (semantic and syntax) are managed, associated with the data, and accessed by the consumer is managed by the platform. In some cases, the evolution of a data product has no impact on its underlying infrastructure resources. For example, a bug fix to the transformation code simply requires rebuilding and redeploying the data product and fixing the generated data. In other cases, the evolution impacts the underlying resources. For example, migrating a data product to a new environment in the case of switching storage vendors requires reallocating the storage resources and migrating the data product’s underlying data. Monitoring the operational health of each data product, and the mesh as a whole, is another key capability that the platform offers during this phase. The operational excellence of the mesh depends on monitoring various aspects of each data product: monitoring its performance, reliability, SLOs, effectiveness of its computational policies, and operational cost based on resource usage. In addition to monitoring individual data products, the mesh experience plane must gather insights and monitor the state of the mesh as a whole. For example, detect and alarm when master data products are detected—data products that aggregate data from too many sources and become a bottleneck as they are used by many. During the lifetime of a data product, there will be times when mesh-level administrative controls must be invoked. For example, the right to be forgotten can be triggered through the mesh-level global controls, delegated to every single data product’s control interface, and implemented through the data eviction function of their underlying storage. Will a data product ever cease to exist? I can think of two scenarios that it would. A data product can retire in the case of migration to a new data product or in the case that all of the data records it has ever generated must be discarded and no new record is ever going to be generated. In both cases the platform enables data product developers to gracefully retire the data product so its downstream consumers can either migrate to the new source over time or themselves retire. As long as someone is using the past data, the data quantum continues to exist, though it may not execute any further transformation nor produce any new data. A dormant data product will continue to serve its old data and enforce its policies, while a fully retired data product is simply extinct. Table 10-4 shows a few of the platform interfaces to support the maintenance, evolution, and retirement phases of data product development. Table 10-4. Example of data platform interfaces to support data product maintenance Platform plane Platform interface Platform interface description Data product experience /{dp}/status Checking the status of a data product. Data product experience /{dp}/logs
/{dp}/traces
/{dp}/metrics The mechanism for a data product to emit its runtime observability information such as logs, traces, and metrics, according to the design of the data quantum’s observability (“Observability Design”). The mesh layer monitoring services utilize the data provided by these mechanisms. Data product experience /{dp}/accesses The logs of all accesses to the data product. Data product experience plane /{dp}/controls Ability to invoke high-privileged administrative controls such as right to be forgotten on a particular data quantum. Data product experience /{dp}/cost Tracking the operational cost of a data product. This can be calculated based on the resource allocation and usage. Data product experience /migrate Ability to migrate a data product to a new environment. Updating the data product revision is simply a function of build and deploy. Mesh experience /monitor Multiple monitoring abilities at the mesh level, logs, status, compliance, etc. Mesh experience /notifications Notification and alerting in response to detected anomalies of the mesh. Mesh experience /global-controls Ability to invoke high-privileged administrative controls such as right to be forgotten on a collection of data products on the mesh. Now let’s move our attention to the data consumer journey and see how platform interfaces may evolve to support such a persona. Data Product Consumer Journey The  persona of a data consumer represents a wide range of users, with various skill sets and responsibilities. In this section, I focus on an example data scientist persona who is consuming existing data products to train an ML model and then is deploying the ML model as a data product to make inferences and generate new data. For example, Daff uses an ML model to produce a curated playlist every Monday for all of its listeners to start their week. The ML recommendation model that generates the Monday playlists is trained using data from existing data products such as listener profiles, listener play events, their recent reactions to the songs they have liked or haven’t, and the playlists of what they have listened to recently.2 Once the ML model is trained, it is deployed as a data product, monday playlists. If you recall from Chapter 7 data mesh creates a tightly integrated operational and analytical plane. It interconnects domain-oriented (micro)services and data products. Despite a tight integration and feedback loop between the two planes, data mesh continues to respect the differences of the responsibility and characteristics of each plane, i.e., microservices respond to online requests or events to run the business, and data products curtate, transform, and share temporal data for downstream analytical usage such as training machine learning models or generating insights. Despite the attempt in clarifying the boundary between the entities of the two planes, ML models blur this boundary. ML models can belong to either plane. For example, an ML model can be deployed as a microservice to make inferences online as the end user makes requests. For example, during the registration of a new listener, given the information provided by the listener, a classifier ML model is called to augment the user’s information with their profile classification. Alternatively, an ML model can be deployed as the transformation logic of a data product. For example, the playlist recommender ML model is deployed as a data product that makes inferences every Monday and generates new playlists. This data then can be fed into the operational service that displays the playlists to the listener. Figure 10-5 shows this example. Figure 10-5. Example of a bimodal deployment of ML models In this section, I explore the data scientist’s journey for an ML model that gets deployed as a data product to demonstrate the overlap of this journey with data mesh platform capabilities. Figure 10-6 demonstrates a high-level journey to continuously deliver an ML model as a data product. This value stream follows the practice of continuous delivery for machine learning (CD4ML) to continuously hypothesize, train, deploy, monitor, and improve the model in rapid feedback loops with a repeatable process.3 Figure 10-6. Example ML model development journey From the data platform perspective, this journey is very similar to the “Data Product Developer Journey” we explored earlier. There are some differences that I call out next. Incept, Explore, Bootstrap, Source In this example, ML model developers start their journey with the inception of the data product that will ultimately embed in their model. The inception of this data product includes the formulation of the hypothesis for an intelligent action or decision based on the existing data. For example, is it possible to curate and recommend a playlist that is played and replayed (loved) by many listeners? To start the exploration to validate the hypothesis, the ML model developers explore and search existing data products and evaluate the sources using their discoverability APIs as well as through sampling output ports data. During this phase, the platform interfaces remain the same as what I presented earlier. The main difference here is perhaps the type of discoverability information the data scientists care about. They must evaluate whether there is bias in the data and fairness of the data. This can be evaluated by sampling and profiling the source data product output port. Build, Test, Deploy, Run The journey continues to build the model and train it using the upstream data product output ports. During the training, the model developer may decide that the training dataset requires change, modifying the upstream data product data into features. For example, the training features may only include a subset of playlists information joined with song profiles. In this case, the data transformation pipeline to generate features becomes itself a data product. For example, monday playlist features becomes a data product that is developed similarly to any other one. During this phase the model developers need to keep track of the data that results in a particular revision of the model. The versioning of the source (training and test) data can use the processing time parameter exposed by all data products as the revision of the data—a timestamp that is unique to each state of data and indicates when a data was processed and published. This removes the need to keep a copy of the data4 for future reuse and repeatability of the model, as the source data product always allows retrieving past data. Chapter 12 will elaborate on this. The process of model training and tracking can be handled by an ML platform—a platform of technologies and tools that satisfy the unique needs of ML model training pipelines. These services work closely and integrate with the data platform. During deployment, the model needs to be packaged into a format that can run as the transformation code of the monday playlists data product.5 During the model run, the infrastructure plane transformation engine can handle the unique needs of executing ML models such as execution on targeted hardware—GPUs or TPUs. Maintain, Evolve, and Retire Similar to other data products, developers continue to monitor the performance of the model and make sure the output data, the playlists, is created as expected. One of the unique needs of model monitoring is observing the efficacy of the models and monitoring actions that the listeners take. For example, do new revisions of the model result in an increase in duration of listening to the playlist, more replays, and more listeners? Monitoring such parameters may utilize the operational plane capability of monitoring metrics. In this case, the player application can provide these metrics. Recap If you take away one thing from this chapter, I wish it to be this: there is no single entity such as a platform. There are APIs, services, SDKs, and libraries that each satisfy a step in the journey of the platform users. A platform user journey, whether the user is a data product developer or a consumer, is never isolated to the data mesh platform services. They often cross boundaries with APIs that serve the operational systems—e.g., using streaming capabilities to consume data from a microservice—or cross boundaries with ML model training capabilities such as experiment tracking or utilization of GPUs. Hence, view your platform as a set of well-integrated services that are there to satisfy a seamless experience for their users. The second point I’d like you to take away is to respect the separation of planes so that you can choose to optimize both for the human experience and machine efficiency. For example, use the data product experience plane to optimize the user experience at the logical level of interacting with a single unit of a data quantum. Create a separate data infrastructure plane to optimize for the machine, e.g., separation of compute from storage to handle independent scaling of each or co-locality of all data to reduce data movement. The lower plane optimization decisions should not impair the developer experience, and vice versa. I’d like to leave you with one last point in the first step in your path to creating a data mesh platform: start with discovering and designing an optimized journey for roles such as data product developer and data product consumer, with minimal friction and maximum delight and engagement. Once you have done that, feel free to continue your search to find the right tools for various stages of their journey.

データメッシュアーキテクチャのレベルコンポーネント
このハイレベルアーキテクチャと実際の実装の間にはまだ多くの問いが残っています
どのような具体的なテクノロジーを購入または構築するのか？これらのテクノロジーを評価するにはどうすればよいですか？それらを統合するにはどうすればよいですか？この章では、組織と利用可能な技術スタックの文脈でこれらの問いに答えるために適用できるフレームワークを紹介します
創造的な旅を始めるのに役立つプラットフォームの例を提供します
あなたの組織にとって真実であるかを発見するための曲がりくねった道を示します
今では、データメッシュプラットフォームの概念についてある程度理解ができ、それを説明するための言葉を持っています
4章では以下のようなプラットフォームとして紹介されました：ドメインに依存しない相互運用可能なサービス、ツール、APIのコレクションで、クロスファンクショナルなドメインチームが認知負荷を低減し、自律性を持ってデータ製品を作成し、利用することができるようにします
9章では、操作範囲に基づいてプラットフォームサービスの異なるクラスを区別するためのマルチプレーンプラットフォームを提案しました
プラットフォームの3つのプレーンには、次のものが含まれます：データインフラ (ユーティリティ) プレーン ストレージ、パイプラインのオーケストレーション、計算などの物理リソースを供給し管理するためのアトミックサービス データ製品エクスペリエンスプレーン データ製品と直接操作する高レベルの抽象化サービスで、データ製品のプロデューサーとコンシューマーがデータ製品の作成、アクセス、保護などの操作を行うためのもの メッシュエクスペリエンスプレーン データ製品のメッシュ上で操作を行うサービス、例えばデータ製品の検索やデータ製品間のデータラインセージの観察
これらのプレーンは、データ製品の開発者、コンシューマー、オーナー、ガバナンス機能など、プラットフォームの利用者で直接アクセスすることができます
図10-1は、プラットフォームのプレーンとそのユーザーのパーソナリティを示しています
図10-1
マルチプレーンのセルフサーブプラットフォームとプラットフォームのユーザー メッシュ体験プレーンは、メッシュ全体を操作、ガバナンス、クエリする必要がある人々の体験を最適化します
たとえば、ガバナンスチームのメンバーやデータ製品オーナーは、メッシュの政策の現状を評価し、メッシュの一般的な運用状況をモニタリングし、既存のデータ製品を検索するために、このプレーンのサービスとやり取りを行います
また、データ製品のコンシューマーとプロバイダーも、検索、ブラウジング、ラインセージやログを介したデバッグなど、データ製品のコレクションを扱う場合に、このプレーンのサービスを使用します
データ製品のコンシューマーとプロバイダーは、主にデータ製品エクスペリエンスプレーンとやり取りし、データ製品の発見、学習、理解、利用、作成、維持を行います
データ製品エクスペリエンスプレーンは、データ製品の提供と消費を最適化するために設計されており、それらの間で交換される価値の単位であるデータ製品の提供と消費を管理します
それらは少ないながらも、インフラストラクチャプレーンのサービスを使用する場合もあります
これは、データ製品エクスペリエンスプレーンがまだ必要なサービスを提供していない場合によく起こります
たとえば、特定のデータ製品は、グラフクエリ言語を介して出力ポートを公開したい場合がありますが、データ製品エクスペリエンスプレーンはまだパッケージ化された出力ポートの実装としてこれをサポートしていません
この場合、データ製品の開発者は、インフラストラクチャサービスを直接使用してグラフクエリエンジンをプロビジョニングし、それを出力ポートの実装に接続することがあります
理想的には、彼らのグラフクエリ出力ポートの実装が後で収穫され、プラットフォームでサポートされている標準的な出力ポートに再組み込まれます
データメッシュプレーンは、サービスへのアクセスに厳密なレイヤリングを課しません
データ製品エクスペリエンスプレーンは、データ製品上の一連の操作を提供し、その基礎となるインフラストラクチャの複雑さを管理します
しかし、インフラストラクチャプレーンへのアクセスを厳密に制限することは意図していません
それはデータ製品との作業を非常に簡単にするだけでなく、ドメインチームにはインフラストラクチャーサービスで直接作業する動機がなくなるだろう
データ製品のエクスペリエンスプレーンは、ユーザーエクスペリエンスに最適化され、データインフラストラクチャーのユーティリティプレーンはリソースのパフォーマンスと利用率を最適化します
ユーティリティプレーンは、ストレージ、コンピュート、オーケストレーションなどの基盤となるインフラストラクチャープロバイダーが提供する最大の利点を引き出すために最適化されます
それはデータ製品のエクスペリエンスプレーンを基礎となる物理的なホスティング環境に適応させます
それは基礎となるリソースとそのユーザーの周りで組織されており、その大部分は他のプレーンを構築しているプラットフォームエンジニアです
データインフラストラクチャープレーンの多くのサービスは、運用システムと共有されています
この章では、ユーザーエクスペリエンスを中心に、データメッシュプラットフォームの機能設計をさらに進めていきます
データ製品の開発者と消費者のいくつかの主要な旅を進み、プラットフォームが旅を効率化するためにどのようなインターフェースを提供するかを発見します
ユーザージャーニーに基づいたプラットフォームの設計 プラットフォームの最終的な目的は、クロスファンクションのドメインチームにデータ製品を提供または消費させることです
したがって、プラットフォームの設計を開始する最良の方法は、プラットフォームユーザーの主要な旅を理解し、プラットフォームを使用して旅を成功裏に完了するための方法を評価することです
データメッシュエコシステムにはいくつかの主要な高レベルのパーソナがあります
 データ製品の開発者 このパーソナは、一般的なプログラミングスキルを持つジェネラリスト開発者から既存の分析データ処理技術に精通した専門のデータエンジニアまで、幅広いスキルセットを考慮したデータ製品開発者の役割をカバーしています
 データ製品の消費者 このパーソナは、仕事をするためにデータにアクセスし使用する必要がある複数の役割をカバーしています
彼らの仕事は次のようなものです
 - データサイエンティストとしてのMLモデルのトレーニングまたは推論 - データアナリストとしてのレポートやダッシュボードの作成 - 既存のデータからデータを消費する新しいデータ製品の開発 - オペレーションドメインでのデータ駆動型のサービスやアプリケーションの開発 これはかなり幅広い範囲であり、プラットフォームのインターフェースとサービスは役割によって異なる場合があります
 データ製品の所有者 データ製品の所有者は、特定のドメインに対して成功したデータ製品を提供し普及させる責任を負っています
データ製品の成功は、採用と価値の提供、およびメッシュの広範なポリシーへの準拠性と他のデータ製品との相互運用性に依存します
彼らはプラットフォームを使用してデータ製品の最新動向を把握します
 データガバナンスメンバー データガバナンスは連邦構造に基づいているため、特定のデータガバナンスの役割はありません
それにもかかわらず、ガバナンスチームのメンバーは、メッシュ全体の最適で安全な運用を保証するための一連の責任を持っています
セキュリティや法的な問題についての専門知識を持つ専門家や特定の責任を持つデータ製品の所有者など、ガバナンスグループ内にはさまざまな役割があります
プラットフォームは彼らの集団的な旅をサポートし、個々の役割もサポートします
 データプラットフォームの製品オーナーデータプラットフォームのプレーンとそのサービスは、他のすべての役割を再生するユーザーのものです
データプラットフォームの製品オーナーは、プラットフォームが最高のユーザーエクスペリエンスとなるようにプラットフォームサービスを提供する責任を持っています
プラットフォームユーザーのニーズと制約を理解した上で、データプラットフォームの製品オーナーは、プラットフォームが提供するサービスを優先します
 データプラットフォーム開発者 プラットフォーム開発者は、データプラットフォームを構築・運用し、それを使用します
データプラットフォームのエクスペリエンスプレーンサービスで働くデータプラットフォーム開発者は、ユーティリティプレーンサービスのユーザーです
したがって、彼らのスキルセットと経験はユーティリティプレーンサービスの設計に重要です
この章では、データプロダクトの開発者とデータ科学者としてのデータプロダクトの消費者の2つの役割を使用して、プラットフォームの設計にアプローチする方法を示しています
 データプロダクトデベロッパーのジャーニー データプロダクトデベロッパーの主要なジャーニーの一つは、データプロダクトを作成し運用することです
これは、連続的な配信の原則「誰もが責任を持つ」という原則を受け入れた全体的な責任であり、長期的な責任です
データプロダクトの構築と運用は他人の問題ではありません
図10-2は、データプロダクトの作成プロセスの高レベルのフェーズを示しています
これらの段階は反復的で、継続的なフィードバックループを持っています
シンプルなデモンストレーションのため、私は旅を段階的に描いています
この旅は、データプロダクトの発想やソースの探索など、データプロダクトの発端に関連する活動から始まります
それは、実際のビルドとデータプロダクトの提供に進み、監視し、継続的に進化し、必要に応じて廃止されます
データプロダクトの開発の旅は、他の旅とも相互作用します
オペレーショナルシステムからデータを消費するソースに整列したデータプロダクトの場合、データプロダクトの開発者はソースアプリケーションの開発者と緊密に連携します
彼らは、アプリケーションがデータプロダクトへの入力として操作データを共有する方法を共同設計し、実装します
これらの人々は同じドメインチームに所属していることに注意してください
たとえば、プレイイベントデータプロダクトの開発は、生データを生成するプレイヤーアプリケーションとの緊密な協力が必要です
 同様に、データプロダクトが既知の分析アプリケーションにビルドされている場合、消費者に整列したデータプロダクトの場合、データプロダクトの開発者と分析アプリケーションの開発者との早い段階での共同作業が発生します
データプロダクトの開発者は、分析アプリケーションの目標やデータフィールドのニーズ、データプロダクトがサポートする必要がある保証について理解する必要があります
たとえば、DaffではMLを利用したプレイリストの推奨アプリケーションは、音楽プロファイルデータプロダクトからデータを必要とします
音楽プロファイルデータプロダクトの開発者は、プレイリストの推奨アプリケーションチームと密接に連携して、プロファイルの予想分類に合わせる必要があります
データプロダクト開発の高レベルなステージと、それらをサポートするために設計されたプラットフォームインターフェースを見てみましょう（図10-3）
 注：インターフェースの表記とフレーズは最終的な実装の代表ではありません
機能に焦点を当てるために、命令形の言語が選択されていますが、実際には、APIの宣言的な形式の方が望ましいです
たとえば、/deployインターフェースは、展開されたデータプロダクトのための宣言的リソースである/deploymentsとして非常によく実装されます
これらのインターフェースは模範的であり、操作の意味論を示していますが、実際の構文やインターフェースのメカニズム（CMD、UI、APIなど）とは異なります
 図10-3. プラットフォームを使用したデータプロダクト開発の旅の高レベルな例 インセプト、探索、ブートストラップ、およびソース データプロダクトは、そのデータを必要とする1つまたは複数の特定の分析ユースケースを発見することで存在するようになります
データプロダクト、特にソースに整列したものは、想像を超える多くの異なるユースケースがありますが、それでも我々はまず、現実のユースケースからデータプロダクトの創造（創造）を始める必要があります
消費者に整列したデータプロダクトの始まりは常に消費者との直接の協力を含みます
データプロダクトの開始時に、開発者はデータプロダクトの潜在的なソースを探し求める探索的なフェーズにいます
ソースは、上流データ製品、組織外のシステム、または組織内の運用システムとして発見されることがあります
探索中、開発者は上流ソースの発見情報を調査し、ソースの適切性、データの保証、ドキュメンテーション、データプロファイリング、既存のユーザーなどを評価することがあります
潜在的なソースが指名されたら、開発者は迅速にハローワールドデータ製品を立ち上げることができます
プラットフォームは必要なインフラストラクチャを提供します
これにより、データ製品の開発者はソースに接続し、シンセサイズされたデータ、隠蔽されたデータ、または実際のデータの形式でデータを消費することができます
この時点で、開発者はソースデータとデータ製品の変換ロジックを実験するために必要なすべてのリソースを持っています
これらのステップは、データ製品の構築の探索的ステップを構成します
迅速な発見、ソースの出力データへのアクセス、データ製品の素早い骨組み作り、インフラストラクチャの提供は、すべてインセプションおよびブートストラップフェーズの必要な要素です
たとえば、プレイリストチームは、異なる国での休日シーズンのために新しいプレイリストを生成する計画があります
彼らは地域の休日とその感情、またさまざまなトラックと関連する休日を関連付けるラベル付けされた音楽プロファイルのデータが必要です
彼らは同様の情報を持つメッシュ上の既存のデータ製品を検索して始めます
彼らはいくつかの音楽プロファイリングのソースを見つけ、それらのデータ製品に接続し、データの完全性と関連性を評価するためにデータを操作し始めます
彼らはシンプルな休日プレイリストデータ製品を作成して、音楽休日プロファイリングの最も単純な実装からどのようなプレイリストが生成できるかを確認します
表10-1は、このフェーズで使用されるプラットフォームAPIの概要を提供しています
表10-1
データ製品の開始フェーズをサポートするためのプラットフォームプレーンが提供する例のインターフェース 開発フェーズ プラットフォームプレーン プラットフォームインターフェース プラットフォームインターフェースの説明設立 | 探索 Mesh experience /search 適切なソースを見つけるために既存のデータ製品のメッシュを検索します
ソースの運用システム、ドメイン、およびデータの種類など、さまざまなパラメータに基づいて検索することができます
メッシュの経験/知識グラフ 関連するデータ製品のセマンティックモデルのメッシュを参照します
セマンティック関係をたどって、必要なデータソースを特定します

例：プレイリストチームは、見つけた休日に関連する音楽プロファイリング間のセマンティック関係を参照できます
Mesh experience /lineage メッシュ上の異なるデータ製品間の入出力データの系譜をたどり、データの起源とデータがどのような変換を経てきたかに基づいて必要なソースを特定します

例：プレイリストチームは、既存の音楽休日プロファイリングがいかに作成され、そのソースが何であるか、および何の変換が分類につながったかを調べることができます
これにより、既存の音楽休日プロファイリングの適切性を評価することができます
ブートストラップ|ソース データ製品の経験/{dp}/ディスカバリ ソースが特定されると、ドキュメント、データモデル、利用可能な出力ポートなど、すべてのデータ製品の発見可能性情報にアクセスできます
データ製品の経験/{dp}/観察 ソースが特定されると、データ製品の保証とメトリクス（リアルタイム）にアクセスできます
データがどのくらいの頻度でリリースされているか、最後のリリース日、データ品質メトリクスなどです

例：プレイリストチームが有望な休日音楽プロファイリングを見つけた場合、このデータがいかに頻繁に更新され、プラットフォーム上のすべての音楽と比較して完全性がどうかを評価する必要があります
この評価により、データへの信頼性が高まります
データ製品の経験/初期設定 ソースデータで実験を開始するために、このAPIはインフラストラクチャが利用可能な十分な骨組みのデータ製品をブートストラップします
これにより、ソースに接続し、データにアクセスし、データに変換を実行し、確認のために単一のアクセスモードで出力を提供することができます
データ製品の足場は、継続的な配信パイプライン、初期環境、データ処理クラスタの割り当て、およびデータ製品のインフラリソースの実行とアクセスのためのアカウントを確保します
これは、データ製品のライフサイクルの始まりを示します
メッシュ体験/登録 新しいデータ製品の初期化中、自動的にメッシュに登録されます
一意のグローバル識別子とアドレスが割り当てられ、データ製品がメッシュとガバナンスプロセスに表示されるようになります
データ製品の経験/接続 ソースが見つかると、データ製品はそれに接続することでソースにアクセスできるようになります
このステップでは、ソースを規制するアクセス制御ポリシーが検証されます
これにより、データ製品へのアクセス許可のリクエストがトリガーされる場合があります
ブートストラップ|ソース（続き） データ製品の経験/ {dp} / {output} / 
query
/ {dp} / {output} / 
subscribe ソースデータ製品の特定の出力ポートからデータを読み取ります
プルベースのクエリングモデルまたは変更のサブスクライブ方法で
ビルド、テスト、デプロイ、ラン データ製品開発者は、データ製品の構築、テスト、デプロイ、および運用に対して最後までの責任を負います
私が単に「ビルド、テスト、デプロイ、ラン」と呼んでいるこのステージは、データ製品開発者が成功したデータ製品のために必要なすべてのコンポーネントを提供するために実行する連続的かつ反復的な一連の活動です
このデータ製品は、自律的に発見可能、理解可能、信頼性があり、アドレス可能、相互運用可能かつ構成可能で、安全で、ネイティブにアクセスできるだけでなく、それ自体で価値のあるものである（「データ製品のベースラインの使いやすさ属性」を参照）
表10-2には、データ製品を提供するための高レベルのインターフェースがリストされています
表10-2
データ製品開発のサポートとしてプラットフォーム平面が提供する例のインターフェース 開発フェーズ プラットフォーム平面 プラットフォームインターフェース プラットフォームインターフェースの説明 ビルド データ製品の経験/ビルド データ製品のすべてのコンポーネントをコンパイル、検証、および構成し、展開可能なアーティファクトに組み立てます
これにより、データ製品をさまざまな展開環境で実行するために使用できるようになります

このフェーズで使用されるデータインフラストラクチャ平面のインターフェースについては、表10-3を参照してください
 テスト データ製品の経験/テスト データ製品のさまざまな側面をテストします
データ変換の機能、出力データの整合性、データのバージョニングとアップグレードプロセスの期待値、データプロファイル（期待される統計的特性）、バイアスのテストなどを行います

テストの機能は、異なる展開環境で提供されます
 デプロイ データ製品の経験/デプロイ データ製品の構築されたリビジョンを環境にデプロイします
これには、開発者のローカルマシン、ターゲットホスティング環境（特定のクラウドプロバイダなど）の開発サンドボックス、および事前プロダクションまたはプロダクション環境が含まれます
 実行 データ製品の経験/開始 
/停止 特定の環境でデータ製品インスタンスを実行または停止します
 ビルド/テスト/デプロイ/ラン データ製品の経験/ローカル-
ポリシー データ製品の主要なコンポーネントの1つは、データと機能を制御するポリシーです
ポリシーには、暗号化、アクセス、保持、ローカリティ、プライバシーなどが含まれます

プラットフォームは、データ製品の開発中にこれらのポリシーの設定と作成を容易にし、テスト中のポリシーの検証、およびデータへのアクセス時の実行を支援します
 ビルド/テスト/デプロイ/ラン メッシュ体験/グローバル-
ポリシー 多くの組織では、データ製品を制御するポリシーはグローバル（連邦）ガバナンスチームによって作成されます

したがって、プラットフォームはグローバルポリシーの作成とこれらのポリシーの適用をすべてのデータ製品に可能にします
データインフラストラクチャ平面での詳細をもう少し詳しく見てみましょう
データ製品の構築には多くの要素があります
データ製品開発者は、自律的なデータ製品を実装するためにこれらのすべての部分を構成またはコード化する必要があります
プラットフォームは、これらのコンポーネントのそれぞれをコーディングまたは構成するために必要なコストと労力を削減するための重要なツールです
図10-4は、データプロダクトの異なるコンポーネントの構築と実行を、下位レベルのプラットフォームインターフェース、インフラストラクチャプレーンAPIにマッピングしています
データプロダクトの開発者は、ほとんどの場合、データプロダクトのエクスペリエンスプレーンと対話するだけです
つまり、単一のデータ量子を構築、展開、テストします
データプロダクトのエクスペリエンスプレーンは、データプロダクトコンポーネントの実装をデータインフラストラクチャプレーンに委任します
開発者は、データインフラストラクチャプレーンが提供する技術に基づいて、データプロダクトのさまざまな側面やコンポーネントの理解、コーディング、設定を行わなければなりません
たとえば、データインフラストラクチャプレーンがサポートしている計算エンジンによって、データプロダクトの変換のコーディング方法が異なります
ただし、データプロダクトの開発者は、利用可能な技術を使用して変換のコーディングに焦点を当て、ビルド、展開、変換の実行、および他のすべてのコンポーネントの管理はデータプロダクトのエクスペリエンスプレーンに任せます
このプロセスでは、データプロダクトのエクスペリエンスプレーンは、変換の実行の詳細をデータインフラストラクチャプレーンに委任します
図10-4
データプロダクトの提供をサポートするデータインフラストラクチャプレーンインターフェースの例表10-3では、データプロダクトの開発中に利用されるデータインフラストラクチャプレーンの機能の一部をリストアップしています
表10-3
データプロダクトエクスペリエンスプレーンAPIをサポートするデータインフラストラクチャプレーンインターフェースの例プラットフォームインターフェースプラットフォームインターフェースの説明/ 入力ポートデータプロダクトの設計に応じて、イベントストリーミング入力、リモートデータベースクエリ、ファイル読み取りなど、さまざまなメカニズムを提供します
データが利用可能になったときに変換を実行するトリガーです
入力ポートのメカニズムは、新しいデータが利用可能になるにつれて消費の進捗状況を追跡します
/ 出力ポートデータプロダクトが提供するアクセスモードに応じて、ストリーム、ファイル、テーブルなど、さまざまなメカニズムを提供します
/ 変換データを変換するために必要なすべての計算を実行するためのさまざまな実行メカニズムを提供します
/ コンテナデータプロダクトとその必要なインフラリソースのライフサイクルを自律的なユニットとして管理します
/ コントロールアクセス制御、暗号化、プライバシー、保持、ローカリティなど、多様で進化しているポリシーの設定と実行を可能にするさまざまな技術を提供します
/ ストレージ入力ポート、出力ポート、および変換には、データとそれらのメタデータ（SLO、スキーマなど）への永続的および一時的なストレージへのアクセスが必要です
プラットフォームは、異なるタイプのストレージ（ブロブ、関係、グラフ、時系列など）へのアクセス、フェイルオーバー、リカバリ、バックアップ、リストアなどすべての操作機能を提供する必要があります
/ モデルデータの意味と構文のスキーマモデルを記述、共有、リンクするメカニズム
/ IDデータ製品をメッシュ上で一意に識別およびアドレス指定するメカニズム
メンテナンス、進化、廃止データプロダクトのメンテナンスと進化には、データプロダクト自体の連続的な変更が関与します
つまり、変換ロジックの改善、データモデルの進化、データへの新しいアクセスモードのサポート、ポリシーの拡充などが行われます
データプロダクトの変更により、データの処理と共有を途切れさせずに継続的にイテレーション（ビルド、テスト、デプロイ、実行）が行われます
プラットフォームは、データプロダクトのメンテナンス方法のオーバーヘッドを最小限に抑える必要があります
たとえば、新しいビルドの結果として意味とスキーマの変更がある場合、データプロダクトの開発者は単純にデータモデルの変更と新しいモデルに基づくデータの変換に焦点を当てます
さまざまなバージョンのスキーマ（意味と構文）の管理、データとの関連付け、および消費者によるアクセスの複雑さは、プラットフォームによって管理されます
一部の場合では、データプロダクトの進化は基礎となるインフラストラクチャリソースに影響を与えません
例えば、変換コードのバグ修正は、データ製品の再ビルドや再展開、生成されたデータの修正が必要です
他の場合では、進化は基盤となるリソースに影響を与えます
例えば、ストレージベンダーの切り替えに伴うデータ製品の新環境への移行では、ストレージリソースの再割り当てとデータ製品の基盤データの移行が必要です
各データ製品とメッシュ全体の運用状態のモニタリングは、このフェーズでプラットフォームが提供するもう一つの重要な能力です
メッシュの運用の優れた性能は、各データ製品の様々な側面のモニタリングに依存します：パフォーマンス、信頼性、SLA、計算ポリシーの有効性、リソース使用に基づいた運用コストのモニタリング
個々のデータ製品のモニタリングに加えて、メッシュの全体的な状態や洞察も収集し、モニタリングする必要があります
例えば、マスターデータ製品が検出されたときにアラートを検出することがあります
これらは多くのソースからデータを集約するデータ製品であり、使用されることでボトルネックになります
データ製品の寿命中には、メッシュレベルの管理コントロールを呼び出す必要がある場合があります
例えば、「忘れられる権利」は、メッシュレベルのグローバルコントロールを介してトリガーされ、個々のデータ製品の制御インターフェースを通じて委任され、基盤ストレージのデータ排除機能を介して実装されることがあります
データ製品は消滅することはあるのでしょうか？私は2つのシナリオを考えることができます
データ製品は、新しいデータ製品への移行の場合や、それが生成したデータレコードをすべて廃棄し、新しいレコードが生成されなくなる場合には廃止される可能性があります
いずれの場合でも、プラットフォームはデータ製品開発者がデータ製品を優雅に廃止できるようにするため、下流の消費者が新しいソースに徐々に移行するか、それ自体が退職することができます
過去のデータを使用している限り、データの量子は存在し続けますが、それ以上の変換は行われず、新しいデータは生成されません
休眠中のデータ製品は古いデータを引き続き提供し、ポリシーを強制しますが、完全に廃止されたデータ製品は単に存在しません
表10-4には、データ製品のメンテナンス、進化、退職のフェーズをサポートするためのプラットフォームインターフェースのいくつかが示されています
表10-4
データプラットフォームインターフェースの例　データ製品のメンテナンスをサポートするためのプラットフォーム平面のプラットフォームインターフェースの説明　データ製品の体験/{dp}/status データ製品のステータスを確認するためのもの
データ製品の体験/{dp}/logs
/{dp}/traces
/{dp}/metrics データ製品が実行時の可観測性情報（ログ、トレース、メトリクス）を出力するための仕組み
メッシュレイヤーのモニタリングサービスは、これらのメカニズムが提供するデータを利用します
データ製品の体験/{dp}/accesses データ製品へのアクセスの全てのログ
データ製品の体験平面/{dp}/controls 特定のデータの量子上で「忘れられる権利」などの高権限の管理コントロールを呼び出す機能
データ製品の体験/{dp}/cost データ製品の運用コストを追跡する機能
これはリソースの割り当てと使用状況に基づいて計算されることができる
データ製品の体験/migrate データ製品を新しい環境に移行する機能
データ製品のリビジョンの更新は、ビルドと展開の単なる関数です
メッシュ体験/monitor メッシュレベルでの複数のモニタリング機能、ログ、ステータス、コンプライアンスなど
メッシュ体験/notifications メッシュの異常を検出したときの通知とアラート機能
メッシュ体験/global-controls メッシュ上の複数のデータ製品に対して「忘れられる権利」などの高権限の管理コントロールを呼び出す機能
それでは、データ消費者の旅に注目して、プラットフォームインターフェースがどのように進化してそのようなペルソナをサポートするかを見てみましょう
データプロダクトの消費者ジャーニー

データ消費者のパーソナは、さまざまなスキルセットと責任を持つユーザーの広範な範囲を表しています
このセクションでは、既存のデータ製品を消費して機械学習モデルをトレーニングし、その後機械学習モデルをデータ製品として展開して推論を行い、新しいデータを生成する、データサイエンティストのパーソナに焦点を当てています
 例えば、Daffは、リスナーが週の初めにリスナー全員のためにキュレートされたプレイリストを生成するためにMLモデルを使用します
月曜日のプレイリストを生成するためのML推薦モデルは、既存のデータ製品のデータ（リスナープロファイル、リスナープレーリベント、彼らが好きな曲やそうでない曲に対する最近の反応、そして最近聞いたプレイリスト）を使用してトレーニングされます
 MLモデルがトレーニングされたら、月曜のプレイリストというデータ製品として展開されます
第7章で述べたように、データメッシュは緊密に統合された運用と分析のプレーンを作り出します
ドメイン指向の（マイクロ）サービスとデータ製品を相互に接続します
 2つのプレーン間の緊密な統合とフィードバックループにもかかわらず、データメッシュは各プレーンの責任と特性の違いを尊重し続けます
つまり、マイクロサービスはビジネスを実行するためのオンラインのリクエストやイベントに応答し、データ製品は一時データをキュレート、変換、共有してトレーニング用の機械学習モデルや洞察の生成などの下流の分析用に使用します
 2つのプレーンのエンティティの境界を明確にする試みにもかかわらず、MLモデルはこの境界をぼかします
 MLモデルはどちらのプレーンにも属することができます
 例えば、MLモデルはマイクロサービスとして展開され、最終ユーザーがリクエストを行う際にオンラインで推論を行うために使用されることがあります
 たとえば、新しいリスナーの登録中に、リスナーが提供した情報を元に分類器MLモデルが呼び出され、ユーザーの情報にプロファイル分類を追加します
 また、MLモデルはデータ製品の変換ロジックとして展開されることもあります
 たとえば、プレイリストの推奨MLモデルは、毎週推論を行い新しいプレイリストを生成するデータ製品として展開されます
このデータは、プレイリストをリスナーに表示する運用サービスに供給することができます
 図10-5はこの例を示しています
 図10-5
 MLモデルのバイモーダルな展開の例 このセクションでは、データサイエンティストがデータメッシュプラットフォームの機能との重なりを示すために、データ製品として展開されるMLモデルのジャーニーに焦点を当てています
 図10-6は、継続的にMLモデルをデータ製品として提供するための高レベルなジャーニーを示しています
この価値ストリームは、機械学習のための継続的なデリバリー（CD4ML）の実践に従って、反復可能なプロセスでモデルを仮説化、トレーニング、展開、監視、改善することで、迅速なフィードバックループで行われます
 3 図10-6
 モデル開発のジャーニーの例 データプラットフォームの観点からは、このジャーニーは以前に探求した「データプロダクト開発者のジャーニー」と非常に似ています
 次に、いくつかの違いがあります
 インセプト、エクスプロア、ブートストラップ、ソース この例では、MLモデル開発者は、最終的にモデルに埋め込まれるデータ製品のインセプトから旅を始めます
このデータ製品の発端には、既存のデータに基づいた賢明なアクションや意思決定のための仮説の策定が含まれます
 たとえば、多くのリスナーが再生して繰り返し再生（好き）するプレイリストをキュレーションして推奨することは可能でしょうか？ 仮説を検証するための探索を開始するために、MLモデル開発者は既存のデータ製品を探索し、ディスカバリAPIやサンプリング出力ポートデータを使用してソースを評価します
 このフェーズでは、プラットフォームのインターフェイスは以前に紹介したものと同じです
ここでの主な違いは、データサイエンティストが関心を持つ探索情報のタイプかもしれません
 データのバイアスやデータの公平性を評価する必要があります
これは、ソースデータプロダクトの出力ポートのサンプリングとプロファイリングによって評価することができます
ビルド、テスト、デプロイ、実行
モデルの構築とトレーニングは、上流データプロダクトの出力ポートを使用して行われます
トレーニング中、モデル開発者はトレーニングデータセットに変更が必要であると判断する場合、上流データプロダクトのデータを特徴に変換します
例えば、トレーニング特徴は、プレイリスト情報の一部と曲のプロファイルが結合されたものだけを含むことがあります
この場合、特徴を生成するためのデータ変換パイプライン自体がデータプロダクトとなります
例えば、月曜のプレイリスト特徴は他のどのデータプロダクトと同様に開発されるデータプロダクトとなります
このフェーズでは、モデル開発者はモデルの特定のリビジョンに結果をもたらすデータの追跡が必要です
ソース（トレーニングおよびテスト）データのバージョニングは、すべてのデータプロダクトによって公開される処理時間パラメータを使用できます
このパラメータはデータのリビジョンを示すものであり、データの処理および公開が行われた日時を示します
これにより、将来の再利用とモデルの繰り返し可能性のためにデータのコピーを保持する必要がなくなります
なぜなら、ソースデータプロダクトは常に過去のデータを取得することができるからです
第12章でこの詳細を説明します
モデルのトレーニングと追跡のプロセスは、MLプラットフォームによって処理されることがあります
これは、MLモデルのトレーニングパイプラインの特別なニーズを満たすための技術とツールのプラットフォームです
これらのサービスは、データプラットフォームと密接に連携して動作します
デプロイ時には、モデルを月曜のプレイリストデータプロダクトの変換コードとして実行できる形式にパッケージ化する必要があります
モデルの実行中、インフラストラクチャ面の変換エンジンは、MLモデルの実行に関する固有のニーズ（例：ターゲットとなるハードウェアでの実行 - GPUまたはTPU）を処理できます
保守、進化、引退モデルと同様に、開発者はモデルのパフォーマンスを監視し、出力データであるプレイリストが予想どおりに作成されているかを確認します
モデルの監視の特別なニーズの一つは、モデルの実効性を観察し、リスナーが行うアクションを監視することです
例えば、モデルの新しいリビジョンはプレイリストの再生時間の増加、リプレイの増加、リスナーの増加に結果をもたらすでしょうか？そのようなパラメータの監視には、操作面のメトリック監視の機能を利用することがあります
この場合、プレーヤーアプリケーションがこれらのメトリックを提供できます
まとめこの章から一つだけ心に留めていってほしいことは、プラットフォームのような単一のエンティティは存在しないということです
各ステップで必要なAPI、サービス、SDK、およびライブラリがあります
データプロダクト開発者や消費者などのプラットフォームユーザーの旅は、データメッシュプラットフォームサービスにとどまることはありません
彼らはしばしば、マイクロサービスからデータを消費するためのストリーミング能力を提供するAPIや、実験の追跡やGPUの利用など、MLモデルトレーニングの能力を持つAPIとの境界を越えます
したがって、プラットフォームを、ユーザーにとってシームレスなエクスペリエンスを提供するためのよく統合されたサービスのセットと見なしてください
2つ目のポイントとして、ヒューマンエクスペリエンスと機械効率の両方を最適化するために、プレーンの分離を尊重することです
例えば、データプロダクトのエクスペリエンス面を使用して、データ量子の単位との論理的なレベルでユーザーエクスペリエンスを最適化する
データの計算とストレージを分離するなど、マシンの最適化のために別個のデータインフラストラクチャプレーンを作成する
下位のプレーン最適化の決定は、開発者エクスペリエンスに支障をきたしてはなりませんし、逆もまた然りです
データメッシュプラットフォームを作成するための最初のステップで、最も小さな摩擦と最大の喜びと関与をもたらすような、データ製品開発者やデータ製品の消費者などの役割のための最適な旅の発見と設計から始めることが理想です
それが完了したら、彼らの旅のさまざまな段階に適したツールを見つけるために、自由に検索を続けてください