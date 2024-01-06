One long-standing challenge of existing analytical data architectures is the high friction and cost of using data: discovering, understanding, trusting, exploring, and ultimately consuming quality data. There have been numerous surveys surfacing this friction. A recent report from Anaconda, a data science platform company, “The State of Data Science 2020”, finds that nearly half of a data scientist’s time is spent on data preparation—data loading and cleansing. If not addressed, this problem only exacerbates with data mesh, as the number of places and teams who provide data, i.e., domains, increases. Distribution of the organization’s data ownership into the hands of the business domains raises important concerns around accessibility, usability, and harmonization. Further data siloing and regression of data usability are potential undesirable consequences of data mesh’s first principle, domain-oriented ownership. The principle of data as a product addresses these concerns. The second principle of data mesh, data as a product, applies product thinking to domain-oriented data to remove such usability frictions and truly delight the experience of the data users—data scientists, data analysts, data explorers, and anyone in between. Data as a product expects that the analytical data provided by the domains is treated as a product, and the consumers of that data should be treated as customers—happy and pleased. Furthermore, data as a product underpins the case for data mesh, unlocking the value of an organization’s data by dramatically increasing the potential for serendipitous and intentional use. In his book INSPIRED, Marty Cagan, a prominent thought leader in product development and management, provides convincing evidence on how successful products have three common characteristics: they are feasible, valuable, and usable. Data as a product principle defines a new concept, called data product, that embodies standardized characteristics to make data valuable and usable. Figure 3-1 demonstrates this point visually. This chapter introduces these characteristics. Chapter 4 describes how to make building data products feasible. Figure 3-1. Data products live at the intersection of Marty Cagan’s characteristics of successful products For data to be a product it adheres to a set of rules and exhibits a set of traits that make it fit right in the intersection of Cargan’s usability, feasibility, and valuable Venn diagram. For data to be a product, it must be valuable to its users—on its own and in cooperation with other data products. It must demonstrate empathy for its users and be accountable for its usability and integrity. I admit that treating data as a product doesn’t simply happen out of good intentions. For that, this principle introduces new roles to domains such as the domain data product owner and data product developer who have responsibility for creating, serving, and evangelizing data products, while maintaining the specific objective measures of data accessibility, quality, usability, and availability over the lifetime of the data products. Chapter 16 will cover further details about these new roles. Compared to past paradigms, data as a product inverts the model of responsibility. In data lake or data warehousing architectures the accountability of creating data with quality and integrity resides downstream from the source and remains with the centralized data team. Data mesh shifts this responsibility close to the source of the data. This transition is not unique to data mesh; in fact, over the last decade we have seen the trend of shift left with testing and operations, on the basis that addressing problems is cheaper and more effective when done close to the source. I go as far as saying that what gets shared on a mesh is not merely data; it is a data product. Note Data as a product is about applying product thinking to how data is modeled and shared. This is not to be confused with product selling. Let’s unpack how we can apply product thinking to data. Applying Product Thinking to Data Over the last decade, high-performing organizations have embraced the idea of treating their internal operational technology like a product, similarly to their external technology. They treat their internal developers as customers and their satisfaction a sign of success. Particularly in two areas, this trend has been strongly adopted: applying product management techniques to internal platforms, which accelerates the ability of internal developers to build and host solutions on top of internal platforms (e.g., Spotify Backstage), and treating APIs as a product to build APIs that are discoverable, understandable, and easily testable to assure an optimal developer experience (e.g., Square APIs). Applying the magic ingredient of product thinking to internal technology begins with establishing empathy with internal consumers (i.e., fellow developers), collaborating with them on designing the experience, gathering usage metrics, and continuously improving the internal technical solutions over time to maintain ease of use. Strong digital organizations allocate substantial resources and attention to building internal tools that are valuable to the developers and ultimately the business. Curiously, the magical ingredient of empathy, treating data as a product and its users as customers, has been missing in big data solutions. Operational teams still perceive their data as a byproduct of running the business, leaving it to someone else, e.g., the data team, to pick it up and recycle it into products. In contrast, data mesh domain teams apply product thinking with similar rigor to their data, striving for the best data user experience. Consider Daff. One of its critical domains is the media player. The media player domain provides essential data such as what songs have been played by whom, when, and where. There are a few different key data users for this information. For example, the media player support team is interested in near-real-time events to quickly catch errors causing a degrading customer experience and recover the service, or respond to incoming customer support calls informedly. On the other hand, the media player design team is interested in aggregated play events that tell a data story about the listener’s journey over a longer period of time, to improve media players with more engaging features toward an overall better listener experience. With an informed understanding of these use cases and what information the other teams need, the media player domain provides two different types of data as its products to the rest of the organization: near-real-time play events exposed as infinite event logs, and aggregated play sessions exposed as serialized files on an object store. This is product ownership applied to data. Figure 3-2 shows the media player domain data products. Figure 3-2. Data products example As you can imagine, you can adopt the majority of the product ownership techniques to data. However, there is something unique about data. The difference between data product ownership and other types of products lies in the unbounded nature of data use cases, the ways in which particular data can be combined with other data and ultimately turned into insights and actions. At any point in time, data product owners are aware or can plan for what is known today as viable use cases of their data, while there remains a large portion of unknown future use cases for the data produced today, perhaps beyond their imagination. This is more true for source-aligned domains and less for consumer-aligned ones. This source-aligned data captures the reality of the business interactions and events as they have happened. They can continue to be used, transformed, and reinterpreted by data users of the future. Source-aligned data products need to balance the immediate known use cases and the unknown ones. They have no choice but to strive to model the reality of the business, as closely as possible, in their data, without too much assumption in how it will be used. For example, capturing all the play events as an infinite high-resolution log is a safe choice. It opens the spectrum of future users who can build other transformations and infer new insights from the data that is captured today. This is the main difference between data product design and software product design. Baseline Usability Attributes of a Data Product In my opinion, there is a set of non-negotiable baseline characteristics that a data product incorporates to be considered useful. These characteristics apply to all data products, regardless of their domain or archetype. I call these baseline data product usability attributes. Every data product incorporates these characteristics to be part of the mesh. Figure 3-3 lists data products’ usability attributes. Figure 3-3. The baseline usability attributes of data products (DAUTNIVS) Note that these usability characteristics are oriented around the experience of the data users. These are not intended to express the technical capabilities. Part IV, “Design of Data Quantum Architecture,” covers the technical characteristics of data products. The baseline characteristics listed in this section are an addition to what has been known as FAIR data in the past—data that meets the principles of findability, accessibility, interoperability, and reusability.1 In addition to these principles I have introduced characteristics that are necessary to make distributed data ownership work. Let’s put ourselves in the shoes of data users and walk through these attributes. Discoverable Two of the most important characteristics of good design are discoverability and understanding. Discoverability: Is it possible to even figure out what actions are possible and where and how to perform them? Understanding: What does it all mean? How is it supposed to be used? What do all the different controls and settings mean?Don Norman The very first step that data users take in their journey is to discover the world of data available to them and explore and search to find “the one.” Hence, one of the first usability attributes of data is to be easily discoverable. Data users need to be able to explore the available data products, search and find the desirable sets, and explore them and gain confidence in using them. A traditional implementation of discoverability is a centralized registry or catalog listing available datasets with some additional information about each dataset, the owners, the location, sample data, etc. Often this information is curated after the fact by the centralized data team or the governance team. Data product discoverability on data mesh embraces a shift-left solution where the data product itself intentionally provides discoverability information. Data mesh embraces the dynamic topology of the mesh, continuously evolving data products and the sheer scale of available data products. Hence, it relies on individual data products to provide their discoverability information at different points of their life cycle—build, deploy, and run—in a standard way. Each data product continuously shares its source of origin, owners, runtime information such as timeliness, quality metrics, sample datasets, and most importantly information contributed by their consumers such as the top use cases and applications enabled by their data. Part IV, “How to Design the Data Product Architecture”, will discuss the technical design of data product discoverability. Addressable A data product offers a permanent and unique address to the data user to programmatically or manually access it. This addressing system embraces the dynamic nature of the data and the mesh topology. It must recognize that many aspects of data products will continue to change, while it assures continuity of usage. The addressing system accommodates the following ever-changing aspects of a data product, among others, while retaining access to the data product through its long-lasting unique address: Semantic and syntax changes in data products Schema evolution Continuous release of new data over time (window) Partitioning strategy and grouping of data tuples associated with a particular time (or time window) Newly supported modes of access to data New ways of serializing, presenting, and querying the data Changing runtime behavioral information For example, service-level objectives, access logs, debug logs The unique address must follow a global convention that helps users to programmatically and consistently access all data products. The data product must have an addressable aggregate root that serves as an entry to all information about a data product, including its documentation, service-level objectives, and the data it serves. Understandable Once a data product is discovered, the next step of the data user’s journey is to understand it, which involves getting to know the semantics of its underlying data, as well as various syntax in which the data is encoded. Each data product provides semantically coherent data: data with a specific meaning. A data user needs to understand this meaning: what kind of entities the data product encapsulates, what the relationships among the entities are, and their adjacent data products. Getting back to the media player event example, a data user should easily understand what constitutes a player event: the user, the play actions they have taken, the time and location of their action, and the feedback the action has resulted in. The data user should easily understand the kinds of actions available and that there is a relationship between a listener triggering a player event and a subscriber from the adjacent listener domain. Data products provide a formal representation of such semantics. In addition to understanding the semantics, data users need to understand how exactly the data is presented to them. How is it serialized, and how can they access and query the data syntactically? What kind of queries can they execute or how can they read the data objects? They need to understand the schema of the underlying syntax of data. Sample datasets and example consumer codes ideally accompany this information. Examples accompanied with a formalized description of the data improve data users’ understanding. Additionally, dynamic and computational documents like computational notebooks are great companions to tell the story of a data product. Computational notebooks include documentation of the data, as well as code to use it, with immediate feedback that visually demonstrates the code’s result. Lastly, understanding is a social process. We learn from each other. Data products facilitate communication across their users to share their experience and how they take advantage of the data product. Understanding a usable data product requires no hand-holding. A self-serve method of understanding is a baseline usability characteristic. Trustworthy and truthful [Trust is] a confident relationship with the unknown.3 ​—​Rachel Botsman No one will use a product that they can’t trust. So what does it mean to trust a data product, and more importantly what does it take to trust? To unpack this, I like to use the concept of trust offered by Rachel Botsman: the bridge between the known and the unknown. A data product needs to close the gap between what data users know confidently about the data, and what they don’t know but need to know to trust it. While the prior characteristics like understandability and discoverability close this gap to a degree, it takes a lot more to trust the data for use. The data users need to confidently know that the data product is truthful, that it represents the fact of the business correctly. They need to confidently know how closely data reflects the reality of the events that have occurred, the probability of truthfulness of the aggregations and projections that have been created from business facts. A piece of closing the trust gap is to guarantee and communicate data products’ service-level objectives (SLOs)—objective measures that remove uncertainty surrounding the data. Data product SLOs include, among others: Interval of change How often changes in the data are reflected Timeliness The skew between the time that a business fact occurs and becomes available to the data users Completeness Degree of availability of all the necessary information Statistical shape of data Its distribution, range, volume, etc. Lineage The data transformation journey from source to here Precision and accuracy over time Degree of business truthfulness as time passes Operational qualities Freshness, general availability, performance In the traditional approaches to data management, it’s common to extract and onboard data that has errors, does not reflect the truth of the business, and simply can’t be trusted. This is where the majority of the efforts of centralized data pipelines are concentrated, cleansing data after ingestion. In contrast, data mesh introduces a fundamental shift that the owners of the data products must communicate and guarantee an acceptable level of quality and trustworthiness—specific to their domain—as an intrinsic characteristic of their data product. This means cleansing and running automated data integrity tests at the point of the creation of a data product. Providing data provenance and data lineage—the data journey, where it has come from and how it got here—as the metadata associated with each data product helps consumers gain further confidence in the data product. The users can evaluate this information to determine the data’s suitability for their particular needs. I’m of the opinion that once the discipline of building trustworthiness in each data product is established, there is less need for establishing trust through investigative processes and applying detective techniques navigating a lineage tree. Having said that, data lineage will remain an important element in a few scenarios, such as postmortem root-cause analysis, debugging, data compliance audits, and evaluation of data’s fitness for ML training. Natively accessible Depending on the data maturity of the organization there is a wide spectrum of data user personas in need of access to data products. The spectrum spans a large profile of users, as displayed in Figure 3-4: data analysts who are comfortable with exploring data in spreadsheets, data analysts who are comfortable with query languages to create statistical models of data as visualizations or reports, data scientists who curate and structure data and consume data frames to train their ML models, and data-intensive application developers who expect a real-time stream of events or pull-based APIs. This is a fairly wide spectrum of users with equally diverse expectations of how to access and read data. Figure 3-4. Example of a wide spectrum of data product users with different data access patterns There is a direct link between the usability of a data product and how easily a particular data user can access it with their native tools. Hence, a data product needs to make it possible for various data users to access and read its data in their native mode of access. This can be implemented as a polyglot storage of data or by building multiple read adapters on the same data. For example, the play events data product needs to natively support reading data through SQL query to satisfy a data analyst’s native mode of access, publishing a stream of events for data-intensive applications, and columnar files for data scientists. Interoperable One of the main concerns in a distributed data architecture is the ability to correlate data across domains and stitch them together in wonderful and insightful ways: join, filter, aggregate. The key for an effective composability of data across domains is following standards and harmonization rules that allow linking data across domains easily. Here are a few things data products need to standardize to facilitate interoperability and composability: Field type A common explicitly defined type system Polysemes identifiers Universally identifying entities that cross boundaries of data products Data product global addresses A unique global address allocated to each data product, ideally with a uniform scheme for ease of establishing connections to different data products Common metadata fields Such as representation of time when data occurs and when data is recorded Schema linking Ability to link and reuse schemas—types—defined by other data products Data linking Ability to link or map to data in other data products Schema stability Approach to evolving schemas that respects backward compatibility For example, let’s look at managing polysemes identifiers. At Daff, artist is a core business concept that appears in different domains. An artist, while remaining the same global entity, has different attributes and likely different identifiers in each domain. The play events data product recognizes the artist differently from the payment domain—taking care of invoices and payments for artists’ royalties. In order to correlate data about an artist across different domain data products Daff needs to agree on how they identify an artist, globally, across all data products. Chapter 5 covers the topic of global standards and protocols applied to each data product. Interoperability is the foundation of any distributed system design, and data mesh is no exception. Valuable on its own I think it’s fairly obvious that a data product must be valuable—it should have some inherent value for the data users in service of the business and customers. After all, if the data product owner can’t envisage any value out of the data product, perhaps it’s best not to create one. Having said that, it’s worth calling out that a data product should carry a dataset that is valuable and meaningful on its own—without being joined and correlated with other data products. Of course, there is always higher-order meaning, insight, and value that can be derived by correlating multiple data products. However, if a data product on its own serves no value on its own, it should not exist. While this may sound obvious, there is a common antipattern when migrating from a warehouse architecture to data mesh: directly mapping warehouse tables to data products can create data products with no value. In the data warehouse, there are glue (aka facts) tables that optimize correlation between entities. These are identity tables that map identifiers of one kind of entity to another. Such identity tables are not meaningful or valuable on their own—without being joined to other tables. They are simply mechanical implementations to facilitate joins. On the contrary, there are no mechanical data products that solely exist to enable the machines to correlate information across the mesh. Machine optimizations such as indices or fact tables must be automatically created by the platform and hidden from the product products. Secure Data users access a data product securely and in a confidentiality-respecting manner. Data security is a must, whether the architecture is centralized or distributed. However, in a distributed architecture like data mesh, the access control is validated by the data product, right in the flow of data, access, read, or write. The access control policies can change dynamically. They continuously get evaluated at every point of access to data products. Additionally, access to a data product is not quite binary—whether the user can see or can’t see the data. In many cases while the user may not be able to see the actual records, it might have sufficient permissions to see the shape of the data and evaluate it using its statistical characteristics. Access control policies can be defined centrally but enforced at runtime by each individual data product. Data products follow the practice of security policy as code. This means to write security policies in a way that they can be versioned, automatically tested, deployed and observed, and computationally evaluated and enforced. A policy described, tested, and maintained as code can articulate various security-related concerns such as the following, among others: Access control Who, what, and how data users—systems and people—can access the data product Encryption What kinds of encryption—on disk, in memory, or in transit—using which encryption algorithm, and how to manage keys and minimize the radius of impact in case of breaches Confidentiality levels What kinds of confidential information, e.g., personally identifiable information, personal health information, etc., the data product carries Data retention How long the information must be kept Regulations and agreements GDPR, CCPA, domain-specific regulations, contractual agreements Transition to Data as a Product In working with my clients, I have found that they are overwhelmingly receptive to data mesh principles, often questioning “Why didn’t I think of it myself?” or occasionally saying, “We have been doing something similar but not quite the same.” The principles appear to be intuitive and rather natural next steps in their organization’s tech modernization journey, an extension to modernization of the operational aspect of organizations, e.g., moving toward domain-oriented ownership of capabilities with microservices and treating internal products like operational APIs as products. However, a sense of discomfort arises when they go deeper into what it actually takes to implement the transformation toward data mesh. I found in my conversations with data mesh early implementers that while they verbalize the principles and their intention to implement them, the implementations remain heavily influenced by the familiar techniques of the past. For this reason, I have decided to include a number of thought-provoking transition statements as well as a few pragmatic steps to crystalize the differences between the existing paradigm and truly owning data as a product. I invite you to think of new transition statements that I likely have missed here. Include Data Product Ownership in Domains In the past decade, teams have continued to move from being functionally divided to becoming cross-functional. The DevOps movement has led to closing the gap between building and operating business services and forming dev and ops cross-functional teams. Customer-obsessed product development has brought the design and product ownership closer to the developers. The introduction of analytical data as a product adds to the list of existing responsibilities of cross-functional domain teams and expands their roles to: Data product developer The role responsible for developing, serving, and maintaining the domain’s data products as long as the data products live and are being used Data product owner The role accountable for the success of a domain’s data products in delivering value, satisfying and growing the data users, and maintaining the life cycle of the data products Define these roles for each domain and allocate one or multiple people to the roles depending on the complexity of the domain and the number of its data products. Reframe the Nomenclature to Create Change One commonly used term in data analytics is ingestion, receiving data from an upstream source—often an untrustworthy source that has egested data as a byproduct of its operation. It’s now the job of the downstream pipeline to ingest, clean, and process the data before it can be consumed to generate value. Data mesh suggests reframing receiving upstream data from ingestion to consumption. The subtle difference is that the upstream data is already cleansed, processed, and served ready for consumption. The change of language creates a new cognitive framing that is more aligned with the principle of serving data as a product. Relatedly, the word extraction used in ETL (extract, transform, load) and its other variations needs to be critically evaluated. Extraction evokes a passive role for the provider and an intrusive role for the consumer. As we know, extracting data from an operational database that is not optimized for external use other than its own application creates all kinds of pathological coupling and a fragile design. Instead, we can shift the language to publish, serve, or share. This implies shifting the implementation of data sharing from accessing raw databases to intentionally sharing domain events or aggregates. By now you probably have picked up on my emphasis on language and the metaphors we use. George Lakoff—professor of cognitive science and linguistics at UC Berkeley—in his book, Metaphors We Live By, elegantly demonstrates the consequence of shifting our language around the concept of argument, from war to dance. Imagine the world we would live in and the relationships we would nurture, if instead of winning an argument, losing and gaining argument ground, and attacking the weak points of an argument, we would, as dancers, perform a balanced and aesthetically pleasing argument, expressing our ideas and emotions through the beautiful and collaborative ritual of dancing. This unexpected reframing of language has a profound behavioral impact. Think of Data as a Product, Not a Mere Asset “Data is an asset.” “Data must be managed like an asset.” These are the phrases that have dominated our vernacular in big data management. The metaphor of asset used for data is nothing new. After all, for decades, TOGAF, a standard of the Open Group for Enterprise Architecture methodologies and frameworks, explicitly has penciled in “Data Is an Asset” as the first principle of its data principles. While this is a rather harmless metaphor on the surface, it has shaped our perceptions and actions toward negative consequences, for example, our actions toward how we measure success. Based on my observations, data as an asset has led to measuring success by vanity metrics—metrics that make us look or feel good but don’t impact performance—such as the number of datasets and tables captured in the lake or warehouse, or the volume of data. These are the metrics I repeatedly come across in organizations. Data as an asset promotes keeping and storing data rather than sharing it. Interestingly, TOGAF’s “Data Is an Asset” principle is immediately followed by “Data Is Shared.” I suggest the change of metaphor to data as a product, and a change of perspective that comes with that, for example, measuring success through adoption of data, its number of users, and their satisfaction using the data. This underscores sharing the data versus keeping and locking it up. It puts emphasis on the continuous care that a quality product deserves. I invite you to spot other metaphors and vocabulary that we need to change to construct a new system of concepts for data mesh. Establish a Trust-But-Verify Data Culture Data as a product principle implements a number of practices that lead to a culture where data users, by default, can trust the validity of the data and put their focus on verifying its fitness for their use cases. These practices include introducing a role for long-term ownership of a data product, accountable for the integrity, quality, availability, and other usability characteristics of the data product; introducing the concept of a data product that not only shares data but also explicitly shares a set of objective measures such as timeliness, retention, and accuracy; and creating a data product development process that automates testing of the data product. Today, in the absence of these data-as-a-product practices, data lineage remains a vital ingredient for establishing trust. Data users have been left with no choice but to assume data is untrustworthy and requires a detective investigation through its lineage before it can be trusted. This lack of trust is the result of the wide gap between data providers and data users due to the data providers’ lack of visibility to the users and their needs, lack of long-term accountability for data, and the absence of computational guarantees. Data-as-a-product practices aim to build a new culture, from presumption of guilt to the Russian proverb of trust but verify. Join Data and Compute as One Logical Unit Let’s do a test. When you hear the word data product, what comes to your mind? What shape? What does it contain? How does it feel? I can guarantee that a large portion of readers imagine static files or tables, columns and rows, some form of storage medium. It feels static. It’s accumulated. Its content is made of bits and bytes that are representative of the facts, perhaps beautifully modeled. That is intuitive. After all, by definition datum—singular form—is a “piece of information.”4 This perspective results in the separation of code (compute) from data, in this case, separation of the code that maintains the data, creates it, and serves it. This separation creates orphaned datasets that decay over time. At scale, we experience this separation as data swamps—a deteriorated data lake. Data mesh shifts from this dual mode of data versus code to data and code as one architectural unit, a single deployable unit that is structurally complete to do its job, providing the high-quality data of a particular domain. One doesn’t exist without the other. Data and code coexisting as one unit is not a new concept for people who have managed microservices architecture. The evolution of operational systems has moved to a model in which each service manages its code and data, schema definition, and upgrades. The difference between an operational system is the relationship between the code and its data. In the case of microservices architecture, data serves the code. It maintains state so that code can complete its job, serving business capabilities. In the case of a data product and data mesh this relationship is inverse: code serves data. The code transforms the data, maintains its integrity, governs its policies, and ultimately serves it. Note that the underlying physical infrastructures that host code and data are kept separate. Recap The principle of data as a product is a response to the data siloing challenge that may arise from the distribution of data ownership to domains. It is also a shift in the data culture toward data accountability and data trust at the point of origin. The ultimate goal is to make data simply usable. The chapter explained eight nonnegotiable baseline usability characteristics of data products including discoverability, addressability, understandability, trustworthiness, native accessibility, interoperability, independently valuable, and security. I introduced the role of data product owner—someone with an intimate understanding of the domain’s data and its consumers—to assure continuity of ownership of data and accountability of success metrics such as data quality, decreased lead time of data consumption, and in general data user satisfaction through net promoter score. Each domain includes a data product developer role, responsible for building, maintaining, and serving the domain’s data products. Data product developers will be working alongside their fellow application developers in the domain. Each domain team may serve one or multiple data products. It’s also possible to form new teams to serve data products that don’t naturally fit into an existing operational domain. Data as a product creates a new system of the world, where data is and can be trusted, built, and served with deep empathy for its users, and its success is measured through the value delivered to the users and not its size. This ambitious shift must be treated as an organizational transformation. I will cover the organizational transformation in Part V of this book. It also requires an underlying supporting platform. The next chapter looks into the platform shift to make data as a product feasible.


