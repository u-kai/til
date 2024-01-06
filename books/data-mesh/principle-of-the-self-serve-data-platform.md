So far I have offered two fundamental shifts toward data mesh: a distributed data architecture and ownership model oriented around business domains, and data shared as a usable and valuable product. Over time, these two seemingly simple and rather intuitive shifts can have undesired consequences: duplication of efforts in each domain, increased cost of operation, and likely large-scale inconsistencies and incompatibilities across domains. Expecting domain engineering teams to own and share analytical data as a product, in addition to building applications and maintaining digital products, raises legitimate concerns for both the practitioners and their leaders. The concerns that I often hear from leaders, at this point in the conversation, include: “How am I going to manage the cost of operating the domain data products, if every domain needs to build and own its own data?” “How do I hire the data engineers, who are already hard to find, to staff in every domain?” “This seems like a lot of overengineering and duplicate effort in each team.” “What technology do I buy to provide all the data product usability characteristics?” “How do I enforce governance in a distributed fashion to avoid chaos?” “What about copied data—how do I manage that?” And so on. Similarly, domain engineering teams and practitioners voice concerns such as, “How can we extend the responsibility of our team to not only build applications to run the business, but also share data?” Addressing these questions is the reason that data mesh’s third principle, self-serve data infrastructure as a platform, exists. It’s not that there is any shortage of data and analytics platforms, but we need to make changes to them so they can scale out sharing, accessing, and using analytical data, in a decentralized manner, for a new population of generalist technologists. This is the key differentiation of a data mesh platform. Figure 4-1 depicts the extraction of domain-agnostic capabilities out of each domain and moving to a self-serve infrastructure as a platform. The platform is built and maintained by a dedicated platform team. Figure 4-1. Extracting and harvesting domain agnostic infrastructure into a separate data platform In this chapter, I apply platform thinking to the underlying infrastructure capabilities to clarify what we mean by the term platform in the context of data mesh. Then, I share the unique characteristics of data mesh’s underlying platform. Later chapters, such as Chapters 9 and 10, will go into further detail about the platform’s capabilities and how to approach its design. For now, let’s discuss how data mesh’s underlying platform is different from many solutions we have today. Note In this chapter, I use the phrase data mesh platform as shorthand for a set of underlying data infrastructure capabilities. A singular form of the term platform does not mean a single solution or a single vendor with tightly integrated features. It’s merely a placeholder for a set of technologies that one can use to achieve the objectives mentioned in “Data Mesh Platform Thinking”, a set of technologies that are independent and yet play nicely together. Data Mesh Platform: Compare and Contrast There is a large body of technology solutions that fall into the category of data infrastructure and are often posed as a platform. Here is a small sample of the existing platform capabilities: Analytical data storage in the form of a lake, warehouse, or lakehouse Data processing frameworks and computation engines to process data in batch and streaming modes Data querying languages, based on two modes of computational data flow programming or algebraic SQL-like statements Data catalog solutions to enable data governance as well as discovery of all data across an organization Pipeline workflow management, orchestrating complex data pipeline tasks or ML model deployment workflows Many of these capabilities are still needed to enable a data mesh implementation. However, there is a shift in approach and the objectives of a data mesh platform. Let’s do a quick compare and contrast. Figure 4-2 shows a set of unique characteristics of a data mesh platform in comparison to the existing ones. Note that the data mesh platform can utilize existing technologies yet offer these unique characteristics. The following section clarifies how data mesh works toward building self-serve platforms further. Figure 4-2. Data mesh platform’s differentiating characteristics Serving Autonomous Domain-Oriented Teams The main responsibility of the data mesh platform is to enable existing or new domain engineering teams with the new and embedded responsibilities of building, sharing, and using data products end to end; capturing data from operational systems and other sources; and transforming and sharing the data as a product with the end data users. The platform must allow teams to do this in an autonomous way without any dependence on centralized data teams or intermediaries. Many existing vendor technologies are built with an assumption of a centralized data team, capturing and sharing data for all domains. The assumptions around this centralized control have deep technical consequences such as: Cost is estimated and managed monolithically and not per isolated domain resources. Security and privacy management assumes physical resources are shared under the same account and don’t scale to an isolated security context per data product. A central pipeline (DAG) orchestration assumes management of all data pipelines centrally—with a central pipeline configuration repository and a central monitoring portal. This is in conflict with independent pipelines, each small and allocated to a data product implementation. These are a few examples to demonstrate how existing technologies get in the way of domain teams acting autonomously. Managing Autonomous and Interoperable Data Products Data mesh puts a new construct, a domain-oriented data product, at the center of its approach. This is a new architectural construct that autonomously delivers value. It encodes all the behavior and data needed to provide discoverable, usable, trustworthy, and secure data to its end data users. Data products share data with each other and are interconnected in a mesh. Data mesh platforms must work with this new construct and must support managing its autonomous life cycle and all its constituents. This platform characteristic is different from the existing platforms today that manage behavior, e.g., data processing pipelines, data and its metadata, and policy that governs the data as independent pieces. However, it is possible to create the new data product management abstraction on top of existing technologies, but it is not very elegant. A Continuous Platform of Operational and Analytical Capabilities The principle of domain ownership demands a platform that enables autonomous domain teams to manage data end to end. This closes the gap organizationally between the operational plane and the analytical plane. Hence, a data mesh platform must be able to provide a more connected experience. Whether the team is building and running an application or sharing and using data products for analytical use cases, the team’s experience should be connected and seamless. For the platform to be successfully adopted with existing domain technology teams, it must remove the barriers to adoption, the schism between the operational and analytical stacks. The data mesh platform must close the gap between analytical and operational technologies. It must find ways to get them to work seamlessly together, in a way that is natural to a cross-functional domain-oriented data and application team. For example, today the computation fabric running data processing pipelines such as Spark are managed on a different clustering architecture, away and often disconnected from the computation fabric that runs operational services, such as Kubernetes. In order to create data products that collaborate closely with their corresponding microservice, i.e., source-aligned data products, we need a closer integration of the computation fabrics. I have worked with very few organizations that have been running both computation engines on the same computation fabric. Due to inherent differences between the two planes, there are many areas of the platform where the technology for operational and analytical systems must remain different. For example, consider the case of tracing for debugging and audit purposes. The operational systems use the OpenTelemetry standards for tracing of (API) calls across distributed applications, in a tree-like structure. On the other hand, data processing workloads use OpenLineage to trace the lineage of data across distributed data pipelines. There are enough differences between the two planes to mind their gap. However, it is important that these two standards integrate nicely. After all, in many cases, the journey of a piece of data starts from an application call in response to a user action. Designed for a Generalist Majority Another barrier to the adoption of data platforms today is the level of proprietary specialization that each technology vendor assumes—the jargon and the vendor-specific knowledge. This has led to the creation of scarce specialized roles such as data engineers. In my opinion there are a few reasons for this unscalable specialization: lack of (de facto) standards and conventions, lack of incentives for technology interoperability, and lack of incentive to make products super simple to use. I believe this is the residue of the big monolithic platform mentality that a single vendor can provide soup to nuts functionality to store your data on their platform and attach their additional services to keep the data there and its processing under their control. A data mesh platform must break this pattern and start with the definition of a set of open conventions that promote interoperability between different technologies and reduce the number of proprietary languages and experiences one specialist must learn to generate value from data. Incentivizing and enabling generalist developers with experiences, languages, and APIs that are easy to learn is a starting point to lower the cognitive load of generalist developers. To scale out data-driven development to the larger population of practitioners, data mesh platforms must stay relevant to generalist technologists. They must move to the background, fit naturally into the native tools and programming languages generalists use, and get out of their way. Needless to say, this should be achieved without compromising on the software engineering practices that result in sustainable solutions. For example, many low-code or no-code platforms promise to work with data, but compromise on testing, versioning, modularity, and other techniques. Over time they become unmaintainable. Note The phrase generalist technologist (experts) refers to the population of technologists who are often referred to as T-shaped or Paint Drip people. These are developers experienced in a broad spectrum of software engineering who, at different points in time, focus and gain deep knowledge in one or two areas. The point is that it is possible to go deep in one or two areas while exploring many others. They contrast with specialists, who only have expertise in one specific area; their focus on specialization doesn’t allow them to explore a diverse spectrum. In my mind, future generalists will be able to work with data and create and share data through data products, or use them for feature engineering and ML training when the model has already been developed by specialist data scientists. Essentially, they use AI as a service. As of now, the majority of data work requires specialization and requires a large amount of effort to gain expertise over a long period of time. This inhibits the entry of generalist technologists and has led to a scarcity of data specialists. Favoring Decentralized Technologies Another common characteristic of existing platforms is the centralization of control. Examples include centralized pipeline orchestration tools, centralized catalogs, centralized warehouse schema, centralized allocation of compute/storage resources, and so on. The reason for data mesh’s focus on decentralization through domain ownership is to avoid organizational synchronization and bottlenecks that ultimately slow down the speed of change. Though on the surface this is an organizational concern, the underlying technology and architecture directly influence organizational communication and design. A monolithic or centralized technology solution leads to centralized points of control and teams. Data mesh platforms need to consider the decentralization of organizations in data sharing, control, and governance at the heart of their design. They inspect every centralized aspect of the design that can result in lockstep team synchronization, centralization of control, and tight coupling between autonomous teams. Having said that, there are many aspects of infrastructure that need to be centrally managed to reduce the unnecessary tasks that each domain team performs in sharing and using data, e.g., setting up data processing compute clusters. This is where an effective self-serve platform shines, centrally managing underlying resources while allowing independent teams to achieve their outcomes end to end, without tight dependencies to other teams. Domain Agnostic Data mesh creates a clear delineation of responsibility between domain teams—who focus on creating business-oriented products, services that are ideally data-driven, and data products—and the platform teams who focus on technical enablers for the domains. This is different from the existing delineation of responsibility where the data team is often responsible for amalgamation of domain-specific data for analytical usage, as well as the underlying technical infrastructure. This delineation of responsibility needs to be reflected in the platform capabilities. The platform must strike a balance between providing domain-agnostic capabilities, while enabling domain-specific data modeling, processing, and sharing across the organization. This demands a deep understanding of the data developers and the application of product thinking to the platform. Data Mesh Platform Thinking Platform: raised level surface on which people or things can stand. ​—​Oxford Languages The word platform is one of the most commonly used phrases in our everyday technical jargon and is sprinkled all over organizations’ technical strategies. It’s commonly used, yet hard to define and subject to interpretation. To ground our understanding of the platform in the context of data mesh I draw from the work of a few of my trustworthy sources: A digital platform is a foundation of self-service APIs, tools, services, knowledge and support which are arranged as a compelling internal product. Autonomous delivery teams can make use of the platform to deliver product features at a higher pace, with reduced coordination. ​—​Evan Bottcher, “What I Talk About When I Talk About Platforms” The purpose of a platform team is to enable stream-aligned teams to deliver work with substantial autonomy. The stream-aligned team maintains full ownership of building, running, and fixing their application in production. The platform team provides internal services to reduce the cognitive load that would be required from stream-aligned teams to develop these underlying services. The platform simplifies otherwise complex technology and reduces cognitive load for teams that use it. ​—​Matthew Skelton and Manuel Pais, Team Topologies Platforms are designed one interaction at a time. Thus, the design of every platform should start with the design of the core interaction that it enables between producers and consumers. The core interaction is the single most important form of activity that takes place on a platform—the exchange of value that attracts most users to the platform in the first place.1 ​—​Geoffrey G. Parker et al., Platform Revolution) A platform has a few key objectives that I like to take away and apply to data mesh: Enable autonomous teams to get value from data A common characteristic that we see is the ability to enable teams who use the platform to complete their work and achieve their outcomes with a sense of autonomy and without requiring another team to get engaged directly in their workflow, e.g., through backlog dependencies. In the context of data mesh, enabling domain teams with new responsibilities of sharing analytical data, or using analytical data for building ML-based products, in an autonomous way, is a key objective of a data mesh platform. The ability to use the platform capabilities through self-serve APIs is critical to enable autonomy. Exchange value with autonomous and interoperable data products Another key aspect of a platform is to intentionally design what value is being exchanged and how. In the case of data mesh, data products are the unit of value exchange, between data users and data providers. A data mesh platform must build in the frictionless exchange of data products as a unit of value in its design. Accelerate exchange of value by lowering the cognitive load In order to simplify and accelerate the work of domain teams in delivering value, platforms must hide technical and foundational complexity. This lowers the cognitive load of the domain teams to focus on what matters; in the case of data mesh, this is creating and sharing data products. Scale out data sharing Data mesh is a solution offered to solve the problem of organizational scale in getting value from their data. Hence, the design of the platform must cater for scale: sharing data across main domains within the organization, as well as across boundaries of trust outside of the organization in the wider network of partners. One of the blockers to this scale is the lack of interoperability of data sharing, securely, across multiple platforms. A data mesh platform must design for interoperability with other platforms to share data products. Support a culture of embedded innovation A data mesh platform supports a culture of innovation by removing activities that are not directly contributing to the innovation, by making it really easy to find data, capture insights, and use data for ML model development. Figure 4-3 depicts these objectives applied to an ecosystem of domain teams sharing and using data products. Figure 4-3. Objectives of the data mesh platform Now, let’s talk about how a data mesh platform achieves these objectives. Enable Autonomous Teams to Get Value from Data In designing the platform, it is helpful to consider the roles of platform users and their journey in sharing and using data products. The platform can then focus on how to create a frictionless experience for each journey. For example, let’s consider two of the main personas of the data mesh ecosystem: data product developers and data product users. Of course, each of those personas includes a spectrum of people with different skill sets, but for this conversation we can focus on the aspects of their journeys that are common across this spectrum. There are other roles such as data product owner whose journey is as important in achieving the outcome of creating successful data products; favoring brevity, I leave them out of this example. Enable data product developers The delivery journey of a data product developer involves developing a data product; testing, deploying, monitoring, and updating it; and maintaining its integrity and security—with continuous delivery in mind. In short, the developer is managing the life cycle of a data product, working with its code, data, and policies as one unit. As you can imagine, there is a fair bit of infrastructure that needs to be provisioned to manage this life cycle. Provisioning and managing the underlying infrastructure for life cycle management of a data product requires specialized knowledge of today’s tooling and is difficult to replicate in each domain. Hence, the data mesh platform must implement all necessary capabilities allowing a data product developer to build, test, deploy, secure, and maintain a data product without worrying about the underlying infrastructure resource provisioning. It must enable all domain-agnostic and cross-functional capabilities. Ultimately, the platform must enable the data product developer to just focus on the domain-specific aspects of data product development: Transformation code, the domain-specific logic that generates and maintains the data Build-time tests to verify and maintain the domain’s data integrity Runtime tests to continuously monitor that the data product meets its quality guarantees Developing a data product’s metadata such as its schema, documentation, etc. Declaration of the required infrastructure resources The rest must be taken care of by the data mesh platform, for example, infrastructure provisioning—storage, accounts, compute, etc. A self-serve approach exposes a set of platform APIs for the data product developer to declare their infrastructure needs and let the platform take care of the rest. This is discussed in detail in Chapter 14. Enable data product users Data users’ journey—whether analyzing data to create insights or developing machine learning models—starts with discovering the data. Once the data is discovered, they need to get access to it, and then understand it and deep dive to explore it further. If the data has proven to be suitable, they will continue to use it. Using the data is not limited to one-time access; the consumers continue receiving and processing new data to keep their machine learning models or insights up to date. The data mesh platform builds the underlying mechanisms that facilitate such a journey and provides the capabilities needed for data product consumers to get their job done without friction. For the platform to enable this journey, autonomously, it must reduce the need for manual intervention. For example, it must remove the need to chase the team that created the data or the governance team to justify and get access to the data. The platform automates the process that facilitates requests for access and grants access based on automated evaluation of the consumer. Exchange Value with Autonomous and Interoperable Data Products An interesting lens on the data mesh platform is to view it as a multisided platform—one that creates value primarily by enabling direct interactions between two (or more) distinct parties. In the case of data mesh, those parties are data product developers, data product owners, and data product users. This particular lens can be a source of unbounded creativity for building a platform whose success is measured directly by exchanging value, i.e., data products. The value can be exchanged on the mesh, between data products, or at the edge of the mesh, between the end products, such as an ML model, a report, or a dashboard, and the data products. The mesh essentially becomes the organizational data marketplace. This particular data mesh platform characteristic can be a catalyst for a culture change within the organization, promoting sharing to the next level. As discussed in the previous section, an important aspect of exchanging value is to be able to do that autonomously, without the platform getting in the way. For data product developers, this means being able to create and serve their data products without the constant need for hand-holding or dependency on the platform team. Create higher-order value by composing data products The exchange of value goes beyond using a single data product and often extends to the composition of multiple data products. For example, the interesting insights about Daff’s listeners are generated by cross-correlating their behavior while listening to music, the artists they follow, their demographic, their interactions with social media, the influence of their friends network, and the cultural events that surround them. These are multiple data products and need to be correlated and composed into a matrix of features. The platform makes data product compatibility possible. For example, platforms enable data product linking—when one data product uses data and data types (schema) from another data product. For this to be seamlessly possible, the platform provides a standardized and simple way of identifying data products, addressing data products, connecting to data products, reading data from data products, etc. Such simple platform functions create a mesh of heterogeneous domains with homogeneous interfaces. I will cover this in Chapter 13. Accelerate Exchange of Value by Lowering the Cognitive Load Cognitive load was first introduced in the field of cognitive science as the amount of working memory needed to hold temporary information to solve a problem or learn.2 There are multiple factors influencing the cognitive load, such as the intrinsic complexity of the topic at hand or how the task or information is presented. Platforms are increasingly considered a way of reducing the cognitive load of developers to get their job done. They do this by hiding the amount of detail and information presented to the developer: abstracting complexity. As a data product developer, I should be able to express what my domain-agnostic wishes are without describing exactly how to implement them. For example, as a developer I should be able to declare the structure of my data, its retention period, its potential size, and its confidentiality class and leave it to the platform to create the data structures, provision the storage, perform automatic encryption, manage encryption keys, automatically rotate keys, etc. This is domain-agnostic complexity that as a data developer or user I should not be exposed to. There are many techniques for abstracting complexity without sacrificing configurability. The following two methods are commonly applied. Abstract complexity through declarative modeling Over the last few years, operational platforms such as container orchestrators, e.g., Kubernetes, or infrastructure provisioning tools, e.g., Terraform, have established a new model for abstracting complexity through declarative modeling of the target state. This is in contrast with other methods such as using imperative instructions to command how to build the target state. Essentially, the former focuses on the what, and the latter focuses on the how. This approach has been widely successful in making the life of a developer much simpler. In many scenarios declarative modeling hits limitations very quickly. For example, defining the data transformation logic through declarations reaches a diminishing return as soon as the logic gets complex. However, systems that can be described through their state, such as provisioned infrastructure, lend themselves well to a declarative style. This is also true about the data mesh infrastructure as a platform. The target state of infrastructure to manage the life cycle of a data product can be defined declaratively. Abstract complexity through automation Removing the human intervention and manual steps from the data product developer journey through automation is another way to reduce complexity, particularly complexity arising from manual errors through the process. Opportunities to automate aspects of a data mesh implementation are ubiquitous. The provisioning of the underlying data infrastructure itself can be automated using infrastructure as code3 techniques. Additionally, many actions in the data value stream, from production to consumption, can be automated. For example, today the data certification or verification approval process is often done manually. This is an area of immense opportunity for automation. The platform can automate verifying the integrity of data, apply statistical methods in testing the nature of the data, and even use machine learning to discover unexpected outliers. Such automation removes complexity from the data verification process. Scale Out Data Sharing One issue I’ve noticed in the existing big data technology landscape is the lack of standards for interoperable solutions that lead to data sharing at scale, for example, lack of a unified model for authentication and authorization when accessing data, absence of standards for expressing and transmitting privacy rights with data, and lack of standards in presenting temporality aspects of data. These missing standards inhibit scaling the network of usable data beyond the boundaries of organizational trust. Most importantly, the data technology landscape is missing the Unix philosophy: This is the Unix philosophy: Write programs that do one thing and do it well. Write programs to work together… ​—​Doug McIlroy I think we got incredibly lucky with very special people (McIlroy, Ritchie, Thompson, and others) seeding the culture, the philosophy, and the way of building software in the operational world. That’s why we have managed to build powerfully scaled and complex systems through loose integration of simple and small services. For some reason, we have abandoned this philosophy when it comes to big data systems, perhaps because of those early assumptions (see “Characteristics of Analytical Data Architecture”) that seeded the culture. Perhaps, because at some point we decided to separate data (the body) from its code (the soul), which led to establishing a different philosophy around it. If a data mesh platform wants to realistically scale out sharing data, within and beyond the bounds of an organization, it must wholeheartedly embrace the Unix philosophy and yet adapt it to the unique needs of data management and data sharing. It must design the platform as a set of interoperable services that can be implemented by different vendors with different implementations yet play nicely with the rest of the platform services. Take observability as an example of a capability that the platform provides—the ability to monitor the behavior of all data products on the mesh and detect any disruptions, errors, and undesirable access, and notify the relevant teams to recover their data products. For observability to work, there are multiple platform services that need to cooperate: the data products emitting and logging information about their operation; the service that captures the emitted logs and metrics and provides a holistic mesh view; the services that search, analyze, and detect anomalies and errors within those logs; and the services that notify the developers when things go wrong. To build this under the Unix philosophy we need to be able to pick and choose these services and connect them together. The key in simple integration of these services is interoperability,4 a common language and APIs by which the logs and metrics are expressed and shared. Without such a standard, we fall back to a single monolithic (but well-integrated) solution that constrains access to data to a single hosting environment. We fail to share and observe data across environments. Support a Culture of Embedded Innovation To date, continuous innovation must arguably be one of the core competencies of any business. Eric Ries introduced the Lean Startup5 to demonstrate how to scientifically innovate through short and rapid cycles of build-measure-learn. The concept has since been applied to the larger enterprise through Lean Enterprise6—a scaled innovation methodology. The point is that to grow a culture of innovation—a culture of rapidly building, testing, and refining ideas—we need an environment that frees its people from unnecessary work and accidental complexity and friction and allow them to experiment. The data mesh platform removes unnecessary manual work, hides complexity, and streamlines the workflows of data product developers and users, to free them to innovate using data. A simple litmus test to assess how effective a data mesh platform is in doing that is to measure how long it takes for a team to dream up a data-driven experiment and get to use the required data to run the experiment. The shorter the time, the more mature the data mesh platform has become. Another key point is: who is empowered to do the experiments? The data mesh platform supports a domain team to innovate and perform data-driven experiments. The data-driven innovations are no longer exclusive to the central data team. They must be embedded into each domain team in developing their services, products, or processes. Transition to a Self-Serve Data Mesh Platform So far, I have talked about the key differences between existing data platforms and data mesh and covered the main objectives of the data mesh platform. Here, I’d like to leave you with a few actions you can take in transitioning to your data mesh platform. Design the APIs and Protocols First When you begin your platform journey, whether you are buying, building, or very likely both, start with selecting and designing the interfaces that the platform exposes to its users. The interfaces might be programmatic APIs. There might be command-line or graphic interfaces. Either way, decide on interfaces first and then the implementation of those through various technologies. This approach is well-adopted by many cloud offerings. For example, cloud blob storage providers expose REST APIs7 to post, get, or delete objects. You can apply this to all capabilities of your platform. In addition to the APIs, decide on the communication protocols and standards that enable interoperability. Taking inspirations from internet—the one example of a massively distributed architecture—decide on the narrow waist8 protocols. For example, decide on the protocols governing how data products express their semantic, in what format they encode their time-variant data, what query languages each support, what SLOs each guarantee, and so on. Prepare for Generalist Adoption I discussed earlier that a data mesh platform must be designed for the generalist majority (“Designed for a Generalist Majority”). Many organizations today are struggling to find data specialists such as data engineers, while there is a large population of generalist developers who are eager to work with data. The fragmented, walled, and highly specialized world of big data technologies have created an equally siloed fragment of hyper-specialized data technologists. In your evaluation of platform technologies, favor the ones that fit better with a natural style of programming known to many developers. For example, if you are choosing a pipeline orchestration tool, pick the ones that lend themselves to simple programming of Python functions—something familiar to a generalist developer—rather than the ones that try to create yet another domain-specific language (DSL) in YAML or XML with esoteric notations. In reality, there will be a spectrum of data products in terms of their complexity, and a spectrum of data product developers in terms of their level of specializations. The platform must satisfy this spectrum to mobilize data product delivery at scale. In either case, the need for applying evergreen engineering practices to build resilient and maintainable data products remains necessary. Do an Inventory and Simplify The separation of the analytical data plane and the operational plane has left us with two disjointed technology stacks, one dealing with analytical data and the other for building and running applications and services. As data products become integrated and embedded within the operational world, there is an opportunity to converge the two platforms and remove duplicates. In the last few years the industry has experienced an overinvestment in technologies that are marketed as data solutions. In many cases their operational counterparts are perfectly suitable to do the job. For example, I have seen a new class of continuous integration and continuous delivery (CI/CD) tooling marketed under DataOps. Evaluating these tools more closely, they hardly offer any differentiating capability that the existing CI/CD engines can’t offer. When you get started, take an inventory of platform services that your organization has adopted and look for opportunities to simplify. I do hope that the data mesh platform is a catalyst in simplification of the technology landscape and closer collaboration between operational and analytical platforms. Create Higher-Level APIs to Manage Data Products The data mesh platform must introduce a new set of APIs to manage data products as a new abstraction (“Managing Autonomous and Interoperable Data Products”). While many data platforms, such as the services you get from your cloud providers, include lower-level utility APIs—storage, catalog, compute—the data mesh platform must introduce a higher level of APIs that deal with a data product as an object. For example, consider APIs to create a data product, discover a data product, connect to a data product, read from a data product, secure a data product, and so on. See Chapter 9 for the logical blueprint of a data product. When establishing your data mesh platform, start with high-level APIs that work with the abstraction of a data product. Build Experiences, Not Mechanisms I have come across numerous platform building/buying situations, where the articulation of the platform is anchored in mechanisms it includes, as opposed to experiences it enables. This approach in defining the platform often leads to bloated platform development and adoption of overambitious and overpriced technologies. Take data cataloging as an example. Almost every platform I’ve come across has a data catalog on its list of mechanisms, which leads to the purchase of a data catalog product with the longest list of features, and then overfitting the team’s workflows to fit the catalog’s inner workings. This process often takes months. In contrast, your platform can start with the articulation of the single experience of discovering data products. Then, build or buy the simplest tools and mechanisms that enable this experience. Then rinse, repeat, and refactor for the next experience. Begin with the Simplest Foundation, Then Harvest to Evolve Given the length of this chapter discussing the objectives and unique characteristics of a data mesh platform, you might be wondering, “Can I even begin to adopt data mesh today, or should I wait some time to build the platform first?” The answer is to begin adopting a data mesh strategy today, even if you don’t have a data mesh platform. You can begin with the simplest possible foundation. Your smallest possible foundation framework is very likely composed of the data technologies that you have already adopted, especially if you are already operating analytics on the cloud. The bottom-layer utilities that you can use as the foundation include the typical storage technologies, data processing frameworks, federated query engines, and so on. As the number of data products grows, standards are developed, and common ways of approaching similar problems across data products are discovered. Then you will continue to evolve the platform as a harvested framework by collecting common capabilities across data products and domain teams. Remember that the data mesh platform itself is a product. It’s an internal product—though built from many different tools and services from multiple vendors. The product users are the internal teams. It requires technical product ownership, long-term planning, and long-term maintenance. Though it continues to evolve and goes through evolutionary growth, its life begins today as a minimum viable product (MVP).9 Recap Data mesh’s principle of a self-serve platform comes to the rescue to lower the cognitive load that the other two principles impose on the existing domain engineering teams: own your analytical data and share it as a product. It shares common capabilities with the existing data platforms: providing access to polyglot storage, data processing engines, query engines, streaming, etc. However, it differentiates from the existing platforms in its users: autonomous domain teams made up primarily of generalist technologists. It manages a higher-level construct of a data product encapsulating data, metadata, code, and policy as one unit. Its purpose is to give domain teams superpowers, by hiding low-level complexity behind simpler abstractions and removing friction from their journeys in achieving their outcome of exchanging data products as a unit of value. And ultimately it frees up the teams to innovate with data. To scale out data sharing, beyond a single deployment environment or organizational unit or company, it favors decentralized solutions that are interoperable. I will continue our deep dive into the platform in Chapter 10 and talk about specific services a data mesh platform could offer.

