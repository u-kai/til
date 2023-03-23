# 基本的な考え方

- sdn の基本的な考え方はネットワークにはコントロールプレーンとデータプレーンがあり，この 2 つのプレーンの分離はオープンなインターフェイスで実現されるべきというもの
- コントロールプレーンはネットワークがどのように振る舞うべきかを決定し，データプレーンは個々のパケットがどのように振る舞うべきかに関する実装責任を負っている
- コントロールプレーンにはルーティングテーブル(Routing Information Base(RIB))，データプレーンにはフォワーディングテーブル(Forwarding Information Base(FIB))を所持させる
- ディスアグリゲーションされるべき
- プログラム可能なデータプレーンも SDN の定義に入りつつある

- 要約すると SDN の定義は以下

```
コントロールプレーンとフォワーディングプレーンが物理的に分離されており，1つのコントロールプレーンで複数のフォワーディングデバイスを制御するネットワークのこと
```