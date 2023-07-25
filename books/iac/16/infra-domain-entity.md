Pattern: Infrastructure Domain Entity A infrastructure domain entity implements a high-level stack component by combining multiple lower-level infrastructure resources. An example of a higher-level concept is the infrastructure needed to run an application. This example shows how a library that implements a Java application infrastructure instance might be used from stack project code: use module: application_server
service_name: checkout_service
application_name: checkout_application
application_version: 1.23
traffic_level: medium The code defines the application and version to deploy, and also a traffic level. The domain entity library code could look similar to the bundle module example (Example 16-5), but includes dynamic code to provision resources according to the traffic_level parameter: ...
switch (${traffic_level}) {
case ("high") {
$appserver_cluster.min_size = 3
$appserver_cluster.max_size = 9
} case ("medium") {
$appserver_cluster.min_size = 2
$appserver_cluster.max_size = 5
} case ("low") {
$appserver_cluster.min_size = 1
$appserver_cluster.max_size = 2
}
}
... Motivation A domain entity is often part of an abstraction layer (see “Building an Abstraction Layer”) that people can use to define and build infrastructure based on higher-level requirements. An infrastructure platform team builds components that other teams can use to assemble stacks. Applicability Because an infrastructure domain entity dynamically provisions infrastructure resources, it should be written in an imperative language rather than a declarative one. See “Declarative Versus Imperative Languages for Infrastructure” for more on why.
Implementation On a concrete level, implementing an infrastructure domain entity is a matter of writing the code. But the best way to create a high-quality codebase that is easy for people to learn and maintain is to take a design-led approach. I recommend drawing from lessons learned in software architecture and design. The infrastructure domain entity pattern derives from Domain Driven Design (DDD), which creates a conceptual model for the business domain of a software system, and uses that to drive the design of the system itself.3 Infrastructure, especially one designed and built as software, should be seen as a domain in its own right. The domain is building, delivering, and running software. A particularly powerful approach is for an organization to use DDD to design the architecture for the business software, and then extend the domain to include the systems and services used for building and running that software. Related patterns A bundle module (see “Pattern: Bundle Module”) is similar to a domain entity in that it creates a cohesive collection of infrastructure resources. But a bundle module normally creates a fairly static set of resources, without much variation. The mindset of a bundle module is usually bottom-up, starting with the infrastructure resources to create. A domain entity is a top-down approach, starting with what’s required for the use case. Most spaghetti modules (see “Antipattern: Spaghetti Module”) for infrastructure stacks are a result of pushing declarative code to implement dynamic logic. But sometimes an infrastructure domain entity becomes overly complicated. A domain entity with poor cohesion becomes a spaghetti module.
