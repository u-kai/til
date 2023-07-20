Server Roles As mentioned earlier, a server role is a way to define a group of server configuration modules to apply to a server. A role may also set some default parameters. For example, you may create an application-server role: role: application-server
server_modules: - tomcat - monitoring_agent - logging_agent - network_hardening
parameters: - inbound_port: 8443 Assigning this role to a server applies the configuration for Tomcat, monitoring and logging agents, and network hardening. It also sets a parameter for an inbound port, presumably something that the network_hardening module uses to open that port in the local firewall while locking everything else down. Roles can become messy, so you should be deliberate and consistent in how you use them. You may prefer to create fine-grained roles that you combine to compose specific servers. You assign multiple roles to a given server, such as ApplicationServer, MonitoredServer, and PublicFacingServer. Each of these roles includes a small number of server modules for a narrow purpose. Alternatively, you may have higher-level roles, which include more modules. Each server probably has only one role, which may be quite specific; for example, ShoppingServiceServer or JiraServer. A common approach is to use role inheritance. You define a base role that includes software and configuration common to all servers. This role might have networking hardening, administrative user accounts, and agents for monitoring and logging. More specific roles, such as for application servers or container hosts, include the base role and then add a few additional server configuration modules and parameters:
role: base-role
server_modules: - monitoring_agent - logging_agent - network_hardening

role: application-server
include_roles: - base_role
server_modules: - tomcat
parameters: - inbound_port: 8443

role: shopping-service-server
include_roles: - application-server
server_modules: - shopping-service-application This code example defines three roles in a hierarchy. The shopping-service-server role inherits everything from the application-server role, and adds a module to install the specific application to deploy. The application-server role inherits from the base-role, adding the Tomcat server and network port configuration. The base-role defines a core set of general-purpose configuration modules.
サーバーロール
前述したように、サーバーロールはサーバに適用されるサーバー設定モジュールのグループを定義する方法です。ロールはいくつかのデフォルトパラメーターを設定することもできます。例えば、アプリケーションサーバーの役割を作成することができます。

役割: アプリケーションサーバー
サーバーモジュール:
- tomcat
- monitoring_agent
- logging_agent
- network_hardening
パラメーター:
- inbound_port: 8443

この役割をサーバーに割り当てると、Tomcat、モニタリングエージェント、ログエージェント、およびネットワークの強化の設定が適用されます。また、入力ポートのパラメーターも設定されます。おそらく、ネットワークの強化モジュールがローカルのファイアウォールでそのポートを開き、他のすべてをロックダウンするために使用するものです。ロールは複雑になる可能性があるため、使用方法については慎重で一貫性を持って使用する必要があります。特定のサーバーを構成するために組み合わせるために、細かく分割された役割を作成することが好ましい場合もあります。ApplicationServer、MonitoredServer、PublicFacingServerなどの複数の役割を特定のサーバーに割り当てることができます。これらの役割のそれぞれは、特定の目的のためのわずかな数のサーバーモジュールを含みます。また、より大規模な役割もあるかもしれません。それぞれのサーバーにはおそらく1つの役割しかないでしょう。たとえば、ShoppingServiceServerやJiraServerなどです。

一般的なアプローチは、役割の継承を使用することです。すべてのサーバーに共通のソフトウェアと設定を含む基本的な役割を定義します。この役割には、ネットワークの強化、管理ユーザーアカウント、モニタリングエージェント、ログエージェントなどが含まれます。より具体的な役割（アプリケーションサーバーやコンテナホストなど）は、基本的な役割を含み、さらにいくつかのサーバー設定モジュールとパラメーターが追加されます。

役割: 基本的な役割
サーバーモジュール:
- monitoring_agent
- logging_agent
- network_hardening

役割: アプリケーションサーバー
include_roles:
- base_role
サーバーモジュール:
- tomcat
パラメーター:
- inbound_port: 8443

役割: ShoppingServiceServer
include_roles:
- application-server
サーバーモジュール:
- shopping-service-application

このコード例では、ヒエラルキーに基づいて3つの役割が定義されます。shopping-service-server役割は、application-server役割のすべてを継承し、特定のアプリケーションを展開するためのモジュールを追加します。application-server役割は、base-roleを継承し、Tomcatサーバーとネットワークポートの設定も追加します。base-roleは一般的な設定モジュールのコアセットを定義します。