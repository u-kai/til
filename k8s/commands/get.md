- -o で出力をさまざまな形式で出力できる
- Custom Column や JSON Path で配列のデータを出力する際には ?(@.filed == '')の形式で指定することで配列内の特定の条件にマッチする要素でフィルタリングできる

```
kubectl get pod sample-pod -o jsonpath="{.spec.containers[?(@.name=="nginx-container")].image}
nginx:1.16
```
