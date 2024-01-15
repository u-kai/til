We have both made it to the last part of this book. Well done! Are you considering adopting data mesh for your organization or helping others to do so? Are you in a position to influence, lead, or manage the execution of data mesh? Do you need some help on how to approach such a transformation organizationally and where to get started? If you answered yes to these questions, this part is for you. I assume that you have a solid understanding of data mesh foundational principles from Part I and understand the motivation and objectives of data mesh from Part II. Scope Data mesh is an element of a data strategy that fosters a data-driven organization to get value from data at scale. Execution of this strategy has multiple facets. It affects the teams, their accountability structure, and delineation of responsibilities between domains and platform teams. It influences the culture, how organizations value and measure their data-oriented success. It changes the operating model and how local and global decisions around data availability, security, and interoperability are made. It introduces a new architecture that supports a decentralized model of data sharing. In this part of the book, I describe a high-level approach to start and guide the execution of this multifaceted change. In Chapter 15, “Strategy and Execution”, I introduce an evolutionary approach to adopt data mesh incrementally, generating value from an increasing number of data products while maturing the platform capabilities. In Chapter 16, “Organization and Culture”, I cover some of the key changes to teams, roles, and performance metrics to create long-term ownership of data products and peer-to-peer data sharing. At the time of writing this book, we are in the early years of adopting data mesh and beginning to create a decentralized organizational model where data sharing and analytics are becoming embedded into business-technology domains. The approach suggested in these chapters is the adaptation of an evolutionary and large-scale transformation model that has been successfully applied in scenarios such as building domain-oriented operational teams and digital platforms. I will share my experience of blockers and enablers specific to data mesh. This is just a starting point, and we have years ahead of us to refine a transformation approach unique to data mesh. Chapter 15. Strategy and Execution The essence of strategy is choosing to perform activities differently than rivals do. ​—​Michael E. Porter Is your organization adopting a data-oriented strategy? Is it planning to offer distinct value to its customers and partners using data and insights? Does your organization have a vision to use data to perform a different set of activities or its current activities differently? Is your organization curious about data mesh as a foundation for implementing this strategy? If you answered yes to these questions, this chapter is for you. In this chapter I describe the relationship between data mesh and an overall data strategy, and how to approach its execution through an iterative, value-driven, and evolutionary approach. But first, a quick self-assessment. Should You Adopt Data Mesh Today? Before you proceed to the rest of this chapter, you may ask yourself, is data mesh the right choice for my organization? Or more importantly, is data mesh the right choice now?1 Figure 15-1 depicts my perspective on these questions. I suggest that organizations that fall in the colored section of the spider graph—scoring medium or high on each of the assessment axes—are best positioned to adopt data mesh successfully at this point in time. You can use the following criteria to self-assess whether data mesh is the right choice for your organization now. Figure 15-1. Criteria to self-assess readiness for data mesh adoption Organizational complexity Data mesh is a solution for organizations that experience scale and complexity, where existing data warehouse or lake solutions have become blockers in their ability to get value from data at scale and across many functions of their business, in a timely fashion and with less friction. If your organization has not reached such a level of complexity and the existing data and analytics architectures meet your needs, data mesh will not offer you much value. Organizations with a large proliferation of data sources and use cases that score moderate or high on data complexity benefit from data mesh. For example, enterprises or scaleups—fast-growing startups—whose business encapsulates many different data domains and a large proliferation of data types collected from many sources and that have a business strategic differentiator heavily dependent on ML and analytical use cases are great candidates for data mesh. Enterprises or scaleups in industries such as retail, banking and finance, insurance, and large tech services fall into this category. For example, consider a healthcare provider or payer. There is an inherent complexity in the diversity of data. It ranges from clinical visits to patient profiles, lab results, pharmaceuticals, and clinical claims. Each of these categories of data can be sourced from multiple parties in an ecosystem of business domains or independent organizations, hospitals, pharmacies, clinics, labs, etc. The data is being used for many use cases: optimized and personalized care models, intelligent insurance authorization, and value-based care. The primary criteria for adopting data mesh are the intrinsic business complexity and proliferation of data sources and data use cases. Data-oriented strategy Data mesh is a solution for an organization planning to get value from data at scale. It requires the commitment of product teams and business units to imagine using intelligent decision making and actions in their applications and services. Such a commitment is possible only when the organization identifies data enabling ML and analytics as a strategic differentiator. A data-oriented strategy then translates down to data-driven applications and services, which then turns into teams’ commitment to data. Without a strategic vision around data usage, it would be difficult to motivate domain teams and business domains to take on the extra responsibility of data sharing. Executive support Data mesh introduces change, and with that comes resistance to change. Any change of such scale experiences detractors and blockers. It will be necessary to make tough decisions between the progress of the platform versus building point solutions to meet a deadline. It demands motivating the organization to change how people work. All this calls for executive support and top-down engagement of leaders. In my experience, the most successful data mesh implementations have had the backing of C-level executives. Data technology at core Implementing data mesh builds on a new relationship between technology and business. It builds on the pervasive application of data in each business domain. Organizations that have data technology at their core use data and AI as a competitive advantage, as the means to expand and reshape their business and not merely as an enabler. Such organizations have the desire and ability to build and create the technology needed to embed data sharing and consumption at the core of each business function. In contrast, organizations that see technology as a supporter of the business and not at the core often delegate and externalize technology capabilities to external vendors and expect to buy and plug in a ready-made solution for their business needs. Such organizations are not yet positioned to adopt data mesh. Early adopter At the time of writing this book, data mesh is enjoying its early years of adoption by innovators and lead adopters. Early adopters of new innovations must have certain characteristics: they are venturesome opinion leaders2 in their industry. They have an appetite for adopting emerging technologies, such as data mesh. Early adopters of a novel approach, particularly one as multidimensional and pervasive as data mesh, demand a spirit of experimentation, taking risks, failing fast, learning, and evolving. Companies that adopt data mesh at this point in time must be willing to begin from the first principles, adapt them to their context, and learn and evolve. On the contrary, organizations that don’t take risks and like to only adopt well-tested, refined, and prescribed playbooks and fall into the late adopter category may need to just wait for some time. Modern engineering Data mesh builds on a foundation grounded in modern software engineering practices. Continuous and automated delivery of software, DevOps practices, distributed architecture, computational policies, and availability of modern data storage and processing stacks—on private or public clouds. Without solid engineering practices and access to open and modern data tooling, it will be challenging to bootstrap data mesh from the ground up. Data mesh shares prerequisites for microservices3 and requires modern data and analytics technology and practices. Additionally, data mesh works with modern technology stacks that are API driven and easy to integrate and aim to enable many smaller teams, as opposed to one large, centralized team. Technologies that enforce centralized data modeling, centralized control of data, and centralized storage of data don’t lend themselves to a successful data mesh implementation. They become a bottleneck and conflict with the distributed nature of data mesh. Domain-oriented organization Data mesh assumes that the adopters either are modern digital businesses with a level of technology and business alignment or are on the journey toward it. It assumes that organizations are designed based on their business domains, where each business domain has dedicated technical teams who build and support the digital assets and products that serve the business domain. It assumes that technology and business teams are in close collaboration to use technology to enable, reshape, and expand the business. For example, a modern bank has dedicated technology groups for digital core banking, digital processing of loans, digital commercial and personal banking, digital credit and risk, etc. There is a close alignment between technology, technical teams, and the business domains. In contrast, organizations with a centralized IT unit that shares people (resources) based on projects across all business units, without continuous and long-term ownership of technical assets for each business domain, are not a good fit for data mesh. Domain-oriented data sharing scales only if there is a domain-oriented tech and business alignment in place. Long-term commitment As you will see in upcoming chapters, the adoption of data mesh is a transformation and a journey. While it can be started with a limited scope and an organizational footprint of a small number of domains, the benefits are compounded when exploited at scale. Hence, you need to be able to create an organizational commitment to the vision and embark on a transformation journey. Data mesh cannot be implemented as one or two one-off projects. If you see yourself aligned with the previous criteria, let’s go ahead and look at how to execute a data strategy using data mesh. Data Mesh as an Element of Data Strategy To understand how data mesh fits into a larger data strategy, let’s look at our hypothetical company, Daff, Inc. Daff’s data strategy is to create an intelligent platform that connects artists and listeners through an immersive artistic experience. This experience is continuously and rapidly improved using data and ML embedded in each and every interaction of the listeners and artists on the platform. This implies capturing data from every single touchpoint on the platform and augmenting every platform feature using ML derived from the engagement data. For example, the platform’s better playlist predictions lead to more listeners with higher engagement, which then leads to more data and hence better predictions—creating a positive self-reinforcing loop for the business. Daff’s ambitious strategy is executed through an organizational-wide data mesh implementation, where the nodes on the mesh can represent data collected and curated from various activities on their platform and the intelligence inferred by machine learning models trained on such data. Figure 15-2 shows the relationship between the data strategy and data mesh. Figure 15-2. The big picture. connecting data mesh execution to the overall data strategy As Figure 15-2 shows, data strategy drives the execution of data mesh. Let’s explore the diagram top down and go through the implementation of data mesh from strategy to the platform, with continuous feedback loops. Data strategy Data strategy defines how uniquely a company takes advantage of data and creates a roadmap of strategic initiatives. For example, Daff’s strategy is based on partnering every content producer’s, artist’s, and listener’s experience with data and ML. Strategic business initiatives and use cases Following the strategy, a set of business initiatives, use cases, and experiments are defined to utilize data and ML. One such business initiative is the curation and presentation of ML-based content—music, podcasts, radio—to the listener. The ML models exploit multidomain data such as the user’s context—location, time of day, latest activities, and the user’s preferences matched with the content’s profile. These initiatives then lead to the implementation of one or multiple intelligent user touchpoints and applications. Intelligent applications and touchpoints Ultimately the user’s experience needs to be enhanced through their touchpoints and the applications they use. Each intelligent business initiative can lead to the augmentation of existing or new user applications with ML. For example, the initiative of “Intelligently curate and present context-aware and personalized content” can lead to changes to multiple applications. It can be offered with a new application such as “Tell me about my preferences” that helps users grasp a deeper understanding of their preferences matched with music profiles and other environmental factors. It can lead to creation of multiple playlists through the “Personalized playlist browser” app that each provide a playlist intelligently curated based on calendar events and holidays, based on the day of the week, based on different genres, based on time of day, etc. It can lead to the creation of a new intelligent exploration tool, “Similar music explorer,” that allows the user to traverse a graph of related music content, based on their profiling and user preferences. Development of these intelligent applications leads to the identification and development of data products that provide the necessary data and insights. Data products One of the core components of data mesh is a set of interconnected data products. Intelligent applications drive the identification and development of the data products on the mesh. While real-world applications initiate the creation of data products, not all data products get directly used by them. Consumer-aligned data products are directly used by the applications, while aggregate or source-aligned ones are likely used by those data products. Over time, a mesh of data products gets created that get reused for new and novel intelligent applications. For example, the “Personalized playlist browser” app depends on various recommended playlist data products, such as the monday playlists, sunday morning playlists, or focus playlists data products that in turn rely on music profiles data product, among others. Multiplane platform The creation and consumption of data products and intelligent applications drive prioritization and development of platform capabilities. As the number and diversity of the data products, data users, and providers grow on the mesh, so do the features and maturity of the platform. The platform features are developed in parallel and sometimes ahead of the data product and intelligent application development. However, the platform value is demonstrated and measured based on its usage through the frictionless and reliable delivery of data products. Over time, platform features and services create a foundation for reuse in the delivery of new data products. Organizational alignment As the organization executes its data strategy, it implements the organizational and social aspect of the data mesh. The formation of a federated governance operating model, formation of cross-functional—business, dev, data, ops—teams, and establishment of the data product ownership role are all organizational aspects of data mesh that require the care and attention of the change management discipline. Organizational alignment happens concurrently and as part of the delivery and execution of data business initiatives, however with explicit intention and organizational planning. I will cover these aspects in Chapter 16. Business-driven, iterative, end-to-end, and evolutionary execution The overall process described in the preceding establishes the high-level steps of the execution. However, these steps are executed iteratively, with built-in feedback and monitoring so that the organization can measure and monitor its progress toward the target state in an evolutionary fashion. Iterations continuously deliver value and execute all the steps involved end to end. See the section “Data Mesh Execution Framework” for more on this. In the rest of this chapter, I share some practical guidelines to use this high-level execution framework. In the next chapter, I cover the cultural and organizational aspect of data mesh implementation. Data Mesh Execution Framework Brilliant strategy puts you on the competitive map, but only a solid execution keeps you there. ​—​Gary L. Neilson, Karla L. Martin, and Elizabeth Powers, “The Secrets to Successful Strategy Execution” (Harvard Business Review) To get started with a data mesh execution, we first need to acknowledge that it’s likely going to create transformational change—culture, organizational structure, and technology. I understand the word transformation carries quite a bit of weight, as if we need a great deal of potential energy to get it into motion. What I want to share with you here is a high-level framework that can help you get started. This framework helps you move forward and make progress one iteration at a time, delivering value and outcomes while maturing the foundation—the platform, the policies, and the operating mode. It allows using the kinetic energy of moving forward to keep going, instead of delaying the execution to gather all the will and energy needed for a big bang push to change. Figure 15-3 shows the high-level elements of the data mesh execution framework, business-driven, end to end, iterative, and evolutionary. Figure 15-3. A high-level data mesh execution framework Let’s look into how these aspects impact the activities, outcomes, and measurements involved in executing a data mesh implementation. Business-Driven Execution Data mesh is a component of a larger data strategy, and as demonstrated in Figure 15-2, its execution is driven by the strategic business cases and user scenarios that require ML or analytics, which then drive the identification and delivery of the data products and the components of the platform to build, manage, and consume them. This flows similar to how Amazon approaches building their major products and initiatives. “Start with the customer and work backwards—harder than it sounds, but a clear path to innovating and delighting customers. A useful Working Backwards tool: writing the press release and FAQ before you build the product.”4 While there are many benefits to this approach, a few challenges need to be managed. Benefits of business-driven execution These are a few key benefits to a business-driven execution: Continuous delivery and demonstration of value and outcome Showing results during a long transformation has a positive impact on morale, getting buy-in, increased commitment, and assuring continuation of investment. A business-driven execution is continuously delivering business value, though initially the pace might be slower or the scope of impact might be smaller. As the platform capabilities and operating model mature, the pace and scope improve. Rapid feedback from the consumers Engaging with the users—internal and external—has a positive influence on the platform and data products’ usability. The organization builds and evolves what is needed based on the feedback of the business, data, and platform consumers. The feedback is incorporated earlier and iteratively. Reducing waste Developing data products and platforms based on real business use cases leads to building what is actually needed when it is needed. The technology investments focus on delivering value to the business, sustainably, in contrast to merely following a build plan. Challenges of business-driven execution There are a few challenges to this approach that need to be closely managed: Building point-in-time solutions A narrow focus on delivering solutions to a single business use case at a time may lead to developing platform capabilities and data products that are not extensible to meet future use cases. This is a possible negative outcome of business-case-driven development. If we only focus on a particular business application, and not beyond that, we likely end up building point solutions narrowly designed for one particular use case, which do not extend to be reused and evolved for future use cases. This is particularly problematic for platform services as well as source-aligned and aggregate data products. The purpose of the platform services and data products on the mesh is to serve and enable a wide set of use cases. How to remediate point-in-time solution building? Product thinking is the key to finding the right balance between point-in-time solutioning and the YAGNI principle—building something ahead of time when “You Aren’t Gonna Need It.” Applying product ownership techniques to the platform services and the data products can deliver value for customers today, with an eye on the horizon for future and upcoming use cases. Data product owners understand the unbounded nature of analytical use cases. Future use cases that we are likely unaware of today will need to reach into past data to discover patterns and trends. Hence, data products need to capture data beyond the exact need of known use cases. This is particularly true for source-aligned data products (“Source-Aligned Domain Data”), where they need to capture the reality of the business and the facts as closely as possible and as they occur. Delivering to tight business deadlines A business-driven execution subjects the development of platform and data products to business deadlines. Business deadlines are often hard to change as they require coordination across many parties inside the organization and sometimes with partners and customers outside. To meet these deadlines, the lower layers of the stack such the platform services can incur technical debt, cutting corners in quality in order to meet the deadline of the business case. This is not a new phenomenon in software development. We almost never build software in a vacuum and almost always are faced with external deadlines, drivers, and unforeseen environmental circumstances that change the priority of features and how they are built. How to remediate incurring technical debt in the face of deadlines? Long-term product ownership of the platform and data products and evergreen engineering practices can assure that there is a continuous improvement process in place to address technical debts that may occur along with new feature delivery. The long-standing platform and data product teams have the knowledge and incentives to maintain their systems and data beyond a particular project delivery. This balances between the extensibility and long-term design of the platform and short-term project delivery. In short, the essence of the product mode of working strikes a balance between short-term technical debt to meet deadlines versus long-term product quality. For example, long-term product owners of platform services or data products know that getting the APIs and interfaces right is a priority, even when in a pinch for time. Often it is easier to fix the tech debt incurred in implementation details than make changes to the interfaces. API changes likely have a higher cost of change as they impact a number of consumers. Project-based budgeting Delivery of business initiatives is often organized as time-bound projects with a fixed allocated budget. Executing a data mesh implementation through business initiatives can have a negative impact on budget and resource allocation for long-standing digital components of the mesh—data products and the platform. When the budget is allocated to a time-bound business initiative, it becomes the sole source of investment for building the platform services or data products. As a result, their long-term ownership, continuous improvement, and maturity suffers. How to remediate project-based budgeting? Business projects are great vehicles for execution of the long-standing platform and data products. But they cannot be their sole funding mechanism. The platform and data products need to have their own allocated long-term resources and budget. The project mentality is at odds with platform and product thinking where product teams are never done. They continue to mature, maintain, and evolve the product in the face of new requirements and support many projects. Guidelines for business-driven execution Here are a few guidelines and practices when starting with a business-driven data mesh execution: Start with complementary use cases Picking a couple of complementary use cases at a time, over a single use, avoids building point solutions. Building the reusable components of the mesh through complementary and concurrent use cases can be a forcing function to avoid the pitfall of point solutioning. For example, the early use cases can be taken from different yet complementary areas such as analytics and reporting, complementing ML-based solutions. This forces the design of the data product to cater to two different modes of access, ML and reporting, over solutioning for one mode in a way that is not extensible to others. Know and prioritize the data consumer and provider personas Focusing on both types of data mesh user personas—data providers and data users—helps prioritize the platform capabilities fairly and for the most end-to-end impact. Get to know the people involved in building the intelligent solutions, the application developers, data product providers, and data product consumers. Understand their native tooling and skill, and then prioritize the work based on this information. For example, an early project is ideally done by the population of developers who are considered advanced users of the platform. They can work with missing features and be part of creating the required platform features. In my experience, many data mesh executions begin with an emphasis on data product creation. As a result, the data consumer experience suffers, and the end-to-end outcome is compromised. Start with use cases that have minimal dependencies on the platform’s missing features The platform evolves in response to its users. It cannot simply satisfy all the features needed by all the use cases at any point in time. In the early phases of execution, it is important that the use cases are carefully selected. The early use cases have just enough dependencies to the platform, so they drive the prioritization and build of necessary platform services without getting completely blocked. Start with use cases and developer personas that don’t get completely blocked by the missing platform capabilities. This creates momentum. As for the missing platform capabilities, they can be custom developed by the domain teams and later get harvested and generalized into the platform. Create long-term ownership and budgeting for platform services and data products While business initiatives and individual use cases can provide investment to develop the mesh, the mesh must be considered a long-standing digital investment and an internal product of the company. It is expected to continue to deliver value beyond its inception. Hence, it requires its own allocated resources and investment to sustain its long-term evolution and operation. Example of business-driven execution Let’s have a closer look at the example I introduced in the opening of this chapter and see how business initiatives can drive the execution of data mesh components. Daff has a strategic initiative to “Intelligently curate and present context-aware personalized content.” This initiative embodies a program of work using ML and insights to present and suggest different types of content such as music, podcasts, and radio channels to listeners. The program is planning to exploit data including users’ interactions, profiles, preferences, their network, locations, and local social events to continuously improve the relevance of the recommendations and curated content. This business initiative leads to the development of a few ML-based applications such as “Tell me about my preferences,” “Personalized playlist browser,” and “Similar music explorer.” These applications depend on curated playlists and music, recommended based on many data dimensions such as mood, activities, time of day, holidays, music genre, etc. To provide the data for these curated playlists, multiple data products must be developed. The list of data products includes consumer-aligned recommended playlists such as monday playlists, sunday morning playlists, sports playlists, etc., that utilize ML. To build these data products, a few source-aligned and aggregate data products are needed, such as listener profiles, music profiles, listener play events, calendar, etc. Depending on the maturity and availability of the platform capabilities, the platform prioritizes the needs of data product developers and consumers. For example, given the growth of the number of data products, the platform may prioritize data product automated life-cycle management, or due to ML-based transformations it will support ML-model transformation and flow-based transformation engines. Given the privacy concerns of users’ profiles and information, the governance and the platform prioritize establishing, codifying, and automating encryption and access to personally identifiable information (PII) for all the data products. This process repeats with the introduction of new business initiatives, e.g., “Intelligently identity, onboard, and promote artists.” End-to-End and Iterative Execution The only thing a Big Bang re-architecture guarantees is a Big Bang! ​—​Martin Fowler Continuous iterations of end-to-end delivery of intelligent business value enriches the mesh. It builds up the number of data products. It enhances the platform services. It increases the coverage and effectiveness of the automated policies. Figure 15-4 visualizes an end-to-end iteration of data mesh execution based on the previous example. At any point in time, there are multiple iterations executing concurrently. Through iterations of delivering use cases, such as the one mentioned earlier, more data products get created, more applications become intelligently augmented, more domains become part of the mesh, further platform services mature, and more of the policies become computationally enabled. Figure 15-4. Maturing and enriching the mesh components through iterations of business use cases Evolutionary Execution The proposed execution model for creating a data mesh operational model and architecture is incremental and iterative, where with each iteration the organization learns, refines the next steps, and incrementally moves toward the target state. I don’t recommend a rigid up-front plan of carefully designed tasks and milestones. From experience, I simply don’t believe such a plan will survive the first encounter with the execution and volatile environment of a large-scale business. Having said that, frameworks such as EDGE—a value-driven digital transformation—that prioritize getting outcomes, lightweight planning, continuous learning, and experimentation founded on agile principles can be very helpful in defining the operational model of the transformation. Regardless of the transformation framework you choose, the process of executing data mesh is treated as an evolutionary one. An evolutionary process recognizes and caters for different stages of maturity of the implementation and has a set of tests and metrics that continuously guide the execution toward the optimal outcome and higher levels of maturity. In this section, I share two tools that help you navigate an evolutionary execution: a multiphase adoption model that guides the execution at the macro level looking at the long-horizon milestones, and a set of objective functions that guide the evolution at the micro level of each iteration. A multiphase evolution model I have found the s-curve of innovation adoption5 is quite a useful tool to frame the evolution of data mesh components. While the shape of the curve might be different in different organizations, some with a steeper or slower slope, or a shorter or a longer tail, the general phases of the curve can be successfully used to guide the approach to evolution. This is useful in guiding the long-horizon evolutionary milestones. Figure 15-5 shows the s-curve of growth for data mesh. Each phase of the curve is mapped to a different population of adopters—innovators, early adopters, majority adopters, late adopters, and laggards—in the organization. Then, each phase of the curve is mapped to a corresponding mode of development—exploration when bootstrapping, expanding when scaling, and extracting6 when sustaining operation at scale. Figure 15-5. Evolutionary adoption of data mesh in an organization This is a de facto model for adopting a new innovation, and it can be adapted to guide the evolutionary phases of the data mesh components and principles. For example, let’s see how this can be applied to the execution of the platform component. The platform implementation will go through phases of bootstrapping with an exploration mindset where the early features of the platform are built in service of the innovator and early adopters—the risk takers and expert users. As the platform matures, it expands and extends its self-serve capabilities to serve a larger population of users, the population of generalist technologists in domains without hand-holding. During the expansion phase, the platform matures its support services to reduce the friction of using the platform and satisfy a more diverse persona of late platform adopters. Once the majority of users and their needs are addressed, platform changes remain minimal, and the organization extracts and exploits the features that are currently in place, at scale. During this phase, the platform APIs and services get reused by new data products and use cases. With the arrival of a major industry innovation in data management, an organization’s attempt to use the platform in new and novel ways, and the vendors’ shifts in the underlying technologies, the platform will go into the next s-curve of evolution, and the process repeats. Similarly, you can use this curve to plan the execution of domain ownership. Depending on the phase of the evolution, identify the corresponding class of domains—innovators, lead adopters, late adopters, and laggards—and prioritize which ones join the mesh first. For example, pick the domains that have more data expertise as the innovator adopter before moving to the majority adopter domains. Use the early adopter domains to pave the path for others. For each phase of the evolution, be intentional about the different capabilities developed and the behavior of the teams. For example, in the early phases, the teams should adopt a more exploratory attitude to establishing policies, the platform, and data products, with a tolerable level of manual intervention. In the expand phase, the mesh needs to mature to minimize manual activities, have extensive observability in place to monitor the load, and scale usage with great policy coverage. Let’s look at the adaptation of the s-curve to each of the data mesh components and establish their characteristics for each evolution phase. Domain ownership evolution phases Figure 15-6 shows an example of dimensions of domain ownership as it evolves through the adoption phases. Figure 15-6. Example of domain ownership characteristics through phases of evolution During the exploration and bootstrapping phase, only a small number of domains are involved in the execution of data mesh. These early domains often act both as the provider of the data and as the consumer in order to participate in end-to-end value delivery with minimal dependencies. Nevertheless, both roles are necessary to be part of this phase. In this phase the domains focus on exploration and setting up a foundation for growth. The early adopter domains are setting sensible practices in terms of defining the data product owner role and setting up the domain team structure. They pave the path by creating tooling and processes to integrate operational data from applications to source-aligned data products. The early domains are among the most advanced in their data and analytics capabilities with a modern technology stack. At this stage, value comes from unblocking a small number of high-value use cases. During the expand and scale phase, in rapid increments, an increasing number of domains join the mesh, the majority middle. The early majority continue to pave the path and establish a set of repeatable patterns, technically and organizationally, so that the execution can rapidly scale out to the majority of the domains. New domains contribute to source- and consumer-aligned data products and aggregate data products. The need for aggregate data products and pure data domains arises with the increased number of domains. Domains have a diverse set of technology stacks, very likely with older and legacy systems that need to be integrated or migrated to the mesh. They introduce a new set of practices and tooling to handle legacy integration. At this stage, value comes from the positive network effect. During the extract and sustain phase, most of the domains have moved to an autonomous data ownership model, and the number of domains stabilizes. Fluctuations to the domains are expected due to optimization of the data ownership model. Domains that have become bottlenecks due to the large number of data products may break up into multiple domains, and domains that users are minimally dependent on may merge with others. In this phase provider domains continue to modify, evolve, and optimize their data products, but the majority of activities are happening in consumer domains taking advantage of the large set of data products available for reuse and building new scenarios. Domains that have been technologically misaligned with the mesh, due to the location of their data or the age of their systems and data technology, may choose to join the mesh at this point where most of the capabilities and organizational practices are well established. At this stage, value comes from consistency and completeness effects across the organization. Data as a product evolution phases Figure 15-7 shows an example of dimensions of data product development as they progress through the phases of evolution. Figure 15-7. Example of data product characteristics through phases of evolution During the explore and bootstrap phase, only a small set of data products is created. The early data products support an essential subset of affordances introduced in Chapter 11. For example, only a subset of output port modes may be available, with a limited retention period of sharing just the latest snapshot, or limited discoverability capabilities. In this phase, data product development has an exploratory nature in order to identify the best approach to implement each of the data product affordances. Data product developers are actively establishing patterns and sensible practices. Early data product developers work collaboratively to share knowledge and learnings. Data product development actively influences the platform features, based on experimentation and early learnings. Data products developed during this phase may not have a standard blueprint and don’t yet fully utilize the platform capabilities. During this phase the chosen data products are among the low-risk ones from the security and reliability perspective, and they are tolerant of minimal standardization. The early data products are selected to be different but complementary in terms of their development model and technology requirements. This keeps the platform dependencies focused. At this time, the majority of data products are source-aligned, with a smaller number of consumer-aligned data products addressing the needs of specific use cases that are bootstrapping the mesh. During the expand and scale phase, a large number of data products are being added to the mesh, at a rapid rate. With early adopters having established the standardized patterns of data product development, the data products can be developed rapidly. The data products create the necessary affordances required to be observed, discovered, and used at scale by a large number of the users without hand-holding. Data products support a diverse set of transformations such as running ML models or stream-based data processing. Data product development is focused on supporting a diversity of sources and consumers. Data products developed during this phase can be of a higher risk category with regard to privacy, uptime, resilience, etc. A diverse class of data products, with an increasing number of aggregate data products get created in this phase. During the extract and sustain phase, the number of data products is likely to stabilize with a slower pace of change. The optimization of data products or new business needs may create new data products or change the existing ones. The data products continue to evolve. Data products continue to optimize for more rapid access to data, higher performance of data sharing, and other cross-functional concerns to address the large scale of usage. They also continue to optimize their life-cycle management and increase the speed of applying changes. Data products with a longer lead time to change, due to transformation complexity, are broken down into smaller and more efficient data products. During this phase the focus of data product development is to generate large-scale returns on the prior investments of data product development. Many intelligent business initiatives can use the existing data products as is. During this phase, laggard legacy data storage or systems are converted to data products. Self-serve platform evolution phases Figure 15-8 shows an example of self-serve platform characteristics as they change through the evolution of data mesh. Figure 15-8. Example of data platform characteristics through phases of evolution During the bootstrap and explore phase, the platform is establishing its essential foundational capabilities, mainly in the utility plane—storage, compute, accounts, access control, transformation. A primitive set of data product experience services are developed, such as basic data product life cycle management and global addressability. There are very few mesh experience plane capabilities such as basic discoverability. In this phase, the platform self-serve implementation is premature and can be as simple as pull-request-based script sharing. The platform users are expert data product developers and consumers that can get by with minimal self-serve capabilities. The platform supports a smaller number of users. The platform team is working closely with early domains to provide their needs without blocking them and harvesting the cross-functional features the domain teams are building back to the platform. In the early phases, the platform is opinionated and has a limited set of supported technologies to focus on what is necessary to get momentum. During the scale and expand phase, the platform supports automatic data product generation, with all affordances in place, though some might have limited functionality. For example, the automatic data metrics sharing foundation is in place, but the metrics set may not include all possible temporal, quality, or volume information. The platform is mainly used through the data product experience plane or the mesh experience plane. Platform self-serve abilities mature rapidly to support a larger portion of the population with common requirements. Automated processes and tools are established. The persona of the platform target users skew toward majority generalist technologists. The platform focuses on serving the majority middle population of users without requiring hand-holding. Teams who wish to use a different set of technologies that the platform offers remain off the platform and will have the end-to-end responsibility of meeting the mesh governance and compatibility requirements. Only a small number of teams should fall into this category. During the extract and sustain phase, the platform has a mature state of automatic observability, automatic healing, and a high degree of resiliency at scale. Any downtime in the mesh can have large-scale consequences. The platform capabilities are matured and exploited for a diverse set of use cases, data products, and teams. The platform continues to optimize the self-serve capabilities for laggards and outliers. In this phase, the platform is utilized by the majority of the teams in the organization. Platform development focuses on optimizing the experience of the mesh and lowering the risk of downtime or security breaches. Federated computational governance evolution phases The federated governance operating model and the computational policies governing the mesh evolve through the adoption phases. Figure 15-9 shows the evolution of governance characteristics such as the number of domains joining the federated governance operation, maturity of the federated operating model, the focus of governance development, and coverage of computational policies. Figure 15-9. Example of governance characteristics through phases of evolution During the explore and bootstrap phase, the new governance model forms with only a few early domains that are participating in the federated governance. The governance operation establishes the principles that guide the decision-making process and decide what policies the organization must implement globally and what can be left to domains. The global incentives of domain data product teams are defined, augmenting their local domain incentives. During this phase, the existing governance structure reshapes into a federated model. The members of the existing governance group are either taking a subject matter expert role, joining the platform team to help with product management of cross-functional policy automation, or joining the domains as data product owners. The governance is focusing on setting the foundation. Data product developers in the early days may build their own implementation of policies, and the platform harvests them later into its self-serve capabilities, a form of crowdsourcing for some of the early policy implementations. The early governance and domain teams pave the path in establishing the operating model, decision making, and policy automation. The governance function establishes the essential policies relevant to early data products and domains. Only a subset of essential policies might be implemented automatically by the platform and enforced consistently in all data products, such as security and privacy. During the expand and scale phase, with the basic operating model and approach to automating policies having been established by early domains and governance teams, the road is paved for scaled governance with a majority of domains joining the mesh. The process of onboarding new domains on the mesh is frictionless. The number of members of the federated governance team increases to include the majority of the domains in the organization. The team focuses on increasing the coverage and diversity of the policies and maturing the operating model for rapid onboarding of new domains. During this time, the majority of policies are automated to support a mesh with a large number of interoperable and secure data products. Basic monitoring of policies is in place. During the extract and sustain phase, all data products on the mesh are embedding automated policies. There is an automated process in place to detect policy risk and lack of conformance to notify data product teams to address concerns with automated tools provided by the platform. The team focuses on optimizing the governance process and enhancing the supported features of automated policies. The mesh is monitoring and optimizing the performance of the policies. Guided evolution with fitness functions While the adoption phases guide the execution of data mesh at a macro level, during a large time scale of months or years, continuous measurement and tracking of a set of objective fitness functions guide the small increments and short iterative execution cycles. Fitness Functions A fitness function is an objective function used to summarize how close a prospective design solution is to achieving its set aims. In evolutionary computing, the fitness function determines whether an algorithm has improved over time. In Building Evolutionary Architectures, fitness functions objectively guide the iterations of an architecture toward its set objectives across multiple dimensions such as performance, security, extensibility, etc. Fitness functions is a concept borrowed from evolutionary computing, which is an objective measure used to assess whether the system as a whole is moving toward its intended design target and is improving over time, or not. In the context of data mesh, these objective functions can determine how “fit” the mesh implementation is at any point in time, and whether every iteration is moving it closer or further from its optimal state of delivering value at scale. Data mesh’s ultimate objective is increasing the ability of organizations to utilize data for analytical purposes and get value from their data at scale, aligned with organizational growth and complexity. Each of the data mesh principles contributes to this high-level outcome. Hence, fitness functions can target the outcomes expected from each principle. Caution Coming up with numerical fitness functions, particularly when we are still figuring out what good looks like, is very hard. Even in the case of well-established DevOps engineering practices, the book Accelerate: Building and Scaling High Performing Technology Organizations,7 after years of running quantitative research and measuring 24 key metrics, demonstrates that only a small number of metrics, 4, had a direct link to the organization’s performance and the majority were simply vanity measurements. The objective measures proposed here are examples and a starting point. They need to be experimented with, quantitatively measured, and compared across organizations to prove that they are in fact just the right ones. You might have noticed that I have used the concept of fitness function over key performance indicators (KPIs). While both concepts are essentially about objective metrics and giving visibility to the state and trend of data mesh execution progress, the fitness function tells you how “fit” the data mesh design and implementation are at any point of time and guides you through the changes to make the right ones and revert the ones with negative impact, with a focus on behaviors and outcomes that really matter. For example, a KPI related to the “number of data products” might sound appealing to many leaders as an indicator of success and growth of the mesh. This is in fact not the right measure, as it does not indicate the “fitness” of the mesh. The target of the mesh is a design to “deliver value.” This is not directly related to the number of data products but related to the “number of links to data products” as a measure of usage. The more links and interconnectivity, the more data products are being used to generate value. Data mesh execution requires automated instrumentation, collection, and tracking of these metrics. The platform services play a key part in tracking and externalizing many of these metrics. For example, to measure the “acceleration of policy adoption” as a function of effectiveness of automated policies as code, the platform can measure the “lead time of a data product to adopt a new policy.” This can be collected as instrumentation of the data products’ continuous delivery pipeline. Data product and mesh experience plane services enable the domain and global governance teams to measure, report, and track these execution metrics. Detecting trends in the wrong direction can trigger a set of actions to investigate the root cause and prioritize the fixes to nudge the evolution of the mesh toward its objective outcomes. I suggest that you define a set of objective fitness functions for your data mesh execution. Automate their implementation as much as possible, and continuously monitor their trends. They can guide you through the execution of data mesh and take the next best step. The fitness functions I have introduced in the following are only examples. I have demonstrated how they are derived from the expected outcomes of each data mesh principle. You can adopt the same approach and come up with your own fitness functions to test and evolve. Domain ownership fitness functions Figure 15-10 shows an example set of fitness functions that measure the progress of domain data ownership. These fitness functions measure a set of objectives that tell us whether the evolution is moving in alignment with the outcome of domain data ownership. Is the implementation generating more value through data sharing in step with organizational growth? Figure 15-10. Example of domain data ownership fitness functions Data as a product fitness functions Figure 15-11 introduces an example set of fitness functions that measure whether the data mesh implementation is objectively progressing toward data as a product outcome. Is the effectiveness and efficiency of data sharing across domains increasing with organizational growth or not? Figure 15-11. Example of data as a product fitness functions Self-serve platform fitness functions Figure 15-12 introduces an example set of fitness functions that measure whether the data mesh implementation is objectively progressing toward the self-serve platform outcome. Is the platform increasing domains’ autonomy in producing and consuming data products and lowering their cost of data ownership? Figure 15-12. Example of self-serve platform fitness functions Federated computational governance fitness functions Figure 15-13 introduces an example set of fitness functions that measure whether the data mesh implementation is objectively progressing toward the governance outcome. Is the mesh generating higher-order intelligence—securely and consistently—through efficient joins and correlations of data products, in step with organizational growth? Figure 15-13. Example of governance fitness functions Migration from legacy We can’t talk about an evolutionary execution of data mesh without talking about legacy. The majority of the organizations adopting data mesh likely have one or multiple implementations of past data architectures such as data warehouses, data lakes, or data lakehouses, and a centralized governance structure with a centralized data team. Data mesh execution must incorporate migration from such legacy organization structures, architecture, and technology in its incremental and evolutionary journey. However, it is difficult to generalize migration from legacy, as the path highly depends on many variables specific to the organization such as the time constraint on migrating to the new infrastructure in the case that establishing the mesh coincides with migration to a new environment, e.g., a new cloud provider; the scale and number of early adopter domains; the available investment for the platform; and so on. Laying out a migration strategy from a particular state to data mesh is out of the scope of this book and requires close collaboration with an organization and understanding its unique context. However, what I can share with you is a list of high-level lessons learned from my experience so far. There is a lot more to be learned here. No centralized data architecture coexists with data mesh, unless in transition I often get asked if the existing warehouses, lakes, and lakehouses that the organizations have created can coexist with the mesh. The short answer is transitionally yes, but ultimately no. Any incremental migration would require a period of time where the old and the new coexist at the same time. However, the goal of data mesh is to remove the bottlenecks of the lake and warehouse. Decentralized data products replace the centralized and shared warehouses and lakes. Hence, the target architecture does not intend for the centralized bottlenecks to coexist with data mesh. For example, retaining a centralized data warehouse or a lake and then feeding a data mesh downstream from it is an antipattern. In this case, the introduction of data mesh is an added complexity, while the bottleneck of a warehouse collecting data from domains remains. Alternatively, as a transitory step, it is possible to have a data mesh implementation that feeds data into a warehouse, which in turn serves a controlled and small number of legacy users whose reports are still using the warehouse. Until all the legacy warehouse users are migrated to the mesh, the mesh and the warehouse coexist. The migration continues, and it reduces the level of centralization of the warehouse bottleneck. The ultimate goal of the architecture is to remove and reduce the overall level of centralization. Hence, ultimately a mesh implementation and a central data warehouse or lake should not coexist. Centralized data technologies can be used with data mesh At the time of writing this book, no data-mesh-native technologies have been created. We are at a stage of evolution when adopters utilize the existing technologies, with custom configurations and solutions built on top. Today, my advice is to utilize the elements of the existing technologies in a way that they don’t get in the way of building and running autonomous and distributed data products. I suggest configuring and operating the existing technologies in a multitenant way. For example, use the basic elements of lake-oriented technologies such as object storage and configure the allocation and access to the storage based on the boundary of each data product. Similarly, you can utilize data warehousing technologies as an underlying storage and SQL access mode but allocate a separate schema to each data product. Not all existing technologies lend themselves to such configuration. You can utilize the ones that do. Bypass the lake and warehouse and go to directly to the source When migrating from an existing lake or warehouse, the path of least resistance is to put the warehouse or lake upstream to the mesh and implement the data products with input ports consuming data from the existing warehouse tables. My advice would be to not take this path and instead discover the existing consumers of the lake or warehouse, recognize their needs and use cases, and directly connect them to new data products that consume data from the source domain systems. Discover the domain owners early, develop the data products according to the domain’s truthful representation of the data, and update the consumers if need be. Creating artificial data products that represent the lake files or warehouse tables does not achieve the outcomes of the mesh, scaling through ownership by domains and closing the gap between the source and consumer. It only adds further technical debt, creates another layer between the source and consumer, and increases the distance between them. Use the data warehouse as the consuming edge node Assessing the consumers of an existing warehouse, you may decide to keep some of them in the warehouse and never migrate them to the mesh. There are multiple reasons for this decision. For example, reports such as financial reports that don’t change that often, rely on a small and specific set of tables, and are defined with complex business logic that are hard to reverse engineer and re-create may remain dependent on the warehouse. In this case, you want the warehouse to act as a consumer of the mesh and become an edge consuming node, with limited reports accessing a small set of its remaining tables. Migrate from a warehouse or lake in atomic evolutionary steps It’s likely that migration from the data warehouse or lake takes many evolutionary iterations and possibly months or years depending on the size of the organization. I strongly suggest executing your migration in atomic evolutionary steps where the state of the system is left with less technical debt and architectural entropy when the step is completed. While this might sound logical and seemingly obvious, it’s often ignored. The teams often keep adding new infrastructure and new data products while leaving the old pipelines or lake files behind. Over time they just add to the architectural complexity and total cost of ownership to maintain the new and leftover old systems. The main reason behind this coexistence is that we often fail to migrate over users and consumers to the new components. For example, if you are creating a new set of data products, make sure you take the time to migrate over the existing consumers—reports or ML models—of the lake or warehouse to them. An atomic step of migration includes adding new source-aligned, aggregate, or consumer-aligned data products, migrating the consumers over to the new data products, and retiring the tables, files, and pipelines of the lake and warehouse. The migration step is complete when all these steps happen or all revert if some don’t complete. At the end of this atomic step, the architectural entropy is reduced. Conversely, the atomic migration step is not complete if we don’t migrate the consumers after creating new data products or if we don’t retire the existing pipelines and lake files. Make sure your project scheduling and resource allocations are aligned with committing to atomic steps of migration. Recap Data mesh execution is part of a larger data strategy and in fact is a major component of the strategy. It needs business-oriented strategic initiatives, defining the hypothesis of how data generates value aligned with business objectives, to drive its execution. ML and analytics-powered business use cases become the vehicles that execute identification and delivery of data products, adoption of data by domains, and establishment of governance and the platform. Balancing between point solutions delivering a data-driven business outcome versus building a long-term shared platform across multiple use cases remains a challenge. Tools such as the evolutionary approach based on the long-term adoption roadmap and functions to measure fitness of the implementation according to its long-term objective help establish an equilibrium and guide the execution in an evolutionary fashion toward a mature state. On rare occasions an organization has the opportunity to create a greenfield data mesh implementation. In the majority of cases, a brownfield execution must account for migration of existing architectures such as a data warehouse, lake, or lakehouse. In this case, migrating the old architecture and organization depends on executing atomic evolutionary steps that each migrate the data consumers and providers and everything in between the two toward the mesh. These initial guidelines and approach should suffice to get you started while you learn and refine the process along the way. What remains is a few notes on how to organize your teams and evolve the culture of an organization during this evolutionary execution, which I cover in the last chapter of this book, Chapter 16. See you there.

