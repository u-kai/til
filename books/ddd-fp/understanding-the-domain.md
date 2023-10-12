# Understanding the Domain

## Interview with a Domain Expert

- 最初は大枠の workflow，input と output について聞くと良い

### Thinking About Inputs and Outputs

- workflow の output を考える時に，最終的に何かの状態を変えた副作用だったとしても，workflow としては常に event を産むことになるので，output は event ということになる
- そしてその event が別のコンテキスト境界のトリガーになる

## Fighting the Impulse to Do Database Driven Design

- 低レベルな実装のスケッチを始めないこと
- DB の設計ではなく DDD をまずはやるべき
- DB は中央的存在ではない,実際の利用者は DB の存在に気を使わない
- DDD では persistence ignorance とよぶ

## Fighting the Impulse to Do Class Driven Design

- OOP は Domain ではなく objects のデザインを行う
- これは DB 設計の話と同様に危険
- 本当の要求を聞いていない
- まずは要求を集めることが重要で，開発者たちの技術アイデアを Domain に押し付けることでない

## Documenting the Domain

- UML とかよりも普通のテキストベースのもので domain model をキャプチャする
- 下のように Place Order workflow を書く

```
Bounded context : Order-Taking

Workflow: "Place Order"
    triggered by:
        "Order form received" event (when Quote is not checked)
    primary input:
        An order form
    other input:
        Product catalog
    output events:
        "Order Placed" event
    side-effects:
        An acknowledgment is sent to the customer,
        along with the placed order
```

- データ構造に関しては以下

```
bounded context :Order-Taking

data Order =
    CustomerInfo
    And ShippingAddress
    And BillingAddress
    And list of OrderLines
    And AmountToBill

data OrderLine =
    Product
    And Quantity
    And Price
data CustomerInfo = ??? // don't know yet
data BillingAddress = ??? // don't know yet
```

## Diving Deeper into the Order-Taking Workflow

- ここら辺まできたら詳しく workflow について尋ねてみる

### Representing Constraints

- new type まで作ると変更が厳しくなるが，自由に作りすぎると一生 design されない
- 何が正解かはいつだってコンテキストに依存する
- domain expert のポイントをキャプチャすること
- document 化していく感じ？

```
data UnitQuantity = integer between 1 and 1000
```

### Representing the Life Cycle of an Order

- validated や priced といったデータにはライフサイクルがある

  - 状態遷移のようなものか？
  - と思ったけど，少し違う気もする
  - 全く違うデータ構造として定義している感じ

- validated や unvalidate といった情報は異なる
- これによって，不整合が起きないようにしている

### Fleshing out the Steps in the Workflow
