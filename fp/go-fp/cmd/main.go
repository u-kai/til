package main

import (
	"example/pkg"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	game := pkg.NewGame(pkg.NewPlayer("John"))
	for {
		game = pkg.Play(game, func() pkg.HitPin {
			_, err := fmt.Scan(&s)
			if err != nil {
				panic(err)
			}
			d, err := strconv.Atoi(strings.Trim(s, "\n"))
			if err != nil {
				println("input error")
			}
			println(d)
			return pkg.HitPin(d)
		})
		println(game.String())
	}

}
