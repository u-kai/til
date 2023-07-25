Stack Topologies for Application Clusters An application cluster is comprised of different moving parts. One set of parts is the applications and services that manage the cluster. Some of these cluster management services include: Scheduler, which decides how many instances of each application to run, and where to run them Monitoring, to detect issues with application instances so that they can be restarted or moved if necessary
Configuration registry, for storing information needed to manage the cluster and configure applications Service discovery, enabling applications and services to find the current location of application instances Management API and UI, enabling tools and users to interact with the cluster Many cluster deployments run management services on dedicated servers, separate from the services that host application instances. These services should probably also run clustered, for resilience. The other main parts of an application cluster are the application hosting nodes. These nodes are a pool of servers where the scheduler runs application instances. It’s common to set these up as a server cluster (see “Compute Resources”) to manage the number and location of server instances. A service mesh (see “Service Mesh”) may run sidecar processes on the host nodes, alongside applica​tion instances. An example application cluster (Figure 14-2) includes servers to run cluster management services, a server cluster to run application instances, and network address blocks.
Networking structures for an application cluster may be flat. The cluster services assign network addresses to application instances. It should also handle network security, including encryption and connection management, often using a service mesh for this. You use infrastructure stacks to provision this infrastructure. Chapter 5 explained that the size of a stack and the scope of its contents have implications for the speed and risk of changes.
“Patterns and Antipatterns for Structuring Stacks” in particular lists patterns for arranging resources across stacks. The following examples show how these patterns apply to cluster infrastructure. Monolithic Stack Using Cluster as a Service The simplest design is to define all of the parts of your cluster in a single stack, following the monolithic stack antipattern (see “Antipattern: Monolithic Stack”). Although monoliths become an antipattern at scale, a single stack can be useful when starting out with a small and simple cluster. Example 14-1 uses a cluster as a service, similar to AWS EKS, AWS ECS, Azure AKS, and Google GKE. So the code defines the cluster, but it doesn’t need to provision servers to run the cluster management, because the infrastructure platform handles that behind the scenes. Example 14-1. Stack code that defines everything for a cluster address_block:
name: cluster_network
address_range: 10.1.0.0/16"
vlans: - vlan_a:
address_range: 10.1.0.0/8 - vlan_b:
address_range: 10.1.1.0/8 - vlan_c:
address_range: 10.1.2.0/8

application_cluster:
name: product_application_cluster
address_block: $address_block.cluster_network

server_cluster:
name: "cluster_nodes"
min_size: 1
max_size: vlans: $address_block.cluster_network.vlans
each_server_node:
source_image: cluster_node_image
memory: 8GB This example leaves out many of the things you’d have for a real cluster, like network routes, security policies, and monitoring. But it shows that the significant elements of networking, cluster definition, and server pool for host nodes are all in the same project. Monolithic Stack for a Packaged Cluster Solution The code in Example 14-1 uses an application cluster service provided by the infrastructure platform. Many teams instead use a packaged application cluster solution (as described in “Packaged Cluster Distribution”). These solutions have installers that deploy the cluster management software onto servers. When using one of these solutions, your infrastructure stack provisions the infrastructure that the installer needs to deploy and configure the cluster. Running the installer should be a separate step. This way, you can test the infrastructure stack separately from the application cluster. Hopefully, you can define the configuration for your cluster as code. If so, you can manage that code appropriately, with tests and a pipeline to help you to deliver updates and changes easily and safely. It may be useful to use server configuration code (as in Chapter 11) to deploy your cluster management system onto your servers. Some of the packaged products use standard configuration tools, such as Ansible for OpenShift, so you may be able to incorporate these into your stack building process. Example 14-2 shows a snippet that you would add to the monolithic stack code in Example 14-1 to create a server for the cluster management application. Example 14-2. Code to build a cluster management server virtual_machine:
name: cluster_manager
source_image: linux-base
memory: 4GB
provision:
tool: servermaker
parameters:
maker_server: maker.shopspinner.xyz
role: cluster_manager The code configures the server by running the fictitious servermaker command, applying the cluster_manager role.