これまでに、私はデータメッシュへの2つの基本的な変革を提案してきました
まず、ビジネスドメインを中心に配置された分散データアーキテクチャと所有モデル、そして使いやすく価値のある製品として共有されるデータです
しかし、時間の経過とともに、これらの2つの簡単で直感的な変化には望ましくない結果が生じることがあります
各ドメインでの作業の重複、運用コストの増加、およびドメイン間での大規模な不整合や非互換性が生じる可能性があります
ドメインエンジニアリングチームに、アプリケーションの構築とデジタル製品の維持だけでなく、分析データを製品として所有して共有することを期待することは、実践者と彼らのリーダーの両方にとって合理的な懸念です
この会話の段階で、リーダーからしばしば聞かれる懸念事項は次のようなものです
「各ドメインが独自のデータを構築・所有する必要がある場合、ドメインのデータ製品の運用コストはどのように管理すればよいですか？」
「すでに見つけるのが難しいデータエンジニアをどのように雇えばよいですか？」
「これは各チームでのオーバーエンジニアリングと重複作業が非常に多いように思えます
」
「すべてのデータ製品の使いやすさの特性を提供するためにどの技術を購入すればよいですか？」
「混乱を避けるために分散形式でガバナンスを施行するにはどうすればよいですか？」
「コピーされたデータについてどう管理すればよいですか？」などです
同様に、ドメインエンジニアリングチームと実践者からは、「ビジネスを実行するためのアプリケーションを構築するだけでなく、データも共有するためにチームの責任を拡大する方法はありますか？」という懸念もあります
これらの質問に答えるために、データメッシュの3番目の原則である、自己提供型のデータ基盤をプラットフォームとして捉えることがあります
これは、データと分析データの共有、アクセス、使用のスケーラビリティを分散型の汎用技術者の新しい人口向けに変えるために、既存のデータこれまでに、データメッシュに向けた2つの基本的な変更を提案してきました
ビジネスドメインを中心とした分散データアーキテクチャと所有モデル、そして使いやすく価値のあるデータとして共有される製品であるデータ
しかし、これらの2つのシンプルで直感的な変更は、時間の経過と共に望ましくない結果をもたらす可能性があります
各ドメインでの作業の重複、運用コストの増加、およびおそらくドメイン間の大規模な不整合や非互換性です
分析データを製品として所有し共有することをドメインエンジニアリングチームに期待することは、アプリケーションの構築とデジタル製品のメンテナンスに加えて、実践者と彼らのリーダーの両方にとって合理的な懸念を引き起こします
この会話のこの時点でリーダーからよく聞かれる懸念事項には、以下が含まれます
「ドメインごとにデータを構築し、所有する必要がある場合、ドメインデータ製品の運用コストをどのように管理すればよいですか？」「データエンジニアを採用するにはどうすればよいですか？彼らはすでに見つけにくい存在ですが、各ドメインに配置する必要がありますか？」「これは各チームでの過度のエンジニアリングと重複した取り組みのように思えますが、どのような技術を購入してすべてのデータ製品の使いやすさの特性を提供すればよいですか？」「混乱を避けるために分散型でガバナンスを強制するにはどうすればよいですか？」「コピーされたデータはどのように管理すればよいですか？」などです
同様に、ドメインエンジニアリングチームと実践者も次のような懸念を表明しています
「ビジネスを運営するためのアプリケーションを構築するだけでなく、データも共有するためにチームの責任をどのように拡張することができますか？」これらの質問に対応するために、データメッシュの第三の原則であるセルフサービスのデータインフラストラクチャとしてのプラットフォームが存在しています
データや分析データの共有、アクセス、使用を分散型に拡大するために、新しい種類の汎用技術者向けに変化を与えるために、データメッシュプラットフォームは変更を加える必要があります
これがデータメッシュプラットフォームの鍵となる違いです
図4-1は、各ドメインからドメインに依存しない機能を抽出し、セルフサービスのインフラストラクチャとしてのプラットフォームに移動する様子を示しています
このプラットフォームは、専任のプラットフォームチームが構築および維持しています
図4-1. ドメイン非依存のインフラストラクチャの抽出と収穫-別のデータプラットフォームへ
この章では、データメッシュのコンテキストでプラットフォームという用語がどのように意味されるのかを明確にするために、プラットフォーム思考を基盤となるインフラストラクチャ機能に適用します
その後、データメッシュの基盤となるプラットフォームのユニークな特徴を共有します
第9章や第10章などの後の章では、プラットフォームの機能や設計のアプローチについてさらに詳しく説明します
今のところ、データメッシュの基盤となるプラットフォームが今日の多くのソリューションとは異なる点について議論しましょう
注: この章では、データメッシュプラットフォームというフレーズを、基盤となるデータインフラストラクチャ機能のセットの省略形として使用しています
単数形のプラットフォームという用語は、緊密に統合された機能を備えた単一のソリューションや単一のベンダーを意味するものではありません
それは単に、独立していてうまく連携するために使用することができる一連の技術のプレースホルダーです
データメッシュプラットフォーム: 比較と対比
データインフラストラクチャのカテゴリに含まれ、しばしばプラットフォームとして提案される技術ソリューションの大量が存在します
既存のプラットフォームの機能の一部を以下に示します:データレイク、データウェアハウス、またはレイクハウスの形式での分析データストレージ
バッチおよびストリーミングモードでデータを処理するためのデータ処理フレームワークと演算エンジン
計算データフロープログラミングまたは代数的なSQLのようなステートメントに基づくデータクエリ言語
データガバナンスを可能にし、組織全体でのすべてのデータの発見を支援するデータカタログソリューション
複雑なデータパイプラインのタスクやMLモデルの展開ワークフローを組織するパイプラインワークフロー管理これらの機能の多くは、データメッシュの実現にはまだ必要です
しかし、アプローチとデータメッシュプラットフォームの目的には変化があります
比較してみましょう
図4-2は、既存のプラットフォームと比較して、データメッシュプラットフォームの固有の特徴を示しています
データメッシュプラットフォームは既存の技術を利用できますが、これらの固有の特徴を提供します
次のセクションでは、データメッシュが自己サービスプラットフォームの構築に向けてどのように機能するかを詳しく説明します
図4-2
データメッシュプラットフォームの差別化特性 オートノマスなドメイン指向チームのサポート データメッシュプラットフォームの主な責任は、既存または新しいドメインエンジニアリングチームがエンドツーエンドでデータ製品を構築、共有、使用する新しい組み込まれた責任を担うことを可能にすることです
運用システムやその他の情報源からデータを収集し、データを製品として変換して共有し、エンドデータユーザーに提供します
このプラットフォームは、各ドメインのチームが集中管理されたデータチームや仲介者に依存せずに自律的に行うことを可能にする必要があります
多くの既存のベンダーテクノロジーは、すべてのドメインのためにデータを収集し共有する集中制御を前提として構築されています
この集中制御に関する仮定には、以下のような重要な技術的影響があります
 コストはドメインリソースごとではなく、一括で推定および管理されます