既存の分析データアーキテクチャの長期的な課題の一つは、データの使用に伴う高い摩擦とコストです
品質の高いデータを見つけたり理解したり信頼したり探求したり結局は利用するためには、多くの時間と労力がかかります
この摩擦については、数々の調査結果も示唆しています
データサイエンスプラットフォーム企業であるAnacondaの最新のレポート「データサイエンスの現状2020」では、データサイエンティストの時間のほぼ半分がデータの準備に費やされていることが明らかになりました
つまり、データの読み込みやクレンジング作業です
この問題が解決されなければ、データメッシュではドメイン（データを提供する場所やチームの数）が増えることでますます悪化します
組織のデータ所有権をビジネスドメインに委ねることは、アクセシビリティ、利用性、調和に関する重要な懸念を引き起こします
さらなるデータの隔離化とデータの利用性の低下は、データメッシュの最初の原則であるドメイン指向の所有権の潜在的な望ましくない結果です
データとしての製品の原則は、これらの懸念に対応しています
データメッシュの2番目の原則であるデータとしての製品は、ドメイン指向のデータに製品思考を適用して、そのような利用性の摩擦を取り除き、データユーザー（データサイエンティスト、データアナリスト、データエクスプローラーなど）の体験を本当に楽しいものにします
データとしての製品は、ドメインで提供される分析データを製品として扱い、そのデータの消費者を顧客として扱うことを期待します
さらに、データとしての製品は、データメッシュの価値を解き放ち、偶然的な使用や意図的な使用の可能性を大幅に増加させる組織のデータの価値を実現します
プロダクト開発とマネジメントの著名な思想リーダーであるマーティ・ケーガンの著書「INSPIRED」で、成功した製品には3つの共通の特徴があることを納得できる証拠として提供されています
それは、実現可能性、価値性、利用性です
データとしての製品原則は、データを貴重で使いやすいものにするための標準化された特性を具現化した新しい概念であるデータプロダクトを定義しています
図3-1は、このポイントを視覚的に示しています
この章では、これらの特性を紹介します
第4章では、データプロダクトの構築を実現する方法について説明します
図3-1. データ製品はマーティ・ケーガンの成功した製品の特性の交差点に存在する データが製品となるためには、それがCarganの利用性、実現可能性、価値に関するヴェン図の交差点にうまく合致する一連のルールを守り、一連の特性を備えている必要があります
データが製品として成立するためには、それ自体だけでなく他のデータ製品との協力も含めて、そのユーザーにとって価値がある必要があります
また、データはユーザーに対して共感を示し、利用性と完全性に対して責任を負う必要があります
データを製品として扱う原則が単なる良い意図だけでは実現されないと認めます
そのため、この原則では、ドメインに新しい役割であるドメインデータ製品オーナーやデータ製品開発者を導入し、データ製品の作成、提供、普及に責任を持ちながら、データ製品のアクセシビリティ、品質、利用性、可用性の特定の目標尺度を維持します
第16章では、これらの新しい役割の詳細についてさらに説明します
過去のパラダイムと比べて、データとしての製品は責任のモデルを逆転させています
データレイクやデータウェアハウスのアーキテクチャでは、品質と完全性を持ったデータの作成の責任はソースから下流に位置し、集中的なデータチームによって持たれています
データメッシュはこの責任をデータのソースに近づけます
この移行はデータメッシュに特有のものではありません
実際、過去10年間にわたり、テストやオペレーションにおいても同様の傾向が見られました
問題に取り組むのは、ソースに近い場所で行う方が費用が安く効果的だからです
私は、メッシュで共有されるものが単なるデータではなく、データ製品であるとさえ言いたいとまで思っています
補足 データとしての製品は、データのモデリングと共有に製品思考を適用することを意味します
これは製品の販売とは混同されるべきではありません
データに製品思考を適用する方法について詳しく見ていきましょう
データへのプロダクト思考の適用

