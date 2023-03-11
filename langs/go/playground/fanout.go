package main

import (
	"math/rand"
	"time"
)

func main() {
	rand := func() interface{} { return rand.Intn(500000000) }
	done := make(chan interface{})
	defer close(done)

	start := time.Now()
}
