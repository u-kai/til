# ServiceAccount

- ServiceAccount を作成すると Secret が作成される
- ServiceAccount はこのトークンを元にして k8s API への認証情報として利用することが可能

  - そのため pod に割り当てた ServiceAccount にある権限がそのまま Pod に割り当てられる権限となる

- ServiceAccount を指定したコンテナを作成後，起動している Pod の情報を確認すると，トークンが Volume として自動的に埋め込まれている
  - マウントされている領域にはトークンや証明書が配置されており，Pod 上のアプリケーションはこれらを利用することで，指定された Service Account の権限で K8s を操作できる

## クライアントライブラリと認証

- k8s のシステムコンポーネントはクライアントライブラリとして k8s/client-go を利用して API と通信している