セキュリティおよびプライバシー管理は、物理リソースが同じアカウントで共有され、データ製品ごとに分離されたセキュリティコンテキストにはスケーリングされないという前提に基づいています
中央のパイプライン（DAG）オーケストレーションは、すべてのデータパイプラインの管理を中央的に行います
これには中央のパイプライン設定リポジトリと中央のモニタリングポータルがあります
これは独立したパイプライン、各パイプラインが小規模でデータ製品の実装に割り当てられているということに矛盾しています
これらは、既存の技術がドメインチームの自律性を妨げる例の一部です
 オートノマスでインターオペラブルなデータ製品の管理 データメッシュは、アプローチの中心にドメイン指向のデータ製品という新しい概念を置きます
これは自律的に価値を提供する新しいアーキテクチャの概念です
データ製品はお互いにデータを共有し、メッシュ状に相互接続されています
データメッシュプラットフォームは、この新しい概念と共に動作し、その自律的なライフサイクルと全体を管理することをサポートしなければなりません
このプラットフォームの特性は、現在のプラットフォームとは異なり、動作全体を管理するための振る舞い、つまりデータ処理パイプライン、データおよびそのメタデータ、およびデータを独立した構成要素として統治するポリシーなどを管理するものとは異なります
ただし、既存の技術の上に新しいデータ製品管理の抽象化を作成することは可能ですが、それは非常に洗練されていません
操作と分析の連続的なプラットフォーム ドメイン所有の原則により、自律ドメインチームがデータをエンドツーエンドで管理できるプラットフォームが必要とされます
これにより、運用面と分析面の組織的なギャップが埋まります
したがってただし、データメッシュプラットフォームのアプローチと目標には変化があります
以下では、既存のプラットフォームとの比較を行います
図4-2は、既存のプラットフォームと比較してデータメッシュプラットフォームの固有の特徴を示しています
データメッシュプラットフォームは既存の技術を利用できるが、これらの固有の特徴を提供することができます
次のセクションでは、データメッシュが自己サービスプラットフォームの構築に向けてどのように機能するかを説明します
図4-2
データメッシュプラットフォームの差別化特性 自律的なドメイン指向チームへの提供 データメッシュプラットフォームの主な責任は、既存のまたは新しいドメインエンジニアリングチームがデータ製品のビルド、共有、使用の完全なプロセス、運用システムやその他のソースからのデータのキャプチャ、データの変換と共有を埋め込む責任を持つことを可能にすることです
プラットフォームは、集中的なデータチームや仲介者に依存せずに、チームがこれを自律的に行うことを可能にする必要があります
多くの既存のベンダーテクノロジーは、すべてのドメインのためにデータをキャプチャし共有するという前提で構築されています
この集中制御に関する仮定には、次のような技術的な影響があります
コストは分離されたドメインリソースごとではなく、モノリシックに見積もられ管理されます
セキュリティとプライバシー管理は、物理的なリソースが同じアカウントで共有され、データ製品ごとに分離されたセキュリティコンテキストにスケーリングされないという仮定に基づいています
中央パイプライン（DAG）のオーケストレーションは、すべてのデータパイプラインを中央で管理し、中央のパイプライン構成リポジトリと中央のモニタリングポータルを使用します
これは、データ製品の実装ごとに小さく割り当てられた独立したパイプラインとは矛盾します
これは、既存のテクノロジーがドメインチームが自律的に行動する障害となるいくつかの例です
 自律的かつ相互運用可能なデータ製品の管理 データメッシュは、アプローチの中心にドメイン指向のデータ製品という新しい概念を置いています
