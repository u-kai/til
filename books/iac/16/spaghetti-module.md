Antipattern: Spaghetti Module A spaghetti module is configurable to the point where it creates significantly different results depending on the parameters given to it. The implementation of the module is messy and difficult to understand, because it has too many moving parts (see Example 16-6). Example 16-6. Example of a spaghetti module declare module: application-server-infrastructure
variable: network_segment = {
if ${parameter.network_access} = "public"
id: public_subnet
else if ${parameter.network_access} = "customer"
id: customer_subnet
else
id: internal_subnet
end
}

switch ${parameter.application_type}:
"java":
virtual_machine:
origin_image: base_tomcat
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
"NET":
virtual_machine:
origin_image: windows_server
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
"php":
container_group:
cluster_id: ${parameter.container_cluster}
container_image: nginx_php_image
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
end

switch ${parameter.database}:
"mysql":
database_instance: my_database
type: mysql
...
... This example code assigns the server it creates to one of three different network segments, and optionally creates a database cluster and passes a connection string to the server configuration. In some cases, it creates a group of container instances rather than a virtual server. This module is a bit of a beast.
Motivation As with other antipatterns, people create a spaghetti module by accident, often over time. You may create a facade module (“Pattern: Facade Module”) or a bundle module (“Pattern: Bundle Module”), that grows in complexity to handle divergent use cases that seem similar on the surface. Spaghetti modules often result from trying to implement an infrastructure domain entity (“Pattern: Infrastructure Domain Entity”) using a declarative language. Consequences A module that does too many things is less maintainable than one with a tighter scope. The more things a module does, and the more variations there are in the infrastructure that it can create, the harder it is to change it without breaking something. These modules are harder to test. As I explain in Chapter 8, better-designed code is easier to test, so if you’re struggling to write automated tests and build pipelines to test the module in isolation, it’s a sign that you have a spaghetti module. Implementation A spaghetti module’s code often contains conditionals that apply different specifications in different situations. For example, a database cluster module might take a parameter to choose which database to provision. When you realize you have a spaghetti module on your hands, you should refactor it. Often, you can split it into different modules, each with a more focused remit. For example, you might decompose your single application infrastructure module into different modules for different parts of the application’s infrastructure. An example of a stack that uses decomposed modules in this way, rather than using the spaghetti module from Example 16-6, might look like Example 16-7. Example 16-7. Example of using decomposed modules rather than a single spaghetti module
application: "shopping_app"
application_version: "4.20"
network_segment: customer_subnet
server_configuration:
database_connection: ${module.mysql-database.outputs.connection_string}

use module: mysql-database
cluster_minimum: 1
cluster_maximum: 3
allow_connections_from: customer_subnet Each of the modules is smaller, simpler, and so easier to maintain and test than the original spaghetti module. Related patterns A spaghetti module is often an attempt at building an infrastructure domain entity using declarative code. It could also be a facade module (“Pattern: Facade Module”) or a bundle module (“Pattern: Bundle Module”) that people tried to extend to handle different use cases.
アンチパターン: スパゲッティモジュール
スパゲッティモジュールは、与えられたパラメータによって大幅に異なる結果を作り出すまで設定可能です。モジュールの実装は乱雑で理解しにくく、移動部品が多すぎるためです（例16-6を参照）。例16-6. スパゲッティモジュールの例
declare module: application-server-infrastructure
variable: network_segment = {
if ${parameter.network_access} = "public"
id: public_subnet
else if ${parameter.network_access} = "customer"
id: customer_subnet
else
id: internal_subnet
end
}

switch ${parameter.application_type}:
"java":
virtual_machine:
origin_image: base_tomcat
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
"NET":
virtual_machine:
origin_image: windows_server
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
"php":
container_group:
cluster_id: ${parameter.container_cluster}
container_image: nginx_php_image
network_segment: ${variable.network_segment}
server_configuration:
if ${parameter.database} != "none"
database_connection: ${database_instance.my_database.connection_string}
end
...
end

switch ${parameter.database}:
"mysql":
database_instance: my_database
type: mysql
...
... この例のコードは、作成するサーバを3つの異なるネットワークセグメントの1つに割り当て、必要に応じてデータベースクラスタを作成し、接続文字列をサーバ構成に渡します。一部の場合では、仮想サーバではなく、コンテナインスタンスのグループを作成します。このモジュールは少し大きなものです。
動機
他のアンチパターンと同様に、スパゲッティモジュールは偶然に作成されることが多くあります。表面的には似ているが異なるユースケースを処理するために複雑さが増していくファサードモジュール（「パターン：ファサードモジュール」）やバンドルモジュール（「パターン：バンドルモジュール」）を作成することがあります。スパゲッティモジュールは、宣言型言語を使用してインフラストラクチャドメインエンティティ（「パターン：インフラストラクチャドメインエンティティ」）を実装しようとすることによって生じることがよくあります。
結果
1つのモジュールが行うことが多すぎる場合、スコープが狭いモジュールよりも保守性が低下します。モジュールがすることが多く、作成するインフラストラクチャにバリエーションがあるほど、変更する際に何かを壊さずに済ますのは難しくなります。これらのモジュールはテストしづらいです。第8章で説明するように、設計がより良いコードはテストしやすいため、モジュールを分離して単体でテストするための自動化されたテストとビルドパイプラインを作成するのに苦労している場合、スパゲッティモジュールである可能性があります。
実装
スパゲッティモジュールのコードには、異なる状況において異なる仕様が適用される条件分岐がよく含まれます。例えば、データベースクラスタモジュールは、どのデータベースをプロビジョニングするかを選択するためのパラメータを受け取るかもしれません。スパゲッティモジュールになっていることに気付いた場合、リファクタリングする必要があります。多くの場合、より焦点を絞ったリミットを持つ異なるモジュールに分割することができます。例えば、アプリケーションのインフラストラクチャの異なる部分ごとに異なるモジュールに分解することができます。スパゲッティモジュールの例（例16-6から）ではなく、このように分解されたモジュールを使用するスタックの例は、例16-7のようになります。

関連するパターン
スパゲッティモジュールは、宣言的なコードを使用してインフラストラクチャドメインエンティティを構築しようとする試みであることがよくあります。また、ファサードモジュール（「パターン：ファサードモジュール」）やバンドルモジュール（「パターン：バンドルモジュール」）であり、異なるユースケースを処理するために拡張しようとしたモジュールでもある場合があります。