過去10年間、高いパフォーマンスを発揮する組織は、外部の技術と同様に、内部のオペレーション技術をプロダクトとして扱うという考え方を取り入れてきました
彼らは内部の開発者を顧客とし、その満足度を成功の指標としています
特に2つの領域で、このトレンドは強く受け入れられています
まず、内部プラットフォームにプロダクトマネジメントの手法を適用することで、内部開発者が内部プラットフォームの上にソリューションを構築・ホストする能力が加速されます（例：Spotify Backstage）
そして、APIをプロダクトとして扱い、APIを発見可能で理解しやすく、簡単にテストできるようにし、開発者の最適なエクスペリエンスを保証するためにAPIを構築する（例：Square APIs）
内部技術へのプロダクト思考の魔法の要素の適用は、内部の消費者（つまり、開発者仲間）との共感の確立から始まり、彼らと協力して体験を設計し、使用量の指標を収集し、使いやすさを維持するために内部の技術ソリューションを持続的に改善することです
強力なデジタル組織は、開発者や最終的にはビジネスに価値のある内部ツールの構築に大量のリソースと注意を割り当てます


興味深いことに、共感やデータを製品として扱うこと、およびその利用者を顧客とみなすことという魔法の要素は、ビッグデータのソリューションでは見落とされています
オペレーションチームはまだデータをビジネスの副産物として捉えており、それを他の誰か（たとえばデータチーム）が受け取り、製品に再利用することを任せています
一方、データメッシュのドメインチームは、データに対して同様の厳密さでプロダクト思考を適用し、最良のデータユーザーエクスペリエンスを追求しています
例えば、ダフというプロダクトの重要なドメインの1つはメディアプレイヤーです
メディアプレイヤードメインは、誰がいつ、どこでどの曲を再生したかといった重要なデータを提供します
この情報に対していくつかの異なるキーデータユーザーがいます
例えば、メディアプレイヤーサポートチームは、ほぼリアルタイムのイベントに関心を持ち、迅速にカスタマーエクスペリエンスの劣化を引き起こすエラーを見つけ出し、サービスを回復させたり、問い合わせに的確に対応したりするために使用します
一方、メディアプレイヤーデザインチームは、長期間にわたるリスナーの旅をデータストーリーとして伝える集計された再生イベントに興味を持ち、総合的なより良いリスナーエクスペリエンスに向けてより魅力的な機能を持つメディアプレイヤーを改善します
これらのユースケースと他のチームが必要とする情報を理解することで、メディアプレイヤードメインは組織全体に対して2つの異なるタイプのデータを製品として提供します
ほぼリアルタイムの再生イベントは無限のイベントログとして公開され、集計された再生セッションはオブジェクトストア上のシリアライズされたファイルとして公開されます
これがデータに対するプロダクト所有権の適用です
図3-2はメディアプレイヤードメインのデータ製品の例を示しています
図3-2. データ製品の例
想像できるように、データには他の製品と異なる独特の要素があります
データ製品の所有権と他のタイプの製品の違いは、データの使用ケースの非制約性、特定のデータが他のデータと組み合わさり、最終的に洞察とアクションに転換される方法にあります
データ製品所有者は、常に現在のデータの有効な使用ケースを把握し、計画することができますが、今日生成されるデータの未知の将来の使用ケースも存在し、想像を超える可能性があります
これは、ソースに合わせたドメインではより当てはまり、消費者に合わせたドメインではそれほど当てはまりません
ソースに合わせたデータは、ビジネスの相互作用やイベントの現実をキャプチャし続けます
これらのデータは、未来のデータユーザーによって使用され、変換され、再解釈されることがあります
ソースに合わせたデータ製品は、即時に知られている使用ケースと未知の使用ケースとのバランスを取る必要があります
それらは、自分のデータを使用する方法についてあまり仮定せずに、できるだけビジネスの現実をモデリングするよう努力する必要があります
例えば、プレイイベントを無限の高解像度のログとしてキャプチャすることは安全な選択肢です
これにより、今日キャプチャされるデータから他の変換を作成し、新たな洞察を導き出すことができる将来のユーザーの範囲が広がります
これがデータ製品の設計とソフトウェア製品の設計の主な違いです
 データ製品の基本的な使いやすさの特徴 私の意見では、データ製品が有用とみなされるためには、譲れない基本的な特性のセットが含まれていると考えます
