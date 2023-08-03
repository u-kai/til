Expand and Contract Infrastructure teams use the expand and contract pattern (also called Parallel Change) for changing interfaces without breaking consumers. The idea is that changing a provider’s interface involves two steps: change the provider, then change the consumers. The expand and contract pattern decouples these steps.
The essence of the pattern is to first add the new resource while keeping the existing one, then change the consumers over to the new resource, and finally, remove the old unused resource. Each of these changes is delivered using a pipeline (see “Infrastructure Delivery Pipelines”), so it’s thoroughly tested. Making a change by expanding and contracting is similar to a backward compatible transformation (see “Backward Compatible Transformations”). That technique replaced the old resource and re-pointed the old interface to one of the new resources. However, applying the new code to a running instance would attempt to destroy the old resource, which could either disrupt any consumers attached to it or fail to complete. So a few extra steps are required. The first step for the ShopSpinner team to use expand and contract for its VLAN change is to add the new VLANs to the shared-networking-stack, while leaving the old main_vlan in place: vlans:

- main_vlan
  address_range: 10.2.0.0/8
- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- main_vlan: main_vlan.id
- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id Unlike the parallel instances technique (“Parallel Instances”) and infrastructure surgery (“Infrastructure Surgery”), the ShopSpinner team doesn’t add a second instance of the stack, but only changes the existing instance.
  After applying this code, the existing consumer instances are unaffected — they are still attached to the main_vlan. The team can add new resources to the new VLANs, and can make changes to the consumers to switch them over as well. How to switch consumer resources to use new ones depends on the specific infrastructure and the platform. In some cases, you can update the definition for the resource to attach it to the new provider interface. In others, you may need to destroy and rebuild the resource. The ShopSpinner team can’t reassign existing virtual server instances to the new VLANs. However, the team can use the expand and contract pattern to replace the servers. The application-infrastructure-stack code defines each server with a static IP address that routes traffic to the server: virtual_machine:
  name: appserver-${SERVICE}-A
  memory: 4GB
  vlan: external_stack.shared_network_stack.main_vlan

static_ip:
name: address-${SERVICE}-A
  attach: virtual_machine.appserver-${SERVICE}-A The team’s first step is to add a new server instance attached to the new VLAN: virtual_machine:
name: appserver-${SERVICE}-A2
memory: 4GB
vlan: external_stack.shared_network_stack.appserver_vlan_A

virtual_machine:
name: appserver-${SERVICE}-A
memory: 4GB
vlan: external_stack.shared_network_stack.main_vlan

static_ip:
name: address-${SERVICE}-A
  attach: virtual_machine.appserver-${SERVICE}-A The first virtual_machine statement in this code creates a new server instance named appserver-${SERVICE}-A2. The team’s pipeline delivers this change to each environment. The new server instance isn’t used at this point, although the team can add some automated tests to prove that it’s running OK. The team’s next step is to switch user traffic to the new server instance. The team makes another change to the code, modifying the static_ip statement: virtual_machine:
  name: appserver-${SERVICE}-A2
memory: 4GB
vlan: external_stack.shared_network_stack.appserver_vlan_A

virtual_machine:
name: appserver-${SERVICE}-A
memory: 4GB
vlan: external_stack.shared_network_stack.main_vlan

static_ip:
name: address-${SERVICE}-A
  attach: virtual_machine.appserver-${SERVICE}-A2 Pushing this change through the pipeline makes the new server active, and stops traffic to the old server. The team can make sure everything is working, and easily roll the change back to restore the old server if something goes wrong. Once the team has the new server working OK, it can remove the old server from the stack code:
virtual_machine:
name: appserver-${SERVICE}-A2
memory: 4GB
vlan: external_stack.shared_network_stack.appserver_vlan_A

static_ip:
name: address-${SERVICE}-A
  attach: virtual_machine.appserver-${SERVICE}-A2 Once this change has been pushed through the pipeline and applied to all environments, application-infrastructure-stack no longer has a dependency on main_vlan in shared-networking-stack. After all consumer infrastructure has changed over, the ShopSpinner team can remove main_vlan from the provider stack code: vlans:

- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id The VLAN change is complete, and the last remnants of main_vlan have been swept away.8

