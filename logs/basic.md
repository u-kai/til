# ログ出力すべき内容を決めるには

## どんな目的でログを利用するのか

- ログ出力にあたっては，どのような目線で利用するかを意識することが大切
- 開発者目線

  - 不具合が発生した時の問題追従
  - パフォーマンス改善ポイントの発見

- 運用目線

  - サーバ停止などの確認
  - 監査
  - ユーザの分析

- セキュリティ監査目線

  - ユーザや管理者の操作を後から追従できるように
  - 怪しげなリクエストの発見

- 運用の観点では大量に出力されるログの中から異常な状態が発生したかどうかを簡単に見分けられることが重要
  - 一般的にはエラーレベルを使って切り分けられるようにする

## ログ出力の項目

- エラー発生時はエラーの発生箇所ごとにユニークな ID を発番してそれを出力しておくと，分析がしやすくなる
- エラーメッセージとしてフロントエンドに返すかどうかは別としてもそれぞれの発生箇所や原因が明確にわかる情報をエラーコードとして付与しておくと良い
- 起動直後に環境変数などのサーバの設定をダンプすれば，正しくデプロイできたかどうかの検証に使える
- ユーザーの行動分析の場合，分析しやすいように補助情報を出しても良い
  - ログイン，データ閲覧などのユーザー行動のイベントはログ的には散発的な情報になりやすいので一連のセッションで行われている操作がまとめて取り出せるようにセッション ID などを入れるのも手
  - ユーザー行動に大きく影響のあるレコード，特に時間の経過とともに変化する状態は出力しておくと良い
  - 通常の業務アプリケーションでは同じ情報をなるべく 1 箇所に集約するようにデータベースを変更し，全ての箇所が矛盾なく変更されやすくする正規化するが，変化しうる情報のスナップショットを扱う場合はあえて非正規化する

## 出力頻度