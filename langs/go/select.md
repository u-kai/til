# select

- チャネルをまとめるのり
  - chan はゴルーチンをまとめるのり
- プログラム内の接合部でのキャンセル処理，タイムアウト，待機，デフォルト値といった概念とチャネルを安全にまとめられている
- 読み込みチャネルの場合はチャネルに書き込みがあったり閉じられたりしたのか，書き込みの場合は，キャパシティがいっぱいになっていないのかを確認する
- そのチャネルも準備ができていない場合は select 文全体がブロックする

  - チャネルが一つでも準備完了したらその操作が行われ，対応する文が実行される

- Go のランタイムは case 文全体に対して疑似乱数による一様選択をしているため，それぞれの条件が等しく選択される可能性がある

  - 回数の決まっている処理の中に select 文を書いてしまうと，何が実行されるのかわからないって感じか？
    - そうっぽい

- select が空だと(case 節がない)と永遠にブロックする

```go
package main

import (
	"fmt"
)

func main() {

	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	var c1Count, c2Count int
	select {
	case <-c1:
		c1Count++
	case <-c2:
		c2Count++
	}
	fmt.Printf("c1 %d, c2 %d", c1Count, c2Count)
}
```