私たちはこの本の最後の部分まで到達しました
おめでとうございます！データメッシュを組織に採用するか、他の人々がそれを行うのを助けることを考えていますか？データメッシュの実行に影響を与える、リード、または管理の立場にいますか？組織的な変革のアプローチや最初のステップについての助けが必要ですか？これらの質問に「はい」と答えた場合、この部分はあなたのためです
私は、第I部でデータメッシュの基礎的な原則を理解し、第II部でデータメッシュの動機と目標を理解していると想定しています
 スコープ データメッシュは、データ駆動型の組織においてデータからの価値を規模化するためのデータ戦略の要素です
この戦略の実行には複数の要素があります
それはチームに影響を与え、彼らの責任体制やドメインとプラットフォームチームの責任範囲を明確化します
それは、文化に影響を与え、組織がデータ指向の成功をどのように評価し、測定するかを変えます
それは、データの利用可能性、セキュリティ、相互運用性に関するローカルとグローバルな意思決定の運営モデルを変えます
それはデータ共有の分散型モデルをサポートする新しいアーキテクチャを導入します
この本のこの部分では、この多面的な変革の実行を始め、ガイドするためのハイレベルなアプローチを説明します
第15章「戦略と実行」では、段階的にデータメッシュを採用し、増加する数のデータ製品から価値を生み出しながら、プラットフォームの能力を向上させるための進化的なアプローチを紹介します
第16章「組織と文化」では、データ製品の長期的なオーナーシップとピアツーピアのデータ共有を作成するためのチーム、役割、パフォーマンスメトリクスの主要な変更について説明します
この本を書いている時点では、データメッシュを採用し、データ共有と分析がビジネステクノロジードメインに組み込まれる分散型の組織モデルを作り始めている初期の段階です
これらの章で提案されているアプローチは、ドメイン指向の運用チームやデジタルプラットフォームの構築などのシナリオで成功裏に応用されてきた進化的かつ大規模な変革モデルの適応です
私はデータメッシュに特有のブロッカーやエンエーブラーについての経験を共有します
これはあくまでスタート地点であり、データメッシュに特有の変革アプローチを詳細に洗練するためにはまだ数年の時間がかかるでしょう
 第15章 戦略と実行 戦略の本質は、競争相手と異なる方法で活動を行うことを選択することです
 ― マイケル・E・ポーター あなたの組織はデータ指向の戦略を採用していますか？データと洞察を使用して顧客やパートナーに独自の価値を提供する計画を立てていますか？あなたの組織は、異なる一連の活動を実行するためにデータを使用するビジョンを持っていますか、または現在の活動を異なる方法で行うためのデータを使用するビジョンを持っていますか？データメッシュをこの戦略の実現の基盤として興味を持っていますか？これらの質問に「はい」と答えた場合、この章はあなたのためです
