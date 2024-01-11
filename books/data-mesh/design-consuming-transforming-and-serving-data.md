A data product’s primary job is to consume data from upstream sources using its input data ports, transform it, and serve the result as permanently accessible data via its output data ports. In this chapter, I go through the design characteristics of these three basic functions that all data products implement: consume data (“Consume Data”), transform data (“Transform Data”), and serve data (“Serve Data”). Let’s begin with the one that has the most unique properties to a data mesh approach. Serve Data A data product serves domain-oriented data to a diverse set of analytical consumers. It does so through its output data ports (interfaces), introduced in Chapter 9. The output data ports have explicitly defined contracts and APIs. This seemingly simple affordance of “serving domain-driven data” has interesting properties when considering its relationship with the agents in the ecosystem and their capabilities and needs. Let’s look at the relationship between a data product and its data users. The Needs of Data Users Figure 12-1 shows the needs of data users and how a data product serves them. Figure 12-1. Serve data properties to satisfy data users and their needs The requirements driven by data users impose a set of design considerations on how data products serve their data: Analytical data users have a diverse profile Data users—clients accessing and reading data—fall on a wide spectrum of personas and application types: people such as data analysts, data scientists, and data-driven application developers; and systems such as reports, visualizations, statistical models, and machine learning models, etc. Recalling Chapter 3, a data product serves its data in a manner that feels native to such diverse personas. We called this baseline usability characteristic of a data product natively accessible. The design implication of this requirement is to serve data with multimodal access—serving the same data semantic in different formats and modes of access. Analytical data users need longitudinal data The mesh maintains a longitudinal view of the global state of data, wholly, for analytical use cases, and most importantly without an off-the-mesh data lake, warehouse, or any external system to maintain the global state. This continuously changing global state of the data is stored and maintained by a connected graph of data products, and no other architectural elements. This is the meaning of decentralization of architecture. Insights, retrospective or futurespective, are most powerful when they account for the passage of time. Only by having access to continuous data changes over time can we formulate trends, make predictions, and discover correlations between different events across multiple domains. Data mesh assumes time as an ever-present parameter in both presenting and querying data. The design implication to access longitudinal data, data that represents events and state changes over time, is for each data quantum to serve bitemporal data. Analytical data users need a consistent view of multiple domains at a point in time The majority of the analytical use cases process data from multiple data products. Such use cases correlate multiple data products at a consistent point in time. For example, when Daff trains a machine learning model on 2021-07-01 to predict subscriber growth for the next month, it does that based on the data from the last three years, known and processed by multiple data products on 2021-07-01. To support the repeatability of this version of the growth model, processed on 2021-07-21, the mesh maintains an unchangeable state of data across multiple data products on 2021-07-21. Providing point-in-time consistent data across multiple data products, combined with versioning of data for repeatability, introduces multiple design considerations for serving data: bitemporality, immutability, and read-only access. Serve Data Design Properties Let’s go a bit deeper in each of the properties we discovered earlier: multimodal, immutable, bitemporal, and read-only access. These characteristics are integral to the working of data mesh. Multimodal data The job of a data product is to serve the analytical data of a particular domain with a specific and unique domain semantic. However, for a data product to serve its diverse consumers natively, it must share the same domain semantic, in different syntaxes; the semantic can be served as columnar files, as relational database tables, as events or in other formats. The same semantic is served without compromising the experience of the data consumer. People who write reports consume the data as relational tables. People who train machine learning models consume data in columnar files, and real-time app developers consume the events. I personally find it helpful to visualize the nature of analytical data as the dimensions of space and time(s). I use the space dimension to represent the different syntactic materialization of the data, the formats of data. Any data product may present its data with multiple formats, or in a multimodal form, for example: Semi-structured files, e.g., columnar files Entity relationships, e.g., relational tables Graph, e.g., property graph Events When a data user accesses a data product, using its top-level data product discovery API (“Discover, Understand, Trust, and Explore”), it first gets to learn about the data product’s semantic—what domain information the product is serving e.g., podcasts, podcast listeners, etc.—and then it accesses one of the data product’s output APIs (“Output data ports”) for a particular mode of access to the data. How the data is served depends on the underlying (physical) technology: subscribing to an event log, reading distributed columnar files, running SQL queries over relational tables. Domain-oriented semantics is the top-level concern, while formats and the mode of access are second-class concerns. This is an inverse model to existing architectures, where the storage and encoding technology dictates how the data is organized first and then it is served. Figure 12-2 shows the diverse modes of access with an example. The play events data product provides access to its play events data through three modes of access: subscribing to its “play events” topic (capturing the state changes of listeners as they log in, play a podcast, stop playing, etc.), accessing the same play events through SQL queries (over tables with rows of attributes of events), and columnar object files (a file for each attribute of all events). Figure 12-2. Example of a data product’s multimodal access These three modes of access and topologies of the data can satisfy diverse personas of consumers: a data-intensive app developer watching for error events to improve the quality of the player, the data analyst using the synchronous SQL access model to create a daily report of top player usage patterns, and a data scientist using the file access mode to train its ML model discovering classification of play patterns.1 All this is enabled by the same data product, play events, through its individual output data ports. Note Today, the complexity of supporting multimodal access is the symptom of the absence of higher-order and abstraction APIs to present and query data semantically and agnostic to its underlying encoding and format. Immutable data No man can cross the same river twice. ​—​Heraclitus Immutable data, once it is created, doesn’t change. Data products serve their data in a way that once a piece of data is processed and made available to the data users, that specific piece of data does not change; it cannot be deleted or updated. Changing data often leads to complexities and errors, something that is known by any experienced programmer. This is why there’s considerable interest in functional programming, which holds as a key axiom that data is never changed. This is particularly relevant to analytical use cases. With immutable data, data users can rerun analytics in a repeatable way; retraining models or regenerating reports on a particular point-in-time dataset gives the same repeatable result. Repeatability is needed because often results yield a striking observation that requires the analyst to dig deeper. If the data they are using changes, then they may not be able to reproduce the striking result and may wonder if this is due to a data change or a programming error. If there is a bug in the analysis, it’s much harder to track down when working with an unstable data source, where the analyst cannot get the same answers from repeated runs of the same code. The confusion and complexity of mutable data is made worse in data mesh, where source data can be used by multiple data products, and multiple data products can be the sources of a particular analysis where multiple data products each keep a slice of truth contributing to the larger understanding of a state of business. The distributed nature of data mesh demands immutability to give confidence to data users that (1) there is consistency between multiple data products for a point-in-time piece of data and (2) once they read data at a point in time, that data doesn’t change and they can reliably repeat the reads and processing. Figure 12-3 shows a simple example: a listener demographics data product provides geographic demographics of the listeners on a daily basis, where the listeners connect. It has two downstream data products, artists regional popularity (where most listeners for an artist are located) and regional market size (number of listeners in different regions). These two data products are part of the source for regional marketing data product, which recommends regional targeted marketing activities. It uses regional market size to identify low-risk countries with the least market impact to do A/B testing or experiments, and artists regional popularity to promote artists based on their popularity. Figure 12-3. Example of correlating inconsistent data Each of these data products has its own cadence in processing and serving data. Hence, there is no way to guarantee that all their data updates at the same time. This can lead to a deadly diamond. If listener demographics updates its data at the time that regional marketing runs its analysis, it can consume inconsistent data because listener demographics could serve data before the update, and regional market size after. Worse, regional marketing doesn’t even know about this inconsistency. Data mesh addresses this problem by removing the possibility of updated data going undetected by data users. Changes to data are always presented as new pieces of data with identifying time-variant attributes so data users never get to correlate inconsistent pieces of data, each coming from a different data product. The most general way to do this is to use bitemporal data, which I cover in the next section. For example, listener demographics shares its data as tuples like {listener_id: ‘123’, location: ‘San Francisco’, connection_time: ‘2021-08-12T14:23:00, processing_time: ‘2021-08-13T01:00’}. Each piece of information has two time-variant identifying fields, the “connection_time,” e.g., when the listener connected to listen to music, and the “processing_time,” when listener demographics processed this information. Once this tuple is processed and made available to data users, it never changes. Of course, the same listener may listen to content the next day from a different location, and that would manifest itself as an appendance of a new data entity {listener_id: ‘123’, location:’New York’, connection_time: ‘2021-08-13T10:00’, processing_time: ‘2021-08-14T01:00’}. These are two different pieces of time, while they still convey an update, a new location of the listener. More generally, immutability of data reduces opportunity for incidental complexities, complexity in dealing with the side effects of shared states’ updates across a distributed mesh of consumers and solving the intractable computer science problem of distributed transactions. Imagine the complexity that goes into the code of each downstream data reader on the mesh if the mesh allows that the data they have already consumed keeps changing, or replaying the same reads yields a different result. Immutability is another key enabler of design for change and maintaining the eventual consistency2 of the mesh through the propagation of new processing times and their immutable states. Keeping data immutable is important any time but is particularly vital for the data mesh. But it’s also true that past data changes as new information becomes available or data processing bugs are fixed; hence, we need to be able to make retracted changes to data. Let’s look at bitemporality to see how data mesh implements immutability and retracted changes and how it allows for data to change even though it must remain immutable. Bitemporal data It seems that change and time are inseparable: changes take time; are located and ordered in time; and they are separated by time. The inseparability of time and change is a kind of logical truth. ​—​Raymond Tallis Bitemporal data is a way of modeling data so that every piece of data records two timestamps: “when an event has actually occurred or a state actually was,” actual time, in combination with “when the data was processed,” processing time. Bitemporal data modeling makes it possible to serve data as immutable entities, i.e., as tuples of {data product fields, actual time, processing time}, where processing time monotonically increases every time a new data entity is processed,3 and to perform temporal analysis and time travel, i.e., look at past trends and predict future possibilities. Both of these outcomes are essential to data mesh. For example, consider a typical temporal analytical use case that predicts the growth of Daff’s business: a growth prediction model. This model uses subscription changes over a long period of time to discover patterns and trends. It uses the actual time of the subscriptions data product—the time when subscribers actually join or leave their memberships—correlated with other actual times of other related events from the marketing activities, calendar, and support issues data products. These data products process and serve data with a different cadence. To train a repeatable version of the model, using consistent data from different data products, growth prediction selects a common processing time—when the data products have processed and become aware of an event. For example, the version of the model trained on 2022-01-02T12:00 uses the subscriber information, support issues, and marketing events of the last three years (actual time), known and processed, as of 2022-01-02T12:00 (processing time). Actual time and processing time are the two entangled axes of time that a data product preserves and serves: Flux of actual time To represent the data in a form that meets the needs of analytical use cases, the data product captures and shares the state (or events) of the domain over time, over an infinite4 span of time. For example, the podcast listeners data product can share “daily podcast listeners information from a year ago to now.” The actual time is the time when an event occurs or the time when a particular state exists. For example, 2021-07-15 is the actual time of podcast listeners (data) who actually listened to Daff’s podcasts on this date. The presence of actual time matters, as predictive or diagnostic analysis is sensitive to the time when something actually happens. This is not the case for many of the operational functions. Most of the time the operational functions deal with the current state of the data, e.g., “give me the current address of present podcast listeners to send them printed marketing material.” Actual times fluctuate. A data product can observe out-of-order actual times or receive new data about the same actual time after a correction is made on the source data. Continuum of processing time The processing time is the time at which a data product observes, processes, records, and serves its knowledge or understanding of the state or event, for a particular actual time. For example, on 2021-08-12T01:00, the podcast listeners data product processes all data about people who listened to podcasts on 2021-08-11 and gains an understanding of the state of the world then. 2021-08-12T01:00 is the processing time. Providing processing time as a mandatory attribute of data is the key to design for change. Our understanding of the past changes; either we fix errors that occurred in the past or become aware of new pieces of information that improve our understanding of the past. We can’t change the past, but what we can change is the processing time of the present. We serve a new piece of data with a new processing time that reflects fixes of the past actual times. Processing time is the only time that can be trusted to move forward monotonically. Note I use processing time to collapse four different times into one: Observation time When a data product becomes aware of the event or state Processing time When the data product processes and transforms the data observed Record time When a data product stores the processed data Publish time When the data becomes available to data users for access These subtle time differences are most relevant to the internals of a data product and not the data users. Hence, I have collapsed them into one processing time that matters most to the data users. Martin Fowler presents this bitemporality in a simple and brilliant post, “Bitemporal History”. In this section, I summarize how data products adopt this concept, in a unified model, agnostic to the latency of the processing and the shape of the data (events or snapshots). Impact of bitemporality Let’s briefly discuss the positive impact of bitemporality in a few scenarios: Retractions Our understanding of the world continuously evolves. Our understanding of the world can have errors or missing information. In this case at a later processing time, we fix our misunderstandings. For example, at 2021-08-12T1:00 we process all the information about “People who listened to podcasts on 2021-08-11,” but we make a mistake in counting the number of listeners and calculate 3,000. The next time we process the information at 2021-08-13T10:00, we fix the error and create a new count, 2,005 listeners on 2021-08-11. 3,000 and 2,005 become two different values, for the same identity, “Podcast listeners count on 2021-08-11,” captured and shared as two separate states, the state processed at 2021-08-12T10:00 and the state processed at 2021-08-13T10:00. Using bitemporality builds in change, change in the state of data, its data model, SLO metrics, etc. The continuous processing of change becomes a default behavior built into all consumers and the mesh as a whole. This greatly simplifies the logic of data users. Updates to the past are no longer a special case and a surprise. They are just newly processed data that tell a new story about the past. The consumers can keep track and pin the data they have processed at a point in time in the past and travel back to it to access a past revision of data. Of course, building such a system is no easy engineering task. Skew between processing time and actual time Skew is the time difference between actual time and processing time. Only true real-time systems have a negligible skew: the event gets processed, and our understanding of the world gets formed, almost exactly when an event happens. This is very rare in data processing systems, particularly in analytical data processing with a chain of data products, each processing data multiple hops away from the actual source of an event. The presence of the two times informs the data users of the skew so that they can make a decision on how to process the data based on their timeliness needs. The further away from the source-aligned data product, the larger the skew can become. Windowing It is common for a data product to aggregate upstream data over a window of time. For example, play sessions aggregates all the events during an engagement of a listener with the player device—e.g., a series of play events when the listener navigates from one podcast to another, finally picks one, and then drops out after a few minutes of listening and puts the player away. It helps with the behavioral analysis of the listeners engaging with the player. In this case, the actual times of play events can span multiple minutes, a time window. Data users with the knowledge of the window can perform time-aware operations on the data. Reactive processing of continuous change Data mesh assumes a world in constant change. This constant change can be in the form of new data arriving or our understanding of past data evolving, marked by the processing time. Data products can continuously process change by reacting to the event that their upstream data product has changed: a new processing time has become available. Processing time becomes a primitive mechanism for creating reactive and asynchronous data processing between connected data products. Temporality versions all aspects of a data product Processing time, a data product’s notion of time, is built into all attributes and properties of a data product that change over time: data schema, data semantic, relationships between data, SLOs, and other meta. They automatically become versioned by time. Processing time becomes a primitive for versioning permanent properties of a data product, i.e., data schema. The processing time becomes a parameter to correlate time-variant information such as SLOs of the data products. Example Let’s look at the previous example, visually. Figure 12-4 shows the two dimensions of time and their relationships to each other, for a single data product. There are a few notable pieces demonstrated in Figure 12-4: Skew It is inevitable that the data quantum processes an event and gains and formulates its understanding of the state of its data, later than the point in time when an event actually occurs. For example, the state of the daily listeners during 2021-08-11 is known by the system later than its occurrence time at 2021-08-12T01:00. There is a skew of at least 1 hour to know the daily state of podcast listeners, and up to 25 hours. Processing error The total listeners processed at 2021-08-12T01:00 had an error in its calculation and captured the total daily listeners of 2021-08-11 as 3,000 people instead of 2,005, a bit more optimistic. Retractions The error that occurred on 2021-08-12T01:00 was caught by the data product developers, a fix was applied to the processing code, and on the next processing interval at 2021-08-13T01:00 the correct value of 2,005 is reported. Hence, the data served at processing time 2021-08-13T01:00 includes data for daily podcast listeners of 2021-08-11 and 2021-08-12. Figure 12-4. Actual time and record time relationship What Figure 12-4 shows is that the data serving interface and APIs are a function with two time parameters, a processing time and an actual time. For simplicity of use, they can take special and default parameters like latest for data users who are only interested in the latest state of the data. In this example, I have used a low-resolution wall time for simplicity. The governance standardizes how time is standardized, and the platform implements that consistently. The processing time guarantees ordered reads. It can be implemented as an internal system time, an incremental counter since an epoch, or a monotonically increasing number. Data consumers use the processing time as an index to know what data they have consumed.5 The actual time can follow a DateTime standard like ISO 8601. The data users of analytical data can perform time travel and go back and forward in time when accessing data. The upper bound of how far they can travel back can vary: to the first data ever processed or just to the latest one. This depends on the data product’s retention policy. Figure 12-5 demonstrates a simple visualization of this time travel across the axis of processing time. Figure 12-5. Serving data on the axis of processing time States, events, or both How systems encode and process data falls into two camps: states and events.6 State captures the condition of a system at a point in time, e.g., “number of podcast listeners today.” Event captures the occurrence of a particular change of state, e.g., “a new podcast listener connected.” While state or (change) events require two very different ways of storing or serving data, I consider that an orthogonal concern to the time dimensions we discussed here. It’s up to the data product’s logic whether it makes sense to capture or serve a stream of change events, or the stream of inferred states of the system as snapshots, or do both. It’s likely that data readers would like to see both. Nevertheless, the presence of the two time axes remains. Reduce the opportunity for retracted changes As discussed early, changing past data (retractions)—due to fixes or arrival of new data—is handled like any other newly processed data. Fixed data is presented as a new data entity with an actual time that belongs to the past, but with a new processing time. Additionally, in the next section, “Read-only access”, I will present a special-case handling of data updates such as exercising the right to be forgotten at the mesh level. Despite the methods to handle retractions, data products must strive to reduce errors and the need for retractions. Here are some strategies for reducing the need for fixes: Increase the processing time interval to introduce quality control Imagine the play events data product. It captures the signals coming from player devices. The data product can occasionally miss events or receive delayed ones. This is due to network interruptions, no access to the on-device cache, etc. However, the data product transformation code can introduce a processing delay to correct the missing signals through synthetically predicted/generated ones, or aggregate the signals to a median or other statistical representation of the same data. This means that the data product is introducing a longer skew between the actual time and the processing time to have an opportunity to fix the errors in time and avoid retractions. This is often a suitable technique for business processes that have an inherent need for reconciliation. For example, while receiving payments transactions near real time, there is a daily dump of corrected and reconciled payment accounts that get served later. This would require increasing the processing interval. In the case of payments, the correctness of accounts is preferred over the timeliness of transactions. Adjust the data product’s SLOs to reflect the expected errors If we continue the preceding example, some consumers might be perfectly tolerant of the errors in the near-real-time data. For example, an app detecting player errors, on best effort, doesn’t care if there are missing events here and there. In this case, the play events data product can publish the data without reconciliation for a category of consumers like the “player error detection” app. However in this case, the data product communicates its quality objectives according to the expected range of errors that may occur in terms of missing signal. Read-only access Unlike other analytical data management paradigms, data mesh does not embrace the concept of the mythical single source of truth. Every data product provides a truthful portion of the reality—for a particular domain—to the best of its ability, a single slice of truth. Data can be read from one data product, morphed and transformed, and then served by the next. The mesh maintains its integrity through the propagation of change as appends of bitemporal immutable data from an upstream data product to the downstream consumers, automatically. As such, the mesh maintains a state of eventual consistency as new data propagates through the graph. By now, it might seem obvious that the mesh or individual data products do not present a direct update function; updates are the indirect result of the data product processing its input. Change to a data product, in the form of appends of newly processed data, is only carried out by the transformation code of a data product. This guarantees immutability and also maintains the eventual consistency of the mesh. However, there are cases such as the global governance administrative functions executing the right to be forgotten, in accordance with regulations such as GDPR, that change the data. You can think of these as a specific administrative function, triggered by the mesh experience plane, on all data products’ control ports (“Control Port”), executing the command, in this case a crypto shredding. Data products always encode the user information encrypted; by destroying the encryption keys—which live on the platform and not in the data product—the user information is essentially rendered unreadable. If you do find yourself with other special-case operations that require updating data that is already processed, this operation is implemented as a global control port function, and not a function of an output port. Keep the output ports for data users to just read data. Serve Data Design Let’s bring it together and look at the design of the data product to serve data. Figure 12-6 shows a likely logical architecture, where data products own and maintain a core representation of their domain and serve that in multiple spatial modalities, using the concept of output data port adapters. Each port, always, provides data in a bitemporal, immutable, and read-only mode. The data retention’s duration depends on the policy of the data product. It may retain and make data accessible for many observations (processing times), just the latest, or somewhere in between. Figure 12-6. Data product’s high-level components to design serve data Table 12-1 summarizes the data product components involved in serving data. Table 12-1. High-level data product components to serve data Serve data component Description Output data port Interface (API) to serve data according to a specific mode of access for a particular spatial format of data (syntax). This implementation could simply be an API that redirects access to a particular physical technology, e.g., a bitemporal table in a warehouse storage, a file in a lake storage, or a topic of an event log. Output (data) port adapter The code responsible for presenting the data for a particular output port. The implementation could simply be a step in the data product’s transformation code that stores data in a particular syntax, or as sophisticated as a runtime gateway for adapting core data stored to many modes of access on read. Core data semantic Expression of the data semantic—agnostic to the modes of access or its spatial syntax. Consume Data In the majority of cases, data in organizations originates from the internal or external operational systems, in interaction with people or other operational agents, such as devices. In some cases, data is received as purchased or freely downloaded archives. For example, the operational data in Daff is generated by listeners subscribing and listening to different content, content providers publishing music, and the artist management team paying and handling the affairs of the artists, and so on. Data products consume this data and transform and serve it in a manner suitable for analytical use cases. Hence, in the majority of the cases, data products consume data from one or multiple sources. Architecturally, input data ports (“Input data ports”) implement the mechanisms necessary for a data product to define and execute the consumption of its source data. The input data port is a logical architectural construct that allows a data product to connect to a data source, execute queries, and receive data—events or snapshots—as a continuous stream or a one-off payload. The choice of the underlying technology is an implementation-specific concern and is left to the data platform. Figure 12-7 shows a high-level picture of how data products consume data. A data product consumes data from one or multiple sources. The sources could be a collaborating operational system or other data products. The consumed data is then transformed into the core data model and served in multiple modalities via its output ports. Figure 12-7. Data product’s design to consume data Each data product can have one or multiple sources. Source-aligned data products (“Source-Aligned Domain Data”) mostly consume their data from the operational system. As we get into aggregate data products (“Aggregate Domain Data”) we see other data products becoming the source, and consumer-aligned data products (“Consumer-Aligned Domain Data”) often include smart logic or machine learning models to provide locally sourced data for particular use cases. Let’s dive into some of the notable characteristics that impact the design of a data product’s input data. Archetypes of Data Sources The design of data input capability must support multiple archetypes of sources. Here are a few top-level categories: Collaborating operational systems Other data products Self Collaborating operational systems as data sources Source-aligned data products consume data from their collaborating operational systems, one or multiple applications that generate data as a byproduct of running the domain’s business. They consume data optimized for the operational system and transform it to a format suitable for analytical usage. They decouple downstream analytics use cases from the internal details of the operational application. The phrase collaborating here implies a tight coupling between the data product and its source operational system. They both must belong to the same domain. The operational data contract between the source application and the data product is often tightly coupled; changes to the operational system’s data and model impact the data product’s data and model. That’s why I highly encourage keeping the collaborating operational system sources and their collaborating data products in the same domain. This allows the domain team in charge of the changes to the operational systems to collaborate closely with the data product developers to keep the two in sync. There are cases where an operational system is not aligned to a single domain, e.g., CRM COTS (customer relationship management as a commercial off-the-shelf) products encapsulate multiple domains of products, customers, sales, etc., into one. In this case the COTS system can expose multiple domain-aligned interfaces, each for a particular source-aligned data product. Figure 12-8 shows an example of a source-aligned data product input. The listener subscriptions data product consumes data from a microservice in its domain, listener subscription service. It receives the changes in listener subscriptions as near-real-time events published on an event log. The subscription events log is controlled and maintained by the operational system, as a short-retention medium. The data product consumes the events, performs transformations on them, and ultimately serves them as a long-retention temporal view of listener subscription information. Figure 12-8. Example of a data product with input from a collaborating operational system Common mechanisms for implementing the input port for consuming data from collaborating operational systems include asynchronous event-driven data sharing in the case of modern systems, and change data capture for legacy systems that are difficult to change. Modern operational systems sharing their domain events is increasingly becoming a common practice and is a great model for feeding data to a collaborative data product. Change data capture is a set of integration patterns to discover and track changes to an application database, which then can be sourced as the input to the data product. This is the least desirable way to receive data from a collaborating operational system. It exposes the internal implementation of a database transaction and doesn’t correspond to a business domain. But sometimes it is the only available option in the case of legacy systems. How a data product consumes data from an operational system is heavily influenced by the ability of the team to extend the operational system. The design ultimately depends on the agreement within the domain’s team on how to integrate their operational system with their data product. Other data products as data sources Data products can consume data from other data products on the mesh. For example, the podcast demographics data product receives attributes about the listeners from listener profiles and podcast listeners data products. It correlates the listener profile information with the podcast listeners and applies classification transformation to these sources. It then serves information about the podcast demographics. This example is depicted in Figure 12-9. Figure 12-9. Example of a data product with input from other data products In this case, the input port consumes data from another data product’s output ports. The source data product can belong to either the same domain or other domains. Nevertheless, the implementation of input data ports and how they consume data from other data product’s output ports is standardized. A data product utilizes its upstream output’s actual time and processing time to select its desired data. In the case of a streaming input, the input port mechanism will keep track of the processing time in the source data to process new data as it arrives. Self as a data source In some cases, a data product’s local computation can be the source of data. For example, a local transformation such as a machine learning model inference generates new data. Consider the artists data product as shown in Figure 12-10. Its transformation runs a machine learning model that adds new data, such as “emerging” or “declining” data for artist classification, to the information it receives from the artist onboarding microservices. Additionally, the data product uses locally stored data such as a list of possible classifications as its input. Figure 12-10. Example of a data product with local input Locality of Data Consumption Data mesh leaves many of the implementation decisions to the underlying technology. Data mesh architecture is agnostic of the underlying technology and infrastructure—as much as possible. For example, where the data is physically located and whether a consuming data product physically copies data from one location to another are implementation details built by the platform. However, data mesh architecture components such as input and output ports are great interfaces to abstract the internal implementation of data or processing movement across physical boundaries. A simple set of APIs that connect a data product’s input port to another’s output port can hide the physical movement of data from one physical storage and trust boundary to another. Similarly, executing the data consumption request can execute a remote query issued from one computational environment to another. The implication of this design is that a mesh can span one or multiple physical infrastructures, multiple clouds, and on-prem hosting environments. This means that when a data product consumes data from its source, in that process, it can physically move data from one underlying hosting environment to another. This seemingly simple capability can have a profound impact on data migration to the cloud and a multicloud data platform. For example, Daff is moving all their analytical data and processing to a cloud environment. Today, the podcast service runs on an on-prem infrastructure, along with its underlying operational database. The podcast listeners data product runs on the cloud. It consumes data through an asynchronous input data port interface. Once a new podcast listener is registered, the podcast listeners data product essentially copies that information from an on-prem stream to its cloud storage and then makes it available through its cloud-based output data interfaces. This implements a continuous data migration through data products from one environment to another, without a big bang data migration strategy. The same mechanism can move data from one cloud provider to another. For example, if a company has a multicloud strategy and they want to keep data across multiple cloud providers, implementation of the input port can move data from the source’s cloud provider to the consumer’s. Of course, for this to become a reality, the underlying platform infrastructure must support a few key features such as internet-compatible addressability of data products, identity authentication and authorization across trust boundaries, and internet-accessible endpoints for output data ports. Figure 12-11 demonstrates data consumption between two different data products across two environments. Figure 12-11. Multi-environment data consumption model Data Consumption Design Data products intentionally specify what data, from which sources, and how they consume data. The data product specifically defines and controls the capability to consume data. This is contrary to other data architecture where sources dump data on the processors without knowledge of where and how the data gets to them. For example, in previous architectures an external directed acyclic graph (DAG) specification defines how the processors get connected to each other centrally, rather than each processor defining how it receives data and serves it. Figure 12-12 demonstrates the data product’s high-level design components to consume data through input data ports. Figure 12-12. Data product’s high-level components to design consume data Table 12-2 summarizes a data product’s high-level components to design consuming data. Table 12-2. Summary of data product’s high-level components to design consume data Consume data component Description Input data port(s) The mechanism by which a data product receives its source data and makes it available for further internal transformation. Input data port(s) specification A declarative specification of input data ports that configures from where and how data is consumed. Asynchronous input data port Asynchronous input port implementation reactively calls the transformation code when the necessary source data becomes available. Subscription to event streams or asynchronous I/O read on files are examples of async input ports. An asynchronous input port keeps track of the source data product’s processing time pointer and when it receives data for a new processing time it reactively executes the transformation. Synchronous input data port Synchronous input ports pull data from the source and call the transformation code when the data is fetched.
For example, “Daily podcasts summary” pulls data from podcast listeners synchronously and calculates the count and other summaries when the data is fetched. It pulls data daily at midnight. Remote query An input port specification can include the query executed on the source to receive the desired data. This capability reduces the amount of data that is otherwise fetched redundantly. The query is expressed in a query language that the source understands, such as SQL, GraphQL, Flux, etc., and is defined by the input port specification. Input port synchronizer and temporary storage Input ports often interdependently consume data. For example, the artist classification transformation code can’t start running until data from two independent sources of artists and listeners becomes available. Temporary storage is necessary to keep track of observed and unprocessed data, until all observations needed become available for processing. Transform Data Almost all data products perform a transformation, no matter how minimal. We create a data product because we see value in sharing a new analytical model of existing data. Creating this new model requires transformation. Creating and maintaining the transformation code is what a data product developer pays most attention to. Traditionally this transformation has occurred in the data pipelines moving data from an input source to an output sink. In data mesh design, the transformation—whether implemented as a pipeline or not—is encoded in and abstracted by the data product. The transformation is an internal implementation of a data product and is controlled by it. Since it’s an internal concern, I don’t intend to be specific in how it must be designed. In my opinion, it’s up to the taste, capabilities, and needs of the data product developer to choose how to implement the transformation. It is helpful to look at a few different ways of implementing the data product’s transformation. Programmatic Versus Nonprogrammatic Transformation Data processing and transformation falls into two main camps: nonprogrammatic transformations (e.g., SQL, Flux, GraphQL) and programmatic data processing (e.g., Apache Beam, Apache Spark, Metaflow). Nonprogrammatic approaches either use relational algebra performing set operations such as SQL, or flow-based functions such as Flux. Either way, the idea is that we can capture the intent of how data is transformed from one set to another with a statement. It is simple and approachable for many data product developers, but also limited to the features of the statement. For more complex transformations the statements become difficult to understand, hard to modularize, and challenging to test automatically. In practice we should not find many data products that simply perform nonprogrammatic transformations. Any other downstream data product can run the same remote queries themselves, and there is no need for intermediary data products. Figure 12-13 demonstrates an example of a nonprogrammatic transformation. The intention is to create demographic information for top listeners. The top listeners data product uses its input ports on played songs and listener profiles to capture the profile information of listeners who have listened to songs today. Then it produces various statistics on demographics of listeners, e.g., age groups who have listened most or least today, countries that have listeners most or least today, etc. Figure 12-13. Example of a nonprogrammatic transformation On the other hand, programmatic data processing uses code logic, conditions, and statements to transform data. Programmatic data processing libraries such as Apache Beam or Apache Spark can be used in different programming languages such as Python, Java, etc. The transformation code has access to the full power of the hosting programming language, imperative or declarative. It can be modularized and tested. This approach is more complex for noncoding data product developers but more extensible. The advantage of this approach is that the statistics can be calculated incrementally as more records about played songs arrive. Data mesh doesn’t take a position on what approach a data product adopts for its transformation and relies on common sense: use programmatic approaches for more complex use cases, and use a nonprogrammatic approach when the transformation is trivial and simple. Or even better, if the transformation is trivial and nonprogrammatic, do nothing and don’t create an intermediate data product. Just leave the end consumers to run the queries themselves. Note that, even in the case of programmatic transformations, the input port may evoke a nonprogrammatic query on the source before processing it. This reduces the amount of data moved to the transformation code and pushes the processing to where upstream data resides. Dataflow-Based Transformation The dataflow programming paradigm, introduced in the 1960s, defines a computer program as a directed graph of the data flowing between operations. This programming paradigm has inspired many modern data pipeline designs. A data pipeline is a series of transformation steps (functions) executed as the data flows from one step to another. Data mesh refrains from using pipelines as a top-level architectural paradigm and in between data products. The challenge with pipelines as currently used is that they don’t create clear interfaces, contracts, and abstractions that can be maintained easily as the pipeline complexity grows. Due to lack of abstractions, single failure in the pipeline causes cascading failures. The pipeline—or dataflow-based programming model in general—within the boundary of a data product is a natural paradigm for implementing the transformation code. In this scenario, pipelines tend to get less complex, since they are bounded by the context and transformation of a single data product. They also get upgraded, tested, and deployed as a single unit with a data product. Hence, pipeline stage tight coupling is less of a concern. In short, using data pipelines within the boundary of a data product for transformation is OK, as long as the pipeline stages don’t extend beyond the boundary of a data product. No transformation occurs between data products, beyond providing and consuming data through read-only output and input data ports. I call this the principle of dumb pipes and smart filters. Figure 12-14 shows the high-level components involved in a pipeline transformation. Figure 12-14. Pipeline transformation ML as Transformation The third transformation category is model based; deploying and running a machine learning or statistical model as a data product’s transformation. For example, imagine Daff using a TensorFlow recommender to recommend songs to extend the listener’s existing playlists. The TensorFlow model can be serialized and deployed as the playlist recommender data product. The data product consumes the listeners’ playlists, makes predictions on what songs are recommended next, and stores those as playlist extension recommendations that are read and used when the listeners play their lists. The recommender model is executed within a program and a desired programming language, but its computation is mainly performed by the model. ML models can be deployed in many different contexts such as microservices and applications, as well as data products. Time-Variant Transformation The common characteristic of all transformations is respecting the axes of time—processing time and actual time. The transformation code is aware of the time parameters when it processes the input and when it generates output with the respective times. The input port mechanism keeps track of the processing time of each source. The transformation code generates the actual time of its output based on the calculation performed on the source data. The output port serves the transformed data accompanied with the data product’s internal count of processing time. Transformation Design The design of the transformation—build-time definition, deployment, and execution—is heavily dependent on the framework of choice and the underlying technology. Whether the transformation is implemented by a declarative statement or programming code, platform capabilities are required to build, test, deploy, and run the transformation. Figure 12-15 shows a few high-level design elements involved in a data product’s transformation. Figure 12-15. High-level components to design transformation for data products Table 12-3 summarizes the high-level components involved in designing the data product transformation affordance. Table 12-3. Summary of high-level components to transform data for a data product Transform data component Description Transformation artifact(s) The code, configuration, statement, or model that defines the transformation. This artifact is executed on the input data to produce the output, for a particular record time. Transformation runtime environment Transformations are invoked based on the configuration of the transformation, for example on a regular time basis or availability of the necessary input data. Once invoked they require a computational environment to execute, bounded by the data product container. The underlying platform provides this environment. Temporary storage Steps of a transformation code may require access to temporary storage that persists state across different phases of transformation. This is provided by the underlying platform. Recap In this chapter I presented a few core design elements and characteristics of data products as an architecture quantum, to autonomously consume, transform, and serve data. To serve data in a distributed mesh architecture, I identified a few properties for the data products’ output ports: serve bitemporal, immutable data via read-only and multimodal APIs. These constraints were driven by the objectives and assumptions of data mesh: Data products serve their domain-oriented data natively to a wide set of data users (multimodal access). Data products can be used in temporal analytics use cases (bitemporal). Data products can safely consume data from other data products, as well as transform and serve it, while maintaining a global state of eventual consistency at a point in time (bitemporal). Data product users process data for analytics and ML (read-only). Data products have control over where and how they receive their upstream data. They use their input ports to get data from collaborating operational systems in the same domain or upstream data products. The input ports can implement data movement between different hosting environments in support of a multicloud deployment or an incremental cloud migration. Data products almost always perform some kind of transformation in order to serve higher-value data. The transformation implementation can have multiple forms, including data pipeline functions, a sophisticated query, or ML model inference. This chapter sets a high-level frame of thinking and approach to the design of data products in their core function of data sharing. I expect that future data mesh implementations evaluate and refine this.

