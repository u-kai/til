Configuring Servers in a Stack Infrastructure codebases for systems that aren’t fully container-based or serverless application architecture tend to include much code to provision and configure servers. Even container-based systems need to build host servers to run containers. The first mainstream Infrastructure as Code tools, like CFEngine, Puppet, and Chef, were used to configure servers. You should decouple code that builds servers from code that builds stacks. Doing this makes the code easier to understand, simplifies changes by decoupling them, and supports reusing and testing server code. Stack code typically specifies what servers to create, and passes information about the environment they will run in, by calling a server configuration tool. Example 5-2 is an example of a stack definition that calls the fictitious servermaker tool to configure a server.
スタック内のサーバーの設定 完全にコンテナベースやサーバーレスのアプリケーションアーキテクチャではないシステムのためのインフラストラクチャコードベースには、サーバーのプロビジョニングや設定のための多くのコードが含まれる傾向があります。コンテナベースのシステムでも、コンテナを実行するためのホストサーバーを構築する必要があります。CFEngine、Puppet、Chefなどの最初の主流のインフラストラクチャコードツールは、サーバーの設定に使用されました。サーバーを構築するコードとスタックを構築するコードを切り離すべきです。これにより、コードが理解しやすくなり、変更が切り離されることで簡素化され、サーバーコードの再利用やテストがサポートされます。スタックコードは通常、どのサーバーを作成するかを指定し、サーバーの設定ツールを呼び出して実行環境に関する情報を渡します。例5-2は、架空のservermakerツールを呼び出してサーバーを設定するスタックの定義の例です。