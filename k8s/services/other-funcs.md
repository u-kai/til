## セッションアフィニティ

- Service にはセッションアフィニティを有効化することができる
- 例えば ClusterIP Service で有効にした場合，Pod から ClusterIP に向けて送られたトラフィックは Service に紐づくいずれか一つの Pod に対して転送された後，その後のトラフィックもずっと同じ Pod に対して送られるようになる
- Service に対しての設定は，spec.sessionAffinity と spec.sessionAffinityConfig を利用して行う
  - ロジック的にはクライアント IP アドレスを元に転送先を決定している
  - 秒数指定で同じ Pod に振り分けることができる

## externalTrafficPolicy

- この値を Local にすることでノードをまたぐロードバランシングが行われなくなるため，無駄な通信を防ぐことができる
- nat しなくても良くなるため送信元 IP アドレスの取得なども可能になる
