# StatefulSet

- StatefulSet は ReplicaSet の特殊な形と言えるリソース
- ReplicaSet との違いは，以下

  - 作成される Pod 名のサフィックスは数字のインデックスが付与されたものになる
    - Pod 名が変わらない
  - データを永続化するための仕組みを有している
    - PersistentVolume を使っている場合は，Pod の再起動時に同じディスクを利用して再作成される

- spec.volumeClaimTemplates を指定することができる
  - これにより，StatefulSet によって作成される各 Pod に対して，PersistentVolumeClaim を設定できる
  - PersistentVolumeClaim を利用することでクラスタ外のネットワーク越しに提供される PersistentVolume を Pod にアタッチできるため，Pod の再起動時や別ノードへの移動時にデータを保持した状態でコンテナが再作成されるようになる
  - PersistentVolume は 1 つの Pod が占有することも，PersistentVolume の種類によっては複数の Pod で共有することも可能

## 永続化領域のデータ保持の確認

- コンテナ内から永続化領域がマウントされていることを確認する
- 永続化領域を利用している場合には/dev/sdb などの別ディスク(PersistentVolume)がマウントされている

```
kubectl exec -it sample-statefulset-0 --df -h | grep /dev/sd
```

## StatefulSet の削除と PersistentVolume の削除

- StatefulSet を作成すると，Pod に対して PersistentVolumeClaim を設定可能なため，PersistentVolume も同時に確保できる
- この確保した永続化領域は StatefulSet が削除されても同時に解放されることはあり得ない
- これは StatefulSet が永続化領域を解放する前にボリュームからデータをバックアップする猶予を与えるため
- 逆に言えば StatefulSet を削除後 StatefulSet が Persistent VolumeClaim で確保した PersistentVolume を解放せずに再度 StatefulSet を作成した場合には永続化領域のデータはそのまま Pod が起動することになる