この章では、データメッシュと全体的なデータ戦略の関係、およびそれを実行するための繰り返し、価値志向、進化的なアプローチについて説明します
しかし、まず、簡単な自己評価を行ってください
今日データメッシュを採用すべきですか？この章の残りを読む前に、自分自身に問いかけることができます
データメッシュは私の組織にとって適切な選択肢ですか？それ以上に重要なのは、今データメッシュが適切な選択肢であるかどうかです
1 図15-1は、これらの質問に対する私の見解を示しています
スパイダーグラフの色付きセクションに属する組織は、各評価軸で中程度または高得点を獲得し、現時点でデータメッシュを成功裏に採用するために最適な位置にあると言えます
以下の基準を使用して、自己評価を行い、データメッシュがあなたの組織に今適切な選択肢であるかどうかを判断できます
 図15-1.データメッシュ導入のための自己評価基準 組織の複雑さ データメッシュは、既存のデータウェアハウスやデータレイクの解決策が大規模かつ複雑な組織で妨げとなり、ビジネスの多くの機能にわたってデータからの価値を迅速かつスムーズに得ることができなくなっている組織向けのソリューションです
もし組織がまだそのような複雑さのレベルに達しておらず、既存のデータおよび分析のアーキテクチャが要件を満たしている場合、データメッシュはあまり価値を提供しません
データの複雑さが中程度または高いスコアを獲得している多くのデータソースとユースケースを持つ組織は、データメッシュの恩恵を受けます
例えば、多くの異なるデータドメインと多数のソースから収集されたさまざまなデータタイプを含むビジネスが、機械学習と分析のユースケースに大いに依存する早期成長のスタートアップや企業などがデータメッシュの有力な候補です
小売業、銀行業、金融業、保険業、そして大手テックサービスなどの産業での企業やスケールアップも該当します
例えば、医療提供者や支払者などを考えてみましょう
データの多様性には固有の複雑さがあります
それは、臨床訪問から患者のプロフィール、検査結果、薬剤、臨床請求までさまざまなカテゴリのデータで構成されています
これらのデータの各カテゴリは、ビジネスドメインや独立した組織、病院、薬局、クリニック、研究所などのエコシステム内の複数の関係者から入手できます
このデータは、最適化された個別化されたケアモデル、インテリジェントな保険承認、そして価値に基づくケアなど、多くのユースケースで使用されています
データメッシュを採用するための主な基準は、固有のビジネスの複雑さとデータソースとデータユースケースの増加です
 データ指向戦略 データメッシュは、大規模なデータから価値を得ることを計画している組織向けのソリューションです
