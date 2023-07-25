Monolithic Stack Using Cluster as a Service The simplest design is to define all of the parts of your cluster in a single stack, following the monolithic stack antipattern (see “Antipattern: Monolithic Stack”). Although monoliths become an antipattern at scale, a single stack can be useful when starting out with a small and simple cluster. Example 14-1 uses a cluster as a service, similar to AWS EKS, AWS ECS, Azure AKS, and Google GKE. So the code defines the cluster, but it doesn’t need to provision servers to run the cluster management, because the infrastructure platform handles that behind the scenes. Example 14-1. Stack code that defines everything for a cluster address_block:
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
クラスターをサービスとして使用するモノリシックスタック
最もシンプルなデザインは、クラスターのすべての部分を単一のスタックで定義することであり、モノリシックスタックのアンチパターン（「アンチパターン：モノリシックスタック」を参照）に従います。モノリシックはスケール時にアンチパターンとなりますが、小さくてシンプルなクラスターで始める際には単一のスタックは有用です。例 14-1 では、AWS EKS、AWS ECS、Azure AKS、および Google GKE と同様のクラスターをサービスとして使用しています。したがって、コードはクラスターを定義しますが、クラスター管理を実行するためのサーバーをプロビジョニングする必要はなく、インフラストラクチャプラットフォームがその背後で処理します。例 14-1。クラスターを定義するスタックコードの例：

address_block:
name: cluster_network
address_range: 10.1.0.0/16" vlans:

- vlan_a:
  address_range: 10.1.0.0/8
- vlan_b:
  address_range: 10.1.1.0/8
- vlan_c:
  address_range: 10.1.2.0/8

application_cluster:
name: product_application_cluster
address_block: $address_block.cluster_network

server_cluster:
name: "cluster_nodes"
min_size: 1
max_size:
vlans: $address_block.cluster_network.vlans
each_server_node:
source_image: cluster_node_image
memory: 8GB
Monolithic Stack for a Packaged Cluster Solution The code in Example 14-1 uses an application cluster service provided by the infrastructure platform. Many teams instead use a packaged application cluster solution (as described in “Packaged Cluster Distribution”). These solutions have installers that deploy the cluster management software onto servers. When using one of these solutions, your infrastructure stack provisions the infrastructure that the installer needs to deploy and configure the cluster. Running the installer should be a separate step. This way, you can test the infrastructure stack separately from the application cluster. Hopefully, you can define the configuration for your cluster as code. If so, you can manage that code appropriately, with tests and a pipeline to help you to deliver updates and changes easily and safely. It may be useful to use server configuration code (as in Chapter 11) to deploy your cluster management system onto your servers. Some of the packaged products use standard configuration tools, such as Ansible for OpenShift, so you may be able to incorporate these into your stack building process. Example 14-2 shows a snippet that you would add to the monolithic stack code in Example 14-1 to create a server for the cluster management application. Example 14-2. Code to build a cluster management server virtual_machine:
name: cluster_manager
source_image: linux-base
memory: 4GB
provision:
tool: servermaker
parameters:
maker_server: maker.shopspinner.xyz
role: cluster_manager The code configures the server by running the fictitious servermaker command, applying the cluster_manager role.
クラスターをモノリシックスタックのアンチパターンに従って単一のスタックで定義することで、クラスターのすべてのパーツを定義する最もシンプルなデザインになります。クラスターがスケールするとモノリシックはアンチパターンになりますが、小規模なシンプルなクラスターで始めるときには単一のスタックが役立ちます。例14-1では、AWS EKS、AWS ECS、Azure AKS、およびGoogle GKEと同様のクラスターをサービスとして使用しています。コードはクラスターを定義しますが、クラスター管理を実行するためにサーバーをプロビジョニングする必要はありません。インフラストラクチャプラットフォームがそれを裏で処理します。例14-1は、すべてのネットワーキング要素、クラスターの定義、およびホストノードのサーバープールの重要な要素が同じプロジェクトに含まれていることを示しています。

モノリシックスタックはパッケージ化されたクラスターソリューションにも使用されます。これらのソリューションは、クラスター管理ソフトウェアをサーバーに展開するインストーラーを持っています。これらのソリューションを使用する場合、インフラストラクチャスタックは、インストーラーがクラスターを展開および構成するために必要とするインフラストラクチャをプロビジョニングします。インストーラーの実行は別のステップで行うべきです。これにより、アプリケーションクラスターとは別にインフラストラクチャスタックをテストすることができます。クラスターの設定をコードとして定義できれば、最新の更新や変更を簡単かつ安全に配信するためのテストやパイプラインなどの適切な方法でそのコードを管理できるでしょう。サーバーの構成コード（第11章の章で説明されているコード）を使用して、クラスター管理システムをサーバーに展開すると便利です。いくつかのパッケージ化製品では、OpenShift向けのAnsibleなどの標準的な構成ツールが使用されているため、これらをスタックビルディングプロセスに組み込むことができるかもしれません。例14-2は、クラスター管理アプリケーションのためのサーバーを作成するために、例14-1のモノリシックスタックコードに追加するスニペットを示しています。コードは、架空のservermakerコマンドを実行し、cluster_managerの役割を適用してサーバーを構成します。