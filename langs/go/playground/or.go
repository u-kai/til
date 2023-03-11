package main

import (
	"fmt"
	"time"
)

// orチャネルパターン
// 1つ以上のdoneチャネルを1つのdoneチャネルにまとめて，まとめているチャネルのうちのどれかひとつのチャネルが閉じられたらまとめてチャネルを閉じたい時に使う
// 以下の関数によって，任意の数のチャネルを1つのチャネルにまとめることができる
// まとめているどれかひとつのチャネルが1つでも閉じたり，書き込まれたりしたらすぐに合成されたチャネルが閉じるようになっている
// x/2の後ルーチンの生成がされるが，それを気にするのは早すぎる最適化らしい
// コンパイル時に扱うdoneチャネルがいくつあるかわからないのであれば，そもそも他にdoneチャネルをまとめる方法はない
func main() {
	// この関数はチャネルの可変長引数のスライスを受け取り，1つのチャネルを返す
	sum := 0
	var or func(channels ...<-chan interface{}) <-chan interface{}
	// 再帰関数のため，停止条件を決める必要がある
	// 最初の条件は可変長スライスが殻の場合
	// 要素が一つの場合はその要素を返す
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			sum++
			println(sum)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				// ここに全てが展開されて，ここのcaseに達した時点で，このgoroutineが終了して，or関数が返すorDoneがcloseされてchanのブロックが解除されるって感じか
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}

		}()
		return orDone
	}
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