それには、製品チームとビジネスユニットが、アプリケーションやサービスでインテリジェントな意思決定とアクションを行うことを想像するためのコミットメントが必要です
そのようなコミットメントは、組織が機械学習と分析をデータを活用した戦略的な差別化要素と認識する場合にのみ可能です
データ指向戦略は、データに基づくアプリケーションとサービスに具現化され、それが各チームのデータへのコミットメントにつながることになります
データの共有に関する追加の責任を担うという事実に対して、ドメインチームやビジネスドメインが動機付けられるのは戦略的なビジョンがない場合は困難です
 エグゼクティブの支援 データメッシュは変革を導入し、それには変革への抵抗が伴います
このような規模の変革では反対者や妨害者が現れるものです
プラットフォームの進捗と期限を満たすためのポイントソリューションの構築の間で困難な決断をする必要があります
組織の働き方を変えるために、組織を動機付けることが求められます
私の経験では、最も成功したデータメッシュの導入は、Cレベルの幹部の支持を得ています
 コアとしてのデータテクノロジー データメッシュの導入は、テクノロジーとビジネスの新しい関係に基づいています
それは各ビジネスドメインでデータの普及的な活用に基づいています
データテクノロジーを自社のコアとして活用する組織は、データとAIを競争力の源泉としており、ビジネスを拡大し再構築する手段として使用しています
こうした組織は、データの共有と利用をビジネスの各機能のコアに組み込むために必要な技術を構築・創造する意欲と能力を持っています
それに対してテクノロジーをビジネスのサポーターと見なし、コアではない組織は、テクノロジーの能力を外部ベンダーに委任し、ビジネスニーズに対する既成のソリューションを購入して接続することを期待しています
このような組織はまだデータメッシュを導入する準備が整っていません
 早期導入者 本書を執筆した時点では、データメッシュは革新者や主要な導入者による早期の導入を享受しています
新しいイノベーションの先駆者は、特定の特徴を持つ必要があります
彼らは自分たちの業界での冒険的な意見リーダーです
彼らはデータメッシュなどの新興技術を積極的に導入する意欲があります
特にデータメッシュのような多面的で浸透力のある斬新なアプローチの先駆者は、実験の精神、リスクの取り扱い、早期の失敗、学び、進化を求めます
この時点でデータメッシュを採用する企業は、最初の原則から始め、それらを自分たちの状況に適応し、学び、進化する意欲がなければなりません
それに対して、リスクを取らず、検証済みで洗練されたプレイブックのみを採用しようとする組織や、後発採用者のカテゴリーに属する組織は、しばらく待つ必要があるかもしれません


モダンなエンジニアリングデータメッシュは、モダンなソフトウェアエンジニアリングの基礎に基づいて構築されています
ソフトウェアの継続的かつ自動化されたデリバリー、DevOpsの実践、分散アーキテクチャ、計算ポリシー、モダンなデータストレージおよび処理スタック（プライベートまたはパブリッククラウド上）が利用可能です
堅固なエンジニアリングのプラクティスとオープンでモダンなデータツールへのアクセスを備えていないと、データメッシュはゼロから立ち上げることが難しいでしょう
データメッシュは、マイクロサービスの前提条件を共有しており、モダンなデータおよび分析技術とプラクティスが必要です
さらに、データメッシュはAPI駆動型で統合が容易であり、大規模かつ中央集権的なチームではなく、多くの小規模チームを可能にするモダンなテクノロジースタックと連携します
データのモデリングの集中管理、データの集中的な制御、データの集中的な保存を強制する技術は、成功したデータメッシュの実装には向いていません
それらはボトルネックとなり、データメッシュの分散性と衝突します


ドメイン指向の組織データメッシュは、採用者がテクノロジーとビジネスの調和を持つモダンなデジタルビジネスであるか、またはその旅を進行中であることを前提としています
それは組織がビジネスドメインに基づいて設計されていると仮定し、それぞれのビジネスドメインに専任の技術チームがあり、ビジネスドメインに役立つデジタルアセットと製品を構築およびサポートしていると仮定します
テクノロジーとビジネスチームが密接に協力し、ビジネスを可能にし、再形成し、拡大するためにテクノロジーを利用すると仮定します
例えば、モダンな銀行は、デジタルのコアバンキング、ローンのデジタル処理、デジタルの法人・個人バンキング、デジタルのクレジットとリスクなど、専任のテクノロジーグループを持っています
テクノロジー、技術チーム、およびビジネスドメインとの密接な調和があります
これに対して、プロジェクトごとに全ての事業部門で人々（リソース）を共有する集中型のIT部門を持つ組織は、各ビジネスドメインの技術的な資産を継続的かつ長期的に所有しないため、データメッシュには適していません
ドメイン指向のデータ共有は、ドメイン指向のテクノロジーとビジネスの調和がある場合にのみスケールします