拡張／更新と収縮／削除インフラストラクチャーのチームは、インターフェースを変更しても消費者が壊れないようにする「拡張／更新と収縮／削除」（またはパラレル変更）パターンを使用しています。この考え方は、プロバイダーのインターフェースを変更するためには 2 つのステップが必要であることを示しています：プロバイダーを変更し、その後消費者を変更する。拡張／更新と収縮／削除パターンは、これらのステップを分離します。
このパターンの本質は、新しいリソースを最初に追加し、既存のものを保持したままにすることです。次に、消費者を新しいリソースに切り替え、最後に古い未使用のリソースを削除します。これらの変更は、それぞれの変更ごとにパイプライン（「インフラストラクチャーのデリバリーパイプライン」を参照）を使用して提供されるため、徹底的にテストされます。拡張／更新と収縮は、後方互換の変換（「後方互換の変換」を参照）に似ています。このテクニックは、古いリソースを置き換え、古いインターフェースを新しいリソースの 1 つに再指定します。ただし、実行中のインスタンスに新しいコードを適用すると、古いリソースを破壊しようとするか、それに接続された消費者を中断するか、完了しない可能性があります。そのため、いくつかの追加の手順が必要です。ShopSpinner チームが VLAN の変更に拡張／更新と収縮を使用するための最初のステップは、古い main_vlan をそのままにして新しい VLAN を共有ネットワーキングスタックに追加することです：vlans:

- main_vlan
  address_range: 10.2.0.0/8
- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- main_vlan: main_vlan.id
- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id「パラレルインスタンス」テクニック（「パラレルインスタンス」）やインフラストラクチャーサージェリー（「インフラストラクチャーサージェリー」）とは異なり、ShopSpinner チームはスタックの 2 つ目のインスタンスを追加せず、既存のインスタンスのみを変更します。
  このコードを適用した後、既存の消費者インスタンスは影響を受けません。まだ main_vlan に接続されています。チームは新しいリソースを新しい VLAN に追加し、消費者を切り替えるための変更を行うことができます。消費者リソースを新しいものに切り替える方法は、特定のインフラストラクチャーとプラットフォームに依存します。場合によっては、リソースの定義を更新して新しいプロバイダーインターフェースに接続することができます。他の場合、リソースを削除して再構築する必要があるかもしれません。ShopSpinner チームは既存の仮想サーバーインスタンスを新しい VLAN に再割り当てすることはできません。ただし、チームは拡張／更新と収縮パターンを使用してサーバーを置き換えることができます。アプリケーションインフラストラクチャースタックのコードでは、各サーバーに静的 IP アドレスが定義されており、トラフィックをサーバーにルーティングしています：virtual_machine:
  name: appserver-${SERVICE}-A
  memory: 4GB
  vlan: external_stack.shared_network_stack.main_vlan

static_ip:
name: address-${SERVICE}-A
attach: virtual_machine.appserver-${SERVICE}-A このコード内の最初の virtual_machine ステートメントは、appserver-${SERVICE}-A2という名前の新しいサーバーインスタンスを作成します。チームのパイプラインは、この変更を各環境に提供します。この時点では新しいサーバーインスタンスは使用されませんが、チームは実行されていることを証明するためにいくつかの自動テストを追加することができます。次に、チームはユーザートラフィックを新しいサーバーインスタンスに切り替えるために、コードを変更してstatic_ipステートメントを修正します：virtual_machine:
name: appserver-${SERVICE}-A2
memory: 4GB
vlan: external_stack.shared_network_stack.appserver_vlan_A

virtual_machine:
name: appserver-${SERVICE}-A
memory: 4GB
vlan: external_stack.shared_network_stack.main_vlan

static_ip:
name: address-${SERVICE}-A
attach: virtual_machine.appserver-${SERVICE}-A2 この変更をパイプラインに通すことにより、新しいサーバーがアクティブになり、古いサーバーへのトラフィックが停止します。チームはすべてが正常に動作していることを確認し、何か問題が発生した場合には簡単に変更を元に戻すことができます。新しいサーバーが正常に動作するようになったら、チームはスタックコードから古いサーバーを削除することができます：virtual_machine:
name: appserver-${SERVICE}-A2
memory: 4GB
vlan: external_stack.shared_network_stack.appserver_vlan_A

static_ip:
name: address-${SERVICE}-A
attach: virtual_machine.appserver-${SERVICE}-A2 この変更がパイプラインを通り、すべての環境に適用されると、アプリケーションインフラストラクチャースタックは共有ネットワーキングスタックの main_vlan に依存しなくなります。すべての消費者インフラストラクチャーが変更された後、ShopSpinner チームはプロバイダースタックコードから main_vlan を削除することができます：vlans:

- appserver_vlan_A
  address_range: 10.1.0.0/16
- appserver_vlan_B
  address_range: 10.2.0.0/16
- appserver_vlan_C
  address_range: 10.3.0.0/16

export:

- appserver_vlan_A: appserver_vlan_A.id
- appserver_vlan_B: appserver_vlan_B.id
- appserver_vlan_C: appserver_vlan_C.id VLAN の変更は完了し、main_vlan の最後の残骸もなくなりました。

- めっちゃ漸近的な取り組みや
- 読む気失せたわら
