Feature Toggles When making a change to a component, you often need to keep using the existing implementation until you finish the change. Some people branch the code in source control, working on the new change in one branch and using the old branch in production. The issues with this approach include: It takes extra work to ensure changes to other areas of the component, like bugfixes, are merged to both branches.  It takes effort and resources to ensure both branches are continuously tested and deployed. Alternatively, the work in one branch goes with less rigorous testing, increasing the chances of errors and rework later on.  Once the change is ready, changing production instances over is more of a “big bang” operation, with a higher risk of failure.  It’s more effective to work on changes to the main codebase without branching. You can use feature toggles to switch the code implementation for different environments. Switch to the new code in some environments to test your work, and switch to the existing code for production-line environments. Use a stack configuration parameter (as described in Chapter 7) to specify which part of the code to apply to a given instance. Once the ShopSpinner team finishes adding VLANs to its shared-networking-stack, as described previously, the team needs to change the application-infrastructure-stack to use the new VLANs. The team members discover that this isn’t as simple a change as they first thought.
The application stack defines application-specific network routes, load balancer VIPs, and firewall rules. These are more complex when application servers are hosted across VLANs rather than in a single VLAN. It will take the team members a few days to implement the code and tests for this change. This isn’t long enough that they feel the need to set up a separate stack, as described in “Parallel Instances”. But they are keen to push incremental changes to the repository as they work, so they can get continuous feedback from tests, including system integration tests. The team decides to add a configuration parameter to the application-infrastructure-stack that selects different parts of the stack code depending on whether it should use a single VLAN or multiple VLANs. This snippet of the source code for the stack uses three variables — appserver_A_vlan, appserver_B_vlan, and appserver_C_vlan — to specify which VLAN to assign to each application server. The value for each of these is set differently depending on the value of the feature toggle parameter, toggle_use_multiple_vlans: input_parameters:
  name: toggle_use_multiple_vlans
  default: false

variables:
  - name: appserver_A_vlan
    value:
      $IF(${toggle_use_multiple_vlans} appserver_vlan_A ELSE main_vlan)
  - name: appserver_B_vlan
    value:
      $IF(${toggle_use_multiple_vlans} appserver_vlan_B ELSE main_vlan)
  - name: appserver_C_vlan
    value:
      $IF(${toggle_use_multiple_vlans} appserver_vlan_C ELSE main_vlan)
virtual_machine:
  name: appserver-${SERVICE}-A
  memory: 4GB
  address_block: ${appserver_A_vlan}

virtual_machine:
  name: appserver-${SERVICE}-B
  memory: 4GB
  address_block: ${appserver_B_vlan}

virtual_machine:
  name: appserver-${SERVICE}-C
  memory: 4GB
  address_block: ${appserver_C_vlan} If the toggle_use_multiple_vlans toggle is set to false, the appserver_X_vlan parameters are all set to use the old VLAN identifier, main_vlan. If the toggle is true, then each of the variables is set to one of the new VLAN identifiers. The same toggle parameter is used in other parts of the stack code, where the team works on configuring the routing and other tricky elements.
Advice on Feature Toggles See “Feature Toggles (aka Feature Flags)” by Pete Hodgson for advice on using feature toggles, and examples from software code. I have a few recommendations to add. Firstly, minimize the number of feature toggles you use. Feature toggles and conditionals clutter code, making it hard to understand, maintain, and debug. Keep them short-lived. Remove dependencies on the old implementation as soon as possible, and remove the toggles and conditional code. Any feature toggle that remains past a few weeks is probably a configuration parameter. Name feature toggles according to what they do. Avoid ambiguous names like new_networking_code. The earlier example toggle, toggle_use_multiple_vlans, tells the reader that it is a feature toggle, to distinguish it from a configuration parameter. It tells the reader that it enables multiple VLANs, so they know what it does. And the name makes it clear which way the toggle works. Reading a name like toggle_multiple_vlans, or worse, toggle_vlans, leaves you uncertain whether it enables or disables the multiple VLAN code. This leads to errors, where someone uses the conditional the wrong way around in their code.