データ製品の主な役割は、入力データポートを使用して上流ソースからデータを取り込み、それを変換し、出力データポートを介して恒久的にアクセス可能なデータとして提供することです
この章では、すべてのデータ製品が実装する3つの基本機能であるデータの取り込み（「データの取り込み」）、データの変換（「データの変換」）、データの提供（「データの提供」）の設計特性について説明します
まず、データメッシュのアプローチに最も固有の特性を持つ機能から始めましょう
データの提供 データ製品は、さまざまな分析的な消費者に対して、ドメイン指向のデータを提供します
これは、第9章で紹介された出力データポート（インターフェース）を通じて行われます
出力データポートには明示的に定義された契約とAPIがあります
これは「ドメイン指向のデータの提供」というシンプルな機能ですが、そのエコシステム内のエージェントとの関係や彼らの能力とニーズとの関係を考慮すると、興味深い特性を持ちます
データ製品とデータユーザーの関係を見てみましょう
データユーザーのニーズ 図12-1は、データユーザーのニーズとデータ製品がそれを満たすための特性を示しています
図12-1
データユーザーとそのニーズを満たすためのデータ提供の特性 データユーザーによって要求される要件は、データ製品がデータを提供する方法に一連の設計上の考慮事項を課します
分析的なデータユーザーは多様なプロフィールを持っています データユーザー（データにアクセスして読み取るクライアント）は、データアナリスト、データサイエンティスト、データ駆動型アプリケーション開発者などの人物、およびレポート、可視化、統計モデル、機械学習モデルなどのシステムを含む、幅広いスペクトラムのペルソナとアプリケーションタイプです
第3章を思い出すと、データ製品は、そのデータを異なる形式とアクセス方法で提供することで、これらの多様なペルソナにとってネイティブにアクセス可能な方法でデータを提供します
これを基本的な使いやすさの特性と呼びました
この要件の設計上の含意は、データに複数のモードでアクセスすることでデータを提供することです
分析的なデータユーザーは、長期間のデータが必要です メッシュは、分析用途においてデータのグローバルな状態を完全に長期間保持し、それもメッシュ外のデータレイク、データウェアハウス、または他の外部システムなしで保持します
データの連続的に変化するグローバルな状態は、データ製品の接続されたグラフによって保存および維持される唯一の要素です
これがアーキテクチャの分散化の意味です
洞察は、時間の経過を考慮に入れることで最も強力になります
連続的なデータ変更を時間の経過にわたってアクセスできることにより、トレンドを形成し、予測を行い、複数のドメインにまたがるさまざまなイベント間の相関関係を発見できます
データメッシュは、データのプレゼンテーションおよびクエリングの両方で常に存在するパラメータとして時間を前提としています
時系列データ、つまりイベントと状態の変化を表すデータにアクセスするための設計上の含意は、各データ量子が二時的なデータを提供することです
分析的なデータユーザーは、特定の時点で複数のドメインの一貫したビューが必要です 大部分の分析的なユースケースでは、複数のデータ製品からデータを処理します
このようなユースケースでは、一貫した時点で複数のデータ製品を相関させます
例えば、ダフが2021年7月1日に次の月の加入者増加を予測するために機械学習モデルをトレーニングする場合、それは2021年7月1日に複数のデータ製品によって既知のデータと処理されたデータに基づいて行われます
この成長モデルの特定バージョンの再現性をサポートするためには、2021年7月21日に処理されたデータ製品の複数にわたるデータの変更不能な状態を維持する必要があります
複数のデータ製品にまたがる時点毎の一貫したデータの提供と、再現性のためのデータのバージョン管理を組み合わせることで、データの提供には複数の設計上の考慮事項が導入されます: 二時変性、変性しなさ、読み取り専用アクセス
 データの提供設計特性 それでは、先ほど発見した各特性について少し深く掘り下げてみましょう: マルチモーダル、変更不可能、二時的、読み取り専用アクセスという特性です
