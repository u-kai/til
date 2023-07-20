Creating a New Server Instance The basic server life cycle described earlier starts with creating a new server instance. A more extended life cycle includes steps for creating and updating server images that you can create server instances from, but we’ll leave that topic for another chapter (Chapter 13). For now, our life cycle starts with creating an instance, as shown in Figure 11-4. Figure 11-4. Creating a server instance The full server creation process includes provisioning the server instance so that it is fully ready for use. Some of the activities involved in creating and provisioning a server instance are: Allocate infrastructure resources to instantiate the server instance. To do this, a physical server is selected from a pool or a host server is selected to run it as a virtual machine. For a virtual machine, the hypervisor software on the host allocates memory and other resources. The process may also allocate storage for disk volumes for the server instance. Install the operating system and initial software. The OS may be copied onto a disk volume, as when booting from a server image like an AMI. Or the new instance may execute an installation process that selects and copies files to the new instance, maybe using a scriptable installer. Examples of scripted OS installers
include Red Hat Kickstart, Solaris JumpStart, Debian Preseed, and the Windows installation answer file. Additional configuration is applied to the server instance. In this step, the process runs the server configuration tool, following a pattern described in “How to Apply Server Configuration Code”. Configure and attach networking resources and policies. This process can include assigning the server to a network address block, creating routes, and adding firewall rules. Register the server instance with services. For example, add the new server to a monitoring system. These aren’t mutually exclusive — your server creation process may use one or more methods for running these different activities. There are several mechanisms that you might use to trigger the creation of a server instance. Let’s go over each of these mechanisms. Hand-Building a New Server Instance Infrastructure platforms provide tools to create new servers, usually a web UI, command-line tool, and sometimes a GUI application. Each time you create a new server, you select the options you want, including the source image, resources to allocate, and networking details: $ mycloud server new \
 --source-image=stock-linux-1.23 \
 --memory=2GB \
 --vnet=appservers While it’s useful to play with UIs and command-line tools to experiment with a platform, it’s not a good way to create servers that people need. The same principles for creating and configuring stacks discussed previously (see “Patterns for Configuring Stacks”) apply for servers. Setting options manually (as described in “Antipattern:
Manual Stack Parameters”) encourages mistakes, and leads to inconsistently configured servers, unpredictable systems, and too much maintenance work. Using a Script to Create a Server You can write scripts that create servers consistently. The script wraps a command-line tool or uses the infrastructure platform’s API to create a server instance, setting the configuration options in the script code or configuration files. A script that creates servers is reusable, consistent, and transparent. This is the same concept as the scripted parameters pattern for stacks (see “Pattern: Scripted Parameters”): mycloud server new \
 --source-image=stock-linux-1.23 \
 --memory=2GB \
 --vnet=appservers This script is the same as the previous command line example. But because it’s in a script, other people on my team can see how I created the server. They can create more servers and be confident that they will behave the same as the one I created. Before Terraform and other stack management tools emerged, most teams I worked with wrote scripts for creating servers. We usually made the scripts configurable with configuration files, as with the stack configuration pattern (see “Pattern: Stack Configuration Files”). But we spent too much time improving and fixing these scripts. Using a Stack Management Tool to Create a Server With a stack management tool, as discussed in Chapter 5, you can define a server in the context of other infrastructure resources, as shown in Example 11-1. The tool uses the platform API to create or update the server instance. Example 11-1. Stack code that defines a server server:
