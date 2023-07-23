Use Infrastructure as Code to Optimize for Change Given that changes are the biggest risk to a production system, continuous change is inevitable, and making changes is the only way to improve a system, it makes sense to optimize your capability to make changes both rapidly and reliably.4 Research from the Accelerate State of DevOps Report backs this up. Making changes frequently and reliably is correlated to organizational success.5 There are several objections I hear when I recommend a team implement automation to optimize for change. I believe these come from misunderstandings of how you can and should use automation. Objection: “We don’t make changes often enough to justify automating them” We want to think that we build a system, and then it’s “done.” In this view, we don’t make many changes, so automating changes is a waste of time. In reality, very few systems stop changing, at least not before they are retired. Some people assume that their current level of change is temporary. Others create heavyweight change request processes to discourage people from asking for changes. These people are in denial. Most teams that are supporting actively used systems handle a continuous stream of changes. Consider these common examples of infrastructure changes: An essential new application feature requires you to add a new database. A new application feature needs you to upgrade the application server.
Usage levels grow faster than expected. You need more servers, new clusters, and expanded network and storage capacity. Performance profiling shows that the current application deployment architecture is limiting performance. You need to redeploy the applications across different application servers. Doing this requires changes to the clustering and network architecture. There is a newly announced security vulnerability in system packages for your OS. You need to patch dozens of production servers. You need to update servers running a deprecated version of the OS and critical packages. Your web servers experience intermittent failures. You need to make a series of configuration changes to diagnose the problem. Then you need to update a module to resolve the issue. You find a configuration change that improves the performance of your database. A fundamental truth of the Cloud Age is: Stablity comes from making changes. Unpatched systems are not stable; they are vulnerable. If you can’t fix issues as soon as you discover them, your system is not stable. If you can’t recover from failure quickly, your system is not stable. If the changes you do make involve considerable downtime, your system is not stable. If changes frequently fail, your system is not stable.
変更を最適化するためにインフラストラクチャのコードを使用する
プロダクションシステムへの変更が最も大きなリスクであることを考慮すると、持続的な変化は避けられず、システムを改善する唯一の方法は変更を行うことです。したがって、変更を迅速かつ信頼性を持って行う能力を最適化することは合理的です。Accelerate State of DevOps Reportの調査結果もこれを裏付けています。頻繁かつ信頼性のある変更は組織の成功と相関しています。私が変更の最適化のために自動化を導入することを勧めると、いくつかの異議が出されます。これらは自動化の使用方法に関する誤解から生じていると私は考えています。異議：「変更を自動化するのに十分な頻度で変更を行っていないため、自動化することを正当化できない」私たちはシステムを作り上げてから「完了」だと思いたいものです。この視点では、変更をあまり行わないため、変更の自動化は時間の無駄です。実際には、ほとんどのシステムは変化をやめません。いくつかの人々は、現在の変化のレベルが一時的であると仮定しています。他の人々は、変更の要求を desu、退けるために重い変更依頼プロセスを作成します。これらの人々は拒絶しています。現在使用中のシステムをサポートしているほとんどのチームは、連続する変更ストリームを処理しています。以下はインフラストラクチャの変更の一般的な例です。新しく必要なアプリケーション機能により新しいデータベースを追加する必要が生じる。新しいアプリケーション機能により、アプリケーションサーバーをアップグレードする必要がある。利用レベルが予想よりも速く増加しています。サーバー、新しいクラスター、およびネットワークおよびストレージ容量の拡張が必要です。パフォーマンスプロファイリングにより、現在のアプリケーションデプロイメントアーキテクチャがパフォーマンスを制限していることが判明しました。アプリケーションを異なるアプリケーションサーバーに再デプロイする必要があります。これにはクラスタリングおよびネットワークアーキテクチャの変更が必要です。自分のOSのシステムパッケージに新しく発表されたセキュリティの脆弱性が存在します。プロダクションサーバーのパッチを適用する必要があります。廃止予定のバージョンのOSおよび重要なパッケージを実行しているサーバーを更新する必要があります。Webサーバーに断続的な障害が発生しています。問題を診断するために一連の設定変更を行う必要があります。その後、問題を解決するためにモジュールを更新する必要があります。データベースのパフォーマンスを改善する設定変更が見つかりました。クラウド時代の基本的な真理は、変更を行うことで安定性が生まれるということです。パッチ適用されていないシステムは安定していません。それらは脆弱です。問題が発見されてすぐに修正できない場合、システムは安定していません。障害からの迅速な回復ができない場合、システムは安定していません。行われる変更によって長時間ダウンタイムが発生する場合、システムは安定していません。変更が頻繁に失敗する場合、システムは安定していません。