フィーチャートグルー モーダニゼーションを行う際、既存の実装を変更が完了するまで引き続き使用する必要があります。一部の人々は、ソースコードをブランチして、新しい変更を1つのブランチで作業し、古いブランチを本番環境で使用します。このアプローチの問題点は次のとおりです。コンポーネントの他の領域（バグ修正など）への変更が両方のブランチにマージされるようにするために、追加の作業が必要です。両方のブランチが継続的にテストおよび展開されるようにするには、労力とリソースが必要です。また、1つのブランチでの作業は厳密なテストが行われず、エラーと再作業の可能性が高まります。変更が完了した後、本番インスタンスを変更する操作は「ビッグバン」のようなものであり、障害のリスクが高くなります。ブランチを作成せずにメインのコードベースで変更を行う方が効果的です。フィーチャートグルを使用して、異なる環境においてコードの実装を切り替えることができます。一部の環境では新しいコードに切り替えて作業をテストし、本番ラインの環境では既存のコードに切り替えます。指定されたインスタンスに適用するコードのどの部分を指定するために、スタック構成パラメータ（第7章で説明）を使用します。 ShopSpinnerチームが以前に説明したように、共有ネットワーキングスタックにVLANを追加する作業を終えると、チームはアプリケーションインフラストラクチャスタックを新しいVLANを使用するように変更する必要があります。チームメンバーは、これは最初に考えていたほど単純な変更ではないと気付きます。
アプリケーションスタックは、アプリケーション固有のネットワークルート、ロードバランサVIP、およびファイアウォールルールを定義します。アプリケーションサーバが単一のVLANではなく複数のVLANにホストされる場合、これらはより複雑になります。チームメンバーには、この変更のためのコードとテストを実装するために数日かかります。これは、「並列インスタンス」で説明したように別のスタックを設定する必要はないほど長くありませんが、コンティニュアスフィードバックを得るために、テスト（システム統合テストを含む）からの継続的なフィードバックを得るために、リポジトリにインクリメンタルな変更を加えたいと考えています。チームは、アプリケーションインフラストラクチャスタックに設定パラメータを追加して、スタックコードの異なる部分を選択するようにします。このスタックのソースコードのスニペットは、appserver_A_vlan、appserver_B_vlan、およびappserver_C_vlanという3つの変数を使用して、各アプリケーションサーバに割り当てるVLANを指定します。これらの各値は、フィーチャートグルパラメータtoggle_use_multiple_vlansの値に応じて異なるように設定されます。 input_parameters：
  name：toggle_use_multiple_vlans
  default：false

variables：
  - name：appserver_A_vlan
    value：
      $IF(${toggle_use_multiple_vlans} appserver_vlan_A ELSE main_vlan)
  - name：appserver_B_vlan
    value：
      $IF(${toggle_use_multiple_vlans} appserver_vlan_B ELSE main_vlan)
  - name：appserver_C_vlan
    value：
      $IF(${toggle_use_multiple_vlans} appserver_vlan_C ELSE main_vlan)
virtual_machine：
  name：appserver-${SERVICE}-A
  memory：4GB
  address_block：${appserver_A_vlan}

virtual_machine：
  name：appserver-${SERVICE}-B
  memory：4GB
  address_block：${appserver_B_vlan}

virtual_machine：
  name：appserver-${SERVICE}-C
  memory：4GB
  address_block：${appserver_C_vlan} toggle_use_multiple_vlansがfalseに設定されている場合、appserver_X_vlanパラメータはすべて古いVLAN識別子であるmain_vlanを使用するように設定されます。toggleがtrueの場合、それぞれの変数は新しいVLAN識別子の1つに設定されます。同じトグルパラメータは、チームがルーティングやその他のトリッキーな要素を設定するところで、スタックコードの他の部分でも使用されます。
フィーチャートグルのアドバイス フィーチャートグルの使用に関するアドバイスについては、Pete Hodgsonの「フィーチャートグル（またはフィーチャーフラグ）」を参照してください。私はいくつかの推奨事項を追加します。まず、使用するフィーチャートグルの数を最小限に抑えてください。フィーチャートグルと条件文は、コードを乱雑にし、理解、維持、デバッグが困難にします。それらを短期間に保持してください。できるだけ早く古い実装に依存しないようにし、トグルと条件付きのコードを削除します。数週間以上残るフィーチャートグルはおそらく設定パラメータです。フィーチャートグルの名前は、その機能に基づいて付けてください。new_networking_codeのような曖昧な名前は避けてください。前の例のようなトグル、toggle_use_multiple_vlansは、それがフィーチャートグルであることを読者に伝え、設定パラメータと区別します。多重VLANを有効にすることを読者に伝えるので、それが何をするかわかります。名前によってトグルがどちらの方向で機能するかも明確になります。toggle_multiple_vlansのような名前、またはさらに悪いtoggle_vlansのような名前を読むと、多重VLANコードを有効または無効にするのか不確定です。これにより、誰かがコード内の条件を間違った方向に使用するエラーが発生する可能性があります。