これらの特性は、ドメインやアーキタイプに関係なく、すべてのデータ製品に適用されます
これらの基本的なデータ製品の使いやすさの特性を、基本的なデータ製品の使いやすさの属性と呼んでいます
 フィギュア3-3には、データ製品の使いやすさの属性がリストされています
 フィギュア3-3
データ製品の基本的な使いやすさの属性（DAUTNIVS） これらの使いやすさの特性は、データユーザーの体験を中心にとらえています
これらは技術的な能力を表現するものではありません
 パートIV「データクォンタムアーキテクチャの設計」では、データ製品の技術的特性について説明します
 このセクションにリストされている基本的な特性は、これまでFAIRデータとして知られていたものに追加されたものです
FAIRデータとは、見つけやすさ、アクセス可能性、相互運用性、再利用性の原則たとえば、すべての再生イベントを無限の高解像度ログとしてキャプチャすることは、安全な選択肢です
これにより、将来のユーザーは、今日キャプチャされたデータから他の変換を作成し、新しい洞察を得ることができます
これは、データ製品の設計とソフトウェア製品の設計の主な違いです
 データ製品のベースラインの使いやすさ属性 私の意見では、有用とされるためには、データ製品が備えるべき議論の余地のないベースラインの特性があります
これらの特性は、ドメインや原型に関係なく、すべてのデータ製品に適用されます
私はこれらをベースラインのデータ製品の使いやすさ属性と呼びます
すべてのデータ製品は、これらの特性をメッシュの一部として取り込んでいます
 図3-3にはデータ製品の使いやすさ属性がリストアップされています
 図3-3. データ製品のベースラインの使いやすさ属性（DAUTNIVS） これらの使いやすさの特性は、データユーザーの体験を重視しています