これは、自律的に価値を提供する新しいアーキテクチャの概念です
データ製品はお互いにデータを共有し、メッシュで相互接続されています
データメッシュプラットフォームは、この新しい概念と連動して動作し、その自律ライフサイクルとすべての構成要素を管理する必要があります
このプラットフォームの特性は、現在のプラットフォームとは異なり、データ処理パイプライン、データおよびそのメタデータ、データを管理するポリシーなど、それぞれが独立した要素としてデータを管理する既存のプラットフォームとは異なります
ただし、既存の技術の上に新しいデータ製品の管理抽象化を作成することは可能ですが、非常にエレガントではありません
 操作と分析の継続的なプラットフォーム ドメイン所有の原則は、ドメインチームがデータを一貫して管理できるプラットフォームを要求します
これにより、運用面と分析面の組織的なギャップが埋まります
したがって、データメッシュプラットフォームはよりつながりのあるエクスペリエンスを提供できる必要があります
チームがアプリケーションを構築および実行するか、分析用途でデータ製品を共有および使用する場合、チームのエクスペリエンスはつながりがあり、シームレスである必要があります
プラットフォームが既存のドメインテクノロジーチームと成功裏に採用されるためには、採用の障壁、運用面と分析面の間の亀裂を取り除く必要があります
データメッシュプラットフォームは、分析と運用技術のギャップを埋める必要があります
それらがシームレスに連携して作業できる方法を見つける必要があります
たとえば、現在、Sparkなどのデータ処理パイプラインを実行する計算ファブリックは、別のクラスタリングアーキテクチャで管理されており、Kubernetesなどの操作サービスを実行する計算ファブリックからは離れていることがよくあります
データ製品をその対応するマイクロサービスと密接に連携させるためには、これらのギャップを埋める必要があります
ソースに整列したデータ製品では、計算ファブリックをより密接に統合する必要があります
私は非常に少数の組織と共に働いてきましたが、両方の計算エンジンを同じ計算ファブリック上で実行しているということはほとんどありませんでした
二つのプレーンの間には固有の違いがあるため、運用システムと分析システムの技術は異なるままでいる必要があります
例えば、デバッグや監査のためのトレースの場合を考えてみましょう
運用システムでは、分散アプリケーション間の（API）呼び出しのトレースに対してOpenTelemetryの標準を使用し、ツリー状の構造で行います
一方、データ処理のワークロードでは、分散データパイプライン上のデータの系譜をトレースするためにOpenLineageを使用します
二つのプレーンの間には十分な違いがあるため、その差を埋める必要があります
しかし、これらの二つの標準がきちんと統合されることは重要です
結局のところ、多くの場合、データの旅はユーザーアクションに応答するアプリケーションの呼び出しから始まります
 ジェネラリストの多数のために 起きているひとつのバリアーは、技術ベンダーが持つ専門的な特殊化のレベルです——専門用語やベンダー固有の知識です
