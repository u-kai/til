# Evolving a Design and Keeping It Clean

- DDD は静的なものではない
- 継続的な開発者とドメインエキスパートとのコラボレーションを意味している
- 要件が変わったら常に実装の前にドメインモデルを再度評価する必要がある
- 4 つの変更に対してみていく
  - workflow の追加
  - workflow の input を変更
  - キードメインの種類の定義を変更し、それがシステム全体にどのように波及するのか
  - ワークフロー全体をビジネスルールに適合させる

## Change 1 Adding Shipping Charges

- カテゴライズされたデータの処理に対してはパターンマッチではなく，Active Patter が良い
- ビジネスロジックとカテゴライズを分けると、コードがクリアになる
- また、ドキュメント性が上がる

- 確かにこうしておかないと、ドキュメント性は低く、今度値段を変えようとしても意味合いがわからずに苦労しそう
- カテゴライズと、料金を別々に更新することもできる
  - カテゴライズだけ変更や料金だけ変更も簡単にある
- GPT 曰く、Active pattern は F#特有のもので、(|pattern|)というバナナクリップ構文によってパターンマッチよりもパワフルな表現力に富んだものらしい

```fs
// pattern match
let calculateShippingCost validatedOrder =
    let shippingAddress = validatedOrder.ShippingAddress
    if shippingAddress.Country = "US" then
        match shippingAddress.State with
        |"CA"|"OR"|"AZ"|"NV"|->
        5.0
        |_ ->
        10.0
    else
        20.0

// active pattern
// ここでUsLocalStateなどを定義している？よくわからん
// この定義によって、addressの型をパターンマッチさせると、UsLocalStateかUsRemoteStateかInternationalが出てくる感じだと思う
// つまり、enumではない型をパターンマッチに使えるようにしている
// 普通だったら、ある型を引数としてenumを返す関数になりそうだが、それをF#ではactive patternを使える感じか
// 何がいいのか？
// "アクティブ パターン" を使用すると、入力データを分割する名前付きパーティションを定義できます。これにより、判別共用体の場合と同様に、パターン マッチング式でこれらの名前を使用できるようになります。 アクティブ パターンを使用すると、パーティションごとにカスタマイズした方法でデータを分解できます。by MS docs
// 関数のシンタックスシュガー？
// と思ったけど、おそらく結構使える
// enumであれば生成に対してのルールはいつくでも作成できる
// しかしactive patternはある型をある列挙子に変換する関数のようなものであるため、その列挙子はおそらくその一つの作成方法しか存在しないはず
// そして、その列挙子の元となる型はその列挙子のことを知らずに済むので依存関係も少なそう
// ある型を何かのパターンに紐づける時は良いものな気がしてきた
let (UsLocalState|UsRemoteState|International|) address =
    if address.Country = "US" then
        match address.State with
        |"CA"|"OR"|"AZ"|"NV" ->
        UsLocalState
        |_->
        UsRemoteState
    else
        International

let calculateShippingCost validatedOrder
    match validateOrder.ShippingAddress with
    | UsLocalState -> 5.0
    | UsRemoteState -> 10.0
    | International -> 20.0
```

## Creating a New Stage in the Workflow

- 要件が変わったことで、その要件専用に新しく型を作る(元の型にフィールドを増やすのではなく)ことはやりすぎのように思えるが、型を作らないと

  - もし、他の step で新しい型を予想していたらうまくいかない->これは意味わからん。想定すな
  - もし、新しいフィールドに対する値が型の生成後に埋め込まれるような値だと、ミュータブルにするか、初期化時には DefaultValue を使って、後から生成する必要があるのでよくない

- 他のステージから隔離されており、要求に型が一致していると、安全に新しい要求に対いする stage やコンポーネントを追加、削除できる

- 要件以外で pipeline に追加するものとしては
  - logging,performance metrics auditing などは簡単に追加できる
  - 認可チェック、やその失敗や、失敗を下に伝えることや、他のパイプラインをスキップする機能
  - 動的に集約 Root に設定値や context を入力として入れることができる
    - おそらくこれらは型を要件ごとに分けていれば、あるところを変更することや、その変更による変更はとても簡単だよねってことかな？

## Adding Support for VIP Customers
