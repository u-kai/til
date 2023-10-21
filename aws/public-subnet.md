# Public Subnet について

- Public Subnet を監視しようと思ったら以下の Event をとる必要がある

  - CreateRouteTable

    - Main かつ IGW と接続していて，Subnet に RouteTable が紐づいてない場合は RouteTable と明示的な紐付けがない全ての Subnet が PublicSubnet になる

  - CreateRoute

    - Create された Route が igw かどうか & 左の Route が紐づいている RouteTable と Subnet が紐づいているかどうか
    - RouteTable への Route 追加がコンソール上の動き

  - ReplaceRouteTableAssociation
    - RouteTable の関連付けを変更したとき
    - 新しい AssociationId が Event で拾える
    - おそらく新しい RouteTableId も拾える
  - ReplaceRoute
    - コンソールでは CreateRoute が実行されたが，API ドキュメントとして存在するのでおそらく必要