これがデータエンジニアなどの専門的な役割の創造につながっています
この持続不能な専門化の原因にはいくつかの理由があると考えています
具体的な標準や規約の欠如、技術間の相互運用性に対するインセンティブの欠如、製品を簡単に使用できるようにするためのインセンティブの欠如などです
これは、単一のベンダーがデータをプラットフォーム上に格納し、追加のサービスを添付してデータとその処理を制御するという、ビッグモノリシックプラットフォームのメンタリティの残りだと考えています
データメッシュプラットフォームはこのようなパターンを打破し、異なる技術間の相互運用性を促進し、専門家が学び、データから価値を生み出すために覚えなければならない専門的な言語や経験の数を減らすための一連のオープンな規約の定義から始める必要があります
経験豊かなジェネラリスト開発者を奨励し、可能な限り学習が容易な言語やAPIを提供することは、ジェネラリスト開発者の認知負荷を低減するための出発点です
データ駆動型開発を実践する人口の多い人々にスケールアップするために、データメッシュプラットフォームは一般的な技術者に適しています
それらは背景に移り、一般の技術者が自然に使うネイティブツールやプログラミング言語に適応し、じゃまになることはありません
言うまでもなく、これは持続可能なソリューションの結果となるソフトウェアエンジニアリングの慣行を犠牲にすることなく達成されるべきです
例えば、多くの低コードやノーコードのプラットフォームがデータを扱うことを約束していますが、テスト、バージョン管理、モジュール性などが犠牲になります
時間が経つと、メンテナンスが困難になります
 注意 一般的な技術者（エキスパート）というフレーズは、T字型またはPaint Dripの人々と呼ばれることが多い技術者の集団を指します
これらは、ソフトウェアエンジニアリングの広範なスペクトルに精通した開発者であり、ある時点で特定の分野に焦点を当て、深い知識を獲得するという特徴があります
ポイントは、多くの他の分野を探求しながら、1つまたは2つの分野で深く学ぶことができるということです
 彼らは専門家とは対照的であり、彼らの専門化の焦点は多様なスペクトルを探求することを許しません
私の考えでは、将来の一般技術者はデータとデータ製品を作成し共有することができ、また、専門的なデータサイエンティストによって既に開発されたモデルを使用して特徴エンジニアリングやMLトレーニングに役立てることができます
基本的に、彼らはAIをサービスとして使用します
現時点では、データの取り扱いの多くは専門的な知識と長期にわたる努力を要するため、一般技術者の参入を妨げ、データ専門家の不足を招いています
 分散化技術の優遇 既存のプラットフォームのもう一つの共通点は、制御の中央集権化です
例として、中央集権的なパイプラインオーケストレーションツール、中央集権的なカタログ、中央集権的なデータウェアハウスのスキーマ、計算/ストレージリソースの中央集権的な割り当てなどがあります
データメッシュがドメイン所有を通じた分散化に焦点を当てている理由は、組織の同期とボトルネックを回避し、変更のスピードを遅くすることです
表面的には組織の懸念事項ですが、基本的な技術とアーキテクチャは組織のコミュニケーションと設計に直接影響を与えます
一枚岩または中央集権的な技術ソリューションは、中央集権的な制御ポイントとチームを生み出します
データメッシュプラットフォームは、データ共有、制御、ガバナンスの組織の分散化を設計の中心に考慮する必要があります
これらのプラットフォームは、各ドメインチームがデータの共有と使用において行う必要のないタスクを中央管理するために、インフラストラクチャの多くの側面を中央的に管理する必要があります
これが効果的なセルフサーブプラットフォームの出番であり、独立したチームが他のチームへの依存関係なしに最終的な成果を達成できるようにします
ドメインに対して関係のないデータメッシュは、ドメインチーム（ビジネス指向のプロダクト、データ主導のサービス、データ製品を作成することに焦点を当てる）とプラットフォームチームとの間で責任の明確な区分を作り出します
これは、データチームがしばしば分析用途のドメイン固有データの統合、および基礎となる技術インフラの管理を担当するという、現在の責任の区分とは異なります
この責任の区分は、プラットフォームの機能に反映される必要があります
プラットフォームは、ドメインに関係ない機能を提供する一方で、組織全体でのドメイン固有のデータモデリング、処理、共有を可能にするバランスを取る必要があります
これには、データ開発者の深い理解とプロダクト思考の適用が必要です
データメッシュプラットフォームの考え方プラットフォーム：人々や物が立つことができる高い水平の表面
-オックスフォードランゲージプラットフォームという言葉は、私たちの日常のテクニカル用語の中で最も一般的に使われるフレーズの1つであり、組織の技術戦略全体に散りばめられています
一般的に使われるが、定義が難しく、解釈によるものです
データメッシュの文脈でプラットフォームを理解するために、私はいくつかの信頼できる情報源から引用しています
デジタルプラットフォームとは、魅力的な内部製品として配置されたセルフサービスのAPI、ツール、サービス、知識、サポートとしての基盤です
自律的なデリバリーチームは、このプラットフォームを利用して、協調を減らしながらより高いペースで製品の機能を提供することができます
-エバン・ボトッチャー、「私がプラットフォームについて語るときに話すこと」プラットフォームチームの目的は、ストリームに整列したチームが重要なオートノミーを持って作業を配信することを可能にすることです
ストリームに整列したチームは、製品を建設し、稼働させ、修正するために完全な所有権を持ちます
プラットフォームチームは、ストリームに整列したチームがこれらの基礎となるサービスを開発するために必要な認知負荷を軽減するための内部サービスを提供します
プラットフォームは、複雑な技術を簡素化し、それを利用するチームの認知負荷を軽減します
-マシュー・スケルトンとマヌエル・ペイス、チームトポロジーPlatformは、一度に1つの相互作用を設計する
したがって、すべてのプラットフォームの設計は、生産者と消費者の間で可能にされるコアの相互作用の設計から始まるべきです
コアの相互作用は、最初に最も多くのユーザーをプラットフォームに引き付ける価値の交換、つまりプラットフォーム上で最も重要な活動の一形態です
-ジェフリーG.パーカーら, プラットフォームの進化) プラットフォームには、データメッシュにも適用してみたいいくつかの重要な目標があります：データから価値を得るために自律的なチームを可能にする プラットフォームを使用するチームが作業を完了し、その成果を自律的に達成できる能力は、バックログの依存関係などを通じて別のチームが直接に関与することなく、プラットフォームを使用するチームに可能にすることができるという共通の特徴です
データメッシュの文脈では、ドメインチームに分析データの共有の新たな責任を持たせること、または分析データを使用してMLベースの製品を構築することを自律的な方法で可能にすることが、データメッシュプラットフォームの主要な目的です
自己提供のAPIを介してプラットフォームの機能を使用する能力は、自律性を可能にするために重要です
 自律的かつ相互運用可能なデータ製品との価値の交換 プラットフォームのもう一つの重要な側面は、交換される価値とその方法を意図的に設計することです
