Data Continuity in a Changing System Many Cloud Age practices and techniques for deploying software and managing infrastructure cheerfully recommend the casual destruction and expansion of resources, with only a hand wave to the problem of data. You can be forgiven for thinking that DevOps hipsters consider the whole idea of data to be a throwback to the Iron Age — a proper twelve-factor application is stateless, after all. Most systems out in the real world involve data, and people can be touchingly attached to it. Data can present a challenge when incrementally changing a system, as described in “Pushing Incomplete Changes to Production”. Running parallel instances of storage infrastructure may create inconsistencies or even corrupt your data. Many approaches to incrementally deploying changes rely on being able to roll them back, which might not be possible with data schema changes. Dynamically adding, removing, and rebuilding infrastructure resources that host data is particularly challenging. However, there are ways to manage it, depending on the situation. Some approaches include lock, segregate, replicate, and reload. Lock Some infrastructure platforms and stack management tools allow you to lock specific resources so they won’t be deleted by commands that would otherwise destroy them. If you specify this setting for a storage element, then the tool will refuse to apply changes to this element, allowing team members to make the change manually. There are a few issues with this, however. In some cases, if you apply a change to a protected resource, the tool may leave the stack in a partially modified state, which can cause downtime for services.
But the fundamental problem is that protecting some resources from automated changes encourages manual involvement in changes. Manual work invites manual mistakes. It’s much better to find a way to automate a process that changes your infrastructure safely. Segregate You can segregate data by splitting the resources that host them from other parts of the system; for example, by making a separate stack (an example of this is given in “Pattern: Micro Stack”). You can destroy and rebuild a compute instance with impunity by detaching and reattaching its disk volume. Keeping data in a database gives even more flexibility, potentially allowing you to add multiple compute instances. You still need a data continuity strategy for the stack that hosts the data, but it narrows the problem’s scope. You may be able to offload data continuity completely using a hosted DBaaS service. Replicate Depending on the data and how it’s managed, you may be able to replicate it across multiple instances of infrastructure. A classic example is a distributed database cluster that replicates data across nodes. With the right replication strategy, data is reloaded to a newly rebuilt node from other nodes in the cluster. This strategy fails if too many nodes are lost, which can happen with a major hosting outage. So this approach works as the first line of defense, with another mechanism needed for harder failure scenarios. Reload The best-known data continuity solution is backing up and restoring data from more reliable storage infrastructure. When you rebuild the infrastructure that hosts data, you first back the data up. You reload the data to the new instance after creating it. You may also take periodic backups, which you can reload in recovery scenarios, although you will lose any data changes that occurred between the backup and the recovery. This can be minimized and possibly eliminated by streaming data changes to the backup, such as writing a database transaction log.
Cloud platforms provide different storage services, as described in “Storage Resources”, with different levels of reliability. For example, object storage services like AWS S3 usually have stronger guarantees for the durability of data than block storage services like AWS EBS. So you could implement backups by copying or streaming data to an object storage volume. You should automate the process for not only backing up data but also for recovering it. Your infrastructure platform may provide ways to easily do this. For example, you can automatically take a snapshot of a disk storage volume before applying a change to it. You may be able to use disk volume snapshots to optimize the process of adding nodes to a system like a database cluster. Rather than creating a new database node with an empty storage volume, attaching it to a clone of another node’s disk might make it faster to synchronize and bring the node online. “Untested backups are the same as no backups” is a common adage in our industry. You’re already using automated testing for various aspects of your system, given that you’re following Infrastructure as Code practices. So you can do the same thing with your backups. Exercise your backup restore process in your pipeline or as a chaos experiment, whether in production or not. Mixing Data Continuity Approaches The best solution is often a combination of segregate, replicate, and reload. Segregating data creates room to manage other parts of the system more flexibly. Replication keeps data available most of the time. And reloading data is a backstop for more extreme situations.

