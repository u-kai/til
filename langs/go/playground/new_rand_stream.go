package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan interface{})
	randStream := newRandStreamDoneAble(done)
	fmt.Println("3 rand ints")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(1 * time.Second)
}

func newRandStreamDoneAble(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return randStream
}
func newRandStream() <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			randStream <- rand.Int()
		}
	}()
	return randStream
}