データメッシュの場合、データ製品はデータユーザーとデータ提供者の間で価値の単位として交換されます
データメッシュプラットフォームは、データ製品の摩擦のない交換を設計に組み込む必要があります
 認知負荷を低減することによる価値の交換の加速 ドメインチームが価値を提供するための作業を簡素化し、加速するためには、プラットフォームは技術的および基礎的な複雑さを隠蔽する必要があります
これにより、ドメインチームの認知負荷が軽減され、データ製品の作成と共有に集中することができます
 データ共有のスケールアウト データメッシュは、組織のデータから価値を得るための解決策です
したがって、プラットフォームの設計はスケールに対応する必要があります：組織内の主要ドメイン間でデータを共有するだけでなく、組織外の信頼関係の範囲を超えてパートナーのネットワーク全体でデータを共有することもできるようにする必要があります
このスケールの障害の一つは、複数のプラットフォーム間で安全にデータ共有が相互運用可能でないことです
 データメッシュプラットフォームは、他のプラットフォームとの相互運用性のためにデータ製品を共有するように設計する必要があります
 埋め込みイノベーションの文化をサポートする データメッシュプラットフォームは、イノベーションに直接貢献していない活動を排除することにより、イノベーションの文化をサポートします
データの検索、洞察の把握、およびMLモデル開発にデータを使用することを非常に簡単にすることで、この目標を達成します
 図4-3は、ドメインチームがデータ製品を共有および使用するエコシステムにこれらの目標が適用されたものを描いています
 図4-3. データメッシュプラットフォームの目標 さて、データメッシュプラットフォームがこれらの目標を達成する方法について話しましょう
 データを通じて価値を得るために自律的なチームを可能にする プラットフォームを設計する際には、プラットフォームのユーザーの役割とデータ製品を共有および使用する過程を考慮することが役立ちます
プラットフォームは、それぞれの過程に対して摩擦のない体験を作り出す方法に焦点を当てることができます
 例えば、データメッシュエコシステムの主要なペルソナであるデータ製品の開発者とデータ製品のユーザーを考えてみましょう
もちろん、それぞれのペルソナには異なるスキルセットを持つ人々のスペクトルが含まれますが、今回の話では、それらのスペクトルをまたいで共通した過程に焦点を当てることができます
データ製品の開発者、データ製品のユーザーという2つの主要なペルソナを考えてみましょう
もちろん、それぞれのペルソナには異なるスキルセットを持つ人々のスペクトルが含まれますが、今回の話では、それらのスペクトルをまたいで共通した過程に焦点を当てることができます
データ製品オーナーなどの他の役割もありますが、本例では簡潔さを重視し、省略します
 データ製品の開発者を可能にする データ製品開発者の配信過程は、データ製品の開発、テスト、デプロイ、モニタリング、更新、完全な連続配信などを含みます
簡単に言えば、開発者はコード、データ、ポリシーを1つのユニットとして、データ製品のライフサイクルを管理しています
想像できるように、このライフサイクルを管理するためにはかなりのインフラストラクチャが用意される必要があります
データ製品のライフサイクル管理のための基礎となるインフラストラクチャの提供と管理には、今日のツールの専門知識が必要であり、それを各ドメインで複製することは困難です
したがって、データメッシュプラットフォームは、データ製品開発者が基盤となるインフラリソースのプロビジョニングを心配することなく、データ製品を構築、テスト、展開、保護、および維持するために必要なすべての機能を実装する必要があります
ドメインに依存しない機能とクロスファンクションの機能をすべて可能にする必要があります
最終的に、プラットフォームはデータ製品開発者がデータ製品開発のドメイン固有の側面に集中することを可能にする必要があります
データの生成と維持を行うドメイン固有のロジックである変換コード、ドメインのデータの整合性を検証および維持するためのビルド時のテスト、データ製品が品質保証を満たしていることを常に監視するランタイムテスト、スキーマ、ドキュメントなどのデータ製品のメタデータの開発、必要なインフラリソースの宣言
それ以外のことはデータメッシュプラットフォームが面倒を見る必要があります
例えば、ストレージ、アカウント、計算などのインフラストラクチャのプロビジョニングです
セルフサービスのアプローチでは、データ製品開発者が必要なインフラのニーズを宣言し、プラットフォームが残りの面倒を見るための一連のプラットフォームAPIを公開します
これについては、詳細には第14章で説明されています
データ製品の利用者を可能にするデータユーザーの旅は、データを分析して洞察を作成したり、機械学習モデルを開発したりすることから始まります
データが発見されると、それにアクセスして理解し、さらに探求する必要があります
データが適切であることがわかった場合、それを引き続き使用し続けます
データの使用は一度限りではありません
消費者は、機械学習モデルや洞察を最新の状態に保つために、新しいデータを受け取り続けて処理します
データメッシュプラットフォームは、そのような旅を容易にする基盤となるメカニズムを構築し、データ製品の消費者が摩擦なく仕事を完了するために必要な機能を提供します
この旅を自律的に可能にするために、手動介入の必要性を低減する必要があります
例えば、データを作成したチームやガバナンスチームにデータへのアクセスを正当化して取得する必要がないようにする必要があります
プラットフォームは、アクセスのリクエストを容易にし、消費者の自動評価に基づいてアクセスを付与するプロセスを自動化します
自律性と相互運用性のあるデータ製品による価値の交換データメッシュプラットフォームを多面的なプラットフォームと見なす興味深い視点は、主に2つ（またはそれ以上）の異なる当事者間の直接的な相互作用を可能にするものです
データメッシュの場合、当事者はデータ製品開発者、データ製品オーナー、およびデータ製品利用者です
この特定の視点は、価値、つまりデータ製品の交換によって直接測定されるプラットフォームを構築するための非常に創造的な源泉になりえます
価値はメッシュ上でデータ製品間、またはメッシュのエッジでエンド製品（MLモデル、レポート、ダッシュボードなど）とデータ製品間で交換されます
メッシュは基本的に組織のデータマーケットプレイスとなります
このようなデータメッシュプラットフォームの特徴は、組織内の文化変革の触媒となり、共有を次のレベルに促進します
前述のように、価値の交換の重要な側面は、プラットフォームが妨げることなく自律的に行えることです
データ製品開発者にとっては、プラットフォームチームへの依存や常に手を引かれることなく、自分のデータ製品を作成し提供できることを意味します
複数のデータ製品を組み合わせて高次の価値を創造する価値の交換は、単一のデータ製品の使用を超えて広がることがあります
例えば、Daffのリスナーに関する興味深い洞察は、彼らが音楽を聴く際の行動、フォローしているアーティスト、人口統計、ソーシャルメディアとの関わり、友人ネットワークの影響、周囲の文化イベントなどを相互相関させることで生成されます
これらは複数のデータ製品であり、相関および特徴行列へと組み合わされる必要があります
プラットフォームは、データ製品の互換性を実現します
例えば、プラットフォームはデータ製品のリンクを可能にし、あるデータ製品が他のデータ製品からデータやデータ型（スキーマ）を使用する場合に使用されます
これがシームレスに可能となるために、プラットフォームはデータ製品を識別し、アドレスを指定し、データ製品に接続し、データ製品からデータを読み取るなどの標準化された簡単な方法を提供します
このような簡単なプラットフォームの機能により、異種のドメインが均質なインターフェースで結びついたメッシュが作成されます
これについては第13章で説明します


認知負荷の軽減による価値の迅速な交換
認知負荷は、認知科学の分野で最初に導入された用語であり、問題解決や学習に必要な作業メモリの量を意味します
認知負荷には、対象の固有の複雑さや、タスクや情報の提示方法など、さまざまな要素が影響します
プラットフォームは、開発者の認知負荷を軽減する手段として、ますます注目されています
これは、開発者に提示される詳細や情報の量を隠すことによって実現されます
つまり、複雑さを抽象化することです
データ製品の開発者として、私はドメインに依存しない希望を表現できるはずであり、具体的な実装方法を説明する必要はありません
例えば、開発者として、データの構造、保持期間、潜在的なサイズ、機密性クラスを宣言し、データ構造を作成し、ストレージを提供し、自動暗号化を実行し、暗号化キーを管理し、キーを自動的にローテーションするなどの作業をプラットフォームに任せることができるべきです
これは、データ開発者やユーザーが直接関与する必要のないドメインに依存しない複雑さです
設定可能性を犠牲にすることなく複雑さを抽象化するための多くの技術があります
以下の2つの方法が一般的に適用されています


宣言的なモデリングによる複雑さの抽象化
過去数年間、KubernetesなどのコンテナオーケストレータやTerraformなどのインフラストラクチャプロビジョニングツールなど、運用プラットフォームは、対象の状態を宣言的なモデリングによって複雑さを抽象化する新しいモデルを確立しました
これは、ターゲットの状態を構築する方法に関する命令型の手法とは対照的です
前者は「何をするか」に焦点を当て、後者は「どのようにするか」に焦点を当てています
このアプローチは、開発者の生活を大幅に簡素化することに非常に成功しています
多くのシナリオでは、宣言的なモデリングによる複雑さの抽象化はすぐに限界に達します
例えば、データ変換ロジックを宣言を通じて定義することは、ロジックが複雑になると収益が減少する傾向があります
ただし、プロビジョニングされたインフラストラクチャなどの状態によって記述できるシステムは、宣言的なスタイルに適しています
これは、データメッシュインフラストラクチャがプラットフォームの場合も同様です
データ製品のライフサイクルを管理するためのインフラストラクチャのターゲット状態は、宣言的に定義することができます


