Online Image Building Process The online method for building server images involves booting a new, clean server instance, configuring it, and then converting it into the infrastructure platform’s server image format, as shown in Figure 13-2.
The process starts by booting a new server instance from an origin image, as shown in Example 13-1. See “Origin Content for a Server Image” for more about the options for origin images. If you use a tool like Packer to build your image, it uses the platform API to create the server instance. Example 13-1. Example image builder code for booting from an origin image image:
name: shopspinner-linux-image
platform: fictional-cloud-service
origin: fci-12345678
region: europe
instance_size: small This example code builds an FCS image named shopspinner-linux-image, booting it from the existing FCI whose ID is fci-12345678.3 This is a stock Linux distribution provided by the cloud vendor. This code doesn’t specify any configuration actions, so the resulting image is a fairly straight copy of the origin image. The code also defines a few characteristics of the server instance to boot for building the image, including the region and instance type. Infrastructure for the builder instance You probably need to provide some infrastructure resources for your server image building instance — for example, an address block. Make sure to provide an appropriately secure and isolated context for this instance. It might be convenient to run the builder instance using existing infrastructure, but doing this might expose your image building process to compromise. Ideally, create disposable infrastructure for building your image, and destroy it afterward. Example 13-2 adds a subnet ID and SSH key to the image builder instance. Example 13-2. Dynamically assigned infrastructure resources image:
name: shopspinner-linux-image
origin: fci-12345678
region: europe
size: small
subnet: ${IMAGE_BUILDER_SUBNET_ID}
ssh_key: ${IMAGE_BUILDER_SSH_KEY} These new values use parameters, which you can automatically generate and pass to the image builder tool. Your image building process might start by creating an instance of an infrastructure stack that defines these things, then destroys them after running the image builder tool. This is an example of an infrastructure tool orchestration script (see “Using Scripts to Wrap Infrastructure Tools” for more on this topic). You should minimize the resources and access that you provision for the image builder server instance. Only allow inbound access for the server provisioning tool, and only do this if you are pushing the configuration (see “Pattern: Push Server Configuration”). Only allow outbound access necessary to download packages and configuration. Configuring the builder instance Once the server instance is running, your process can apply server configuration to it. Most server image building tools support running common server configuration tools (see “Server Configuration Code”), in a similar way to stack management tools (see “Applying Server Configuration When Creating a Server”). Example 13-3 extends Example 13-2 to run the Servermaker tool to configure the instance as an application server. Example 13-3. Using the Servermaker tool to configure the server instance
image:
name: shopspinner-linux-image
origin: fci-12345678
region: europe
instance_size: small
subnet: ${IMAGE_BUILDER_SUBNET_ID}
configure:
tool: servermaker
code_repo: servermaker.shopspinner.xyz
server_role: appserver The code applies a server role, as in “Server Roles”, which may install an application server and related elements.
オンラインイメージビルディングプロセス サーバーイメージを構築するためのオンラインメソッドは、新しいクリーンなサーバーインスタンスを起動し、それを設定して、インフラストラクチャプラットフォームのサーバーイメージ形式に変換することから始まります（図13-2参照）。
プロセスは、起源イメージから新しいサーバーインスタンスを起動することから始まります（例13-1参照）。起源イメージのオプションについては、「サーバーイメージの起源コンテンツ」を参照してください。イメージを構築するためにPackerのようなツールを使用する場合、プラットフォームAPIを使用してサーバーインスタンスを作成します。例13-1. 起源イメージから起動するための例のイメージビルダーコード：
name: shopspinner-linux-image
platform: fictional-cloud-service
origin: fci-12345678
region: europe
instance_size: small この例のコードは、FCSイメージをshopspinner-linux-imageという名前で構築し、IDがfci-12345678である既存のFCIから起動します。これはクラウドベンダーが提供する標準のLinuxディストリビューションです。このコードは構成アクションを指定していないため、結果のイメージは起源イメージのかなりのコピーです。コードはまた、イメージを構築するために起動するサーバーインスタンスのいくつかの特性（リージョンとインスタンスタイプを含む）も定義しています。ビルダーインスタンスのインフラストラクチャ サーバーイメージのビルディングインスタンスには、おそらくいくつかのインフラリソースを提供する必要があります。たとえば、アドレスブロックなどです。このインスタンスに適切なセキュリティと隔離されたコンテキストを提供するようにしてください。既存のインフラストラクチャを使用してビルダーインスタンスを実行することは便利かもしれませんが、それによってイメージビルディングプロセスが危険にさらされる可能性があります。理想的には、イメージを構築するための使い捨てのインフラストラクチャを作成し、それを後で破棄することです。例13-2では、サブネットIDとSSHキーをイメージビルダーインスタンスに追加しています。例13-2. ダイナミックに割り当てられたインフラストラクチャリソースイメージ：
name: shopspinner-linux-image
origin: fci-12345678
region: europe
size: small
subnet: ${IMAGE_BUILDER_SUBNET_ID}
ssh_key: ${IMAGE_BUILDER_SSH_KEY} これらの新しい値はパラメータを使用しており、イメージビルダーツールに自動的に生成および渡すことができます。イメージビルディングプロセスは、これらの要素を定義するインフラストラクチャスタックのインスタンスを作成し、イメージビルダーツールを実行した後にそれらを破棄することから始まる場合があります。これはインフラストラクチャツールのオーケストレーションスクリプトの例です（このトピックの詳細については、「インフラストラクチャツールをラップするためのスクリプトの使用」を参照してください）。イメージビルダーサーバーインスタンスに対して割り当てるリソースとアクセスを最小限にすべきです。サーバーのプロビジョニングツールの受信アクセスのみを許可し、これを行う場合は構成をプッシュする必要があります（「パターン：サーバーの構成プッシュ」を参照）。パッケージと構成のダウンロードに必要なアウトバウンドアクセスのみを許可するようにしてください。 ビルダーインスタンスの構成 サーバーインスタンスが稼働している場合、プロセスはサーバー構成を適用することができます。ほとんどのサーバーイメージビルディングツールは、共通のサーバー構成ツールを実行することをサポートしています（「サーバー構成コード」を参照）、スタック管理ツールと同様の方法です（「サーバーを作成するときにサーバー構成を適用する」を参照）。例13-3は、Example 13-2を拡張してServermakerツールを使用してインスタンスをアプリケーションサーバーとして構成する方法を示しています。例13-3. Servermakerツールを使用してサーバーインスタンスを構成する
image:
name: shopspinner-linux-image
origin: fci-12345678
region: europe
instance_size: small
subnet: ${IMAGE_BUILDER_SUBNET_ID}
configure:
tool: servermaker
code_repo: servermaker.shopspinner.xyz
server_role: appserver コードは、アプリケーションサーバーや関連要素をインストールする場合がある「サーバーロール」のようにサーバーロールを適用します。