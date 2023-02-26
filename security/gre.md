# GRE(Generic Routing Encapsulation)

- ネットワーク層のプロトコルのパケットをカプセル化するトンネリングプロトコル
- GRE のトンネル区間の両端は通常ルータ
- GRE トンネルは仮想的な専用線の役割を果たす
- 普通の IP の上に IP ヘッダと GRE ヘッダでカプセル化する

## GRE over IPsec

- IP マルチキャストもカプセル化可能

  - IP パケットをカプセル化した GRE パケットはトンネル区間の両端を送信もと／宛先とする IP ユニキャストパケットになる

- この特徴を活かし，GRE は IP マルチキャストパケットをカプセル化できないプロトコルと組み合わせて使用される
- 具体例としては GRE over IPsec
- トンネリングに GRE を使い，トンネル区間の暗号化に IPsec を用いる技術
- IPsec は IP マルチキャストパケットをカプセル化できないが，IP ユニキャストパケットをカプセル化できる
- そのため IP マルチキャストパケットをカプセル化した GRE パケットを IPsec でカプセル化することでインターネット経由で拠点間の IP マルチキャスト通信を実現することができ，かつ，インターネット区間を暗号化することができる

- GRE over IPsec を使用する典型的な例は IPsec を利用したインターネット VPN 回線において OSPF による動的な経路制御を行うネットワーク
- OSPF は IP マルチキャストパケットを用いて経路のリンクを交換しており，IPsec を利用したい場合は GRE でカプセル化する必要がある