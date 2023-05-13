kubectl run --image=nginx:1.16 --restart=Never --rm -it sample-debug --command -- /bin/sh

kubectl run コマンドを使って，実際にコンテナのシェル上でデバッグすることができる

上記コマンドで起動したコンテナは停止すると Pod の実行も呈する

- このコマンドで，コンテナ内の環境やアプリケーションに問題がある場合のデバッグに役に立つ