これらの特徴は、データメッシュの動作に不可欠です
マルチモーダルデータ データ製品の役割は、特定のドメインの分析データを特定かつ一意のドメインセマンティクスで提供することです
しかし、データ製品が多様なコンシューマーにネイティブにサービスを提供するためには、同じドメインセマンティクスを異なる構文で共有する必要があります
セマンティクスはカラム形式のファイル、関係データベースのテーブル、イベントなどの形式で提供することができます
同じセマンティクスは、データコンシューマーの体験を損なうことなく提供されます
レポートを作成する人は、データを関係テーブルとして消費します
機械学習モデルをトレーニングする人は、カラム形式のファイルでデータを消費し、リアルタイムのアプリ開発者はイベントを消費します
分析データの本質を空間および時間の次元として視覚化すると、個人これらの特性は、データメッシュの動作に必要不可欠です
 マルチモーダルデータ データ製品の役割は、特定のドメインの分析データを特定のユニークなドメインセマンティクスで提供することです
 ただし、データ製品が多様な消費者にネイティブに提供するためには、異なる構文で同じドメインセマンティクスを共有する必要があります
 セマンティクスは、列指向ファイル、関連データベーステーブル、イベントなどの形式で提供することができます
 同じセマンティクスがデータ消費者の体験を損なうことなく提供されます
 報告書を書く人は関連テーブルとしてデータを消費します
 マシンラーニングモデルをトレーニングする人は列形式のファイルでデータを消費し、リアルタイムのアプリ開発者はイベントを消費します
 私は、分析データの性質を空間と時間の寸法として視覚化することが役立つと考えています
 空間の次元を使用してデータの異なる構文的具現化、データの形式を表します
 任意のデータ製品は、複数の形式でデータを表示することができます
