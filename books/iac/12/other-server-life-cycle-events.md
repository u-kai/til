Other Server Life Cycle Events Creating, changing, and destroying server instances makes up the basic life cycle for a server. But there are other interesting phases in the extended life cycle, including stopping and restarting a server (Figure 12-1), replacing a server, and recovering a failed server. Stopping and Restarting a Server Instance When most of our servers were physical devices plugged into an electrical supply, we would commonly shut them down to upgrade hardware, or reboot them when upgrading certain operating system components.
People still stop and reboot virtual servers, sometimes for the same reasons — to reconfigure virtual hardware or upgrade an OS kernel. We sometimes shut servers down to save hosting costs. For example, some teams shut down development and test servers during evenings or weekends if nobody uses them. If servers are easy to rebuild, however, many teams simply destroy servers when they aren’t in use, and create new servers when they’re needed again. They do this partly because it’s easy and may save a bit more money than simply shutting them down. But destroying and rebuilding servers rather than stopping and keeping them around also plays to the philosophy of treating servers like “cattle” rather than “pets” (as mentioned in Cattle, Not Pets). Teams may stop and restart rather than rebuilding because they aren’t able to rebuild their servers with confidence. Often the challenge is retaining and restoring application data. See “Data Continuity in a Changing System” for ways to handle this challenge. So having a policy not to stop and restart servers forces your team to implement reliable processes and tools to rebuild servers and to keep them working. Having servers stopped can also complicate maintenance activities. Configuration updates and system patches won’t be applied to stopped servers. Depending on how you manage those types of updates, servers may receive them when starting again, or they may miss out. Replacing a Server Instance One of the many benefits of moving from physical to virtual servers is how easy it is to build and replace server instances. Many of the patterns and practices described in this book, including immutable servers (see “Pattern: Immutable Server”) and frequently updating server images, rely on the ability to replace a running server by building a new one (Figure 12-2).
The essential process for replacing a server instance is to create a new instance, validate its readiness, reconfigure other infrastructure and systems to use the new instance, test that it’s working correctly, and then destroy the old instance. Depending on the applications and systems that use the server instance, it may be possible to carry out the replacement without downtime, or at least with minimal downtime. You can find advice on this in “Changing Live Infrastructure”. Some infrastructure platforms have functionality to automate the server replacement process. For example, you might apply a configuration change to the definition of servers in an autoscaling server cluster, specifying that it should use a new server image that includes security patches. The platform automatically adds new instances, checks their health, and removes old ones. In other cases, you may need to do the work to replace servers yourself. One way to do this within a pipeline-based change delivery system is to expand and contract. You first push a change that adds the new server, then push a change that removes the old server afterward. See “Expand and Contract” for more details on this approach. Recovering a Failed Server Cloud infrastructure is not necessarily reliable. Some providers, including AWS, explicitly warn that they may terminate server instances without warning — for example, when they decide to replace the underlying hardware.5 Even providers with stronger availability guarantees have hardware failures that affect hosted systems. The process for recovering a failed server is similar to the process for replacing a server. One difference is the order of activities — you create the new server after destroying the old one rather than before (see Figure 12-3). The other difference is that you typically replace a server on purpose, while failure is less intentional.6 As with replacing a server instance, recovery may need a manual action, or you may be able to configure your platform and other services to detect and recover failures automatically.
他のサーバーのライフサイクルイベント サーバーの作成、変更、破棄はサーバーの基本的なライフサイクルを構成します。しかし、停止や再起動、サーバーの置き換え、故障したサーバーの回復など、拡張されたライフサイクルには他の興味深いフェーズもあります（図12-1）。 サーバーインスタンスの停止と再起動 物理的なサーバーが電源に接続されていた時代には、ハードウェアのアップグレードのために頻繁にシャットダウンしたり、特定のオペレーティングシステムコンポーネントのアップグレード時に再起動することがよくありました。
仮想サーバーでもそれと同じ理由でサーバーを停止したり再起動することがあります。仮想ハードウェアの再構成やOSカーネルのアップグレードのために停止させることもあります。ホスティング費用を節約するためにサーバーをシャットダウンすることもあります。たとえば、開発やテストのサーバーは利用されていない夜間や週末にチームがシャットダウンする場合もあります。ただし、再構築が容易な場合はサーバーをシンプルに破棄し、必要な時に新しいサーバーを作成することもあります。これは、シンプルに停止させるよりも少し節約になる可能性があるためです。ただし、サーバーを停止して保持する代わりに破棄して再構築することは、「ペット」ではなく「家畜」としてサーバーを扱うという考え方に合致するものです（Cattle, Not Petsで説明されています）。再構築できないために停止と再起動を選択する場合もあります。問題はアプリケーションデータの保持と復元です。この問題を処理する方法については、「変化するシステムでのデータの継続性」を参照してください。したがって、サーバーを停止して再起動しない方針を持つことで、チームは信頼性のあるプロセスとツールを実装し、サーバーを再構築して稼働させ続ける必要があります。 サーバーを停止させておくとメンテナンス作業が複雑化することもあります。設定の更新やシステムのパッチが停止したサーバーに適用されません。これらのタイプの更新をどのように管理するかによって、サーバーは再起動時に更新を受け取るか、あるいは受け取らないかが異なります。 サーバーインスタンスの置き換え 物理的なサーバーから仮想サーバーに移行した際の利点の一つは、実行中のサーバーを簡単に作成して置き換えることができることです。この本で説明されている多くのパターンとプラクティス、例えばイミュータブルサーバー（「パターン：イミュータブルサーバー」を参照）や頻繁なサーバーイメージの更新は、新しいサーバーを作成して実行中のサーバーを置き換える能力に依存しています（図12-2）。
サーバーインスタンスを置き換えるための基本的なプロセスは、新しいインスタンスを作成し、その準備が完了したことを検証し、他のインフラストラクチャやシステムを新しいインスタンスを使用するように再構成し、正しく機能していることをテストし、古いインスタンスを破棄することです。使用するアプリケーションやシステムによって異なりますが、置き換えをダウンタイムなし、あるいは少なくとも最小限のダウンタイムで実行できる場合もあります。これについてのアドバイスは「ライブインフラストラクチャの変更」で見つけることができます。一部のインフラストラクチャプラットフォームには、サーバーの置き換えプロセスを自動化する機能があります。たとえば、オートスケーリングサーバークラスタのサーバーの定義に構成変更を適用し、セキュリティパッチを含む新しいサーバーイメージを使用するよう指定することができます。プラットフォームは自動で新しいインスタンスを追加し、健全性をチェックし、古いインスタンスを削除します。その他の場合では、サーバーを置き換えるために自分で作業を行う必要があります。パイプラインベースの変更配信システム内でこれを行う方法の一つは、拡張と収縮です。まず新しいサーバーを追加する変更をプッシュし、その後に古いサーバーを削除する変更をプッシュします。このアプローチの詳細については、「拡張と収縮」を参照してください。 故障したサーバーの回復 クラウドインフラストラクチャは必ずしも信頼性があるわけではありません。一部のプロバイダ（AWSを含む）は、基礎となるハードウェアを置き換えることを決定した場合など、事前の警告なしにサーバーインスタンスを終了する可能性が明示的に警告しています。より高い可用性の保証を提供するプロバイダでも、ホストされたシステムに影響を与えるハードウェアの故障が発生することがあります。故障したサーバーの回復プロセスは、サーバーの置き換えのプロセスと類似しています。違いはアクティビティの順序です。旧サーバーを破棄した後に新サーバーを作成します（図12-3を参照）。もう一つの違いは、通常はサーバーを意図的に置き換えるが、故障は意図的ではないという点です。サーバーインスタンスの置き換えと同様に、回復には人の手による手動操作が必要な場合もありますし、プラットフォームや他のサービスを構成して故障を検出し自動的に回復させることもできる場合もあります。