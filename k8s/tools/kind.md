- Docker コンテナを複数個起動し，そのコンテナを Kubernetes Node として利用することで，複数台構成の Kubernetes クラスタを構築する
- ローカル環境でマルチノードクラスタを構築するには一番良い

- クラスタを構築するためには yaml ファイルで管理することが可能

```yaml
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
  - role: control-plane
    image: kindest/node:v1.21.1
  - role: control-plane
    image: kindest/node:v1.21.1
  - role: worker
    image: kindest/node:v1.21.1
  - role: worker
    image: kindest/node:v1.21.1
```

```
kind create cluster --config kind.yaml --name kindcluster
kubectl config use-context kind-kindcluster
```

- docker container ls でコンテナを確認すると，node の台数分コンテナが確認できる
  - 上にプラスして k8s master の冗長後世のために使われている HAProxy も確認できる