長期的なコミットメントデータメッシュの採用は、変革と旅であるということが後の章でおわかりいただけると思います
それは限られた範囲と少数のドメインの組織的な足跡から始めることができますが、スケールで活用することで利点が蓄積されます
したがって、ビジョンに対する組織的なコミットメントを作り出し、変革の旅に乗り出すことができる能力が必要です
データメッシュは、1つまたは2つの一時的なプロジェクトとして実装することはできません
前述の基準に合致している場合は、データメッシュを使用してデータ戦略を実行する方法を見ていきましょう


データストラテジーの一部としてのデータメッシュデータメッシュが大規模なデータストラテジーにどのように組み込まれるかを理解するために、架空の企業であるDaff, Inc.の事例を見てみましょう
Daffのデータストラテジーは、アーティストとリスナーを没入した芸術体験を通じてつなぐインテリジェントなプラットフォームを作成することです
この体験は、リスナーとアーティストがプラットフォーム上での各インタラクションに埋め込まれたデータと機械学習を使用して、連続的かつ迅速に改善されます
これには、プラットフォーム上のすべての接点からデータを収集し、エンゲージメントデータを使用してMLを利用してすべてのプラットフォーム機能を強化することが含まれます
例えば、プラットフォームのより良いプレイリスト予測は、エンゲージメントが高いリスナーを増やし、その結果より多くのデータを得てより良い予測が可能になります
これにより、ビジネスへのポジティブなセルフリインフォースメントループが作成されます
Daffの野心的な戦略は、組織全体のデータメッシュの実装によって実現されます
メッシュ上のノードは、プラットフォーム上で収集および管理されるさまざまな活動からのデータや、そのようなデータに基づいて訓練された機械学習モデルによって推論されるインテリジェンスを表すことができます
図15-2は、データ戦略とデータメッシュの関係を示しています
図15-2
全体像
データメッシュの実行をデータ戦略全体につなげるデータメッシュの実行をデータ戦略が駆動することを示す図15-2の通り、データ戦略はデータメッシュの実行を駆動します
図を上から下に探索し、戦略からプラットフォームへのデータメッシュの実装、そして継続的なフィードバックループを経ていきましょう
データ戦略データ戦略は、企業がデータを活用し、戦略的なイニシアチブのロードマップを作成する方法を定義します
たとえば、Daffの戦略は、すべてのコンテンツプロデューサーやアーティスト、リスナーの体験をデータとMLと関連付けることに基づいています
戦略的なビジネスイニシアチブとユースケース戦略に基づいて、データとMLを活用するために、一連のビジネスイニシアチブ、ユースケース、および実験が定義されます
そのようなビジネスイニシアチブの1つは、リスナーに対してMLに基づいたコンテンツ（音楽、ポッドキャスト、ラジオ）のキュレーションとプレゼンテーションです
MLモデルは、ユーザーのコンテキスト（場所、時間帯、最新の活動）やユーザーの好みとコンテンツのプロフィールとのマッチングなど、多ドメインデータを活用しています
これらのイニシアチブは、最終的に1つまたは複数のインテリジェントユーザータッチポイントとアプリケーションの実装につながります
インテリジェントなアプリケーションとタッチポイント最終的には、ユーザーの体験をタッチポイントと使用するアプリケーションを通じて向上させる必要があります
各インテリジェントビジネスイニシアチブは、既存のアプリケーションまたは新しいアプリケーションに対してMLを用いて拡張されることができます
たとえば、「知的なコンテキストに基づいた個別のコンテンツのキュレーションとプレゼンテーション」というイニシアチブは、複数のアプリケーションに変更をもたらすことができます
ユーザーの好みと音楽のプロフィールや他の環境要素とのマッチングに基づいて、より深い理解をユーザーに提供する「私の好みについて教えて」などの新しいアプリケーションとして提供することができます
また、カレンダーのイベントや休日に基づいて、曜日に基づいて、異なるジャンルに基づいて、時間帯に基づいてなど、それぞれのプレイリストをインテリジェントに厳選した「パーソナライズドプレイリストブラウザー」という新しいアプリケーションの作成にもつながります
プロファイリングとユーザーの好みに基づいて、関連する音楽コンテンツのグラフをたどることができる新しいインテリジェントな探索ツール「類似音楽エクスプローラー」の作成にもつながります
これらのインテリジェントなアプリケーションの開発により、必要なデータとインサイトを提供するデータ製品の特定と開発が行われます
データ製品データメッシュの主要な要素の1つは、相互に連携したデータ製品のセットです
インテリジェントなアプリケーションは、メッシュ上のデータ製品の特定と開発を促進します
実世界のアプリケーションがデータ製品の作成を開始するものの、すべてのデータ製品が直接それらによって使用されるわけではありません
消費者に合わせたデータ製品はアプリケーションに直接使用され、集約型またはソースに合わせたデータ製品はそれらのデータ製品によって使用される可能性があります
時間の経過とともに、新しいユニークなインテリジェントなアプリケーションに再利用されるデータ製品のメッシュが作成されます
たとえば、「パーソナライズドプレイリストブラウザー」アプリは、月曜のプレイリスト、日曜の朝のプレイリスト、フォーカスプレイリストなど、さまざまなおすすめのプレイリストデータ製品に依存しており、それらはさらに音楽プロフィールデータ製品などに依存しています
マルチプレーンプラットフォーム

データ製品とインテリジェントアプリケーションの作成と利用は、プラットフォームの機能の優先順位付けと開発を推進します
データ製品、データ利用者、プロバイダーの数と多様性がメッシュ上で成長するにつれて、プラットフォームの機能と成熟度も向上していきます
プラットフォームの機能は、データ製品とインテリジェントアプリケーションの開発と並行して、時にはそれらよりも先行して開発されます
しかし、プラットフォームの価値は、摩擦のない信頼性の高いデータ製品の提供を通じてその使用状況に基づいて示され、測定されます
時間の経過とともに、プラットフォームの機能とサービスは、新しいデータ製品の提供における再利用の土台を作ります


組織の整合性

組織は、データ戦略を実行するにつれて、データメッシュの組織的な側面と社会的な側面を実装します
分散型統治オペレーティングモデルの形成、クロスファンクショナルなビジネス、開発、データ、オペレーションのチームの形成、データ製品のオーナーシップ役割の確立は、すべてデータメッシュの組織的な側面であり、変革管理の視点からの注意と注意を必要とします
組織の整列は、データビジネスイニシアチブの提供と実行と並行して行われますが、明示的な意図と組織的な計画が必要です
これらの側面については、第16章で説明します


ビジネス主導、反復的、エンドツーエンド、進化的な実行

先行する全体的なプロセスは、実行の高レベルな手順を確立します
しかし、これらの手順は反復的に実行され、フィードバックとモニタリングが組み込まれているため、組織は進化的な方法で目標状態への進捗を測定し、モニタリングすることができます
反復は、常に価値を提供し、すべてのステップをエンドツーエンドで実行します
詳細については、「データメッシュ実行フレームワーク」のセクションを参照してください
この章の残りでは、この高レベルな実行フレームワークを使用するための実用的なガイドラインを共有します
次の章では、データメッシュ実装の文化的および組織的側面について説明します


データメッシュ実行フレームワーク

優れた戦略は競争力を高めますが、堅実な実行で持続させることができます

​—​ゲイリー・L・ニールソン、カーラ・L・マーティン、エリザベス・パワーズ、「成功する戦略の実行の秘訣」（ハーバード・ビジネス・レビュー）

データメッシュの実行を始めるには、まず、それが変革的な変化を引き起こす可能性が高いことを認識する必要があります
文化、組織の構造、技術などです
私たちがここで共有したいのは、進むための高レベルなフレームワークです
このフレームワークは、あなたが一回ずつ進んで進行し、価値と成果を提供し、基盤（プラットフォーム、ポリシー、オペレーティングモード）を成熟させるのに役立ちます
大きな変化を推進するために必要な意志とエネルギーを集めるために実行を遅らせるのではなく、前に進む動きの運動エネルギーを利用して継続的に進むことができます
図15-3は、データメッシュ実行フレームワークの高レベルな要素を示しています
ビジネス主導、エンドツーエンド、反復的、進化的な方法です


これらの側面がデータメッシュの実装における活動、成果、および測定にどのように影響するかについて見ていきましょう


ビジネス主導の実行

データメッシュは、より大きなデータ戦略の一部であり、図15-2で示されているように、戦略的なビジネスケースとMLまたは分析が必要とするユーザシナリオによって実行されます
これにより、データ製品の特定と提供、およびそれらを構築、管理、消費するためのプラットフォームのコンポーネントが推進されます
これは、Amazonが主要な製品とイニシアチブの構築に取り組む方法に類似しています
「顧客から始めて後戻りする - これは思ったよりも難しいですが、顧客への革新と喜びの明確な道です
効果的なワーキングバックツール: 製品を構築する前にプレスリリースとFAQを作成することです
この手法には多くの利点がありますが、いくつかの課題が管理される必要があります
ビジネス主導の実行の利点 ビジネス主導の実行には以下のようないくつかの主要な利点があります
 価値と結果の継続的な提供とデモ 長期的な変革の中で結果を示すことは、士気への肯定的な影響を与え、賛同を得たり、コミットメントを高めたり、投資の継続を保証する効果があります
ビジネス主導の実行は、最初の段階ではペースが遅くなるか、影響範囲が小さくなるかもしれませんが、常にビジネス価値を提供し続けます
プラットフォームの機能と運用モデルが成熟すると、ペースと範囲が向上します
 利害関係者からの迅速なフィードバック ユーザーとの関係（内部および外部）は、プラットフォームとデータ製品の使いやすさに肯定的な影響を与えます
組織は、ビジネス、データ、およびプラットフォームの利害関係者のフィードバックに基づいて必要なものを構築し、進化させます
フィードバックは早期かつ反復的に組み込まれます
 無駄の削減 ビジネスの使用事例に基づいてデータ製品とプラットフォームを開発することは、必要な時に必要なものを構築するため、無駄を減らすことにつながります
技術投資はビジネスに価値を提供することに焦点を当てており、単に構築計画に従うだけではありません
 ビジネス主導の実行の課題 この手法にはいくつかの課題があり、注意深く管理する必要があります
時間切れの解決策の構築 一度に1つのビジネス使用事例に対する解決策を提供する狭い焦点は、将来の使用事例に対応できないプラットフォームの機能とデータ製品の開発につながる可能性があります
これはビジネスケースに基づく開発の可能な負の結果です
特定のビジネスアプリケーションに焦点を当てず、その先を見越さずに、特定の使用事例に狭く設計されたポイントソリューションの構築に終始してしまうことがあります
これは特にプラットフォームサービスやソースに整列した集計データ製品にとって問題です
 ポイントソリューション構築の是正方法は？ ポイントソリューションの構築とYAGNI原則（「実際に必要ではない場合」事前に作る）の適切なバランスを見つけるために、製品思考が重要です
プラットフォームサービスとデータ製品に製品所有権の技術を適用することで、今日の顧客に価値を提供すると同時に、将来や今後の使用事例に目を向けることができます
データ製品の所有者は、分析的使用事例の無限性を理解しています
私たちが今日把握していない将来の使用事例では、過去のデータにアクセスしてパターンとトレンドを発見する必要があります
したがって、データ製品は既知の使用事例の正確なニーズを超えてデータを収集する必要があります
これは特にソースに整列したデータ製品（「ソースに整列したドメインデータ」）に当てはまります
プロダクトと事実をできるだけ正確に捉える必要があります
 タイトなビジネスの締め切りへの対応 ビジネス主導の実行では、プラットフォームとデータ製品の開発をビジネスの締め切りに従わせます
ビジネスの締め切りは、組織内の多くの関係者、パートナー、および顧客との調整を必要とするため、変更が困難です
これらの締め切りに対応するために、プラットフォームサービスなどの下位レイヤーでは、ビジネスケースの締め切りを守るために品質を犠牲にするテクニカルデットを発生させることがあります
これはソフトウェア開発における新しい現象ではありません
私たちはほとんど常にソフトウェアを独自の空間で開発せず、外部の締め切り、ドライバー、予期せぬ環境的状況に直面し、機能の優先順位と構築方法が変わります
締め切りに直面した技術的負債の是正方法はどうですか？ プラットフォームやデータ製品の長期的な所有権と常緑的なエンジニアリングの実践は、新機能の提供に伴う技術的負債に対処するための持続的な改善プロセスが存在することを保証することができます
 長年のプラットフォームとデータ製品チームには、特定のプロジェクトの提供を超えて、システムとデータを維持するための知識とインセンティブがあります
 これにより、プラットフォームの拡張性と長期的な設計、短期のプロジェクトの提供のバランスが取られます
 良い例として、プラットフォームサービスやデータ製品の長期的な所有者は、時間的に追い込まれた状況であっても、APIやインターフェースを正しくすることが優先事項であることを知っています
 実装の詳細に発生した技術的負債を修正する方が、インターフェースを変更するよりも簡単です
 APIの変更は、多くの利用者に影響を与えるため、変更のコストが高くなる可能性があります
 プロジェクトベースの予算化 ビジネスイニシアチブの提供は、しばしば時間の制約のあるプロジェクトと固定された予算で組織されます
 データメッシュの実装をビジネスイニシアチブを通じて行うと、メッシュの長期間のデジタルコンポーネント、データ製品、およびプラットフォームの予算とリソース配分に否定的な影響を与えることがあります
 予算が時間的に制約のあるビジネスイニシアチブに割り当てられると、プラットフォームサービスやデータ製品の構築に対する投資の唯一の源となります
 その結果、それらの長期的な所有権、持続的な改善、および成熟度は低下します
 プロジェクトベースの予算化をどのように是正すればよいでしょうか？ ビジネスプロジェクトは、長年のプラットフォームやデータ製品の実行に適した手段です
 しかし、それらが唯一の資金調達メカニズムではありえません
 プラットフォームやデータ製品には、独自の長期的なリソースと予算が必要です
 プロジェクト志向は、プラットフォームと製品思考とは相反します
