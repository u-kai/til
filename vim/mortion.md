# モーション

-   gj,jk,g0,g^,g$ は表示行に対して移動することができる

    -   練習

        ```
        aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

        ```

-   ge で前の単語の末尾にいける - 練習
    `rust let data = String; `

-   f コマンドで行きすぎたら,コマンドで戻る

## 場所をマークしてそこに戻る

-   m{a-xA-Z}でマークをつける
-   小文字のマークはここのバッファにローカルなマーク
-   大文字のマークはグローバルにアクセス可能
-   \`{mark}でマークしたところに戻ることができる
-   mm にすると楽で早い
-   Vim が自動的に設定してくれるマークは便利

        | key | buffer content                             |
        | --- | ------------------------------------------ |
        | ''  | 直前に行われたジャンプ以前にいた場所       |
        | '.  | 直前に変更された場所                       |
        | '^  | 直前に挿入された場所                       |
        | '[  | 直前に変更またはヤンクが行われたされた先頭 |
        | ']  | 直前に変更またはヤンクが行われたされた末尾 |

-   練習

```
{name : kai , age : 20}
```

## ジャンプ(ファイル間の移動)

-   Vim の戻るボタンのようなものはジャンプリストを使う
-   ctl-o は戻るボタン
-   ctl-i は進むボタン
-   :jumps でジャンプリストの内容を確認できる
    -   現在のウィンドウでアクティブなファイルを変更するコマンドは全てジャンプとして扱われる

## 変更リストを辿る

-   Vim はドキュメントに変更があるたびに変更後のカーソル位置を記録する
-   :changes で確認可能
-   g;で直前に変更があった場所に戻ることができる
-   \`.マークは直前に行われた変更の位置を参照する
-   \`^マークは直前の挿入を参照する.これは gi でも同じで，挿入から抜けて他のところにうろうろして戻る時に有効
