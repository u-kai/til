Pipeline for a Monolithic Application Cluster Stack Since there is only one stack, a single pipeline can test and deliver code changes to instances of the application cluster. However, there are other elements involved, including the server image for the host nodes and the applications themselves. Figure 14-3 shows a potential design for these pipelines.
The top pipeline (Figure 14-4) builds a server image for the host nodes, as described in “Using a Pipeline to Test and Deliver a Server Image”. The result of this pipeline is a server image that has been tested in isolation. The tests for that image probably check that the container management software is installed and that it complies with security policies. Figure 14-4. The pipeline for the host node server image The bottom pipeline (Figure 14-5) is for an application that deploys onto the cluster. In practice, you’ll have a number of these, one for every separately deployed application. This pipeline includes at least one early stage to build and test the application on its own. It then has stages that deploy the application to the cluster in each environment. The application can be tested, reviewed, and made available for production use in these environments. The pipelines for the applications are very loosely coupled with the pipelines for the cluster instances. You may choose to trigger application testing stages after cluster updates. Doing this helps you to find any issues that the change to the cluster causes for the application by running the application-specific tests.
The pipeline for the application cluster stack in Figure 14-6 starts with an offline stage (“Offline Testing Stages for Stacks”) that runs some syntax checking, and applies the stack code to a local mock of the infrastructure platform (“Testing with a Mock API”). These tests can catch problems at the coding level, without needing to use infrastructure platform resources, so they run quickly.
The second stage of this pipeline is an online stage (see “Online Testing Stages for Stacks”), and creates an instance of the stack on the infrastructure platform. The instance might be persistent (see “Pattern: Persistent Test Stack”) or ephemeral (see “Pattern: Ephemeral Test Stack”). The tests in this stage can check that the cluster management services are correctly created and accessible. You can also test for security issues — for instance, making sure that access to cluster management endpoints is locked down.3 Because this monolithic cluster stack includes the code to create the host node servers, the online testing stage can test these as well. A test could deploy a sample application to the cluster and prove that it’s working. The advantage of using a sample application rather than a real application is that you can keep it simple. Strip it down to a minimal set of dependencies and configurations, so you can be sure any test failures are caused by issues with the cluster provisioning, rather than due to complexities of deploying a real-world application.
Note that this pipeline stage is heavy. It tests both the cluster configuration and the host node server cluster. It also tests the server image in the context of the cluster. There are plenty of different things that can cause this stage to fail, which complicates troubleshooting and fixing failures. Most of the elapsed time for this stage will almost certainly be taken up by provisioning everything, far more than the time to execute the tests. These two issues — the variety of things tested in this one stage, and the time it takes to provision — are the main drivers for breaking the cluster into multiple stacks.
モノリシックなアプリケーションクラスタースタックのパイプライン
スタックが1つしかないため、単一のパイプラインでアプリケーションクラスターのインスタンスに対するコードの変更をテストしてデリバリーすることができます。ただし、ホストノードのサーバーイメージやアプリケーション自体など、他の要素も関与しています。図14-3は、これらのパイプラインの潜在的なデザインを示しています。
トップのパイプライン（図14-4）は、ホストノード用のサーバーイメージをビルドします。これは、「パイプラインを使用してサーバーイメージをテストおよびデリバリーする」という項で説明されています。このパイプラインの結果は、独立してテストされたサーバーイメージです。そのイメージのテストでは、コンテナ管理ソフトウェアがインストールされているかどうか、セキュリティポリシーに準拠しているかどうかなどが確認されるでしょう。図14-4. ホストノードサーバーイメージのパイプライン
下のパイプライン（図14-5）は、クラスターに展開されるアプリケーション向けです。実際には、個別に展開されるアプリケーションごとに複数のパイプラインがあります。このパイプラインには、アプリケーションを単体でビルドしてテストするための少なくとも1つの初期ステージが含まれています。その後、各環境にアプリケーションを展開するステージがあります。これらの環境でアプリケーションをテストし、レビューし、本番使用に利用できるようにすることができます。アプリケーションのパイプラインは、クラスターインスタンスのパイプラインと非常に疎結合です。クラスターの更新後にアプリケーションのテストステージをトリガーすることを選択することもできます。これにより、クラスターの変更によりアプリケーションに何らかの問題が生じるかどうかを、アプリケーション固有のテストを実行することで見つけることができます。
図14-6のアプリケーションクラスタースタックのパイプラインは、「スタックのオフラインテストステージ」と呼ばれるオフラインステージで始まります。このステージでは、構文チェックが行われ、スタックコードがインフラストラクチャプラットフォームのローカルモックに適用されます。「モックAPIを使用したテスト」として、インフラストラクチャプラットフォームのリソースを使用せずにコーディングレベルで問題を検出することができるため、実行時間が短縮されます。
このパイプラインの2つ目のステージはオンラインステージです（「スタックのオンラインテストステージ」を参照）。このステージでは、スタックのインスタンスがインフラストラクチャプラットフォーム上に作成されます。このインスタンスは永続的な場合もあります（「パターン：永続的なテストスタック」）、一時的な場合もあります（「パターン：一時的なテストスタック」）。このステージのテストでは、クラスター管理サービスが正しく作成され、アクセス可能であることを確認できます。また、セキュリティの問題もテストできます。たとえば、クラスター管理エンドポイントへのアクセスが制限されていることを確認します。このモノリシックなクラスタースタックには、ホストノードサーバーを作成するためのコードも含まれているため、オンラインテストステージでこれらのテストも行うことができます。テストでは、サンプルアプリケーションをクラスターに展開し、動作を確認することができます。実際のアプリケーションではなく、サンプルアプリケーションを使用する利点は、シンプルに保つことができるためです。依存関係と設定の最小限のセットに削減することで、テストの失敗がクラスタープロビジョニングの問題に起因するか、リアルワールドのアプリケーションの複雑さに起因するかを確実にすることができます。
このパイプラインステージは非常に重いことに注意してください。この1つのステージでテストされる項目はさまざまであり、障害のトラブルシューティングと修正を複雑化させます。このステージの経過時間のほとんどは、おそらくテストの実行時間よりもプロビジョニングにかかるでしょう。これら2つの問題、つまりこの1つのステージでテストされるさまざまな要素とプロビジョニングにかかる時間、これがクラスターを複数のスタックに分割する主な要因です。