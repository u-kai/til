-   :

    -   リストの先頭に値を挿入する

    ```haskell
    "A":["B","C"]
    ["A","B","C"]
    ```

-   ++

    -   リスト同士を連結する

    ```haskell
    [1,2,3,4,5] ++ [6]
    [1,2,3,4,5,6]
    ```

-   !!

    -   リストの要素を先頭からの位置で取得する
    -   領域を超えて指定すると index too large error

    ```haskell
    "Steve" !! 3
    v
    ```

-   head
-   tail
-   init
-   last

-   null

    -   リストが空かどうかを調べる

-   elem

    -   要素とリストを受け取り，リストの中にあるか判定

    ```haskell
    elem 3 [1,2,3]
    True

    4 `elem` [1,2,3]
    False
    ```