プロダクトチームは決して完成せず、新しい要件に対応し、多くのプロジェクトをサポートするために製品を成熟、維持、進化させ続けます
 ビジネス主導の実行のためのガイドライン ビジネス主導のデータメッシュの実行を開始する際のいくつかのガイドラインと慣行を以下に示します： 補完的なユースケースから始める 1つのユースケースではなく、いくつかの補完的なユースケースを選ぶことで、ポイントソリューションの構築を避けることができます
 補完的で同時進行のユースケースを通じてメッシュの再利用可能なコンポーネントを構築することは、ポイントソリューションの落とし穴を回避するための強制機能になります
 たとえば、初期のユースケースは、分析とレポーティングなど、異なるが補完的な領域から選ばれることがあります
これにより、データ製品の設計が、MLとレポーティングの2つの異なるアクセスモードを対象にし、他のモードには拡張性のない方法でソリューションを提供することを強制します
 データコンシューマーとデータプロバイダーのペルソナを把握し、優先順位をつける データメッシュのユーザーペルソナであるデータプロバイダーとデータコンシューマーの両方に焦点を当てることで、プラットフォームの機能を公平に評価し、最も終端までの影響を考慮することができます
 インテリジェントなソリューションの構築に関与する人々、アプリケーション開発者、データ製品プロバイダー、データ製品コンシューマーを知ること
 彼らのツールとスキルを理解し、その情報に基づいて仕事の優先順位を付けます
 たとえば、初期のプロジェクトは、プラットフォームの上級ユーザーと見なされる開発者の集団によって実施されるのが理想的です
 彼らは欠落している機能と必要なプラットフォーム機能の作成に参加することができます
 私の経験では、多くのデータメッシュの実行は、データ製品の作成を重視して始まります
その結果、データコンシューマーの体験が悪化し、全体的な結果が妥協されます
 プラットフォームの欠落した機能に最小限の依存性を持つユースケースから始める プラットフォームはユーザーに応じて進化します
ある時点で必要なすべてのユースケースの特徴を満たすことはできません
実行の初期段階では、ユースケースの選択が慎重に行われることが重要です
初期のユースケースは、プラットフォームに十分な依存性があるため、完全にブロックされることなく、必要なプラットフォ実行の初期段階では、ユースケースの選択が慎重に行われることが重要です
初期のユースケースはプラットフォームへの依存関係が十分にあり、完全にブロックされることなく必要なプラットフォームサービスの優先順位付けと構築を推進します
欠落しているプラットフォームの機能に完全にブロックされないユースケースと開発者のペルソナから始めてください
これにより、勢いが生まれます
欠落しているプラットフォームの機能は、ドメインチームによってカスタム開発され、後で収集され、プラットフォームに一般化されることができます
プラットフォームサービスとデータ製品の長期的な所有権と予算を作成しながら、ビジネスイニシアチブと個々のユースケースはメッシュを開発するための投資を提供することができますが、メッシュは長期にわたるデジタル投資であり、会社の内部製品として考慮する必要があります
それは創設以降に価値を提供し続けることが期待されています
したがって、長期的な進化と運営を維持するためには、独自の割り当てられたリソースと投資が必要です
ビジネス主導の実行の例は、この章の冒頭で紹介した例を詳しく見て、ビジネスイニシアチブがデータメッシュのコンポーネントの実行を推進する方法を見てみましょう
ダフは「コンテキストに応じたパーソナライズされたコンテンツをインテリジェントに展示する」という戦略的イニシアチブを持っています
このイニシアチブは、MLと洞察を使用して音楽、ポッドキャスト、ラジオチャンネルなどのさまざまなコンテンツをリスナーに提示・提案するために、プログラムを使用する取り組みを具体化しています
このプログラムでは、ユーザーのインタラクション、プロフィール、好み、ネットワーク、位置、地元の社会的イベントなどのデータを活用して、推奨とキュレーションされたコンテンツの関連性を継続的に向上させる予定です
このビジネスイニシアチブにより、「好みについて教えて」「パーソナライズされたプレイリストブラウザ」「似たような音楽の探索者」など、いくつかのMLベースのアプリケーションの開発が行われます
これらのアプリケーションは、気分、活動内容、時間帯、休日、音楽ジャンルなどといった多くのデータ次元に基づいておすすめのプレイリストと音楽に依存しています
これらのキュレーションされたプレイリストのためのデータを提供するには、複数のデータ製品を開発する必要があります
データ製品のリストには、月曜プレイリスト、日曜の朝のプレイリスト、スポーツプレイリストなど、MLを利用したコンシューマーに合わせたおすすめのプレイリストが含まれます
これらのデータ製品を構築するには、リスナープロファイル、音楽プロファイル、リスナープレイイベント、カレンダーなどのいくつかのソースに合わせた集計データ製品が必要です
プラットフォームの機能の成熟度と可用性に応じて、プラットフォームはデータ製品の開発者と消費者のニーズを優先します
例えば、データ製品の数が増えてきたため、プラットフォームはデータ製品の自動ライフサイクル管理を優先するか、MLベースの変換をサポートし、フローベースの変換エンジンをサポートするかもしれません
ユーザープロファイルと情報のプライバシーの問題に応じて、ガバナンスとプラットフォームは、個人を特定できる情報（PII）の暗号化とアクセスの確立、コード化、自動化を優先します
このプロセスは、新しいビジネスイニシアチブの導入に伴って繰り返されます
たとえば、「アーティストの知的な識別、オンボード、プロモーション」などです
 エンドツーエンドでの反復的な実行 大掛かりな再設計はただの大爆発を保証するだけです
-Martin Fowler エンドツーエンドで知的なビジネスの価値を継続的に提供することにより、メッシュを豊かにするための連続した反復が行われます
データ製品の数が増えます
プラットフォームサービスが向上します
自動化ポリシーのカバレッジと効果が向上します
図15-4は、前述の例に基づいたデータメッシュ実行のエンドツーエンドの反復を視覚化しています
いつでも複数の反復が同時に実行されています
前述のようなユースケースの実行を通じて、さらに多くのデータ製品が作成され、より多くのアプリケーションが知的に拡張され、さらに多くのドメインがメッシュの一部となり、より高度なプラットフォームサービスが成熟し、より多くのポリシーが計算可能になります（図15-4）
成熟し進化するメッシュのコンポーネントは、ビジネスユースケースの反復を通じて形成されます
進化的実行データメッシュの操作モデルとアーキテクチャを作成するための提案された実行モデルは、増分的で反復的です
各反復ごとに、組織は学び、次のステップを洗練させ、徐々に目標状態に向かって進んでいきます
私は、賢明に設計されたタスクとマイルストーンの厳格な事前計画を推奨しません
経験から言えるのは、そのような計画はスケールの大きなビジネスの実行と不安定な環境で最初の試みには耐えられないと考えているからです
言い換えれば、EDGEなどのフレームワーク（価値志向のデジタルトランスフォーメーション）は、成果の取得、軽量な計画立案、継続的な学習、アジャイルな原則に基づく実験を重視することで、変革の操作モデルを定義する際に非常に役立ちます
選択する変革フレームワークに関係なく、データメッシュの実行処理は進化的なプロセスとして扱われます
進化的プロセスは実装の成熟度の異なる段階を認識し、実行を連続的に最適な結果とより高い成熟度に導く一連のテストとメトリクスを持っています
このセクションでは、進化的実行をナビゲートするのに役立つ2つのツールを共有します
1つは、マクロレベルでの実行をガイドする多段階採用モデルであり、長期のマイルストーンに目を向けたものです
もう1つは、各反復のマイクロレベルでの進化をガイドする目標関数のセットです
多段階進化モデルイノベーションの採用のsカーブは、データメッシュのコンポーネントの進化をフレームワークするためにかなり便利なツールだとわかりました
カーブの形状は組織によって異なるかもしれませんが、急な斜面や緩やかな斜面、短いテールや長いテールを持つものがあります
ただし、カーブの一般的な段階を使用して進化へのアプローチをガイドすることは可能です
これは、長期の進化的なマイルストーンをガイドするのに役立ちます
図15-5は、データメッシュの成長のsカーブを示しています
曲線の各フェーズは、組織内の異なる採用者（イノベーター、早期採用者、大多数の採用者、後期採用者、遅れた採用者）にマッピングされます
そして、曲線の各フェーズは、開拓をするときの探索、スケーリングをするときの拡大、スケールで持続する運用をするときの抽出の対応する開発モードにマッピングされます
これは新しいイノベーションを採用するための事実上のモデルであり、データメッシュのコンポーネントと原則の進化のフェーズをガイドするために適応されることができます
例えば、プラットフォームコンポーネントの実行にこれがどのように適用されるか見てみましょう
プラットフォームの実装は、イノベーターや早期採用者のリスクを冒す人や専門ユーザーにサービスを提供するためにプラットフォームの初期機能が構築されるという探索意識のもと、起動の段階を経て進行します
プラットフォームが成熟するにつれて、それは自己サーブ能力を拡大し、ハンドホールディングのないドメインの一般的な技術者の人口にサービスを提供します
拡大フェーズでは、プラットフォームは使用する際の摩擦を減らし、遅いプラットフォーム採用者のさまざまなペルソナを満足させるために、サポートサービスを成熟させます
ユーザーやそのニーズの大部分が対応された後、プラットフォームの変更は最小限にとどまり、組織は現在の状態で利用可能な機能を抽出し活用します
このフェーズでは、プラットフォームのAPIとサービスが新しいデータ製品とユースケースで再利用されます
大規模なデータ管理の業界革新の到来と、組織がプラットフォームを新しい方法や斬新な方法で利用しようとする試み、そしてベンダーが基盤となる技術を変更することで、プラットフォームは次のS字カーブの進化に進み、このプロセスが繰り返されます
同様に、このカーブをドメイン所有の実行計画に使用することができます
進化のフェーズに応じて、対応するドメインのクラス（イノベーター、リーダーアダプター、レイトアダプター、ラガード）を特定し、最初にメッシュに参加するドメインを優先することができます
たとえば、データの専門知識があるドメインをイノベーターとして選び、その後に多数のアダプターを持つドメインに移行します
早期のアダプタードメインを他のドメインのための道筋にするために使用します
進化の各フェーズにおいて、開発される異なる機能とチームの振る舞いについて意図的にする必要があります
たとえば、初期のフェーズでは、チームはより探索的な姿勢をとり、ポリシー、プラットフォーム、データ製品を確立するためのマニュアル操作の許容範囲を持つべきです
拡張フェーズでは、メッシュは手動の活動を最小限に抑えるために成熟する必要があり、負荷を監視するための包括的な観測が行われ、優れたポリシーカバレッジで使用量をスケールする必要があります
それでは、S字カーブをデータメッシュの各コンポーネントに適用し、各進化フェーズの特性を確立してみましょう
ドメイン所有の進化フェーズ図15-6は、ドメイン所有の寸法が採用フェーズを通じて進化する例を示しています
図15-6ドメイン所有の特性の例
探索とブートストラップフェーズでは、データメッシュの実行に関与するドメインはわずかです
これらの初期のドメインは、最小限の依存関係を持つエンドツーエンドの価値提供に参加するために、データの提供者とコンシューマーの両方として行動することが多いです
ただし、両方の役割がこのフェーズに参加するためには必要です
このフェーズでは、ドメインは探索と成長の基礎を築くことに集中しています
早期のアダプタードメインは、データ製品オーナーの役割を定義し、ドメインチームの構造を構築するという合理的な慣行を確立しています
彼らは、アプリケーションからソースに合わせたデータ製品への操作データの統合のためのツールとプロセスを作成することによって道筋をつけます
初期のドメインは、データと分析の能力において最も先進的なものであり、モダンなテクノロジースタックを持っています
この段階では、高い価値のユースケースの数を解除することによって価値が生まれます
拡大とスケールのフェーズでは、急速な増分で多くのドメインがメッシュに参加し、大半が中間的な立場になります
早期の大多数は道筋をつけ続け、技術的にも組織的にも繰り返し可能なパターンを確立し、実行を迅速に多くのドメインにスケールアウトさせる必要があります
新たなドメインは、ソースに合わせたデータ製品とコンシューマーに合わせたデータ製品、および集計データ製品に貢献します
ドメインの数が増えるにつれて、集計データ製品と純粋なデータドメインの需要が生まれます
ドメインは、さまざまなテクノロジースタックを持ち、統合やメッシュへの移行が必要な古い遺産システムを非常に多く持っている可能性があります
彼らは、レガシー統合を処理するための新しい慣行とツールを導入します
この段階では、ポジティブなネットワーク効果から価値が生まれます
抽出と維持のフェーズでは、ほとんどのドメインが自律的なデータ所有モデルに移行し、ドメインの数が安定します
データ製品の数が膨大でボトルネックとなるドメインは複数のドメインに分割される可能性があり、ユーザーが最小限に依存しているドメインは他のドメインと統合される可能性があります
このフェーズでは、プロバイダードメインは引き続きデータ製品を修正し、進化させ、最適化しますが、新たなシナリオの構築と再利用可能なデータ製品の大量の活用がコンシューマードメインで行われる大半の活動です
データの場所やシステムとデータ技術の年齢により、メッシュと技術的に整合しないドメインは、大部分の機能と組織的な実践が確立されているこの時点でメッシュに参加することを選択することができます
この段階では、価値は組織全体にわたる一貫性と完成度の効果から生み出されます
データの製品の進化フェーズ フィギュア15-7は、進化の段階を経てデータ製品の開発の次元の例を示しています
フィギュア15-7
進化の段階を通じたデータ製品の特性の例
探索およびブートストラップのフェーズでは、わずかなデータ製品のセットが作成されます
初期のデータ製品は、第11章で導入された重要なアフィニティの一部をサポートします
たとえば、出力ポートモードのサブセットのみが利用可能な場合や、最新のスナップショットのみを共有する期間限定の保存期間や、限定された探索能力などです
このフェーズでは、データ製品の開発は探索的な性質を持ち、各データ製品のアフィニティを実装するための最適なアプローチを特定するためのものです
データ製品の開発者は、パターンと合理的なプラクティスを主導しています
初期のデータ製品の開発者は、知識と学習を共有するために協力して取り組んでいます
データ製品の開発は、実験と初期の学習に基づいてプラットフォームの機能に積極的に影響を与えます
このフェーズで開発されたデータ製品には標準の設計図がなく、プラットフォームの機能を十分に活用していません
このフェーズでは、セキュリティと信頼性の観点からリスクが低いデータ製品が選択され、最小限の標準化に対応する傾向にあります
初期のデータ製品は、開発モデルと技術要件の点で異なりながらも相補的な選択されます
これにより、プラットフォームの依存関係が焦点を絞られます
この時点では、選択されたデータ製品の大部分がソースに整合し、少数のデータ製品がユーザーのニーズに対応してメッシュをブートストラップしています
拡大およびスケールのフェーズでは、急速なペースでメッシュに多数のデータ製品が追加されています
初期の採用者がデータ製品の標準的なパターンを確立したことで、データ製品は迅速に開発することができます
これらのデータ製品は、多数のユーザーによる手助けなしで観察、発見、使用される必要な環境を作り出します
データ製品は、MLモデルの実行やストリームベースのデータ処理など、様々なデータの場所やシステムおよびデータ技術の年齢により、技術的にメッシュに合わないドメインは、能力や組織の実践がほとんど確立されているこの時点でメッシュに参加することを選択することができます
この段階では、価値は組織全体にわたる一貫性と完全性の効果から生まれます
データとしての製品の進化フェーズ 図15-7は、進化のフェーズを通じてデータ製品の展開の例を示しています
図15-7
進化のフェーズを通じたデータ製品の特徴の例 インプローブとブートストラップのフェーズでは、わずかなデータ製品が作成されます
初期のデータ製品は、第11章で紹介された必須の機能サブセットをサポートします
たとえば、出力ポートモードのサブセットのみが利用可能で、最新のスナップショットのみを共有するような限られた保持期間や、限定的な発見性の能力があります
このフェーズでは、データ製品の開発は探索的な性質を持ち、各データ製品の実装方法を特定するための最良のアプローチを特定するために行われます
データ製品の開発者は、パターンと合理的なプラクティスを積極的に確立しています
初期のデータ製品の開発者は、知識と学びを共有するために協力して作業しています
データ製品の開発は、実験と初期の学びに基づいて、プラットフォームの機能に積極的に影響を与えます
このフェーズで開発されるデータ製品には、標準の設計図がなく、プラットフォームの機能を完全に活用していません
このフェーズでは、選択されたデータ製品はセキュリティと信頼性の観点から低リスクのものであり、最小限の標準化に対して寛容です
初期のデータ製品は、開発モデルと技術要件の観点で異なりながらも補完的であるように選択されています
これにより、プラットフォームの依存関係が集中します
その時点で、選択されたデータ製品は、特定のユースケースのニーズに対応する消費者に合わせて、ソースに整列したデータ製品の大部分から成ります
 拡大とスケールのフェーズでは、大量のデータ製品が急速なスピードでメッシュに追加されています