source_image: stock-linux-1.23
memory: 2GB
vnet: ${APPSERVER_VNET} There are several reasons why using a stack tool to create servers is so handy. One is that the tool takes care of logic that you would have to implement in your own
script, like error checking. Another advantage of a stack is that it handles integration with other infrastructure elements; for example, attaching the server to network structures and storage defined in the stack. The code in Example 11-1 sets the vnet parameter using a variable, ${APPSERVER_VNET}, that presumably refers to a network structure defined in another part of the stack code. Configuring the Platform to Automatically Create Servers Most infrastructure platforms can automatically create new server instances in specific circumstances. Two common cases are auto-scaling, adding servers to handle load increases, and auto-recovery, replacing a server instance when it fails. You can usually define this with stack code, as in Example 11-2. Example 11-2. Stack code for automated scaling server_cluster:
server_instance:
source_image: stock-linux-1.23
memory: 2GB
vnet: ${APPSERVER_VNET}
scaling_rules:
min_instances: 2
max_instances: 5
scaling_metric: cpu_utilization
scaling_value_target: 40%
health_check:
type: http
port: 8443
path: /health
expected_code: 200
wait: 90s This example tells the platform to keep at least 2 and at most 5 instances of the server running. The platform adds or removes instances as needed to keep the CPU utilization across them close to 40%. The definition also includes a health check. This check makes an HTTP request to /health on port 8443, and considers the server to be healthy if the request returns a 200 HTTP response code. If the server doesn’t return the expected code after 90 seconds,

seconds, the platform assumes the server has failed, and so destroys it and creates a new instance. Using a Networked Provisioning Tool to Build a Server In Chapter 3, I mentioned bare-metal clouds, which dynamically provision hardware servers. The process for doing this usually includes these steps: Select an unused physical server and trigger it to boot in a “network install” mode supported by the server firmware (e.g., PXE boot).3 The network installer boots the server from a simple bootstrap OS image to initiate the OS installation. Download an OS installation image and copy it to the primary hard drive. Reboot the server to boot the OS in setup mode, running an unattended scripted OS installer. There are a number of tools to manage this process, including: Crowbar Cobbler FAI, or Fully Automatic Installation Foreman MAAS Rebar Tinkerbell Instead of booting an OS installation image, you could instead boot a prepared server image. Doing this creates the opportunity to implement some of the other methods for preparing servers, as described in the next section.
FaaS Events Can Help Provision a Server FaaS serverless code can play a part in provisioning a new server. Your platform can run code at different points in the process, before, during, and after a new instance is created. Examples include assigning security policies to a new server, registering it with a monitoring service, or running a server configuration tool.
新しいサーバーインスタンスを作成する
先ほど説明した基本的なサーバーライフサイクルは、新しいサーバーインスタンスの作成から始まります。もっと詳細なライフサイクルには、サーバーインスタンスを作成するためのサーバーイメージの作成と更新の手順が含まれますが、それについては別の章（第13章）で説明します。今のところ、我々のライフサイクルはインスタンスの作成とともに始まります（図11-4参照）。図11-4 サーバーインスタンスの作成 
完全なサーバーの作成プロセスには、サーバーインスタンスが使用のために完全に準備されるようにするためにプロビジョニングする作業が含まれます。サーバーインスタンスの作成とプロビジョニングに関わる活動の一部は次のとおりです：

- サーバーインスタンスを起動するためのインフラリソースを割り当てます。これには、プールから物理サーバーを選択するか、ホストサーバーを仮想マシンとして実行するために選択することが含まれます。仮想マシンの場合、ホスト上のハイパーバイザーソフトウェアがメモリやその他のリソースを割り当てます。このプロセスでは、サーバーインスタンスのディスクボリュームにストレージも割り当てられる場合があります。
- オペレーティングシステムと初期ソフトウェアをインストールします。オペレーティングシステムは、AMIのようなサーバーイメージからディスクボリュームにコピーする場合があります。または、新しいインスタンスがスクリプト可能なインストーラーを使用してファイルを選択してコピーするインストールプロセスを実行する場合もあります。スクリプト化されたOSインストーラーの例には、Red Hat Kickstart、Solaris JumpStart、Debian Preseed、Windowsインストール回答ファイルなどがあります。
- サーバーインスタンスに追加の構成が適用されます。このステップでは、プロセスがサーバー構成ツールを実行し、"サーバー構成コードの適用方法"で説明されているパターンに従います。
- ネットワーキングリソースとポリシーを構成してアタッチします。このプロセスには、サーバーをネットワークアドレスブロックに割り当てる、ルートを作成する、ファイアウォールルールを追加するなどが含まれます。
- サーバーインスタンスをサービスに登録します。例えば、新しいサーバーを監視システムに追加します。これらは相互に排他的ではないため、サーバーの作成プロセスではこれらのさまざまなアクティビティを実行する1つ以上のメソッドを使用する場合があります。

