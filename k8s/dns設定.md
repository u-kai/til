# Pod の DNS 設定(dnsPolicy)とサービスディスカバリ

- DNS サーバに関する設定は Pod 定義の spec.dnsPolicy に記述する
- 設定できる値は以下の 4 種類
  - ClusterFirst
    - デフォルトの設定
    - クラスタ内の DNS に問い合わせを行い，解決できなかった場合はアップストリームに問い合わせをおこなう
    - サービスディスカバリやクラスタ内ロードバランサーを利用するため
  - None
    - Pod 定義内で静的に設定を行う
  - Default
    - Pod が起動する K8s Node の/etc/resolve.conf を引き継ぐ
  - ClusterFirstWithHostNet
    - ClusterFirst の挙動と同等

## 静的な名前解決の設定(/etc/hosts)

- k8s には Pod 内のすべてのコンテナの/etc/hosts を書き換える機能が用意されており，spec.hostAliases で指定することが可能
