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

## Creating a New Stage in the Workflow