変化するシステムにおけるデータの継続性 クラウド時代のソフトウェアの展開やインフラストラクチャの管理に関する多くの手法では、データの問題を軽視し、リソースの破壊や拡張をおおらかに推奨しています。データの問題は、12要素アプリケーションは無状態であるべきであるという考えがあるため、デブオプスの極めて新しい考え方に見えるかもしれません。しかし、現実の世界のほとんどのシステムにはデータが関与しており、人々はそれに感情的な愛着を抱くこともあります。データはシステムを段階的に変更する際に課題となる場合があります。「未完了の変更を本番環境にデプロイする」という記事で説明されているように、ストレージインフラストラクチャのパラレルインスタンスを実行することで、データの整合性が失われたり、データが破損したりする可能性があります。変更を段階的に展開するための多くのアプローチは、データスキーマの変更にはロールバックできない場合があります。データをホストするインフラストラクチャリソースの動的な追加、削除、再構築は特に困難です。ただし、状況に応じて管理する方法はあります。ロック 一部のインフラストラクチャプラットフォームやスタック管理ツールでは、特定のリソースをロックして、それらが破壊されるコマンドによって削除されないようにすることができます。ストレージ要素にこの設定を指定すると、ツールはこの要素に対して変更を適用しないよう拒否します。これにより、チームメンバーは手動で変更を行うことができます。しかし、これにはいくつかの問題があります。一部の場合では、保護されたリソースに変更を適用すると、ツールがスタックを部分的に変更した状態にしてしまい、サービスのダウンタイムが発生する可能性があります。ただし、基本的な問題は、一部のリソースを自動変更から保護することで、変更に手動で関与することを奨励してしまうことです。手動作業は手動のミスを招きます。インフラストラクチャを安全に変更するプロセスを自動化する方法を見つける方がずっと良いでしょう。セグリゲート データをシステムの他の部分から分離することで、ホストするリソースを分離することができます。たとえば、「パターン: マイクロスタック」で説明されているように、別のスタックを作成することで、コンピューティングインスタンスを自由に破壊して再構築することができます。データをデータベースに格納することで、より柔軟になり、複数のコンピューティングインスタンスを追加することができる可能性もあります。ただし、データをホストするスタックにはデータの継続性戦略が必要ですが、問題の範囲は狭まります。ホステッドDBaaSサービスを使用してデータの継続性を完全にオフロードすることもできるかもしれません。レプリケート データとその管理方法によっては、複数のインフラストラクチャインスタンス上にデータをレプリケートすることができる場合があります。典型的な例は、データをノード間でレプリケートする分散データベースクラスタです。適切なレプリケーション戦略でデータはクラスタの他のノードから新しく再構築されたノードにロードされます。この戦略は、多数のノードが失われる場合（主要なホスティング障害など）には失敗します。したがって、このアプローチは最初の防御ラインとして機能し、より困難な障害シナリオには別のメカニズムが必要です。リロード 最もよく知られたデータの継続性の解決策は、データをより信頼性の高いストレージインフラストラクチャにバックアップし、データを作成した後に新しいインスタンスにリロードすることです。定期的なバックアップも取ることができ、回復シナリオでリロードすることができますが、バックアップと回復の間に発生したデータ変更は失われます。これは、データ変更をバックアップにストリーミングするなど、データ変更を最小限に抑えることができます。AWS S3などのオブジェクトストレージサービスは、AWS EBSなどのブロックストレージサービスよりもデータの耐久性に強い保証を提供しています。したがって、バックアップをオブジェクトストレージボリュームにデータをコピーしたりストリーミングしたりすることで実装することができます。バックアップだけでなく、データの回復も自動化する必要があります。インフラストラクチャプラットフォームによっては、これを簡単に行う方法が提供される場合があります。たとえば、ディスクストレージボリュームに変更を適用する前にそのスナップショットを自動的に取ることができます。ディスクボリュームのスナップショットを使用して、データベースクラスタのようなシステムへのノードの追加プロセスを最適化することもできます。空のストレージボリュームを持つ新しいデータベースノードを作成する代わりに、別のノードのディスクのクローンに添付することで、同期を迅速化し、ノードをオンラインにすることができます。 "未検証のバックアップはバックアップと同じです"という言葉は、私たちの業界ではよく言われる格言です。インフラストラクチャとしてのコードの実践に従って、さまざまなシステムの側面に対してすでに自動テストを使用しているため、バックアップにも同じことができます。パイプラインやカオス実験の中でバックアップの復元プロセスを実行してみてください。それが本番環境であろうとなかろうと、データの継続性のアプローチを混合する 最良の解決策は、しばしば差別化、レプリケート、リロードの組み合わせです。データの分離により、システムの他の部分をより柔軟に管理する余地が生まれます。レプリケーションにより、データはほとんどの時間利用可能な状態になります。そして、データのリロードは、より極端な状況のための最終手段です。