または、マルチモーダル形式で、例えば： 半構造化ファイル、例：列ファイル エンティティ関係、例：関連テーブル グラフ、例：プロパティグラフ イベント データユーザーがデータ製品にアクセスする際には、最上位のデータ製品発見API（「発見、理解、信頼、探索」）を使用します
最初にデータ製品のセマンティクスについて学びます
どのドメイン情報を提供しているか（例：ポッドキャスト、ポッドキャストのリスナーなど）その後、データの特定のアクセス方法に対してデータ製品の出力API（「出力データポート」）の1つにアクセスします
 データの提供方法は、基礎（物理）テクノロジーによるものです
 イベントログに登録する、分散列ファイルを読み取る、関連テーブル上のSQLクエリを実行するなど
 ドメイン指向のセマンティクスは最上位の関心事であり、形式とアクセスのモードは二次的な関心事です
 これは、既存のアーキテクチャとは逆のモデルであり、ストレージおよびエンコーディング技術がデータが最初にどのように組織化され、それから提供されるかを指示するものです
 図12-2は、多様なアクセスモードを例として示しています
 プレイイベントデータ製品は、3つのアクセスモードを介してプレイイベントデータにアクセスできます： 「プレイイベント」トピックにサブスクライブする（リスナーの状態変化をキャプチャし、ログイン、ポッドキャストの再生、再生停止などを行う）、同じプレイイベントにSQLクエリを介してアクセスする（イベントの属性の行を持つテーブルを使用）、および列オブジェクトファイル（すべてのイベントの属性ごとにファイルを作成）
「図12-2 データ製品のマルチモーダルアクセスの例」 これらの3つのアクセスモードとデータのトポロジーは、さまざまな役割の消費者のニーズを満たすことができます： データ集約的なアプリ開発者はエラーイベントを監視してプレーヤーの品質を向上させるため、同期SQLアクセスモデルを使用してデイリーレポートを作成するデータアナリスト、およびファイルアクセスモードを使用してMLモデルをトレーニングするデータ科学者はプレイパターンの分類を見つけます
 これは、個々の出力データポートを介して可能になります
 注意   現在、マルチモーダルアクセスのサポートの複雑さは、データの上位のアピールと抽象APIの欠如の症状です
これは、データの表現や形式に対して意味論的かつ前向きに問い合わせるためのものです
 不変データ 同じ川を2回渡ることはできない
-ヘラクレイトス 不変データは、一度作成されると変更されません
 データ製品は、データを処理してデータユーザーに利用可能にする際に、その特定のデータが変更されないようにデータを提供します
削除や更新はできません
 データの変更はしばしば複雑さとエラーを引き起こすため、経験豊富なプログラマーなら誰でも知っていることです
 これが関連性のある分析ユースケースに特に関連します