これらは技術上の能力を示すものではありません
第IV部「データ量子アーキテクチャの設計」では、データ製品の技術的特性をカバーしています
 このセクションにリストされているベースラインの特性は、過去に「FAIRデータ」として知られていたものに追加されたものです
つまり、見つけやすさ、アクセス可能性、相互運用性、再利用性の原則に適合するデータです
1また、分散データ所有を実現するために必要な特性も紹介しています
データユーザーの立場に立ち、これらの属性を見ていきましょう
 発見可能な 良いデザインの最も重要な特性のうちの2つは、発見可能性と理解可能性です
 発見可能性：どのようなアクションが可能で、それらをどこで、どのように実行するかを理解できるでしょうか？ 理解可能性：それは全て何を意味しているのでしょうか？ どのように使用するべきなのでしょうか？ 異なる制御や設定は何を意味しているのでしょうか？ ドン・ノーマン データユーザーが旅を始める最初のステップは、彼らに利用可能なデータの世界を発見し、探索して「理想的なもの」を見つけることです
したがって、データの使いやすさの最初の属性は、簡単に見つけられることです
データユーザーは、利用可能なデータ製品を探索し、望ましいデータセットを検索して見つけ、それらを探索し、使用する際の信頼を得る必要があります
 伝統的な発見可能性の実装では、利用可能なデータセットを中央の登録リストやカタログにリストアップし、各データセットに関する追加情報、所有者、場所、サンプルデータなどを提供します
この情報は、通常、中央のデータチームまたはガバナンスチームによって後から手配されます
 データメッシュにおけるデータ製品の発見可能性は、データ製品自体が意図的に発見可能性情報を提供するシフトレフトのソリューションを採用しています