自動化による複雑さの抽象化
データ製品の開発者の作業過程から人間の介入や手動手順を除去することによって、自動化を通じて複雑さを低減することも可能です
データメッシュの実装の一部を自動化する機会は数多く存在します
基礎となるデータインフラストラクチャのプロビジョニング自体も、インフラストラクチャをコードとして自動化するテクニックを使用して自動化することができます
さらに、製造から消費までのデータの価値ストリームにおいて、多くのアクションを自動化することが可能です
例えば、現在はデータの認証や検証の承認プロセスが通常手動で行われています
これは自動化のための非常に大きな機会の領域です
プラットフォームは、データの完全性を自動で検証したり、データの性質を統計的にテストしたり、予期しない外れ値を検出するために機械学習を使用するなど、自動化を通じてデータの検証プロセスから複雑さを取り除くことができます
データの大規模な共有を実現するための相互運用可能なソリューションの標準が存在しないという問題が、既存のビッグデータ技術の分野で見られます
たとえば、データにアクセスする際の認証と承認の統一モデルの欠如、データのプライバシー権利を表現し伝送するための標準の不在、データの時間的側面を提示するための標準の不足などです
これらの標準の不足は、組織内の信頼の枠を超えて利用可能なデータのネットワークを拡大することを妨げています
最も重要なことは、データ技術の分野がUnixの哲学を欠いているということです
「これがUnixの哲学です
一つのことをうまく行うプログラムを書く
プログラムは一緒に動かすために書く…」- Douglas McIlroy
Unixの哲学を普及させる文化、哲学、およびオペレーションスケーラブルなデータ共有に関して、現在のビッグデータ技術の状況にはいくつかの課題があります
例えば、データにアクセスする際の認証と認可のための統一されたモデルがないこと、データのプライバシー権利を表現し伝達するための標準がないこと、データの時系列的な側面を表示するための標準がないことなどです
これらの標準が欠けていることにより、組織の信頼の範囲を超えた利用可能なデータネットワークの拡大が妨げられています
さらに重要なことに、データ技術の領域ではUnixの哲学が欠けています
つまり「一つのことをうまく行うプログラムを書け
そして、それらを連携して動くようにプログラムを書け」という考え方です
特別な人々（McIlroy、Ritchie、Thompsonなど）が文化や哲学、ソフトウェアの構築方法を普及させるための成果を上げてきたため、力強くスケーリングされた複雑なシステムを緩く組み合わせたシンプルで小さなサービスを通じて構築できたと考えています
何らかの理由で、ビッグデータシステムにおいてはこの哲学を忘れてしまったように思います
おそらくそれは初期の仮定（「分析データアーキテクチャの特徴」を参照）が文化を形成したからかもしれません
また、ある時点でデータ（体）とコード（魂）を分離することを決定し、それによって異なる哲学が確立されました
データメッシュプラットフォームが組織の範囲内および範囲外でデータを共有するためにリアリスティックにスケーリングするためには、Unixの哲学を全面的に受け入れ、データ管理とデータ共有の独自のニーズに適応させる必要があります
プラットフォームを異なるベンダーが異なる実装で実装できる相互運用可能なサービスのセットとして設計する必要がありますが、それらが他のプラットフォームサービスとうまく連携するようにする必要があります
例として、データメッシュ上のすべてのデータ製品の動作を監視し、障害、エラー、望ましくないアクセスを検出し、関連チームに通知する能力であるオブザーバビリティを考えてみましょう
オブザーバビリティが機能するためには、複数のプラットフォームサービスが協力する必要があります
データ製品は自身の動作に関する情報を発行し、ログを取得し、ホリスティックなメッシュビューを提供するサービス、これらのログ内の異常やエラーを検索、分析、検出するサービス、問題が発生した場合に開発者に通知するサービスなどがあります
Unixの哲学を基にこれを構築するには、これらのサービスを選択し、つなぎ合わせることができる必要があります
これらのサービスをシンプルに統合するための鍵は、相互運用性、つまりログやメトリクスが表現され共有されるための共通の言語とAPIです
このような標準がなければ、データへのアクセスを単一のホスティング環境に制限する単一のモノリシック（しかし、うまく統合された）ソリューションに退化してしまい、異なる環境間でのデータの共有や観察ができなくなります
埋め込まれたイノベーションの文化をサポートする これまでのところ、継続的なイノベーションはおそらく任意のビジネスの中核的な能力の一つであると言えます
Eric Riesは、ビルド-測定-学習の短期的かつ迅速なサイクルを通じて科学的にイノベーションする方法を示すために「リーンスタートアップ」を紹介しました
その概念は後に大企業にも適用され、「リーンエンタープライズ」というスケールされたイノベーション手法に発展しました
要点は、イノベーションの文化を育むためには、人々を不必要な作業や偶発的な複雑さや摩擦から解放し、実験を行うことができる環境が必要であるということです
データメッシュプラットフォームは、不要な手作業を排除し、複雑さを隠し、データ製品の開発者や利用者のワークフローを効率化することで、彼らにデータを使ってイノベーションする自由を与えます
データメッシュプラットフォームがその役割を果たしているかどうかを評価するための簡単な指標は、チームがデータ駆動の実験を考え、実験を実行するために必要なデータを使用するまでにどれくらい時間がかかるかを測定することです
時間が短ければ短いほど、データメッシュプラットフォームは成熟しています
もう一つの重要なポイントは、誰が実験を行う権限を持っているかです
データメッシュプラットフォームは、ドメインチームが革新を促進し、データに基づいた実験を行うことをサポートしています
データに基づいた革新は、もはや中央のデータチームだけのものではありません
それらは、各ドメインチームに組み込まれなければなりません
彼らが自分たちのサービス、製品、またはプロセスを開発する際に
セルフサーブデータメッシュプラットフォームへの移行これまで、既存のデータプラットフォームとデータメッシュの主な違いについて話し、データメッシュプラットフォームの主な目標を説明しました
ここで、データメッシュプラットフォームへの移行において実行できるいくつかの行動を紹介します
最初にAPIとプロトコルを設計するプラットフォームの旅を始める際には、プラットフォームがユーザーに提供するインターフェースの選択と設計から始めます
インターフェースはプログラム的なAPIであるかもしれません
コマンドラインやグラフィックなインターフェースかもしれません
いずれにしても、まずはインターフェースを決定し、その実装をさまざまな技術を使って行います
このアプローチは、多くのクラウドオファリングで採用されています
たとえば、クラウドのblobストレージプロバイダは、REST APIを公開してオブジェクトの投稿、取得、削除を行います
これはプラットフォームのすべての機能に適用できます
APIに加えて、相互運用性を可能にするコミュニケーションプロトコルと標準を決定します
大規模な分散アーキテクチャの一例であるインターネットからインスピレーションを得て、狭いウェストプロトコルを決定します
たとえば、データ製品がセマンティックをどのように表現し、どの形式で時変データをエンコードするか、それぞれがどのクエリ言語をサポートするか、どのSLOを保証するかなどを規定するプロトコルを決定します
もう1つは、データメッシュプラットフォームは一般的な大多数のために設計されなければならないと述べました（「一般的な大多数のために設計する」）
多くの組織は、データエンジニアのようなデータの専門家を見つけることに苦労していますが、データと一緒に働きたいと思っている多くの一般開発者がいます
ビッグデータテクノロジーの非連続的で閉鎖的で高度に専門化された世界は、同様に超専門化されたデータ技術者のフラグメントを作り出しました
プラットフォームのテクノロジーを評価する際には、多くの開発者にとって自然なプログラミングスタイルに合うものを選ぶようにしましょう
たとえば、パイプラインのオーケストレーションツールを選ぶ場合、一般的な開発者にとって馴染みのあるPython関数の簡単なプログラミングを可能にするものを選ぶべきです
YAMLやXMLでエゾテリックな記法のドメイン固有言語（DSL）を作成しようとするものよりも
現実には、データ製品の複雑さのスペクトルと、データ製品開発者の専門レベルのスペクトルが存在するでしょう
プラットフォームはこのスペクトルを満たし、スケールでのデータ製品の提供を実現する必要があります
いずれにせよ、弾力性のあるメンテナンス可能なデータ製品を構築するために、常に携帯エンジニアリングの実践を適用する必要があります
在庫を調査し、簡素化する分析データプレーンとオペレーションプレーンの分離により、解析データとアプリケーションやサービスの構築と実行に使用される別々の技術スタックが残されました
データ製品が操作的な世界に統合され、埋め込まれるようになると、2つのプラットフォームを統合し、重複を取り除く機会が生まれます
過去数年間、業界ではデータソリューションとしてマーケティングされるテクノロジーへの過剰投資が行われています
多くの場合、これらのツールは既存のCI/CDエンジンでは提供できない差別化能力をほとんど提供していません
たとえば、DataOpsとしてマーケティングされる新しい連続インテグレーションおよび連続デリバリー（CI/CD）ツールのクラスを見てみると、既存のCI/CDエンジンが提供できない差別化能力はほとんどありませんでした
始める際には、組織が採用しているプラットフォームサービスを棚卸し、簡素化の機会を探すことが重要です
データメッシュプラットフォームが技術の簡素化と運用と分析プラットフォーム間のより密な協力の触媒となることを望んでいます
データ製品の管理のための上位APIを作成する データメッシュプラットフォームは、新しい抽象化としてデータ製品を管理するための新しいセットのAPIを導入する必要があります（「自律性と相互運用性のあるデータ製品の管理」）
クラウドプロバイダから提供されるような多くのデータプラットフォームには、ストレージやカタログ、コンピュートなどの下位レベルのユーティリティAPIが含まれていますが、データメッシュプラットフォームでは、データ製品自体をオブジェクトとして扱う上位レベルのAPIを導入する必要があります
例えば、データ製品を作成するためのAPI、データ製品を発見するためのAPI、データ製品に接続するためのAPI、データ製品から読み込むためのAPI、データ製品を保護するためのAPIなどです
データ製品の論理的な設計図は、第9章を参照してください
データメッシュプラットフォームを設立する際には、データ製品の抽象化と連携する高レベルなAPIから始めてください
メカニズムではなく体験を作りましょう 自分は、メカニズムに依存する形でプラットフォームを構築/購入する状況に何度も出会ってきました
これに対して、プラットフォームの定義は、そのプラットフォームが可能にする体験に基づいているのではなく、含まれるメカニズムに基づいていることがよくあります
このアプローチでは、プラットフォームの開発が膨張し、野心的で高価な技術が採用されることがしばしばあります
たとえば、データカタログについて考えてみましょう
私が出会ったほとんどのプラットフォームは、メカニズムのリストにデータカタログが含まれているため、最も多くの機能を持つデータカタログ製品を購入し、チームのワークフローをカタログの内部動作に合わせて調整しようとすることがよくあります
このプロセスには通常数か月かかります
それに対して、プラットフォームはデータ製品の発見という単一の体験の具体化から始めることができます
その後、この体験を可能にする最もシンプルなツールとメカニズムを構築または購入します
そして、次の体験のために繰り返してリファクタリングします
最もシンプルな基盤から始め、進化させる データメッシュプラットフォームの目標と独自の特性について説明するこの章の長さを考慮すると、「データメッシュを今日から採用することはできるのか、それともまずプラットフォームを構築するまで待つべきなのか」と思うかもしれません
答えは、データメッシュ戦略を今すぐ採用し始めることです
データメッシュプラットフォームがなくても、最もシンプルな基盤から始めることができます
最もシンプルな基盤フレームワークは、おそらくすでに採用しているデータ技術で構成されています
特にクラウド上で分析を行っている場合には、ボトムレイヤのユーティリティであるストレージ技術、データ処理フレームワーク、連合クエリエンジンなどを基盤として利用始める際には、組織が採用しているプラットフォームサービスの在庫を把握し、簡素化する機会を探してください
データメッシュプラットフォームが技術の風景を簡素化し、運用と分析のプラットフォーム間でより緊密な協力を生み出すことを望んでいます
データ製品を管理するための上位レベルのAPIを作成するデータメッシュプラットフォームは、「自律的かつ相互運用可能なデータ製品の管理」という新しい抽象化として、データ製品を管理するための新しい一連のAPIを導入する必要があります
多くのデータプラットフォーム（クラウドプロバイダから提供されるサービスなど）には、下位レベルの便益API（ストレージ、カタログ、計算など）が含まれていますが、データメッシュプラットフォームは、データ製品をオブジェクトとして扱うより高いレベルのAPIを導入する必要があります
例えば、データ製品の作成、データ製品の発見、データ製品への接続、データ製品からの読み取り、データ製品の保護などのAPIを考えてみてください
データ製品の論理設計図については、第9章を参照してください
データメッシュプラットフォームを確立する際には、データ製品の抽象化と連携する高レベルのAPIから始めてください
メカニズムではなく体験を構築する私は、さまざまなプラットフォームの構築/購入の状況に出くわしてきましたが、プラットフォームの言明はその含まれるメカニズムに基づいているのではなく、それが可能にする体験に基づいています
このアプローチによるプラットフォームの定義は、しばしば膨張したプラットフォームの開発や野心的かつ高価な技術の採用につながります
例としてデータカタログを挙げてみましょう
私が出会ったほとんどのプラットフォームは、機能リストが最も長いデータカタログ製品を購入し、その内部の動作にチームのワークフローを合わせ込むことになります
このプロセスはしばしば数か月かかります
それに対して、プラットフォームはデータ製品の発見という一つの体験の言明から始めることができます
その後、この体験を可能にする最もシンプルなツールとメカニズムを構築または購入します
そして、次の体験に向けて繰り返し、再設計しましょう
最もシンプルな基盤から始め、進化するための収穫を行う現在の章では、データメッシュプラットフォームの目標と独自の特性についての議論が長くなっているため、「今日からデータメッシュを採用し始めることはできるのか、それともまずプラットフォームを構築するまで待つべきなのか？」と思われるかもしれません
答えは、データメッシュの戦略を今日から採用し始めることです、たとえデータメッシュプラットフォームを持っていなくてもです
最もシンプルな基盤から始めることができます
最も小さい可能な基盤フレームワークは、おそらく既に採用しているデータテクノロジーで構成されています
特にクラウド上での分析を既に行っている場合はそうです
基盤として使用できる下層のユーティリティには、通常のストレージテクノロジー、データ処理フレームワーク、フェデレーテッドクエリエンジンなどが含まれます
データ製品の数が増え、標準が開発され、データ製品間の似た問題へのアプローチが発見されると、共通の機能やドメインチームを横断した類似問題への共通のアプローチを集めることで、プラットフォームを収穫したフレームワークとして進化させていきます
データメッシュプラットフォーム自体が製品であることを忘れないでください
それは内部の製品ですが、多くの異なるツールやサービスから構築されています
製品のユーザーは内部のチームです
それには技始める時は、組織が採用しているプラットフォームサービスの在庫を確認し、簡素化の機会を探しましょう
データメッシュプラットフォームが技術の景色を簡素化し、運用と分析のプラットフォーム間の緊密な協力を促進するきっかけになることを願っています
データ製品を管理するための上位レベルのAPIを作成する データメッシュプラットフォームは、新たな抽象化（「自律的で相互運用可能なデータ製品の管理」）としてデータ製品を管理するための新しい一連のAPIを導入しなければなりません
多くのデータプラットフォーム（クラウドプロバイダーから提供されるサービスなど）では、ストレージ、カタログ、計算などのより低レベルのユーティリティAPIが含まれていますが、データメッシュプラットフォームは、オブジェクトとしてのデータ製品に対処する上位レベルのAPIを導入しなければなりません
例えば、データ製品の作成、データ製品の発見、データ製品への接続、データ製品からの読み取り、データ製品の保護などのAPIを考えてみてください
データ製品の論理的な設計図については、第9章を参照してください
データメッシュプラットフォームを確立する際は、データ製品の抽象化と連携する高レベルのAPIから始めてください
メカニズムではなくエクスペリエンスを構築する プラットフォームの説明が含まれる多くの状況で、私はメカニズムに焦点を当てたプラットフォームのイメージングが行われており、むしろそれが可能にするエクスペリエンスにはフォーカスが当てられていませんでした
このアプローチにより、プラットフォームの定義が肥大化し、野心的で高額な技術の採用が行われることがあります
例えば、データカタログを考えてみましょう
私が出会ったほとんどのプラットフォームは、機構のリストにデータカタログが含まれていました
これにより、機能の一覧が最も長いデータカタログ製品を購入し、チームのワークフローをカタログの内部機能に合わせることになります
このプロセスは通常数ヶ月かかります
それに対して、あなたのプラットフォームは、データ製品の発見という単一のエクスペリエンスの具現化から始めることができます
その後、このエクスペリエンスを可能にする最もシンプルなツールとメカニズムを作成または購入します
その後、次のエクスペリエンスに合わせて繰り返しリファクタリングを行います
最もシンプルな基盤から始め、進化のために収穫する この章でデータメッシュプラットフォームの目的と独自の特徴について議論した長さを考えると、「今日からデータメッシュを採用することはできるのか、それともまずプラットフォームを構築するまで待つべきか」という疑問が生じるかもしれません
答えは、データメッシュ戦略を今日から採用し始めることです
データメッシュプラットフォームはまだなくても構いません
最もシンプルな基盤から始めることができます
最もシンプルな基盤フレームワークは、おそらくすでに採用しているデータ技術で構成されているはずです、特に既にクラウド上で分析を実行している場合は
基盤に使用できる下位ユーティリティには、一般的なストレージ技術、データ処理フレームワーク、連邦クエリエンジンなどが含まれます
データ製品の数が増え、標準が開発され、データ製品間の類似の問題にアプローチするための共通の方法が見つかるにつれて
その後、プラットフォームは、データ製品とドメインチームを横断した共通の機能を収集して進化させるための枠組みとして収穫されたものとして成長を続けます
データメッシュプラットフォーム自体も製品です
それは内部製品であり、複数のベンダーからのさまざまなツールとサービスで構築されています
製品のユーザーは内部のチームです
それには技術的なプロダクトオーナーシップ、長期計画、長期的な保守が必要です
それは進化し続け、進化的成長を遂げるものですが、その生活は今日から最小限の実行製品（MVP）として始まります
9 総括 データメッシュの自己サービスプラットフォームの原則は、既存のドメインエンジニアリングチームに課される認知負荷を軽減するために登場しました：自分の分析データを所有し、製品として共有します
既存のデータプラットフォームとは共通の機能を共有しており、多様なストレージへのアクセス、データ処理エンジン、クエリエンジン、ストリーミングなどを提供しています
しかし、このプラットフォームは既存のプラットフォームと異なり、主に総合技術者から成る自治型のドメインチームがユーザーとなっています
このプラットフォームは、データ、メタデータ、コード、ポリシーを一つのユニットとしてカプセル化したデータ製品のより高度な構成を管理します
その目的は、ドメインチームにスーパーパワーを与えることで、低レベルの複雑さをシンプルな抽象化の背後に隠し、データ製品の値としての交換の目標を達成するための道筋から摩擦を取り除くことです
そして、最終的には、チームがデータでイノベーションを行うことを可能にします
データの共有を単一の展開環境や組織単位、または企業を超えてスケールアウトするために、相互運用可能な分散型のソリューションを優先します
私たちは第10章でプラットフォームについてさらに詳しく掘り下げ、データメッシュプラットフォームが提供できる具体的なサービスについて話をします