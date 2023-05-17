# External IP
- 特定のk8s NodeのIPアドレス:Portで受信したトラフィックをコンテナに転送する形で外部疎通性を確立するService
# NodePort

- ExternalIP に類似したもの
- ExternalIP は指定した k8s Node の IP アドレス:Port で受信したトラフィックをコンテナに転送する形で外部疎通性を確立するものだが，NodePort はすべての k8s Node の IP アドレス:Port で受信したトラフィックをコンテナに転送する形で外部疎通性を確立する
- ExternalIP で全ノードからトラフィックを受け付けるように設定するのに近い機能だが，厳密に言うと Listen する際に 0.0.0.0:Port を使用してすべての IP アドレスで Bind するような形になっている

## 注意点

- NodePort で利用できるポート範囲は多くの k8s 環境で 30000~32767 となっている
- また複数の NodePort で同じポートを利用することはできない