データメッシュは、メッシュの動的なトポロジー、進化し続けるデータ製品、利用可能なデータ製品の規模に頼るため、データ製品自体が標準的な方法でその発見可能性情報を提供する必要があります
各データ製品は、その生涯の異なる段階（ビルド、デプロイ、ラン）で、ソースの元、所有者、タイムリネス、品質メトリクス、サンプルデータなどのランタイム情報、それに応じた情報を共有します
最も重要なのは、データの利用者によって提供された情報、つまりデータによって可能にされる主要なユースケースやアプリケーションです
 「データ製品アーキテクチャの設計方法」の第IV部では、データ製品の発見可能性の技術的な設計について説明します
 アドレス可能 データ製品は、データユーザーに対してプログラム的または手動でアクセスするための永続的で一意のアドレスを提供します
このアドレスシステムは、データとメッシュのトポロジーの動的な性質を受け入れる必要があります
多くのデータ製品の側面が変化し続ける中でも、使用の継続性を保証します
アドレッシングシステムは、データ製品の以下の変化に対応し、データ製品へのアクセスを長期間にわたって維持することができます：データ製品におけるセマンティックと構文の変化、スキーマの進化、時間の経過に伴う新しいデータの継続的なリリース（ウィンドウ）、特定の時間（または時間ウィンドウ）に関連するデータタプルのパーティショニング戦略とグループ化、データへの新しいアクセスモード、データのシリアライズ、表示、およびクエリの新しい方法、ランタイムの動作情報の変更（たとえば、サービスレベル目標、アクセスログ、デバッグログ）
ユニークなアドレスは、ユーザーがすべてのデータ製品にプログラムによって一貫してアクセスできるようにするためのグローバルな規約に従う必要があります
データ製品は、ドキュメンテーション、サービスレベル目標、およびデータを提供するデータ製品に関するすべての情報の入り口となるアドレッサブルな集約ルートを持たなければなりません
理解可能なデータ製品が発見されたら、データユーザーの次のステップはそれを理解することです
これには、基礎データの意味やエンコードされたデータのさまざまな構文など、データ製品の意味を知ることが含まれます
各データ製品は、特定の意味を持つ意味的に一貫したデータを提供します
データユーザーは、この意味を理解する必要があります
データ製品が取り込むエンティティの種類、エンティティ間の関係、および隣接するデータ製品について理解する必要があります
メディアプレーヤーのイベントの例に戻ると、データユーザーはプレーヤーイベントを構成する要素を簡単に理解する必要があります
ユーザー、実行した再生アクション、アクションの発生時刻と場所、およびアクションの結果としてのフィードバックなどです
データユーザーは利用可能なアクションの種類や、リスナーがプレーヤーイベントをトリガーするとリスナードメインからの購読者との関係があることを簡単に理解する必要があります
データ製品は、そのような意味を正式に表現します
セマンティックを理解するだけでなく、データユーザーはデータがどのように提示されるかも理解する必要があります
データがどのようにシリアライズされているのか、データにどのようにアクセスしてクエリを行うことができるのか
基礎データのシンタックスのスキーマを理解する必要があります
サンプルデータセットと例示したコンシューマーコードがこの情報と一緒に提供されると理想的です
具体例はデータユーザーの理解を向上させます
また、計算ノートブックのような動的で計算可能なドキュメントは、データ製品の物語を伝えるための素晴らしい補完となります
計算ノートブックには、データのドキュメンテーションや使用方法のコードが含まれ、コードの結果を視覚的に示す即時フィードバックも提供されます
最後に、理解は社会的なプロセスです
私たちはお互いから学びます
データ製品は、ユーザー間のコミュニケーションを促進し、彼らがデータ製品をどのように活用しているかや経験を共有することができます
使いやすいデータ製品を理解するために、手助けは必要ありません
自己利用可能な理解方法は、ユーザビリティの基本的な特徴です
信頼性と真実性 "[Trust is] a confident relationship with the unknown.（信頼は、未知との確信に基づく関係です
）" - レイチェル・ボツマン 誰もが信頼できない製品を使用しません
では、データ製品を信頼するとはどういうことでしょうか、さらに重要なのは信頼するためにはどうすればいいのでしょうか
この点を解明するために、私はレイチェル・ボツマンが提供する信頼の概念を使って考えてみたいです
信頼は、知っていることと知らないことの間の橋です
データ製品は、データユーザーがデータについて確信を持って知っていることと、それを信頼するために知らなければならないこととのギャップを埋める必要があります
先述の理解可能性や発見可能性はある程度このギャップを埋めますが、実際にはデータを信頼するにはさらに多くの要素が必要です
データユーザーは、データ製品が正確であり、事業の事実を正しく表していると確信する必要があります
データがどれだけ現実の出来事を反映しているか、ビジネス上の事実から作成された集計や予測の真実性の確率を自信を持って知る必要があります
信頼のギャップを埋める一環として、データのサービスレベル目標（SLO）を保証し、伝えることが重要です
SLOは、データに関する不確実性を取り除く客観的な尺度となります
データのSLOには、以下が含まれます：データの変化の間隔、データが変更を反映する頻度、タイムリネス、ビジネス上の事実が発生してからデータユーザーが利用できるようになるまでの時間のズレ、完全性、必要な情報の利用可能度、統計的データの形状、分布、範囲、容量など、アーキテクチャが把握されている、データがソースから配信までどのように変換されたかのデータラインジェ、時間の経過とともにビジネスの真実性がどの程度保たれたか、新鮮さ、一般的な利用可能性、パフォーマンスなどの操作特性
従来のデータ管理手法では、エラーを含むデータを抽出し、取り込んだ後、ビジネスの真実を反映していないデータや信頼できないデータが一般的でした
これが、集中型データパイプラインの大半の取り組みが集中して行われる点です
対照的に、データメッシュは、データ製品のオーナーが、データ製品の固有の特性として、品質と信頼性の許容レベルをコミュニケートし、保証しなければならない、という根本的な変化をもたらします
つまり、データ製品の作成時点で自動的なデータの正確性テストを実行し、クレンジングを行うことを意味します
各データ製品に関連付けられたメタデータとして、データの出所やここに至るまでの経緯を提供することで、データ製品への信頼を得ることができます
ユーザーは、この情報を評価して、データが自分の特定のニーズに適しているかどうかを判断することができます
私は、各データ製品に信頼性を構築するという原則が確立されれば、調査プロセスを通じた信頼を確立する必要性は減ると考えています
そしてラインジーツリーを辿る検出手法を適用する必要性も減るでしょう
ただし、データの利用可能性のポストモーテムのルート原因分析、デバッグ、データコンプライアンス監査、およびMLトレーニング用データの適合性評価など、一部のシナリオでは、データのラインジーンは引き続き重要な要素となるでしょう
ネイティブにアクセス可能なデータ製品の利用者によっては、データ製品の使いやすさと特定のデータユーザーがネイティブツールを使ってアクセスできるかどうかとの直接的な関連性があります
したがって、データ製品は、異なるデータアクセス方法を持つさまざまなデータユーザーがネイティブなアクセス方法でデータにアクセスして読み取ることが可能にする必要があります
これは、データのポリグロットストレージとして実装するか、複数の読み取りアダプタを同じデータに構築することによって実現できます
例えば、プレイイベントのデータ製品では、データアナリストのネイティブなアクセス方法でSQLクエリを介してデータを読み取ること、データ密度の高いアプリケーションにイベントのストリームを公開すること、データサイエンティストに対してカラムファイルを提供することなどがネイティブにサポートされる必要があります
相互運用可能なデータアーキテクチャにおいて、異なるドメイン間のデータを相関させ、結び付けて素晴らしい洞察を提供する能力は、主要な関心事の1つです
結合、フィルタリング、集計などが含まれます
データの相互運用性と組み立て性を実現するためには、データ間で簡単にリンクを作成できるよう、標準と調和ルールに従うことが重要です
以下に、相互運用性と組み立て性を促進するためにデータ製品が標準化する必要のあるいくつかの要素を示します


