# Mediator Pattern

- Mediator とは仲介人のこと
- Mediator Pattern とは、複数のオブジェクト間の相互作用をカプセル化するオブジェクトのことです。by copilot

## Sample Code

- 名前とパスワードを入力するログインダイアログ
- 入力するデータやパターンによっての振る舞いが変わる
  - このような多数のオブジェクト間の調整を行わなければいけない時こそ役にたつ

## class とインターフェイス一覧

| 名前               | 解説                                                                  |
| ------------------ | --------------------------------------------------------------------- |
| Mediator           | 相談役のインターフェイスを定めるインターフェイス                      |
| Colleague          | メンバーのインターフェイスを定めるインターフェイス                    |
| ColleagueButton    | Colleague インターフェイスを実装しているボタンを表すクラス            |
| ColleagueTextField | Colleague インターフェイスを実装しているテキストを表すクラス          |
| ColleagueCheckBox  | Colleague インターフェイスを実装しているチェックボックスを表すクラス  |
| LoginFrame         | Mediator インターフェイスを実装しているログインダイアログを表すクラス |

## 利点

- ロジックが mediator に固まる
  - デバックがしやすい
- インスタンス同士のやり取りを少なくできる
- Colleague に詳細なロジックがないので，その他の用途でも使いやすい

  - 逆に Mediator は各 Colleague に特化したロジックになるのでその他の用途ではまず使えない

- ルールが変わっても Colleague が変わらなくても良い
- 新たな Colleague が増えても Colleague は変わらなくて良い