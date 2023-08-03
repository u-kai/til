With this arrangement, you can direct part of your workload to each of the cluster stacks. You can divide the workload in different ways: Workload percentage Direct part of the workload to each stack. Usually, the old stack handles most of the load at first, with the new stack taking a small percentage to evaluate how well it works. As the new stack beds in, you can dial the load up over time. After the new stack successfully manages 100% of the load and everyone is ready, you can decommission the old stack. This option assumes the new stack has all of the capabilities of the old stack, and that there aren’t any issues with data or messaging being split across the stacks. Service migration Migrate services one by one to the new cluster. Workloads in the main stack, such as network connections or messages, are directed to whichever stack instance the relevant service is running on. This option is especially useful when you need to modify service applications to move them to the new stack. It often requires more complex integration, perhaps even between the old and new cluster stacks. This complexity may be justified for migrating a complex service portfolio.6 User partitioning In some cases, different sets of users are directed to different stack implementations. Testers and internal users are often the first group. They can conduct exploratory tests and exercise the new system before risking “real” customers. In some cases, you might follow this by giving access to customers who opt-in to alpha testing or preview services. These cases make more sense when the service running on the new stack has changes that users will notice. Running new and old parts of the system conditionally, or in parallel, is a type of branch by abstraction. Progressively shifting portions of a workload onto new parts of a system is a canary release. Dark launching describes putting a new system capability into production, but not exposing it to production workloads, so your team can test it.
Consumer stacks, including application-infrastructure-stack, integrate with the single VLAN managed by the networking stack using one of the discovery methods described in “Discovering Dependencies Across Stacks”. The shared-networking-stack code exports the VLAN identifier for its consumer stacks to discover: vlans:

- main_vlan
  address_range: 10.2.0.0/8

export:

- main_vlan: main_vlan.id The new version of shared-networking-stack creates three VLANs and exports their identifiers under new names. It also exports one of the VLAN identifiers using the old identifier: vlans:
- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id

# Deprecated

- main_vlan: appserver_vlan_A.id
  By keeping the old identifier, the modified networking stack still works for consumer infrastructure code. The consumer code should be modified to use the new identifiers, and once all of the dependencies on the old identifier are gone, it can be removed from the networking stack code.

この方法では、ワークロードの一部を各クラスタースタックに直接割り当てることができます。ワークロードは異なる方法で分割することができます：ワークロードの割合、ワークロードの一部を各スタックに直接割り当てます。通常、古いスタックは最初はほとんどの負荷を処理し、新しいスタックはその機能を評価するために少しの割合を担当します。新しいスタックが定着したら、徐々に負荷を増やすことができます。新しいスタックが負荷を 100％正常に管理し、全員が準備ができたら、古いスタックを廃止することができます。このオプションでは、新しいスタックが古いスタックのすべての機能を持っていると仮定し、スタック間でデータやメッセージが分割される問題がないとします。サービスの移行では、1 つずつサービスを新しいクラスターに移行します。ネットワーク接続やメッセージなどのメインスタックのワークロードは、関連するサービスが実行されているスタックインスタンスに直接向けられます。このオプションは、サービスアプリケーションを修正して新しいスタックに移動する必要がある場合に特に便利です。これには、古いスタックと新しいスタック間のより複雑な統合が必要な場合もあります。この複雑さは、複雑なサービスポートフォリオを移行する際に正当化される場合があります。一部の場合では、異なるユーザーセットが異なるスタック実装に誘導されます。テスターや内部ユーザーが最初のグループになることがよくあります。彼らはリスクを冒さずに新しいシステムを試験的にテストし、実際の顧客にリスクを冒した後でそれを行うことができます。一部の場合では、アルファテストやプレビューサービスにオプトインする顧客へのアクセスを提供することができます。この場合、新しいスタック上で実行されているサービスにユーザーが気付く変更がある場合により意味があります。システムの新しい部分と古い部分を条件付けて、または並行して実行することは、抽象化によるブランチの一種です。ワークロードの一部をシステムの新しい部分に逐次的にシフトすることは、カナリアリリースと呼ばれます。ダークランチは、新しいシステムの機能を本番環境に導入するが、本番ワークロードには公開しない状態ですので、チームがそれをテストすることができます。

アプリケーション・インフラストラクチャ・スタックを含むコンシューマースタックは、Networking スタックが管理する単一の VLAN にディスカバリーメソッドの 1 つを使用して統合されます。共有 Networking スタックのコードは、そのコンシューマースタックがディスカバーするための VLAN 識別子をエクスポートします。次のようになります：

vlans:

- main_vlan
  address_range: 10.2.0.0/8

export:

- main_vlan: main_vlan.id

shared-networking-stack の新バージョンでは、3 つの VLAN を作成し、それぞれの識別子を新しい名前の下でエクスポートします。また、そのうち 1 つの VLAN 識別子を古い識別子を使用してエクスポートします：

vlans:

- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id

# Deprecated

- main_vlan: appserver_vlan_A.id

古い識別子を残すことで、修正された Networking スタックは引き続きコンシューマーのインフラストラクチャコードに対して機能します。コンシューマーコードは新しい識別子を使用するように修正する必要があり、古い識別子への依存関係がすべてなくなったら、それが Networking スタックのコードから削除されることができます。

- 考え方まとめ
  - ワークロードっていうのがいまいちわからんのよな
    - 定義されたプロセスを集合的にサポートする IT 資産(サーバー，VM，アプリケーション，データ)などの集合らしい
  - ダークローンチは初めて知った
