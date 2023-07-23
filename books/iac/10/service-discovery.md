Service Discovery Applications and services running in an infrastructure often need to know how to find other applications and services. For example, a frontend web application may send requests to a backend service to process transactions for users. Doing this isn’t difficult in a static environment. Applications may use a known hostname for other services, perhaps kept in a configuration file that you update as needed. But with a dynamic infrastructure where the locations of services and servers are fluid, a more responsive way of finding services is needed. A few popular discovery mechanisms are: Hardcoded IP addresses Allocate IP addresses for each service. For example, the monitoring server always runs on 192.168.1.5. If you need to change the address or run multiple instances of a service (for example, for a controlled rollout of a major upgrade), you need to rebuild and redeploy applications. Hostfile entries Use server configuration to generate the /etc/hosts file (or equivalent) on each server, mapping service names to the current IP address. This method is a messier alternative to DNS, but I’ve seen it used to work around legacy DNS implementations.7 DNS (Domain Name System) Use DNS entries to map service names to their current IP address, either using DNS entries managed by code, or DDNS (Dynamic DNS). DNS is a mature, well-supported solution to the problem. Resource tags
Infrastructure resources are tagged to indicate what services they provide, and the context such as environments. Discovery involves using the platform API to look up resources with relevant tags. Care should be taken to avoid coupling application code to the infrastructure platform. Configuration registry Application instances can maintain the current connectivity details in a centralized registry (see “Configuration Registry”), so other applications can look it up. This may be helpful when you need more information than the address; for example, health or other status indications. Sidecar A separate process runs alongside each application instance. The application may use the sidecar as a proxy for outbound connections, a gateway for inbound connections, or as a lookup service. The sidecars will need some method to do network discovery itself. This method could be one of the other discovery mechanisms, or it might use a different communication protocol.8 Sidecars are usually part of a service mesh ( I discuss service meshes in “Service Mesh”), and often provide more than service discovery. For example, a sidecar may handle authentication, encryption, logging, and monitoring. API gateway An API gateway is a centralized HTTP service that defines routes and endpoints. It usually provides more services than this; for example, authentication, encryption, logging, and monitoring. In other words, an API gateway is not unlike a sidecar, but is centralized rather than distributed.9 Avoid Hard Boundaries Between Infrastructure, Runtime, and Applications In theory, it could be useful to provide application runtimes as a complete set of services to developers, shielding them from the details of the underlying infrastructure. In practice, the lines are much fuzzier than presented in this model. Different people and teams need access to resources at different levels of abstraction, with different levels of control. So you should not design and implement systems with absolute boundaries, but instead define pieces that can be composed and presented in different ways to different users.
インフラストラクチャ内で実行されるアプリケーションやサービスは、他のアプリケーションやサービスの場所を知る必要があります。たとえば、フロントエンドのWebアプリケーションは、ユーザーのトランザクションを処理するためにバックエンドのサービスにリクエストを送信するかもしれません。静的な環境ではこれは難しくありません。アプリケーションは、他のサービスのための既知のホスト名を使用する場合があります。必要に応じて更新するための設定ファイルに保持されているかもしれません。しかし、サービスやサーバーの場所が流動的な動的なインフラストラクチャでは、より迅速にサービスを見つける方法が必要です。いくつかの人気のある検出メカニズムは次のとおりです。ハードコードされたIPアドレス各サービスにIPアドレスを割り当てます。たとえば、モニタリングサーバーは常に192.168.1.5で実行されます。アドレスを変更する必要がある場合や、サービスの複数のインスタンス（たとえば、主要なアップグレードの制御された展開のため）を実行する場合は、アプリケーションを再ビルドして再展開する必要があります。ホストファイルのエントリサーバーの設定を使用して、各サーバー上の/etc/hostsファイル（または同等のファイル）を生成し、サービス名を現在のIPアドレスにマッピングします。この方法はDNSよりも乱雑な代替手段ですが、レガシーのDNS実装を回避するために使用されることがあります。DNS（ドメイン名システム）DNSエントリを使用してサービス名をその現在のIPアドレスにマッピングします。これは、コードによって管理されるDNSエントリまたはDDNS（動的DNS）を使用することができます。DNSは、問題に対する成熟した、よくサポートされた解決策です。リソースタグインフラストラクチャリソースには、提供するサービスや環境などのコンテキストを示すタグが付けられます。検出は、関連するタグを持つリソースを検索するためにプラットフォームAPIを使用して行われます。アプリケーションコードをインフラストラクチャプラットフォームに結合しないように注意する必要があります。設定レジストリアプリケーションインスタンスは、現在の接続詳細を集中管理されたレジストリ（「設定レジストリ」を参照）に保持することができます。これにより、他のアプリケーションがそれを検索できるようになります。アドレス以上の情報が必要な場合、たとえば、ヘルスや他のステータスの表示など、これは役に立つ場合があります。サイドカープロセスは、各アプリケーションインスタンスと一緒に実行される別のプロセスです。アプリケーションは、サイドカーをアウトバウンド接続のプロキシ、インバウンド接続のゲートウェイ、またはルックアップサービスとして使用することができます。サイドカーは、ネットワークの検出を行うための手法が必要です。この方法は他の検出メカニズムの1つかもしれませんし、異なる通信プロトコルを使用するかもしれません。サイドカーは通常、サービスメッシュの一部であり（「サービスメッシュ」でサービスメッシュについて説明しています）、サービスの検出以上の機能を提供することが多いです。たとえば、サイドカーは認証、暗号化、ログ記録、モニタリングを処理する場合があります。APIゲートウェイAPIゲートウェイは、ルートとエンドポイントを定義する集中化されたHTTPサービスです。通常、これ以上のサービスを提供します。たとえば、認証、暗号化、ログ記録、モニタリングなどです。つまり、APIゲートウェイは分散ではなく集中化されたサイドカーのようなものです。インフラストラクチャ、ランタイム、アプリケーション間の固定境界を回避するインフラストラクチャランタイムを開発者に完全なサービスセットとして提供し、彼らを基盤の詳細から保護することは理論的には有用であるかもしれません。しかし、実務では、このモデルで示されているよりもはるかに曖昧です。異なる人やチームは、異なる抽象化レベル、異なる制御レベルでリソースにアクセスする必要があります。したがって、絶対的な境界を持つシステムを設計・実装するべきではありません。代わりに、異なるユーザに異なる方法で構成および提示できるピースを定義するべきです。