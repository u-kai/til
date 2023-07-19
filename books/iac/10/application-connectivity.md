Application Connectivity In addition to compute resources to execute code and storage resources to hold data, applications need networking for inbound and outbound connectivity. A server-oriented application package, such as for a web server, might configure network ports and add encryption keys for incoming connections. But traditionally, they have relied on someone to configure infrastructure outside the server separately. You can define and manage addressing, routing, naming, firewall rules, and similar concerns as part of an infrastructure stack project, and then deploy the application into the resulting infrastructure. A more cloud native approach to networking is to define the networking requirements as part of the application deployment manifest and have the application runtime allocate the resources dynamically. You can see this in part of Example 10-1, shown here: application:
name: nginx
connectivity:
inbound:
id: https_inbound
port: 443
allow_from: $PUBLIC_INTERNET
ssl_cert: $SHOPSPINNER_PUBLIC_SSL_CERT
outbound:
port: 443
allow_to: [ $APPSERVER_APPLICATIONS.https_inbound ] This example defines inbound and outbound connections, referencing other parts of the system. These are the public internet, presumably a gateway, and inbound HTTPS
ports for application servers, defined and deployed to the same cluster using their own deployment manifests. Application runtimes provide many common services to applications. Many of these services are types of service discovery.
アプリケーションの接続性
コードを実行するための計算リソースやデータを保持するためのストレージリソースに加えて、アプリケーションは受信および送信の接続のためにネットワーキングが必要です。Webサーバのようなサーバ志向のアプリケーションパッケージでは、ネットワークポートの設定や受信接続のための暗号化キーを追加することがあります。しかし、従来はサーバの外部のインフラを別途設定することに依存していました。アドレッシング、ルーティング、ネーミング、ファイアウォールルールなどの関連事項をインフラストラクチャスタックプロジェクトの一部として定義し、それに基づいてアプリケーションをデプロイできます。ネットワーキングに対するよりクラウドネイティブなアプローチは、アプリケーションデプロイメントマニフェストの一部としてネットワーキング要件を定義し、アプリケーションランタイムがリソースを動的に割り当てることです。以下に示すExample 10-1の一部でこれを確認できます。

アプリケーション:
  名前: nginx
  接続性:
    受信:
      ID: https_inbound
      ポート: 443
      Allow_from: $PUBLIC_INTERNET
      SSL証明書: $SHOPSPINNER_PUBLIC_SSL_CERT
    送信:
      ポート: 443
      Allow_to: [$APPSERVER_APPLICATIONS.https_inbound]

この例では、他のシステムの一部を参照する受信および送信接続が定義されています。これらは公共インターネット（おそらくゲートウェイ）と、独自のデプロイメントマニフェストを使用して同じクラスタに定義およびデプロイされたアプリケーションサーバの受信HTTPSポートです。アプリケーションランタイムはアプリケーションに多くの共通のサービスを提供します。これらのサービスの多くはサービスディスカバリの一種です。