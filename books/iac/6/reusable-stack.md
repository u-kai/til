Pattern: Reusable Stack A reusable stack is an infrastructure source code project that is used to create multiple instances of a stack (Figure 6-5).
Motivation You create a reusable stack to maintain multiple consistent instances of infrastructure. When you make changes to the stack code, you can apply and test it in one instance, and then use the same code version to create or update multiple additional instances. You want to provision new instances of the stack with minimal ceremony, maybe even automatically. As an example, the ShopSpinner team extracted common code from different stack projects that each use an application server. Team members put the common code into a module used by each of the stack projects. Later, they realized that the stack projects for their customer applications still looked very similar. In addition to using the module to create an application server, each stack had code to create databases and dedicated logging and reporting services for each customer. Changing and testing changes to this code across multiple customers was becoming a hassle, and ShopSpinner was signing up new customers every month. So the team decided to create a single stack project that defines a customer application stack. This project still uses the shared Java application server module, as do a few other applications (Jira and GoCD). But the project also has the code for setting up the rest of the per-customer infrastructure as well. Now, when they sign on a new customer, the team members use the common customer stack project to create a new instance. When they fix or improve something in the project codebase, they apply it to test instances to make sure it’s OK, and then they roll it out one by one to the customer instances. Applicability You can use a reusable stack for delivery environments or for multiple production environments. This pattern is useful when you don’t need much variation between the environments. It is less applicable when environments need to be heavily customized. Consequences The ability to provision and update multiple stacks from the same project enhances scalability, reliability, and throughput. You can manage more instances with less effort, make changes with a lower risk of failure, and roll changes out to more systems more rapidly. You typically need to configure some aspects of the stack differently for different instances, even if it’s just what you name things. I’ll spend a whole chapter talking about this (Chapter 7). You should test your stack project code before you apply changes to business-critical infrastructure. I’ll spend multiple chapters on this, including Chapters 8 and 9. Implementation
You create a reusable stack as an infrastructure stack project, and then run the stack management tool each time you want to create or update an instance of the stack. Use the syntax of the stack tool command to tell it which instance you want to create or update. With Terraform, for example, you would specify a different state file or workspace for each instance. With CloudFormation, you pass a unique stack ID for each instance.
The following example command provisions two stack instances from a single project using a fictional command called stack. The command takes an argument env that identifies unique instances. As a rule, you should use simple parameters to define differences between stack instances — strings, numbers, or in some cases, lists. Additionally, the infrastructure created by a reusable stack should not vary much across instances. Related patterns The reusable stack is an improvement on the copy-paste environment antipattern (see “Antipattern: Copy-Paste Environments”), making it easier to keep multiple instances consistent. The wrapper stack pattern (see “Pattern: Wrapper Stack”) uses stack components to define a reusable stack, but uses a different stack project to set parameter values for each instance.
パターン：再利用可能なスタック
再利用可能なスタックは、複数のスタックのインスタンスを作成するために使用されるインフラソースコードプロジェクトです（図6-5）。

動機
再利用可能なスタックを作成することで、インフラの複数の一貫したインスタンスを維持します。スタックのコードを変更した場合、1つのインスタンスでそれを適用およびテストし、同じコードバージョンを使用して複数の追加のインスタンスを作成または更新できます。最小限の手続きでスタックの新しいインスタンスをプロビジョニングしたい、おそらく自動的にしたいと考えています。例として、ShopSpinnerチームは、アプリケーションサーバーを使用する異なるスタックプロジェクトから共通のコードを抽出しました。チームメンバーは共通のコードを、各スタックプロジェクトで使用されるモジュールに格納しました。後で、彼らは顧客アプリケーションのスタックプロジェクトがまだ非常に似ていることに気付きました。各スタックはアプリケーションサーバーを作成するためのモジュールの使用に加えて、顧客ごとにデータベースと専用のログ記録およびレポートサービスを作成するためのコードも持っていました。このコードの変更とテストを複数の顧客に対して行うのは手間がかかり、ShopSpinnerは毎月新しい顧客を獲得していました。そのため、チームは顧客アプリケーションスタックを定義する単一のスタックプロジェクトを作成することにしました。このプロジェクトは、共有のJavaアプリケーションサーバーモジュールを使用し続けます。また、他のアプリケーション（JiraおよびGoCD）も同様です。しかし、このプロジェクトには、顧客ごとのインフラストラクチャを設定するためのコードも含まれています。今では、新しい顧客がサインアップする際に、チームメンバーは共通の顧客スタックプロジェクトを使用して新しいインスタンスを作成します。プロジェクトのコードベースで何か修正または改善を行った場合は、テストインスタンスに適用してOKかどうかを確認し、次に顧客インスタンスに順次展開します。

適用範囲
再利用可能なスタックは、デリバリー環境または複数のプロダクション環境に使用できます。環境間の変動がほとんどない場合に有用です。環境を大幅にカスタマイズする必要がある場合は、適用範囲が低いです。

結果
同じプロジェクトから複数のスタックをプロビジョニングおよび更新できる能力は、スケーラビリティ、信頼性、およびスループットを向上させます。より少ない労力でより多くのインスタンスを管理し、失敗のリスクを低く変更し、変更をより迅速にシステムに展開できます。通常、さまざまなインスタンスに対してスタックのいくつかの側面を異なるように設定する必要があります。ただし、それが名前の変更である場合でもです。この問題については、後の章で詳しく説明します（第7章）。ビジネスクリティカルなインフラストラクチャに変更を適用する前に、スタックプロジェクトのコードをテストする必要があります。この問題については、第8章および第9章で詳しく説明します。

実装
再利用可能なスタックは、インフラストラクチャスタックプロジェクトとして作成し、スタック管理ツールを実行することで、スタックのインスタンスを作成または更新するたびに使用します。スタックツールコマンドの構文を使用して、作成または更新するインスタンスを指定します。たとえば、Terraformでは、各インスタンスに異なるステートファイルまたはワークスペースを指定します。CloudFormationでは、各インスタンスに一意のスタックIDを渡します。

次の例のコマンドは、stackという架空のコマンドを使用して、単一のプロジェクトから2つのスタックインスタンスをプロビジョニングしています。コマンドは、一意のインスタンスを識別する引数envを取ります。原則として、異なるスタックインスタンスの違いを定義するために、シンプルなパラメーターを使用する必要があります- 文字列、数値、または場合によってはリストです。また、再利用可能なスタックによって作成されるインフラストラクチャは、インスタンスごとにあまり変化しないはずです。

関連パターン
再利用可能なスタックは、コピー＆ペースト環境のアンチパターン（「アンチパターン：コピー＆ペースト環境」参照）の改善です。これにより、複数のインスタンスを一貫させることが容易になります。ラッパースタックパターン（「パターン：ラッパースタック」参照）は、スタックコンポーネントを使用して再利用可能なスタックを定義しますが、各インスタンスのパラメータ値を設定するために異なるスタックプロジェクトを使用します。