サーバーインスタンスの作成をトリガーするために使用できるいくつかのメカニズムがあります。それぞれのメカニズムについて説明しましょう。

新しいサーバーインスタンスを手動で作成する
インフラストラクチャプラットフォームは、通常、Web UI、コマンドラインツール、GUIアプリケーションなど、新しいサーバーを作成するためのツールを提供します。新しいサーバーを作成するたびに、ソースイメージ、割り当てるリソース、ネットワーキングの詳細など、必要なオプションを選択します。

$ mycloud server new \
 --source-image=stock-linux-1.23 \
 --memory=2GB \
 --vnet=appservers

プラットフォームを試行するためにUIやコマンドラインツールを使うことは役立ちますが、人々が必要とするサーバーを作成する良い方法ではありません。以前に説明したスタックの作成と構成の原則（「スタックの構成パターン」を参照）は、サーバーにも適用されます。手動でオプションを設定する（「アンチパターン：手動スタックパラメータ」で説明）と、ミスが生じやすく、設定が一貫しておらず、予測できないシステムと多くのメンテナンス作業を引き起こします。

スクリプトを使用してサーバーを作成する
サーバーを一貫して作成するためのスクリプトを作成することができます。スクリプトは、コマンドラインツールをラップするか、スクリプトコードや設定ファイルで構成オプションを設定するためにインフラストラクチャプラットフォームのAPIを使用します。サーバーを作成するスクリプトは再利用可能で一貫性があり、透明です。これは、スタックのスクリプト化されたパラメータパターンと同じコンセプトです。

mycloud server new \
 --source-image=stock-linux-1.23 \
 --memory=2GB \
 --vnet=appservers

このスクリプトは、前のコマンドラインの例と同じです。しかし、スクリプト内にあるため、私のチームの他のメンバーが私がどのようにサーバーを作成したかを見ることができます。彼らもより多くのサーバーを作成し、自分が作成したものと同じ動作をすることを確信することができます。Terraformや他のスタック管理ツールが登場する前は、私が働いていたほとんどのチームは、サーバーを作成するためのスクリプトを書いていました。通常、我々はこれらのスクリプトを設定ファイルで設定可能にしました（スタック構成パターンと同様）。しかし、これらのスクリプトの改善や修正に時間を費やしすぎました。

サーバーを作成するためのスタック管理ツールを使用する
第5章で説明したようなスタック管理ツールを使用すると、サーバーを他のインフラストラクチャリソースのコンテキストで定義することができます（例11-1参照）。このツールは、プラットフォームのAPIを使用してサーバーインスタンスを作成または更新します。

Example 11-1. サーバーを定義するスタックコード
server:
  source_image: stock-linux-1.23
  memory: 2GB
  vnet: ${APPSERVER_VNET}

サーバーを作成するためにスタックツールを使用するメリットはいくつかあります。1つは、スクリプトで実装する必要があるロジックをツールが処理してくれることです（エラーチェックなど）。スタックのもう1つの利点は、他のインフラストラクチャ要素との統合処理を扱うことです。例えば、スタックで定義されたネットワーク構造やストレージにサーバーをアタッチすることができます。例11-1のコードでは、vnetパラメータを変数${APPSERVER_VNET}を使用して設定し、それはおそらくスタックコードの別の箇所で定義されたネットワーク構造を指していると想定されています。

プラットフォームを自動的にサーバーを作成できるように構成する
ほとんどのインフラストラクチャプラットフォームは、特定の状況で自動的に新しいサーバーインスタンスを作成できます。一般的なケースは、スケーリング（負荷増加に対応するためにサーバーを追加する）とオートリカバリ（サーバーインスタンスが失敗した場合に置き換える