フィールドのタイプ
明示的に定義された共通のタイプシステム

多義性のある識別子
データ製品の境界を越えてエンティティを一意に識別する

データ製品のグローバルなアドレス
各データ製品に割り当てられた一意のグローバルアドレス
異なるデータ製品への接続の容易さのために理想的には統一されたスキームがあるとよい


共通のメタデータフィールド
データが発生する時点やデータが記録される時点の表現など

スキーマのリンキング
他のデータ製品で定義されたスキーマ（タイプ）にリンクして再利用する能力

データのリンキング
他のデータ製品のデータにリンクやマップする能力

スキーマの安定性
進化するスキーマに対するアプローチで、後方互換性を尊重する

例えば、多義性のある識別子の管理を見てみましょう
Daffでは、アーティストはさまざまなドメインで重要なビジネスコンセプトです
アーティストは、同じグローバルエンティティでありながら、各ドメインごとに異なる属性とおそらく異なる識別子を持っています
プレイイベントのデータ製品は、アーティストを支払いドメインとは異なる方法で認識しています（アーティストのロイヤリティのための請求書や支払いに関するものです）
さまざまなドメインのデータ製品に関するアーティストのデータを関連付けるためには、Daffはすべてのデータ製品を通じてアーティストを一意に識別する方法に合意する必要があります
第5章では、各データ製品に適用されるグローバルな標準とプロトコルについて説明します
相互運用性は、分散システム設計の基礎であり、データメッシュにおいても例外ではありません


単体で価値がある
データ製品は、有価値でなければなりません
ビジネスと顧客のために、データユーザーにとって何らかの固有の価値を持つべきです
結局のところ、データ製品のオーナーがデータ製品から何の価値も想像できないのであれば、作成しない方がよいかもしれません
ただし、言うまでもなく、データ製品は他のデータ製品と結びつけることなく、単独で有益かつ意味のあるデータセットを保持する必要があります
もちろん、複数のデータ製品を相関させることで高次の意味、洞察、価値が得られる場合もあります
しかし、データ製品自体が単独で何の価値も持たない場合は存在すべきではありません
これは明らかなことかもしれませんが、データウェアハウスアーキテクチャからデータメッシュに移行する際には、データ製品の値がないデータ製品が作成されるという共通のアンチパターンがあります
データウェアハウスでは、エンティティ間の相関を最適化する接合テーブル（ファクトテーブル）があります
これらは、ある種のエンティティの識別子を別の種類のエンティティにマッピングする識別子テーブルです
このような識別子テーブルは、他のテーブルと結合しない限り意味や価値を持ちません
単に結合を容易にするための機械的な実装です
逆に、メッシュ全体で情報を相関させるために単独で存在する機械的なデータ製品はありません
インデックスやファクトテーブルなどの機械的な最適化は、プラットフォームによって自動的に作成され、製品からは隠されます


セキュア
データユーザーは、データ製品に安全かつ機密性を尊重した方法でアクセスします
データセキュリティは、アーキテクチャが集中型であろうと分散型であろうと必須です
ただし、データメッシュのような分散アーキテクチャでは、アクセス制御はデータ製品によってデータの流れ、アクセス、読み取り、書き込みの過程で検証されます
アクセス制御ポリシーは動的に変更されることがあります
データ製品へのアクセスごとに継続的に評価されます
さらに、データ製品へのアクセスは完全にバイナリではありません
ユーザーは実際のレコードを見ることができないかもしれませんが、統計的な特性を使用してデータの形状を見ることができる十分な許可を持っている場合もあります
アクセス制御ポリシーは中央で定義されますが、個々のデータ製品によって実行時に強制されます
データ製品は、セキュリティポリシーをコードとして記述するという手法に従います
つまり、セキュリティポリシーをバージョン管理、自動テスト、展開および観察、計算的評価および強制実行できるように記述することを意味します
コードとして記述、テスト、および保守されるポリシーは、以下のようなセキュリティに関する懸念事項を明確に述べることができます
 アクセス制御 - 誰が、何によって、どのようにデータ製品にアクセスできるか 暗号化 - ディスク上、メモリ上、または転送中でのどのような種類の暗号化を使用し、どのように鍵を管理し、侵害が発生した場合の影響範囲を最小限に抑えるか 機密性レベル - データ製品が含む個人識別情報、個人保健情報などのような機密情報の種類 データの保持期間 - 情報を保持する期間 GDPR、CCPA、ドメイン固有の規制、契約上の合意 データを製品として扱うための移行 私はクライアントとの作業中に、データメッシュの原則に非常に受け入れられると感じています
彼らはしばしば「なぜ自分自身で考えつかなかったのか？」と問いかけたり、「似たようなことをやっていたが、完全に同じではなかった」と言ったりします
これらの原則は、彼らの組織の技術の近代化の旅における直感的で自然な次のステップのように思われます
ただし、データメッシュへの変換を実装するために実際に必要なことについて詳しく考えると、不快感が生じます
データメッシュの初期導入者との会話の中で、彼らは原則と実装の意思を言葉にする一方で、実装は過去の熟知した手法の影響を強く受けていることがわかりました
そのため、私は既存のパラダイムとデータを本当に製品として所有する違いを明確にするために、いくつかの考えを刺激する移行の文や実用的な手順を追加することにしました
ここで見落としているであろう新しい移行の文を考えるように皆様にお願いします
それぞれのドメインにデータ製品の所有権を含める 過去10年間、チームは機能ごとに分割されていたのから、クロスファンクションのチームになるように移行してきました
DevOpsの運動は、ビジネスサービスの構築と運用のギャップを埋め、DevとOpsのクロスファンクションのチームを形成することにつながりました
顧客中心のプロダクト開発により、設計と製品所有権が開発者に近づいてきました
分析データを製品として導入することは、クロスファンクションのドメインチームの既存の責任のリストに加わり、以下の役割を拡大します
 データ製品開発者 - データ製品の開発、提供、および維持を担当する役割 データ製品オーナー - ドメインのデータ製品の成功、価値の提供、データユーザーの満足と成長、およびデータ製品のライフサイクルの維持に責任を持つ役割 これらの役割を各ドメインに定義し、ドメインの複雑さとデータ製品の数に応じて1人または複数の人物を役割に割り当てる
