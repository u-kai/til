# Serialization

- deserialization は常にクリーンであるべき

  - これはすでに正常な値をコンテキスト境界で提供するべきだから

- input や output の json については単なる json string で良い
  - それはインフラであり，ドメインから独立しているので

## DTOs as a Contract Between Bounded Contexts

- DTO のつながりはゆるいけど，ここだけは慎重にしたい？みたいな？
