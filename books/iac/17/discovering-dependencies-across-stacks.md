Discovering Dependencies Across Stacks The ShopSpinner system includes a consumer stack, application-infrastructure-stack, that integrates with networking elements managed by another stack, shared-network-stack. The networking stack declares a VLAN:
vlan:
name: "appserver_vlan"
address_range: 10.2.0.0/8 The application stack defines an application server, which is assigned to the VLAN: virtual_machine:
name: "appserver-${ENVIRONMENT_NAME}"
vlan: "appserver_vlan" This example hardcodes the dependency between the two stacks, creating very tight coupling. It won’t be possible to test changes to the application stack code without an instance of the networking stack. Changes to the networking stack are constrained by the other stack’s dependence on the VLAN name. If someone decides to add more VLANs for resiliency, they need consumers to change their code implementation in conjunction with the change to the consumer.1 Otherwise, they can leave the original name, adding messiness that makes the code and infrastructure harder to understand and maintain: vlans:

- name: "appserver_vlan"
  address_range: 10.2.0.0/8
- name: "appserver_vlan_2"
  address_range: 10.2.1.0/8
- name: "appserver_vlan_3"
  address_range: 10.2.2.0/8 Hardcoding integration points can make it harder to maintain multiple infrastructure instances, as for different environments. This might work out depending on the infrastructure platform’s API for the specific resources. For instance, perhaps you create infrastructure stack instances for each environment in a different cloud account, so you can use the same VLAN name in each. But more often, you need to integrate with different resource names for multiple environments.
  So you should avoid hardcoding dependencies. Instead, consider using one of the following patterns for discovering dependencies.
  スタック間の依存関係の発見 ShopSpinner システムには、別のスタックで管理されるネットワーキング要素と統合されるアプリケーションインフラストラクチャスタックを含むコンシューマースタックが存在します。ネットワーキングスタックは次の VLAN を宣言します：
  vlan:
  name: "appserver_vlan"
  address_range: 10.2.0.0/8
  アプリケーションスタックはアプリケーションサーバを定義します。このアプリケーションサーバは VLAN に割り当てられます：
  virtual_machine:
  name: "appserver-${ENVIRONMENT_NAME}"
  vlan: "appserver_vlan" この例では、2 つのスタック間の依存関係がハードコードされ、非常に緊密な結合が作成されています。アプリケーションスタックのコードの変更をテストするには、ネットワーキングスタックのインスタンスが必要です。ネットワーキングスタックの変更は、他のスタックが VLAN 名に依存して制約が加えられます。耐障害性のためにさらに多くの VLAN を追加する場合、変更と共にコンシューマのコード実装を変更する必要があります。そうでなければ、元の名前を残して、コードやインフラストラクチャを理解し、保守するのが困難になります。:
  vlans:

- name: "appserver_vlan"
  address_range: 10.2.0.0/8
- name: "appserver_vlan_2"
  address_range: 10.2.1.0/8
- name: "appserver_vlan_3"
  address_range: 10.2.2.0/8
  統合ポイントをハードコードすると、異なる環境に複数のインフラストラクチャインスタンスを保守するのがより困難になる場合があります。具体的なリソースに対して、インフラストラクチャプラットフォームの API に依存する可能性があります。たとえば、各環境のために異なるクラウドアカウントでインフラストラクチャスタックのインスタンスを作成することで、同じ VLAN 名を使用できるようにする場合もあります。しかし、より頻繁に、異なる環境のために異なるリソース名に統合する必要があります。
  そのため、依存関係をハードコードすることは避けるべきです。代わりに、以下のいずれかのパターンを使用して依存関係を発見することを検討してください。
