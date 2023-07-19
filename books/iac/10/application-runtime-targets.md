Application Runtime Targets Implementing an application-driven strategy starts with analyzing the runtime requirements of your application portfolio. You then design application runtime solutions to meet those requirements, implementing reusable stacks, stack components, and other elements teams can use to assemble environments for specific applications. Deployable Parts of an Application A deployable release for an application may involve different elements. Leaving aside documentation and other metadata, examples of what may be part of an application deployment include: Executables The core of a release is the executable file or files, whether they are binaries or interpreted scripts. You can consider libraries and other files used by the executables as members of this category. Server configuration Many application deployment packages make changes to server configuration. These can include user accounts that processes will run as, folder structures, and changes to system configuration files. Data structures When an application uses a database, its deployment may create or update schemas. A given version of the schema usually corresponds to a version of the executable, so it is best to bundle and deploy these together. Reference data An application deployment may populate a database or other storage with an initial set of data. This could be reference data that changes with new versions or example data that helps users get started with the application the first time they install it. Connectivity An application deployment may specify network configuration, such as network ports. It may also include elements used to support connectivity, like certificates or keys used to encrypt or authenticate connections. Configuration parameters An application deployment package may set configuration parameters, whether by copying configuration files onto a server, or pushing settings into a registry.
You can draw the line between application and infrastructure in different places. You might bundle a required library into the application deployment package, or provision it as part of the infrastructure. For example, a container image usually includes most of the operating system, as well as the application that will run on it. An immutable server or immutable stack takes this even further, combining application and infrastructure into a single entity. On the other end of the spectrum, I’ve seen infrastructure code provision libraries and configuration files for a specific application, leaving very little in the application package itself. This question also plays into how you organize your codebase. Do you keep your application code together with the infrastructure code, or keep it separate? The question of code boundaries is explored later in this book (see Chapter 18), but one principle is that it’s usually a good idea to align the structure of your codebase to your deployable pieces.
Deployment Packages Applications are often organized into deployment packages, with the package format depending on the type of the runtime environment. Examples of deployment package formats and associated runtimes are listed in Table 10-1. Table 10-1. Examples of target runtimes and application package formats Target runtime Example packages Server operating system Red Hat RPM files, Debian .deb files, Windows MSI installer packages Language runtime engine Ruby gems, Python pip packages, Java .jar, .war, and .ear files. Container runtime Docker images Application clusters Kubernetes Deployment Descriptors, Helm charts FaaS serverless Lambda deployment package A deployment package format is a standard that enables deployment tools or runtimes to extract the parts of the application and put them in the right places.
アプリケーションのランタイムターゲット
アプリケーション駆動型の戦略を実装するには、まずアプリケーションポートフォリオのランタイム要件を分析します。次に、これらの要件を満たすためにアプリケーションのランタイムソリューションを設計し、特定のアプリケーションのために環境を組み立てるためにチームが使用できる再利用可能なスタック、スタックコンポーネント、その他の要素を実装します。
アプリケーションのデプロイ可能な部分
アプリケーションのデプロイメントには、さまざまな要素が含まれる場合があります。ドキュメンテーションやその他のメタデータを置いておいて、アプリケーションデプロイメントの一部として考えられる例には以下があります。
実行可能ファイル
リリースの中核は、バイナリまたは解釈スクリプトである実行可能ファイルまたはファイル群です。実行可能ファイルが使用するライブラリやその他のファイルも、このカテゴリのメンバーと考えることができます。
サーバーの設定
多くのアプリケーションデプロイメントパッケージは、サーバーの設定を変更します。これには、プロセスが実行されるユーザーアカウント、フォルダー構造、システム設定ファイルの変更などが含まれる場合があります。
データ構造
アプリケーションがデータベースを使用する場合、デプロイメントによってスキーマが作成または更新されることがあります。スキーマの特定のバージョンは通常、実行可能ファイルのバージョンに対応しているため、これらを一緒にバンドルしてデプロイするのが最善です。
リファレンスデータ
アプリケーションのデプロイメントでは、データベースや他のストレージに初期データのセットを入れることがあります。これは、新しいバージョンで変更されるリファレンスデータや、ユーザーが最初にインストールしたときにアプリケーションを始めるのに役立つ例示データなどです。
接続
アプリケーションのデプロイメントでは、ネットワーク構成（ネットワークポートなど）が指定される場合があります。接続の暗号化や認証に使用される証明書やキーなどの要素も含まれる場合があります。
設定パラメータ
アプリケーションのデプロイメントパッケージは、設定パラメータを設定する場合があります。これは、設定ファイルをサーバーにコピーするか、設定をレジストリにプッシュすることによって行われる場合があります。

アプリケーションとインフラストラクチャの境界線を異なる位置に引くことができます。必要なライブラリをアプリケーションのデプロイメントパッケージに含めるか、インフラストラクチャの一部としてプロビジョニングするかは、例です。たとえば、コンテナイメージには通常、ほとんどのオペレーティングシステムとそれが実行されるアプリケーションが含まれます。イミュータブルサーバーやイミュータブルスタックは、さらにこれを進化させ、アプリケーションとインフラストラクチャを1つのエンティティに組み合わせます。スペクトラムのもう一方の端では、インフラストラクチャコードが特定のアプリケーションのライブラリや設定ファイルをプロビジョニングし、アプリケーションパッケージ自体にはほとんど含まれていないこともあります。この質問は、コードベースの組織化方法にも関係しています。アプリケーションコードとインフラストラクチャコードを一緒に保持するか、別々に保持するかは、この本の後半（第18章を参照）で探求される問題ですが、1つの原則として、デプロイ可能な部分に構造を合わせるのが良いアイデアです。

デプロイメントパッケージ
アプリケーションは、デプロイメントパッケージに組織化されることがよくありますが、パッケージの形式はランタイム環境のタイプによって異なります。デプロイメントパッケージの形式と関連するランタイムの例は、以下の表に示されています（表10-1）。

表10-1. ターゲットランタイムとアプリケーションパッケージの形式の例

ターゲットランタイム	例のパッケージ形式
サーバーオペレーティングシステム	Red Hat RPMファイル、Debian .debファイル、WindowsのMSIインストーラーパッケージ
言語ランタイムエンジン	Ruby gems、Python pipパッケージ、Javaの.jar、.war、.earファイル
コンテナランタイム	Dockerイメージ
アプリケーションクラスタ	Kubernetesデプロイメントデスクリプタ、Helmチャート
FaaSサーバーレス	Lambdaデプロイメントパッケージ

デプロイメントパッケージの形式は、デプロイメントツールやランタイムがアプリケーションの部分を抽出し、適切な位置に配置するための標準です。