変化を生み出すために用語のリフレーミングを行う データ分析において一般的に使用される用語の一つは、上流ソース（しばしば信頼性のないソース）からデータを受け取る「ingestion」です
これは、上流データを下流パイプラインが受け入れ、クリーンアップして処理することで、生成価値を生み出すために消費される前に準備する役割です
データメッシュでは、上流データを「ingestion」ではなく「consumption（消費）」という枠組みでとらえることを提案しています
微妙な違いは、上流データが既にクリーンナップ、処理済みであり、消費のためにすでに準備ができているということです
言語の変化により、データを製品として提供する原則とより一致した新しい認識の枠組みが生まれます
関連して、ETL（抽出、変換、読み込み）で使用される単語の抽出も厳密に評価する必要があります
抽出は、プロバイダーの受動的な役割と消費者の侵入的な役割を連想させます
私たちは、自身のアプリケーション以外で使用するために最適化されていない運用データベースからデータを抽出すると、あらゆる種類の病理的な連結と脆弱な設計を引き起こします
代わりに、公開、提供、共有といった言葉に変えることができます
これにより、データ共有の実装を、生のデータベースへのアクセスから故意にドメインイベントや集約の共有へと移行することが示唆されます
今まで言語と私たちが使用する比喩に重点を置いたことに気付いたかもしれません
ジョージ・レイコフ（カリフォルニア大学バークレー校の認知科学と言語学の教授）は、彼の著書『Metaphors We Live By』で、戦争からダンスへの概念の言語を変えることの結果を洗練された形で示しています
もし私たちが議論を勝つのではなく、議論の足場を失ったり得たりし、議論の弱点を攻撃する代わりに、ダンサーとして、バランスの取れた美しい議論を行い、アイデアや感情を美しく協力的な舞踏の儀式を通じて表現するのであれば、私たちが生きる世界や育てる関係はどのようなものになるでしょうか
この予期しない言語の再構築は、深い行動的な影響を与えます
データを単なる資産ではなく製品と考える 「データは資産である」、「データは資産のように管理されなければならない」といったフレーズが、ビッグデータ管理の私たちの話し言葉を支配してきました
データに対する資産という比喩は、新しいものではありません
なぜなら、数十年にわたり、エンタープライズアーキテクチャの方法論やフレームワークのOpen Groupの標準であるTOGAFでは、「データは資産である」という原則を明示的に掲げてきたからです
これは表面上は比較的無害な比喩ですが、例えば成功を測定するための私たちの行動について否定的な結果につながる影響を与えてきました
私の観察によれば、データを資産として捉えることは、データセットやテーブルの数、データのボリュームなど、外観上は良く見えたり良い気分にさせるメトリック（成果指標）である「虚栄のメトリック」による成功を測定することにつながっています
これらのメトリックは、組織内で繰り返し目にするものです
「データは資産である」という考え方は、データを共有することよりも保持し格納することを促進します
興味深いことに、TOGAFの「データは資産である」という原則は直ちに「データは共有される」という原則に続いています
私は、データを製品としての比喩と、それに伴う視点の変化、例えばデータの採用、利用するユーザー数、およびデータを使用した彼らの満足度を通じて成功を測定することを提案します
これにより、データの共有と保持・ロックアップの強調がなされます
品質の製品に１つの関与トラストを築く 「データを製品とする」という原則には、デフォルトでデータ利用者がデータの妥当性を信頼し、それを自身のユースケースに適したものかどうかを検証することに焦点を当てることができる文化を築くためのいくつかの実践が含まれます
これらの実践には、データ製品の長期的な所有権の役割を導入し、データ製品の品質、可用性、その他の使いやすさの特性に対して責任を持つ要素を含めること、データの共有だけでなく、タイムリネス、保持期間、正確性などの客観的な尺度のセットも明示的に共有するデータ製品の概念を導入すること、そしてデータ製品の開発プロセスにおいてデータ製品のテストを自動化することが含まれています
これらの「データを製品とする」実践がない現在、信頼を確立するためには、データの系譜が重要な要素となります
データユーザーは、データが信頼できないと仮定し、信頼できる前にその派生物を検証するために捜査を行うという選択肢しか残されていません
この不信感は、データプロバイダーとデータユーザーの間の広いギャップに起因しています
データプロバイダーはユーザーやそのニーズに対する視界が不足しており、データに対する長期的な責任や計算上の保証が存在しないためです
データとしてのプロダクトの実践は、ギルティプレシャン（有罪推定）からロシアの格言である信頼して検証するという新しい文化を築くことを目指しています
データとコンピュートを1つの論理ユニットとして結び付けましょう
テストをしてみましょう
データプロダクトという言葉を聞いて、あなたは何を思い浮かべますか？どんな形ですか？何を含んでいますか？どう感じますか？私は多くの読者が静的なファイルやテーブル、列と行、ある種のストレージメディアを想像していることを保証します
それは静的なものです
蓄積されます
その内容は事実を表すビットとバイトから成り立っており、おそらく美しくモデル化されています
それは直感的です
なぜなら、定義上データ（単数形）は「情報の一部」を意味するからです
この視点により、コード（コンピュータ）とデータが分離され、この場合はデータを保持し、作成し、提供するコードが分離されます
この分離により、経時的に退化する孤立したデータセットが作成されます
スケール化すると、この分離をデータスワンプとして体験します - 退化したデータ湖
データメッシュは、データ対コードのこの二重モードから、特定のドメインの高品質なデータを提供するために機能的に完全な単一の展開可能なユニットとして、データとコードを含む一つの建築ユニットとして移行します
一方なしで他方は存在しません
データとコードが1つのユニットとして共存するという概念は、マイクロサービスアーキテクチャを管理してきた人々にとっては新しいものではありません
運用システムの進化により、各サービスが独自のコードとデータ、スキーマ定義、アップグレードを管理するモデルに移行しました
運用システムとの違いは、コードとデータの関係です
マイクロサービスアーキテクチャの場合、データはコードに仕えます
コードが完了するために状態を維持し、ビジネス機能を提供します
データプロダクトとデータメッシュの場合、この関係は逆転しています
コードがデータに仕えます
コードはデータを変換し、その整合性を維持し、ポリシーを管理し、最終的にデータを提供します
コードとデータをホストする基礎となる物理インフラは別々に保持されます
要約 データをプロダクトとして扱う原則は、データ所有権をドメインへ分散させることから生じるデータのシロン化の課題への対応策です
また、データの文化をデータの責任と起源地点でのデータの信頼性に向けた変革へのシフトです
究極の目標は、データを単に使いやすくすることです
この章では、データプロダクトに必須の8つのユーザビリティ特性（発見性、アドレス性、理解可能性、信頼性、ネイティブアクセシビリティ、相互運用性、独立して価値のあるもの、セキュリティ）を説明しました
また、データプロダクトオーナーの役割も紹介しました
データプロダクトオーナーは、ドメインのデータとその利用者を深く理解しており、データの所有と成功メトリクス（データの品質、データ消費のリードタイムの短縮、一般的なデータユーザーの満足度）の責任を確実にします
各ドメインには、データプロダクト開発者の役割もあり、ドメインのデータプロダクトの構築、維持、提供を担当します
データプロダクト開発者は、ドメインのアプリケーション開発者と協力して作業を行います
各ドメインチームは1つまたは複数のデータプロダクトを提供する場合もあります
既存の運用ドメインに自然に適合しないデータプロダクトを提供するために新しいチームを形成することも可能です
データをプロダクトとして扱うことで、データが信頼でき、ディープエンパシーを持ってユーザーに提供され、その成功はユーザーへの提供価値によって測定される新しい世界のシステムが作られます
この野心的な変革は組織の変革として取り扱われる必要があります
本書の第V部では、組織の変革について詳しく説明します
これには基盤となる支援プラットフォームが必要です
次の章では、データを製品として実現するためのプラットフォームのシフトについて探求します