不変のデータを使用することで、データユーザーは繰り返し可能な方法で分析を再実行することができます
モデルの再トレーニングや特定の時点のデータセットでのレポートの再生成は、同じ結果を繰り返し得るものです
再現性は必要です
なぜなら、結果が驚くべき観察結果を生み出すことが多く、アナリストがより深く掘り下げる必要があるからです
使用しているデータが変更されると、驚くべき結果を再現することができなくなり、これがデータの変更によるものなのか、プログラミングのエラーなのかを疑問に思うことがあります
分析にバグがある場合、不安定なデータソースで作業する際には、同じコードを繰り返し実行しても同じ回答を得ることができないため、追跡がはるかに難しくなります
データメッシュでは、ソースデータが複数のデータ製品で使用され、複数のデータ製品が特定の分析のソースとなることがあり、複数のデータ製品が1つのビジネス状態のより大きな理解に貢献する真実の一部を保持しているため、可変データの混乱と複雑さはさらに悪化します
データメッシュの分散した性質は、データユーザーに対して次の自信を与えるために、変更不可能性が要求されます：（1）時間刻みのデータの複数のデータ製品間に一貫性があり、（2）ある時点でデータを読み取った後、そのデータが変更されず、信頼性のある読み取りと処理を繰り返すことができます
図12-3は、シンプルな例を示しています：リスナーの人口統計データ製品は、リスナーの地理的な人口統計を日々提供し、リスナーが接続する場所を提供します
これには、2つの下流データ製品があります
アーティストの地域人気（特定のアーティストのリスナーが最も多い場所）と地域の市場規模（異なる地域のリスナー数）
これら2つのデータ製品は、地域マーケティングデータ製品のソースの一部であり、地域に対してターゲットを絞ったマーケティング活動を推奨します
このデータ製品は、A/Bテストや実験を行うために市場リスクの少ない国を特定するために地域の市場規模を使用し、人気に基づいてアーティストをプロモートするためにアーティストの地域人気を使用します
図12-3
相関しないデータの例これらのデータ製品のそれぞれは、データの処理と提供に独自のリズムを持っています
したがって、すべてのデータが同時に更新されることを保証する方法はありません
これによって致命的なダイヤモンドが生じる可能性があります
リスナーの人口統計がデータを更新する時に、地域マーケティングが分析を実行している場合、リスナーの人口統計は更新前にデータを提供し、地域の市場規模は更新後に提供するため、一貫性のないデータを消費する可能性があります
さらに悪いことに、地域マーケティングはこの不一致についてさえ知りません
データメッシュは、データユーザーが更新されたデータを検出できないようにすることで、この問題に対処します
データの変更は常に、異なるデータ製品から来る異なるデータの断片を関連付けることなく、時系列の属性を識別するための新しいデータの断片として提示されます
これを行う最も一般的な方法は、次のセクションで説明する2つの時間軸を持つデータを使用することです
たとえば、リスナーの人口統計は、{リスナーID：'123'、場所：'サンフランシスコ'、接続時間：'2021-08-12T14:23:00、処理時間：'2021-08-13T01:00'}のようなタプルとしてデータを共有します
各情報には、2つの時間変動の識別フィールド、「接続時間」（音楽を聴くためにリスナーが接続した時刻）と「処理時間」（リスナーの人口統計がこの情報を処理した時刻）があります
このタプルが処理され、データユーザーに利用可能になると、絶対に変更されません
もちろん、同じリスナーが翌日には別の場所からコンテンツを聴くことがあり、それは新しいデータエンティティの追加として現れます（リスナーID：'123'、場所：'ニューヨーク'、接続時間：'2021-08-13T10:00'、処理時間：'2021-08-14T01:00'）
これらは2つの異なる時間であり、更新を伴っていますが、リスナーの新しい場所を示しています
データの不変性は、付随的な複雑さの機会を減らし、分散メッシュ上の共有ステートの更新とそれに伴う副作用の取り扱いの複雑さ、および分散トランザクションという扱いにくいコンピュータサイエンスの問題を解決する機会を減らします
メッシュがすでに消費されたデータが変化し続けることを許容する場合、または同じ読み取りを繰り返すことで異なる結果が得られる場合、メッシュ上の各下流データリーダのコードにどれほどの複雑さが組み込まれているかを想像してみてください
不変性は変更の設計とメッシュの最終的な整合性を維持するための新しい処理時間とその不変な状態の伝播を通じて、変更の設計を可能にするもう一つの重要な要素です
データを不変に保つことはいつでも重要ですが、データメッシュにとっては特に重要です
ただし、新しい情報が利用可能になったり、データ処理のバグが修正されるにつれて、過去のデータも変更されることがあります
したがって、データに対して取り消し可能な変更を行うことができる必要があります
データメッシュが不変性と取り消された変更を実装し、変更可能である必要があるにもかかわらず、データが不変であるようにする方法を見るために、バイテンポラリティを見てみましょう
バイテンポラリティデータ変化と時間は切り離せないようです
変化には時間がかかり、時間的に配置および順序付けられ、時間的に区切られています
時間と変化の切り離せなさは、一種の論理的な真実です
--レイモンド・タリスバイテンポラリティデータは、実際のイベントが発生した時間または実際の状態が存在した時間"実際の時間"と、データが処理された時間"処理時間"の2つのタイムスタンプを記録するようにデータをモデリングする方法です
バイテンポラリティデータモデリングにより、不変なエンティティとしてデータを提供することができます
つまり、データの不変性により、偶発的な複雑さが減少し、分散メッシュ内の共有状態の更新の副作用や分散トランザクションの解決における複雑さが軽減されます
メッシュがすでに消費したデータが変更され続けることを許可する場合、または同じ読み取りを繰り返すと異なる結果が得られる場合、メッシュ上の各下流データリーダーのコードに関連する複雑さを想像してみてください
不変性は変更の設計とメッシュの最終的な一貫性の維持を可能にするもう一つの重要な要因です
新しい処理時間とその不変の状態を伝播することによって
データを不変に保つことはいつでも重要ですが、データメッシュでは特に重要です
ただし、新しい情報が利用可能になるか、データ処理のバグが修正されることで、過去のデータが変更されることも事実です
したがって、データに撤回可能な変更を行う必要があります
データメッシュが不変性と撤回可能な変更を実装し、データが不変である必要があるにもかかわらず変更を許可する方法、つまり、バイテンポラリティを見てみましょう
 バイテンポラリティ データ 変更と時間は切り離せない関係にあるようです：変更には時間がかかり、時間的に位置づけられます
そして、時間によって区切られます
時間と変化の切り離せなさは、ある種の論理的な真実です
 バイテンポラリティデータは、データの各部分が2つのタイムスタンプを記録する方法です
つまり、「イベントが実際に発生した時点または状態が実際にあった時点」である実際の時間と、「データが処理された時点」である処理時間です
バイテンポラリティデータモデリングにより、データを不変なエンティティとして提供することが可能になります
つまり、{データプロダクトのフィールド、実際の時間、処理時間}のタプルとして、新しいデータエンティティが処理されるたびに処理時間は単調に増加します
また、時系列分析とタイムトラベル(過去の傾向を見て将来の可能性を予測すること)も行うことができます
これらの結果は、データメッシュにとって重要です
たとえば、Daffのビジネスの成長を予測する典型的な時系列分析のユースケースを考えてみましょう
このモデルは、長期間にわたる加入者の変更を使用してパターンや傾向を発見します
 データプロダクトは、異なる処理速度でデータを処理し提供するという特徴があります
成長予測モデルは、異なるデータプロダクトから一貫したデータを使用して、繰り返し可能なモデルのトレーニングにおいて共通の処理時間を選択します
つまり、データプロダクトがイベントを処理し認識するタイミングです
例えば、2022-01-02T12:00にトレーニングされたモデルでは、2022-01-02T12:00までの過去三年間（実際の時間）の加入者情報、サポート問題、マーケティングイベントを使用します
実際の時間と処理時間は、データプロダクトが保持し提供する時間の2つの絡まった軸です
 実際の時間の変動 分析用途のニーズを満たす形式でデータを表現するため、データプロダクトは時間の経過に伴うドメインの状態（またはイベント）をキャプチャし共有します
たとえば、ポッドキャストのリスナーデータプロダクトは、「1年前から現在までの毎日のポッドキャストのリスナー情報」を共有することができます
実際の時間は、イベントが発生する時点または特定の状態が存在する時間です
たとえば、2021-07-15は、この日にDaffのポッドキャストを実際に聴いたポッドキャストリスナー（データ）の実際の時間です
実際の時間の存在は重要です
なぜなら、予測分析や診断分析は、実際に何が起こったかの時間に敏感だからです
これはオペレーションの機能の場合には当てはまりません
オペレーションの機能のほとんどは、データの現在の状態に対処します
例えば、「現在のポッドキャストリスナーの住所を教えてください
印刷されたマーケティング資料を送るために必要です」といった具体的な要求です
実際の時刻は変動します
データ製品は、順序外の実際の時刻を観測したり、ソースデータの修正後に同じ実際の時刻に関する新しいデータを受け取ることができます
処理時間の連続処理時間は、データ製品が特定の実際の時刻についてその知識または理解を観測し、処理し、記録し、提供する時間です
例えば、2021-08-12T01:00に、ポッドキャストのリスナーデータ製品は、2021-08-11にポッドキャストを聞いた人々に関するすべてのデータを処理し、当時の世界の状態を理解します
2021-08-12T01:00が処理時間です
データの必須属性としての処理時間の提供は、変化に対応するための設計の鍵です
過去の理解は変化します
過去に発生したエラーを修正したり、過去の理解を改善する新しい情報に気付いたりすることがあります
過去を変えることはできませんが、過去の処理時間を変えることはできます
過去の実際の時刻の修正を反映した新しい処理時間で新しいデータを提供します
処理時間は、前進的な時間として信頼できる唯一の時間です
注記処理時間は、4つの異なる時間を1つの時間に統合するために使用します
観測時間は、データ製品がイベントや状態に気付く時の時間です
処理時間は、データ製品がデータを処理し、変換する時間です
記録時間は、データ製品が処理されたデータを保存する時間です
公開時間は、データがデータユーザーにアクセス可能になる時間です
これらの微妙な時差は、データ製品の内部に最も関連があり、データユーザーに関係がありません
したがって、それらをデータユーザーにとって最も重要な処理時間にまとめました
マーティン・ファウラーは、「Bitemporal History」というシンプルで優れた記事でこの二重時間性を紹介しています
このセクションでは、データ製品がこのコンセプトを統一モデルで採用する方法をまとめています
過去への影響の影響それでは、いくつかのシナリオで二重時間性の正の影響について簡単に説明しましょう
取り消し世界の理解は継続的に進化しています
世界の理解にはエラーや情報の欠落があるかもしれません
この場合、後の処理時間で誤解を修正します
例えば、2021-08-12T1:00に「2021-08-11にポッドキャストを聴いた人々」に関する情報をすべて処理しますが、リスナーの数を数える際に誤りがあり、3000人と計算します
次に、2021-08-13T10:00に情報を処理する際に、エラーを修正し、新しい数値である2021-08-11の2,005人のリスナーを作成します
3,000と2,005は、同じアイデンティティ「2021-08-11のポッドキャストリスナー数」のために2つの異なる値として捉えられ、2021-08-12T10:00で処理された状態と2021-08-13T10:00で処理された状態として別々の状態としてキャプチャされ、共有されます
二重時間性を使用することで、データの状態、データモデル、SLOメトリックなどの変更が組み込ま実際の時間は変動します
データ製品は、順序が逆になった実際の時間を観測したり、ソースデータに修正が加えられた後でも同じ実際の時間に関する新しいデータを受け取ることがあります
処理時間の連続 処理時間とは、データ製品が特定の実際の時間における状態やイベントの知識または理解を観測し、処理し、記録し、提供する時間です
例えば、2021年08月12日01:00に、ポッドキャストのリスナーデータ製品は2021年08月11日にポッドキャストを聴いた人々に関するすべてのデータを処理し、その当時の世界の状態を理解します
2021年08月12日01:00が処理時間です
データの必須属性として処理時間を提供することが変更に対する設計の鍵です
過去の理解は変わります
過去に発生したエラーを修正するか、過去の理解を向上させる新しい情報を認識することがあります
過去を変えることはできませんが、変えることができるのは現在の処理時間です
過去の実際の時間の修正を反映した新しい処理時間で新しいデータを提供します
処理時間は、単調に進むことが信頼できる唯一の時間です
注：私は処理時間を使用して、観測時間、データ製品がイベントまたは状態を認識した時点、処理時間、データ製品がデータを処理し変換する時点、記録時間、データ製品が処理されたデータを保存する時点、発行時間、データがデータユーザーに利用可能になる時点、の4つの異なる時点を1つにまとめました
この微妙な時差は、データ製品の内部において重要であり、データユーザーには関係しません
したがって、データユーザーに最も重要な処理時間を1つにまとめました
Martin Fowlerはこの「二重時系列」をシンプルかつ素晴らしい記事「二重時系列の歴史」で紹介しています
このセクションでは、データ製品がこのコンセプトを採用し、処理の遅延やデータの形状（イベントまたはスナップショット）に対して不偏な統一モデルで採用する方法をまとめています
二重時系列の影響 ここでは、二重時系列のいくつかのシナリオにおけるポジティブな影響について簡単に説明します
 誤りの修正 世界の理解は常に進化しています
世界の理解には誤りや欠落した情報がある場合があります
この場合、後の処理時間において、私たちは誤解を修正します
例えば、2021年08月12日01:00に「2021年08月11日にポッドキャストを聴いた人々」に関する情報をすべて処理しますが、リスナーの数を誤って3,000人と計算します
次に、2021年08月13日10:00に情報を再処理すると、エラーを修正し、新しいカウントとして2021年08月11日に2,005人のリスナーを算出します
3,000人と2,005人は、同じ識別子「2021年08月11日のポッドキャストリスナーの数」に対して、2021年08月12日10:00に処理された状態と2021年08月13日10:00に処理された状態として記録および共有される2つの異なる値となります
二重時系列を使用することで、データの状態、データモデル、SLOメトリックなどの変更が組み込まれます
変更の連続的な処理は、すべてのコンシューマーとメッシュ全体にデフォルトの振る舞いとして組み込まれます
これにより、データユーザーのロジックが大幅に簡素化されます
過去の更新は特別なケースや驚きではなくなります
それらは過去について新しいストーリーを伝えるために処理される新しいデータです
コンシューマーは、過去のデータの過去のリビジョンにアクセスするために、過去のある時点で処理したデータを追跡し、固定することができます
もちろん、そのようなシステムを構築することは容易なエンジニアリングの課題ではありません
処理時間と実際の時間のずれ ずれとは、実際の時間と処理時間の時間差です
真のリアルタイムシステムの場合、ずれは無視できるほど小さいです
イベントが処理され、世界の理解が形成されるのは、ほぼイベントが発生した時とほぼ同じです
これは、特にデータ製品の連鎖でデータ製品が実際のイベントのソースから複数のホップ先でデータを処理する分析データ処理では非常に珍しいです
二つの時間の存在は、データの利用者にスキューを通知し、それに基づいてデータの処理方法を決定することができるようにします
ソースに合わせたデータ製品から遠く離れるほど、スキューは大きくなります
時間枠データ製品の上流データを一定の時間枠で集計することは一般的です
例えば、プレイセッションはリスナーがプレーヤーデバイスと関わる間に発生するすべてのイベントを集計します
つまり、リスナーがポッドキャストから別のポッドキャストに移動し、最終的に1つを選び、数分間聞いてからプレーヤーをしまうまでの一連の再生イベントです
これは、リスナーがプレーヤーと関わる際の行動分析を支援します
この場合、再生イベントの実際の時間は複数の分（時間枠）にわたる場合があります
時間枠の知識を持つデータの利用者は、データ上で時間に関連する操作を行うことができます
連続する変更の反応的処理データメッシュでは、常に変化し続ける世界を前提としています
この定常的な変化は、新しいデータの到着や過去のデータの理解の進化の形で現れ、処理時間によって示されます
データ製品は、上流データ製品が変更されたことに反応して、連続的に変更を処理することができます
つまり、新しい処理時間が利用可能になったことを意味します
処理時間は、接続されたデータ製品間の反応的で非同期なデータ処理を作成するための基本的なメカニズムとなります
データ製品のすべての属性とプロパティには、時間の概念が組み込まれており、時間の経過とともに変化します
データスキーマ、データセマンティクス、データ間の関係、SLO（Service Level Objective）などのメタ情報も同様です
これらは、自動的に時間によってバージョン管理されます
処理時間は、データスキーマなどのデータ製品の永続的なプロパティのバージョン管理のための基本的な要素となります
処理時間は、データ製品のSLOなどの時間変動情報と相関付けるためのパラメータとなります
例えば、前の例を見てみましょう
図12-4は、単一のデータ製品における時間の数値とその関係を示しています
図12-4ではいくつかの注目すべき点が示されています
スキュー データの量子がイベントを処理し、そのデータの状態についての理解を形成するのは、イベントが実際に発生した時点よりも遅いことは避けられません
例えば、2021年8月11日のデイリーリスナーの状態は、2021年8月12日の1:00よりも後にシステムによって把握されます
ポッドキャストリスナーのデイリー状態を把握するためには、最低でも1時間から25時間のスキューがあります
処理エラー2021年8月12日の1:00に処理された総リスナー数には、計算上のエラーがあり、2021年8月11日のデイリーリスナーの総数を2,005人ではなく3,000人として捉えています
訂正2021年8月12日の1:00に発生したエラーは、データ製品開発者によって発見され、処理コードに修正が加えられ、次の処理間隔である2021年8月13日の1:00に正しい値である2,005が報告されます
したがって、処理時間2021年8月13日の1:00で提供されるデータには、2021年8月11日と2021年8月12日のデイリーポッドキャストリスナーのデータが含まれます
図12-4では、データの提供インターフェースやAPIが2つの時間パラメータ、処理時間と実際の時間を持つ関数として機能していることが示されています
使用の簡便さのために、最新のデータの状態に関心があるデータ利用者のためにlatestのような特別なデフォルトパラメータを使用することができます
この例では、簡単のために低解像度のウォールタイムを使用しています
ガバナンスは、時間の標準化方法を標準化し、プラットフォームがそれを一貫して実装する役割を果たします
処理時間は、順序付けられた読み取りを保証します
これは、内部システム時間、エポック以降の増分カウンタ、または単調に増加する数値として実装することができます
データ消費者は、処理時間をインデックスとして使用して、自分が消費したデータを知ることができます
実時間はISO 8601などのDateTime標準に従うことができます
分析データのデータユーザーは、データにアクセスする際に、時間を旅行してデータを遡ったり、前へ進んだりすることができます
遡れる限界は、最初に処理されたデータまでか、最新のデータまでかは、データ製品の保持ポリシーによります
図12-5は、この時間軸を横切るタイムトラベルの簡単な可視化を示しています
図12-5
処理時間の軸上でデータを提供する 状態、イベント、またはその両方 システムがデータをエンコードおよび処理する方法は、状態とイベントの2つのキャンプに分かれます
状態は、時間のある時点でのシステムの状態を捉えます
例えば、「今日のポッドキャストのリスナー数」です
イベントは、特定の状態の変化の発生を捉えます
例えば、「新しいポッドキャストのリスナーが接続した」といった具体的な変化の発生です
状態または（変化の）イベントによって、データの格納や提供方法は非常に異なるものになりますが、ここで議論した時間の次元とは異なる関心事だと考えています
ストリーム上の変更イベントをキャプチャしたり提供したりするか、またはシステムの推論された状態のスナップショットのストリームをキャプチャしたり提供したりするかは、データ製品のロジックによりますが、おそらくデータの読み取り手は両方を見たいと考えるでしょう
それにもかかわらず、2つの時間軸の存在は残ります
 取り消された変更の機会を減らす 既に述べたように、過去のデータ（取り消し）は、修正や新たなデータの到着と同様に処理されます
修正されたデータは、過去に属する実際の時間を持つ新しいデータエンティティとして提示されますが、新しい処理時間を持ちます
また、次のセクション「読み取り専用アクセス」では、データ更新の特別なハンドリング（たとえば、メッシュレベルでの削除権の行使）を紹介します
取り消しを処理するための手法に関わらず、データ製品はエラーを減らし、修正の必要性を減らすことを目指す必要があります
以下に、修正の必要性を減らすためのいくつかの戦略を紹介します
 品質管理のために処理時間間隔を増やす プレイイベントのデータ製品を想像してください
これは、プレイヤーのデバイスからのシグナルをキャプチャします
データ製品は、時々イベントを見逃したり遅れて受信したりすることがあります
これは、ネットワークの中断、デバイスキャッシュへのアクセス不可などが原因です
しかし、データ製品の変換コードは、欠落したシグナルを予測または生成したもので補正するための処理遅延を導入することができます
または、同じデータの中央値や他の統計的表現にシグナルを集約することもできます
これは、データ製品がエラーを修正し、取り消しを避けるために、実際の時間と処理時間の間により長いズレを導入していることを意味します
これは、和解が必要なビジネスプロセスに対してしばしば適した技術です
例えば、支払いトランザクションをほぼリアルタイムで受け取る一方で、修正されたおよび和解済みの支払い口座データのデイリーダンプが後で提供される場合があります
これには、処理間隔の増加が必要です
支払いの場合、トランザクションの迅速さよりも口座の正確性が優先されます
 期待されるエラーを反映したデータ製品のSLO（サービスレベル目標）を調整する 前述の例を続けると、一部のコンシューマーは、ほぼリアルタイムのデータにおけるエラーに完全に寛容です
例えば、プレイヤーエラーを検出するアプリは、ここにあることとかなり関係なく、欠落したイベントを気にしません
この場合、プレイイベントのデータ製品は、「プレイヤーエラー検出」といったカテゴリのコンシューマー向けに和解なしでデータを公開することができます
ただし、この場合、データ製品は、欠落するシグナルに関する予想されるエラー範囲に応じてその品質目標を伝えます
 読み取り専用アクセス データメッシュは、神話的な唯一の真実のソースの概念を受け入れない、他の分析データ管理パラダイムとは異なります
各データ製品は、その能力の範囲内で特定のドメインの現実の真実の一部を提供します
データは1つのデータ製品から読み取られ、変形および変換され、次のデータ製品によって提供されることがあります
メッシュは、上流のデータ製品から下流の消費者に対してバイテンポラルな不変のデータを付加することによって、変化の伝播を通じてその完全性を維持します
したがって、新しいデータがグラフを通じて伝搬する際には、メッシュは最終的な整合性の状態を維持します
これまでにも、メッシュまたは個々のデータ製品が直接的な更新機能を持っていないことは明らかになっているかもしれません
更新は、データ製品が入力を処理する結果として間接的に行われます
データ製品への変更は、新しく処理されたデータの追加の形でのみ、データ製品の変換コードによって行われます
これにより、データの不変性が保証され、メッシュの最終的な整合性が維持されます
ただし、GDPRなどの規制に従って削除の権利を行使するためのグローバルガバナンス管理機能などの場合には、データが変更されます
これは、メッシュエクスペリエンスプレーンによってトリガーされる特定の管理機能として考えることができます
すべてのデータ製品の制御ポート（「コントロールポート」）上でコマンドを実行し、この場合は暗号化の破棄を行います
データ製品は常にユーザー情報を暗号化してエンコードします
暗号化キー（プラットフォーム上に存在し、データ製品内には存在しない）を破壊することで、ユーザー情報は実質的に読めなくなります
すでに処理されたデータを更新するために特別な操作が必要な特殊なケースがある場合は、この操作はグローバルコントロールポート機能として実装されます
アウトプットポートの機能ではありません
アウトプットポートは、データユーザーがデータを読むだけに使用します
データの提供デザイン それでは、データ製品のデータ提供のデザインをまとめてみましょう
図12-6は、おそらくの論理アーキテクチャを示しており、データ製品が自らのドメインのコア表現を所有し維持し、それを複数の空間モダリティで提供するために、アウトプットデータポートアダプタの概念を使用しています
各ポートは常にバイテンポラルで不変で読み取り専用のモードでデータを提供します
データの保持期間は、データ製品のポリシーによって異なります
多くの観測（処理時間）のためにデータを保持しアクセス可能にする場合、最新のデータのみを保持する場合、その中間のどこかにする場合があります
表12-1は、データを提供するために関与するデータ製品の構成要素をまとめたものです
表12-1. データを提供するためのデータ製品の高レベル構成要素 データの提供構成要素 説明 出力データポート 特定のアクセスモード（構文）に従ってデータを提供するためのインターフェース（API）
この実装は、単に特定の物理技術（例：データウェアハウスのバイテンポラルテーブル、データレイクのファイル、イベントログのトピックなど）へのアクセスをリダイレクトするAPIである場合もあります
 出力（データ）ポートアダプタ 特定の出力ポートでデータを表示するためのコード
この実装は、データ製品の変換コード内の単純なステップでデータを特定の構文で格納する場合から、読み取り時にコアデータを多くのアクセスモードに適応するためのランタイムゲートウェイとしての高度なものまで、さまざまです
 コアデータの意味 データの意味表現-アクセスモードまたは空間構文に依存しない形での意味表現
 データの消費 多くの場合、組織内または外部の運用システムからのデータは、人々やデバイスなどの他の運用エージェントとの相互作用によって生成されます
場合によっては、データは購入または無料でダウンロードしたアーカイブとして受け取られます
例えば、Daffの運用データは、リスナーが異なるコンテンツに購読し、聴取し、コンテンツプロバイダが音楽を公開し、アーティストマネージメントチームがアーティストの業務を支払い、処理することによって生成されます
データ製品は、このデータを消費し、分析用途に適した方法で変換して提供します
したがって、多くの場合、データ製品は1つまたは複数のソースからデータを消費します
アーキテクチャ的には、入力データポート（「入力データポート」）は、データ製品がソースデータを定義し実行するために必要なメカニズムを実装します
入力データポートは、データ製品がデータソースに接続し、クエリを実行し、データ（イベントまたはスナップショット）を連続したストリームまたはワンオフのペイロードとして受信するための論理的なアーキテクチャ上の構造です
基礎となる技術の選択は、実装固有の問題であり、データプラットフォームに委ねられています
図12-7は、データ製品がデータを消費する方法の高レベルなイメージを示しています
データ製品は、1つまたは複数のソースからデータを消費します
ソースは、共同の運用システムまたは他のデータ製品である可能性があります
消費されたデータは、次のようにしてコアデータモデルに変換され、出力ポートを介して複数の形式で提供されます
図12-7
データ製品がデータを消費するための設計 各データ製品には1つまたは複数のソースがあります
ソースに合わせられたデータ製品（「ソースに合わせられたドメインデータ」）は、大部分のデータを運用システムから消費します
集約データ製品（「集約ドメインデータ」）になると、他のデータ製品がソースになり、消費者に合わせたデータ製品（「消費者に合わせたドメインデータ」）では、特定の使用ケースのためにローカルにソース化されたデータを提供するためにスマートロジックや機械学習モデルがしばしば含まれています
データ製品の入力データの設計に影響を与えるいくつかの注目すべき特性について詳しく見てみましょう
 データソースの原型 入力機能の設計は、複数のデータソースの原型をサポートする必要があります
以下はいくつかの上位カテゴリです： ・共同の運用システム ・他のデータ製品 ・自己 共同の運用システムとしてのデータソース ソースに合わせられたデータ製品は、ドメインの業務を実行する際にデータを生成する1つまたは複数のアプリケーションからデータを消費します
彼らは運用システムに最適化されたデータを消費し、分析目的に適した形式に変換します
彼らは、下流の分析ユースケースを運用アプリケーションの内部の詳細から切り離します
ここでの共同作業というフレーズは、データ製品とそのソース運用システムの密な結合を意味します
両者は同じドメインに属している必要があります
ソースアプリケーションとデータ製品の間の運用データ契約は、しばしば密接に結びついており、運用システムのデータとモデルの変更がデータ製品のデータとモデルに影響を与える場合があります
そのため、共同の運用システムソースとそれらの共同作業データ製品を同じドメインに保つことを強くお勧めします
これにより、運用システムの変更を担当するドメインチームとデータ製品開発者が密接に協力して両者を同期させることができます
CRM COTS（商用の顧客関係管理）製品など、1つのドメインに統合された複数のドメイン（製品、顧客、販売など）をカプセル化した運用システムが単一のドメインに対応していない場合もあります
その場合、COTSシステムは、特定のソースに合わせたデータ製品ごとにドメインに合わせたインターフェースを公開することができます
図12-8は、共同の運用システムからの入力を持つデータ製品の例を示しています
リスナーサブスクリプションデータ製品は、そのドメインのマイクロサービスであるリスナーサブスクリプションサービスからデータを消費します
リスナーサブスクリプションの変更をリアルタイムに近いイベントとしてイベントログに公開します
サブスクリプションイベントログは、運用システムによって制御および保守される短期間の保存媒体です
データ製品はイベントを消費し、それらに変換を行い、最終的にはリスナーサブスクリプション情報の長期間のビューとして提供します
図12-8
共同の運用システムからの入力を持つデータ製品の例 共同の運用システムからデータを消費するための入力ポートを実装するための一般的なメカニズムには、モダンなシステムの場合は非同期のイベント駆動型データ共有、レガシーシステムの場合は変更データキャプチャが含まれます
ドメインイベントを共有するモダンな運用システムは、ますます一般的な実践となり、協調的なデータ製品へのデータ供給には素晴らしいモデルです
データの変更を検出して追跡するための統合パターンのセットである変更データキャプチャは、データ製品の入力ソースとして利用できますが、それは協力している運用システムからデータを受け取る最も望ましくない方法です
データベーストランザクションの内部実装を露出し、ビジネスドメインには対応していません
ただし、旧来のシステムの場合は、これが唯一の利用可能なオプションとなることもあります
データ製品が運用システムからデータを消費する方法は、運用システムの拡張能力に大きく影響を受けます
設計は、ドメインチーム内で、運用システムとデータ製品をどのように統合するかについての合意に最終的に依存します
他のデータ製品としてのデータソース データ製品はメッシュ上の他のデータ製品からデータを消費することができます
例えば、ポッドキャストの人口統計データ製品は、リスナープロファイルとポッドキャストリスナーのデータ製品からリスナーに関する属性を受け取ります
それはリスナープロファイル情報をポッドキャストリスナーと関連付け、これらのソースに分類変換を適用します
それからポッドキャストの人口統計に関する情報を提供します
この例は図12-9に示されています
図12-9. 他のデータ製品からの入力を持つデータ製品の例 この場合、入力ポートは別のデータ製品の出力ポートからデータを消費します
ソースデータ製品は、同じドメインに属するか、他のドメインに属することができます
ただし、入力データポートの実装と、どのようにして他のデータ製品の出力ポートからデータを消費するかは、標準化されています
データ製品は、上流の出力の実際の時間と処理時間を利用して、必要なデータを選択します
ストリーミング入力の場合、入力ポートの仕組みはソースデータの処理時間を追跡して、新しいデータが到着するたびにそのデータを処理します
データ製品としての自己をデータソースとして 一部の場合、データ製品のローカルな計算がデータのソースになることがあります
たとえば、アーティストの分類に関するデータ製品では、アーティストオンボーディングマイクロサービスから受け取った情報に加えて、「新興」や「衰退」などのデータを生成する機械学習モデルの推論など、ローカルな変換が新しいデータを生成します
さらに、データ製品は、入力として可能な分類のリストなど、ローカルに保存されたデータも使用します
図12-10に示すアーティストのデータ製品を考えてみましょう
その変換は、アーティストのオンボーディングマイクロサービスから受け取る情報に加えて、「新興」や「衰退」といったアーティストの分類に関する新しいデータを生成する機械学習モデルを実行します
図12-10. ローカル入力を持つデータ製品の例 データ消費の局所性 データメッシュは、基盤となる技術についてはできるだけ中立です
データメッシュアーキテクチャは、基盤技術やインフラストラクチャに関しては中立であり、可能な限り実装の詳細には立ち入りません
例えば、データが物理的にどこにあるか、データ製品がデータを物理的にコピーして別の場所に移動するかどうかは、プラットフォームによって構築された実装の詳細です
しかし、入出力ポートなどのデータメッシュアーキテクチャの構成要素は、データや処理の移動を物理的な境界を越えて抽象化するための優れたインターフェースです
データ製品の入力ポートを他のデータ製品の出力ポートに接続する一連のシンプルなAPIを使用すると、データが物理的なストレージや信頼境界を越えて移動することを隠すことができます
同様に、データ消費リクエストの実行は、1つの計算環境から別の計算環境に発行されるリモートクエリを実行することができます
この設計の結果、メッシュは1つまたは複数の物理インフラストラクチャ、複数のクラウド、およびオンプレミスのホスティング環境にまたがることができます
つまり、データ製品がソースからデータを消費する場合、そのプロセスでデータを基盤となるホスティング環境から別のホスティング環境に物理的に移動することができます
このように見える単純な機能は、クラウドへのデータ移行やマルチクラウドデータプラットフォームに深い影響を与える可能性があります
たとえば、Daffは現在、分析データと処理データをクラウド環境に移行しています
今日、ポッドキャストサービスはオンプレミスのインフラストラクチャと、その基盤となる運用データベース上で実行されています
ポッドキャストリスナーデータ製品はクラウド上で動作します
非同期の入力データポートインターフェースを介してデータを消費します
新しいポッドキャストリスナーが登録されると、ポッドキャストリスナーデータ製品は基本的に、オンプレミスのストリームからその情報をクラウドストレージにコピーし、その後クラウドベースの出力データインターフェースを介して利用可能にします
これにより、ビッグバンデータ移行戦略なしで、データ製品間での継続的なデータ移行が実行されます
同じメカニズムを使用して、データを1つのクラウドプロバイダから別のクラウドプロバイダに移動することもできます
たとえば、企業がマルチクラウド戦略を持っており、複数のクラウドプロバイダにわたってデータを保持したい場合、入力ポートの実装により、データをソースのクラウドプロバイダから利用者のクラウドプロバイダに移動することができます
もちろん、この実現のためには、基盤となるプラットフォームインフラストラクチャが、データ製品のインターネット対応のデータ利用可能性、トラスト境界を超えたアイデンティティ認証と承認、および出力データポートのインターネットアクセス可能なエンドポイントなど、いくつかの重要な機能をサポートする必要があります
図12-11は、2つの異なる環境間で2つの異なるデータ製品間でのデータ消費を示しています
図12-11
マルチ環境データ消費モデルデータ消費デザインデータ製品は、どのデータをどのソースからどのように消費するかを意図的に指定します
データ製品は、データを消費する能力を明確に定義および制御します
これは、他のデータアーキテクチャとは異なり、ソースがデータがどこにどのように到達するのかを知らずにデータをプロセッサにダンプするものです
たとえば、以前のアーキテクチャでは、外部の有向非循環グラフ（DAG）の仕様が、各プロセッサがデータをどのように受け取り、提供するかを定義するのではなく、各プロセッサが中央的に接続される方法を定義します
図12-12は、入力データポートを介してデータを消費するためのデータ製品の高レベルな設計コンポーネントを示しています
図12-12
データ製品のデータの消費を設計するための高レベルなコンポーネント表12-2は、データ製品のデータの消費を設計するための高レベルなコンポーネントの要約です
表12-2
データ製品のデータの消費を設計するための高レベルなコンポーネントの要約データの消費コンポーネントの説明入力データポートデータ製品がソースデータを受け取り、さらなる内部変換のために利用可能にするメカニズム
入力データポートの仕様入力データポートの宣言的な仕様で、データがどこからどのように消費されるかを構成します
非同期入力データポート非同期入力ポートの実装は、必要なソースデータが利用可能になったときに変換コードを反応的に呼び出します
イベントストリームへのサブスクリプションやファイルへの非同期I/O読み取りなどが、非同期入力ポートの例です
非同期入力ポートは、ソースデータ製品の処理時間ポインタを追跡し、新しい処理時間のデータを受け取ると、変換を反応的に実行します
同期入力データポート同期的な入力ポートは、ソースからデータを取得し、データが取得されたときに変換コードを呼び出します

たとえば、「デイリーポッドキャストサマリー」は、ポッドキャストリスナーからデータを同期的に取得し、データが取得されたときにカウントやその他のサマリーを計算します
データは毎日真夜中に取得されます
リモートクエリ入力ポートの仕様には、受け取る必要のあるデータに対してソース上で実行されるクエリを含めることができます
この機能により、重複して取得されるデータの量を削減することができます
クエリは、ソースが理解する問い合わせ言語（SQL、GraphQL、Fluxなど）で表され、入力ポートの仕様によって定義されます
入力ポートの同期化と一時的なストレージ入力ポートは、しばしばお互いに依存してデータを消費します
たとえば、アーティスト分類の変換コードは、アーティストとリスナーの2つの独立したソースからのデータが利用可能になるまで実行を開始することができません
一時的なストレージは、処理が必要な観測済みおよび未処理のデータを追跡するために必要です
データの変換
ほとんどのデータ製品は、どんなにわずかであっても変換を行います
私たちは新しい分析モデルを既存のデータと共有することで価値を見出してデータ製品を作成します
この新しいモデルを作成するには、変換が必要です
変換コードの作成と保守こそが、データ製品開発者が最も注意を払うものです
従来、この変換は、データパイプライン内でデータを入力元から出力先に移動させることで行われていました
データメッシュの設計では、パイプラインとして実装されているかどうかに関わらず、変換はデータ製品によってエンコードされ、抽象化されます
変換はデータ製品の内部実装であり、それによって制御されます
内部の問題であるため、具体的なデザイン方法について指定するつもりはありません
私の意見では、変換の実装方法は、データ製品開発者の好み、能力、ニーズによって選択されるべきです
データ製品の変換の実装方法をいくつか見ておくと役に立ちます
 プログラマティックな変換と非プログラマティックな変換 データの処理と変換は、非プログラマティックな変換（例：SQL、Flux、GraphQL）とプログラマティックなデータ処理（例：Apache Beam、Apache Spark、Metaflow）の2つの主要な分野に分類されます
非プログラマティックなアプローチでは、SQLなどの関係代数を使用したセット演算、またはFluxなどのフローベースの関数を使用します
いずれの場合も、データがどのように変換されるかの意図をステートメントで捉えることができます
これは多くのデータ製品開発者にとってシンプルでアクセスしやすいアプローチですが、ステートメントの機能に制限があります
より複雑な変換では、ステートメントが理解しにくく、モジュール化が難しく、自動的にテストすることも難しくなります
実際のところ、非プログラマティックな変換のみを行うデータ製品はあまり見つかりません
他の下流データ製品は、同じリモートクエリを自身で実行できるため、中間データ製品は必要ありません
図12-13は、非プログラマティックな変換の例を示しています
意図は、トップリスナーの人口統計情報を作成することです
トップリスナーのデータ製品は、今日の曲の再生履歴とリスナープロファイルの入力ポートを使用して、今日曲を聴いたリスナーのプロファイル情報を取得します
その後、リスナーの人口統計情報（例：今日最もリスナーの多い年齢層、あるいは最も少ない年齢層、今日最もリスナーの多い国、あるいは最も少ない国など）をさまざまな統計情報として生成します
 図12-13. 非プログラマティックな変換の例 一方、プログラマティックなデータ処理では、コードのロジック、条件、ステートメントを使用してデータを変換します
Apache BeamやApache Sparkなどのプログラミング言語（Python、Javaなど）を使用したプログラマティックなデータ処理ライブラリを使用できます
変換コードは、ホスティングプログラミング言語の完全な機能にアクセスできます（命令型または宣言型のいずれか）
モジュール化やテストも可能です
このアプローチは、非プログラミングのデータ製品開発者にとってはより複雑ですが、拡張性があります
このアプローチの利点は、再生された曲のレコードが到着するたびに統計情報を増分で計算できることです
データメッシュは、データ製品が変換のためにどのアプローチを採用するかについて立場を取らず、常識に頼るとされています
より複雑なユースケースにはプログラマティックなアプローチを使用し、変換が些細で非プログラマティックな場合は、何もせずに中間データ製品を作成しないでください
最終的な消費者にクエリを実行させます
なお、プログラマティックな変換の場合でも、入力ポートは処理の前にソースで非プログラマティックなクエリを呼び出す場合があります
これにより、変換コードに移動するデータの量が減少し、処理が上流のデータが存在する場所にプッシュされます
データフローベースの変換データフロープログラミングパラダイムは、1960年代に導入され、データが操作間を流れる有向グラフとしてコンピュータプログラムを定義します
このプログラミングパラダイムは、現代のデータパイプラインデザインに多くの影響を与えています
データパイプラインは、データが一つのステップから別のステップに流れるときに実行される一連の変換ステップ（関数）です
データメッシュは、パイプラインをトップレベルのアーキテクチャパラダイムやデータ製品の間で使用しないようにします
現在使用されているパイプラインの問題は、パイプラインの複雑さが増すにつれて、明確なインターフェース、契約、および抽象化が容易に維持できないことです
抽象化の不足により、パイプライン内の単一の障害が連鎖的な障害を引き起こします
パイプラインまたはデータ製品の範囲内でのパイプラインまたはデータフローベースのプログラミングモデルは、変換コードを実装するための自然なパラダイムです
このシナリオでは、パイプラインは単一のデータ製品のコンテキストと変換によって制約されるため、複雑さが低くなります
また、データ製品とともに単一のユニットとしてアップグレード、テスト、展開されます
したがって、パイプラインステージのタイトカプリングはあまり心配されません
要するに、データ製品の範囲内での変換にはデータパイプラインを使用しても構わないということですが、パイプラインステージがデータ製品の範囲を超えない限りです
データ製品間で変換が行われることはなく、読み取り専用の出力および入力データポートを通じてデータの提供と消費が行われるだけです
私はこれを「ダンブパイプとスマートフィルタの原則」と呼んでいます
図12-14は、パイプライン変換に関与するハイレベルのコンポーネントを示しています
図12-14
パイプライン変換MLとしての変換三つ目の変換カテゴリはモデルベースです
機械学習や統計モデルをデータ製品の変換として展開し、実行することです
例えば、リスナーの既存のプレイリストを拡張するためのTensorFlowリコメンダを使用して、デフが曲を推薦すると想像してください
TensorFlowモデルはシリアライズされ、プレイリストリコメンダデータ製品として展開されます
データ製品はリスナーのプレイリストを消費し、次に推薦される曲を予測し、それらをプレイリストの拡張推薦として保存します
それらはリスナーがリストを再生する際に読み取られて使用されます
リコメンダモデルはプログラムと所望のプログラミング言語内で実行されますが、その計算は主にモデルによって行われます
MLモデルは、マイクロサービスやアプリケーション、データ製品など、さまざまなコンテキストで展開できます
時間変動変換すべての変換の共通の特徴は、時間の軸に敬意を払うことです
処理時間と実際の時間です
変換コードは、入力を処理するときと出力を生成するときに、時間パラメータを意識しています
入力ポートメカニズムは、各ソースの処理時間を追跡します
変換コードは、ソースデータ上で行われた計算に基づいて出力の実際の時間を生成します
出力ポートは、データ製品の内部処理時間のカウントと共に変換されたデータを提供します
変換デザイン変換のデザイン、ビルドタイムの定義、展開、実行は、選択したフレームワークや基礎となる技術に大きく依存しています
変換が宣言的な文またはプログラミングコードによって実装されるかどうかにかかわらず、変換をビルド、テスト、展開、実行するためにプラットフォーム機能が必要です
図12-15は、データ製品の変換に関与するいくつかのハイレベルなデザイン要素を示しています
図12-15
データ製品の変換を設計するためのハイレベルのコンポーネント表12-3は、データ製品の変換機能を設計するために関与するハイレベルコンポーネントをまとめたものです
表12-3
データ製品のデータ変換に関与するハイレベルコンポ―ネントのまとめデータ変換コンポーネントの説明変換アーティファクト変換を定義するコード、設定、文またはモデル
このアーティファクトは、特定のレコード時間のために入力データ上で実行され、出力を生成します
変換ランタイム環境は、変換の設定に基づいて呼び出されます
例えば、定期的な時間基準や必要な入力データの可用性に基づいて呼び出されます
呼び出されると、データ製品コンテナによって制約される計算環境が必要です
基礎となるプラットフォームはこの環境を提供します
一時的なストレージ変換コードのステップは、変換のさまざまなフェーズで状態を保持するために一時的なストレージへのアクセスが必要となる場合があります
これは基礎となるプラットフォームによって提供されます
要約この章では、データ製品のデザイン要素と特性をアーキテクチャの量子として紹介し、データの自律的な消費、変換、提供を行います
分散メッシュアーキテクチャでデータを提供するために、データ製品の出力ポートにはいくつかのプロパティがあります
読み取り専用およびマルチモーダルなAPIを介して、時間的にバイアスを持った不変のデータを提供します
これらの制約は、データメッシュの目標と前提によって駆動されました
データ製品は、ドメイン指向のデータを多様なデータユーザーにネイティブに提供します（マルチモーダルアクセス）
データ製品は、時間的な分析ユースケースで使用できます（時間的バイアス）
データ製品は、他のデータ製品からデータを安全に取得し、変換して提供すると同時に、特定の時点での一貫性のグローバルな状態を維持します（時間的バイアス）
データ製品のユーザーはデータを分析および機械学習のために処理します（読み取り専用）
データ製品は、どこからどのように上流データを受け取るかを制御できます
入力ポートを使用して、同じドメイン内の協力的な運用システムまたは上流データ製品からデータを取得します
入力ポートは、マルチクラウド展開またはインクリメンタルなクラウド移行を支援するために、異なるホスティング環境間でのデータ移動を実装することができます
データ製品はほぼ常に、より高い価値のあるデータを提供するために何らかの形式の変換を行います
変換の具体的な実装方法は、データパイプライン関数、高度なクエリ、または機械学習モデルの推論など、複数の形式を取ることがあります
この章では、データ製品のデザインにおけるデータ共有のコア機能に対する高レベルな思考とアプローチを提示しました
将来のデータメッシュの実装では、この思考を評価し、改善することを期待しています