早期の採用者がデータ製品の開発の標準化されたパターンを確立しているため、データ製品の開発は迅速に行われることがあります
データ製品は、大勢のユーザーによる手助けなしで観察、発見、使用される必要がある機能を作成します
データ製品は、MLモデルの実行やストリームベースのデータ処理など、さまざまな変換をサポートします
データ製品の開発は、さまざまなソースと消費者をサポートすることに重点を置いています
このフェーズで開発されるデータ製品は、プライバシー、稼働時間、耐久性などに関して、より高リスクのカテゴリに属する場合があります
このフェーズでは、多様なクラスのデータ製品とともに、増加している集計データ製品が作成されます
抽出と持続のフェーズでは、データ製品の数が安定する可能性があり、変化のペースが遅くなります
データ製品の最適化や新しいビジネスのニーズにより、新しいデータ製品が作成されたり、既存のものが変更されたりします
データ製品は引き続き進化します
データ製品は引き続き、大規模な使用状況に対処するために、データへの迅速なアクセス、データ共有のパフォーマンスの向上など、他のクロスファンクションの問題を最適化し続けます
また、ライフサイクル管理を最適化し、変更を適用するスピードを向上させ続けます
変換の複雑さにより変更までのリードタイムが長いデータ製品は、より小さな効率的なデータ製品に分割されます
このフェーズでは、データ製品の開発の事前投資の大規模なリターンを生み出すことに焦点が当てられます
数多くのインテリジェントなビジネスイニシアチブは、既存のデータ製品をそのまま使用することができます
このフェーズでは、遅れている旧来のデータストレージやシステムがデータ製品に変換されます
セルフサーブプラットフォームの進化フェーズ 図15-8は、データメッシュの進化に伴いセルフサーブプラットフォームの特徴がどのように変化するかを示しています
図15-8.進化の段階を通じたデータプラットフォームの特性の例
ブートストラップおよび探索フェーズでは、プラットフォームは主にユーティリティ面で、ストレージ、コンピュート、アカウント、アクセス制御、変換などの基本的な能力を確立しています
基本的なデータ製品のライフサイクル管理やグローバルアドレス作成など、データ製品のエクスペリエンスを提供する一連のサービスが開発されています
基本的な発見性など、メッシュエクスペリエンス面の機能はほとんどありません
このフェーズでは、プラットフォームのセルフサーブ実装は未熟であり、プルリクエストベースのスクリプト共有のようにシンプルなものである場合もあります
プラットフォームのユーザーは、最低限のセルフサーブ能力で済むエキスパートなデータ製品の開発者や消費者です
プラットフォームは少数のユーザーをサポートしています
プラットフォームチームは、早期のドメインと密接に連携して、ドメインチームがプラットフォームにビルドバックするための機能クロスを収穫することなく、彼らのニーズを提供しています
早期のフェーズでは、プラットフォームは意見を持ち、必要なものに焦点を当てるため、サポートされる技術の範囲が限られています
スケールアップおよび拡張フェーズでは、プラットフォームはすべてのアフォーダンスを備えた自動データ製品生成をサポートしていますが、一部の機能が制限されている場合もあります
たとえば、自動データメトリクス共有の基盤は整っていますが、メトリクスセットにはすべての可能な時間的、品質、ボリューム情報が含まれていない場合があります
プラットフォームは主にデータ製品エクスペリエンス面またはメッシュエクスペリエンス面を介して使用されます
プラットフォームのセルフサーブ能力は急速に成熟し、共通の要件を持つ人口の大部分をサポートします
自動化されたプロセスとツールが確立されます
プラットフォームのターゲットユーザーのペルソナは、主に多数派のジェネラリスト技術者に偏っています
プラットフォームは、ハンドホールディングを必要とせずに、利用者の大多数の中間層を対象にしています
プラットフォームが提供する異なるセットの技術を使用したいチームは、プラットフォーム外であり、メッシュガバナンスと互換性の要件を満たすためのエンドツーエンドの責任を持つことになります
このカテゴリに属するチームはほとんどいません
抽出および維持フェーズでは、プラットフォームは自動監視、自動回復、およびスケールでの高い信頼性を持つ成熟した状態にあります
メッシュ内のダウンタイムは大規模な影響をもたらす可能性があります
プラットフォームの機能は成熟し、さまざまなユースケース、データ製品、およびチームに利用されます
プラットフォームは引きこもりや外れ値のためにセルフサーブの能力を最適化し続けます
このフェーズでは、組織内のほとんどのチームがプラットフォームを利用しています
プラットフォームの開発は、メッシュのエクスペリエンスを最適化し、ダウンタイムやセキュリティの脆弱性のリスクを低下させることに焦点を当てています
フェデレーテッドな計算ガバナンスの進化フェーズフェデレーテッドなガバナンスの運用モデルとメッシュを規制する計算ポリシーは採用フェーズを通じて進化していきます
図15-9は、ドメインの参加数、フェデレーテッドな運用モデルの成熟度、ガバナンス開発の焦点、計算ポリシーのカバレッジなどのガバナンス特性の進化を示しています
図15-9.進化の段階を通じたガバナンスの特性の例探索およびブートストラップフェーズでは、少数の早期ドメインがフェデレーテッドなガバナンスに参加している形で新しいガバナンスモデルが形成されます
ガバナンスの運用は、意思決定プロセスを導く原則を確立し、組織がグローバルに実施しなければならないポリシーとドメインに任せられるポリシーを決定します
ドメインデータ製品チームのグローバルなインセンティブを定義し、ローカルなドメインのインセンティブを補完します
このフェーズでは、既存のガバナンス構造がフェデレーテッドモデルに再構築されます
現在のガバナンスグループのメンバーは、主題の専門家役を果たすか、プラットフォームチームに参加してクロス機能ポリシーオートメーションの製品管理を支援するか、データ製品オーナーとしてドメインに参加しています
ガバナンスは基盤の設定に焦点を当てています
初期のデータ製品開発者は、初期のポリシーの実装を独自に行うことがあり、プラットフォームはそれらを後で自己サービスの機能に取り込みます
これは、初期のポリシー実装の一部のクラウドソーシング形式です
初期のガバナンスおよびドメインチームは、オペレーティングモデル、意思決定、ポリシーオートメーションの確立に取り組んでいます
ガバナンス機能は、早期のデータ製品とドメインに関連する重要なポリシーを確立します
セキュリティやプライバシーなど、一部の重要なポリシーのみが自動的にプラットフォームによって実装され、すべてのデータ製品で一貫して強制されることがあります
拡大と拡張のフェーズでは、初期のドメインとガバナンスチームによってオペレーティングモデルとポリシーオートメーションのアプローチが確立されたことで、多くのドメインがメッシュに参加できる状態が整っています
メッシュに新しいドメインをオンボーディングするプロセスは摩擦なく行われます
連邦ガバナンスチームのメンバー数は組織内の大多数のドメインを含めるように増加します
チームは新しいドメインの迅速なオンボーディングのためのカバレッジとポリシーの多様性を高めることに焦点を当てます
この時点で、ほとんどのポリシーは自動化され、多くの互換性のある安全なデータ製品をサポートするためのメッシュが形成されます
基本的なポリシーの監視が行われます
抽出と維持のフェーズでは、メッシュ上のすべてのデータ製品は自動ポリシーを埋め込んでいます
ポリシーリスクや整合性の欠如を検出するための自動化プロセスがあり、プラットフォームが提供する自動化ツールでデータ製品チームに懸念点を通知します
チームはガバナンスプロセスを最適化し、自動ポリシーのサポート機能を向上させることに焦点を当てています
メッシュはポリシーのパフォーマンスを監視し、最適化します
フィットネス関数によるガイド付きの進化探索は、データメッシュの実行をマクロレベルで指導する一方、数ヶ月または数年という長い時間スケールでの連続した計測と追跡によって小さな増分と短い反復的な実行サイクルを指針とします
フィットネス関数は、設計ソリューションが目標を達成するかどうかを要約する目的関数であり、進化計算ではアルゴリズムが時間の経過とともに改善したかどうかを決定します
進化的アーキテクチャの構築においては、フィットネス関数は性能、セキュリティ、拡張性などの複数の次元で設定された目標に対してアーキテクチャの反復を客観的にガイドします
フィットネス関数は、進化計算から借用された概念であり、システム全体が意図したデザイン目標に向かっているかどうか、および時間の経過とともに改善しているかどうかを評価するための客観的な尺度です
データメッシュの文脈では、これらの目的関数は、いつでもメッシュの実装がどれだけ「適合」しているか、および各反復が価値をスケールで提供する最適な状態に近づいているかどうかを決定することができます
データメッシュの最終的な目標は、組織がデータを分析目的で活用し、成長と複雑性に適合させる能力を高めることです
データメッシュの各原則がこの高水準の成果に貢献しています
したがって、フィットネス関数は各原則から期待される結果に焦点を当てることができます
ただし、良いものの具体的な基準をまだ把握していない場合、数値のフィットネス関数を作成することは非常に困難です
よく確立されたDevOpsのエンジニアリングプラクティスでも、書籍「Accelerate: Building and Scaling High Performing Technology Organizations」では、数年にわたる定量的な研究と24の主要指標の測定により、わずか4つの指標のみが組織の成果と直接的な関連があり、残りの多くは単なる虚栄心の指標であることが示されています
ここで提案されている客観的な指標は例と出発点です
これらは実験し、定量的に測定し、組織間で比較して、それが正しいものであることを証明する必要があります
私が重要パフォーマンス指標（KPI）よりもフィットネス関数のコンセプトを使用していることに気づいたかもしれません
両方のコンセプトは基本的には客観的な指標についてであり、データメッシュの実行進捗状況の状態とトレンドの可視化を行いますが、フィットネス関数は、データメッシュの設計と実装がどれだけ「適合」しているかを任意の時点で示し、重要な振る舞いと結果に焦点を当てて、適切な変更を行い、負の影響を与える変更を元に戻すためのガイドになります
例えば、「データ製品の数」と関連するKPIは、多くのリーダーにとって成功とメッシュの成長の指標として魅力的に聞こえるかもしれません
これが実際には正しい尺度ではありません、なぜならそれはメッシュの「適合」を示していないからです
メッシュのターゲットは「価値を提供する」ための設計です
これは直接的にはデータ製品の数と関連しておらず、使用量の尺度として「データ製品へのリンクの数」と関連しています
リンクと相互接続性が多ければ多いほど、より多くのデータ製品が使用されて価値が生成されます
データメッシュの実行には、これらの指標の自動計装、収集、トラッキングが必要です
プラットフォームサービスは、これらの指標の追跡と外部化に重要な役割を果たします
例えば、ポリシーの自動化の有効性に関する「ポリシー採用の加速」を、データ製品が新しいポリシーを採用するまでのリードタイムとして計測することができます
これは、データ製品の継続的なデリバリーパイプラインの計装として収集できます
データ製品とメッシュ体験プレーンのサービスは、ドメインおよびグローバルガバナンスチームがこれらの実行指標を測定し、報告し、追跡するのに役立ちます
間違った方向への傾向を検出すると、メッシュの目的成果に向けた進化を促すための一連のアクションがトリガーされ、原因を調査し、修正の優先順位を付けることができます
データメッシュの実行には、一連の客観的なフィットネス関数を定義することをお勧めします
できるだけ自動化して実施し、トレンドを継続的に監視してください
これらは、データメッシュの実行を案内し、次の最善の手段を取ることができます
ここで紹介したフィットネス関数は、あくまで例です
それらが各データメッシュの原則の予想される成果からどのように導き出されるかを示しました
同じアプローチを採用して、自分自身のフィットネス関数をテストし、進化させることができます
ドメイン所有フィットネス関数 図15-10は、ドメインデータ所有の進捗状況を測定する一連のフィットネス関数の例を示しています
これらのフィットネス関数は、進化がドメインデータ所有の成果と一致しているかどうかを示す一連の目標を測定します
実装は、データ共有を通じてより多くの価値を生成しているのか、組織の成長と一致していますか？ 図15-10.ドメインデータ所有のフィットネス関数の例 データとしての製品フィットネス関数 図15-11では、データメッシュの実装がデータとしての製品の結果に客観的に進化しているかどうかを測定する一連のフィットネス関数の例を紹介しています
効果と効率的なデータ共有は、組織の成長とともに増加していますか？ 図15-11.データを製品として扱うフィットネス関数の例 自己サービスプラットフォームのフィットネス関数 図15-12は、データメッシュの実装が自己サービスプラットフォームの成果に客観的に進展しているかを測定する一連のフィットネス関数の例を示しています
プラットフォームは、ドメインがデータ製品の生産と消費における自律性を向上させ、データ所有のコストを下げていますか？図15-12 自己サービスプラットフォームのフィットネス関数の例 フェデレーテッドコンピュータガバナンスのフィットネス関数 図15-13は、データメッシュの実装がガバナンスの成果に客観的に進展しているかを測定する一連のフィットネス関数の例を示しています
メッシュは、データ製品の効率的な結合と相関によって、組織の成長と合わせて、安全で一貫した高次のインテリジェンスを生成していますか？図15-13 ガバナンスのフィットネス関数の例 レガシーからの移行 データメッシュの進化的な実行について語る際に、レガシーのことについて触れないわけにはいきません
データメッシュを採用する多くの組織は、過去のデータアーキテクチャ（データウェアハウス、データレイク、またはデータレイクハウス）や中央集権的なガバナンス構造と中央集権的なデータチームなどの複数の実装を持っている可能性が高いです
データメッシュの実行は、その徐々に進化する旅において、このようなレガシーの組織構造、アーキテクチャ、技術を増分的かつ進化的に移行する必要があります
しかし、レガシーからの移行を一般化することは難しいです
メッシュの設定が新しい環境（例えば、新しいクラウドプロバイダー）への移行と重なる場合の新しいインフラストラクチャへの移行の時間制約など、組織に固有の多くの変数に高度に依存するからです
この本の範囲外で特定の状態からデータメッシュへの移行戦略を展開することは、その組織との緊密な協力と独自のコンテキストの理解が必要です
しかし、私がこれまでの経験から共有できるのは、高レベルの教訓のリストです
ここから学ぶべきことはまだたくさんあります
データメッシュと共存できる集中型のデータアーキテクチャは存在しない、移行中を除いて私はよく尋ねられます
組織が作成してきた既存のウェアハウス、レイク、およびレイクハウスがメッシュと共存できるかどうかです
短い回答は、移行的には「はい」ですが、最終的には「いいえ」です
いかなる増分的な移行でも、古いものと新しいものが同時に共存する期間が必要です
ただし、データメッシュの目標は、レイクとウェアハウスのボトルネックを取り除くことです
分散型のデータ製品が中央集権的かつ共有のウェアハウスとレイクを置き換えます
したがって、目標アーキテクチャでは、中央集中型のボトルネックがデータメッシュと共存することを意図していません
例えば、中央集権的なデータウェアハウスやレイクを保持し、それに沿ってデータメッシュを下流に供給することはアンチパターンです
この場合、データメッシュの導入は追加の複雑さであり、ドメインからデータを収集するウェアハウスのボトルネックは残ります
代わりに、移行の一環として、データメッシュの実装を作成し、データをウェアハウスに供給し、そのウェアハウスを使用しているレガシーユーザーの数を制限した報告書を提供することが可能です
レガシーウェアハウスユーザーがすべてメッシュに移行するまで、メッシュとウェアハウスは共存します
移行は続き、ウェアハウスの集中度合いを減らしていきます
アーキテクチャの究極の目標は、中央集権化のレベルを取り除き、低減することです
したがって、メッシュの実装と中央集権的なデータウェアハウスやレイクは共存すべきではありません
データメッシュと共存できる集中型のデータテクノロジーが利用できる この本を書いている時点では、データメッシュネイティブのテクノロジーは存在していません
私たちは、既存の技術を活用し、それらの上にカスタムの設定と解決策を構築するという進化の段階にいます
今日のアドバイスは、自律分散型データ製品の構築および運用において、既存の技術要素が邪魔にならないように、既存の技術要素を活用することです
私は、既存の技術をマルチテナントの方法で設定および運用することを提案します
たとえば、オブジェクトストレージなどのレイク指向テクノロジーの基本要素を使用し、各データ製品の境界に基づいてストレージへの割り当てとアクセスを設定します
同様に、データウェアハウス技術を基礎ストレージとし、SQLアクセスモードを使用しますが、各データ製品には個別のスキーマを割り当てます
このような設定に適したすべての既存の技術があるわけではありません
適した技術を活用してください
レイクやウェアハウスをバイパスして、ソースに直接アクセスする既存の消費者を見つけ、彼らのニーズやユースケースを認識し、ソースドメインシステムからデータを消費する新しいデータ製品に直接接続することが最善です
ドメインのオーナーを早期に特定し、データ製品をドメインの真実の表現に応じて開発し、必要な場合は消費者を更新してください
レイクのファイルやウェアハウスのテーブルを表す人工的なデータ製品を作成することは、メッシュの目標であるドメインによる所有権を通じた拡張と、ソースと消費者間のギャップの縮小を達成するものではありません
それは技術的な負債をさらに増やし、ソースと消費者の間に別のレイヤーを作成し、その距離を増大させるだけです
データウェアハウスを消費エッジノードとして使用してください
既存のデータウェアハウスの消費者を評価すると、ウェアハウスに一部の消費者を保持し、メッシュに移行しないことを決定するかもしれません
この決定には複数の理由があります
たとえば、頻繁に変更されない財務報告書などの報告書は、小規模で特定の一部のテーブルに依存し、逆工学や再作成が困難な複雑なビジネスロジックで定義されている場合、ウェアハウスに依存し続ける可能性があります
この場合、ウェアハウスをメッシュの消費者として機能させ、残存する一部のテーブルにアクセスする制限付きのレポートがある状態にします
ウェアハウスやレイクからの移行は、原子的な進化的ステップで行うべきです
データウェアハウスやレイクからの移行は、組織の規模によっては数ヶ月または数年かかる場合があります
その場合、ステップが完了した時点でシステムの状態に技術的負債やアーキテクチャのエントロピーが少なくなるように、原子的な進化的ステップで移行を実行することを強くお勧めします
これは論理的で明らかなことのように思われるかもしれませんが、しばしば無視されます
チームはしばしば古いパイプラインやレイクファイルを置き去りにしながら新しいインフラストラクチャや新しいデータ製品を追加し続けます
時間が経つにつれて、彼らはアーキテクチャの複雑さと維持コストを増やすだけです
この共存の主な理由は、新しいコンポーネントにユーザーや消費者を移行することができないことです
たとえば、新しいデータ製品を作成する場合、既存のレイクやウェアハウスの消費者（レポートやMLモデル）をそれらに移行する時間を確保してください
移行の原子的なステップには、新しいソースに整合したデータ製品、集計または消費者に整合したデータ製品の追加、消費者を新しいデータ製品に移行させ、レイクやウェアハウスのテーブル、ファイル、パイプラインを廃止することが含まれます
これらのステップがすべて実行されるか、いずれかが完了しない場合は、移行ステップは完了していません
この原子的なステップの最後には、アーキテクチャのエントロピーが低減されます
逆に、新しいデータ製品を作成した後に消費者を移行しない場合や、既存のパイプラインやレイクファイルを廃止しない場合は、原子的な移行ステップは完了していません
プロジェクトのスケジューリングとリソースの割り当てが、移行のアトミックなステップに合致していることを確認してください
データメッシュの実行は、より大きなデータ戦略の一部であり、実際には戦略の主要な要素です
ビジネス指向の戦略的なイニシアチブが必要であり、ビジネス目標に合致したデータがどのように価値を生成するかの仮説を定義して、実行を推進する必要があります
MLと分析を活用したビジネスユースケースは、データ製品の特定と提供、ドメインによるデータの採用、ガバナンスとプラットフォームの確立を実行するための手段となります
データ駆動型のビジネス成果を提供する点のソリューションと、複数のユースケースにわたる長期的な共有プラットフォームの構築とのバランスを取ることは依然として課題です
長期的な採用ロードマップに基づく進化的なアプローチや、実装の長期的な目標に対する適合度を測定する機能などのツールは、平衡を確立し、進化的な形で成熟した状態に向かって実行をガイドするのに役立ちます
まれな場合には、組織がグリーンフィールドのデータメッシュ実装を作成する機会があります
しかし、ほとんどの場合、ブラウンフィールドの実行では、データウェアハウス、データレイク、またはレイクハウスなどの既存のアーキテクチャの移行を考慮する必要があります
この場合、旧いアーキテクチャと組織をメッシュに移行するために、データの消費者と提供者、およびその間のすべてを一つずつ移行するアトミックな進化的なステップを実行することに依存します
これらの初期のガイドラインとアプローチは、プロセスを学習し洗練する間に始めるのに十分なものでしょう
最後の章で、チームの組織と組織の文化を進プロジェクトのスケジューリングとリソースの割り当てが、移行の原子的なステップに合致していることを確認してください
データメッシュの実行は、より大きなデータ戦略の一部であり、実際には戦略の主要なコンポーネントです
ビジネス志向の戦略的イニシアティブが必要であり、データがビジネス目標と一致する形で価値を生み出す仮説を定義し、それを推進するための実行力が必要です
MLと分析に基づくビジネスユースケースは、データ製品の特定と提供、ドメインによるデータの採用、ガバナンスとプラットフォームの確立を実行するための手段となります
データ駆動型のビジネス成果を提供するポイントソリューションと、複数のユースケースにわたる長期的な共有プラットフォームの構築との間のバランスは、依然として課題です
長期的な採用ロードマップに基づいた進化的なアプローチや、実装の長期目標に対する適合性を測定する機能などのツールは、均衡を確立し、進化的な方法で成熟した状態に向けて実行をガイドします
まれな場合には、組織にはグリーンフィールドのデータメッシュ実装の機会があります
ほとんどの場合、データウェアハウス、レイク、またはレイクハウスなどの既存のアーキテクチャの移行を考慮するブラウンフィールドの実行が必要です
この場合、古いアーキテクチャと組織の移行は、データの消費者と提供者、およびその間のすべてをメッシュに向けて移行するそれぞれの原子的な進化的なステップの実行に依存しています
これらの初期のガイドラインとアプローチは、プロセスを学びながら洗練させるために十分です
残る課題は、この進化的な実行中に組織のチームを組織化し、文化を発展させる方法についてのいくつかのノートです
これについては、本書の最後の章である第16章で議論します
そこでお会いしましょう