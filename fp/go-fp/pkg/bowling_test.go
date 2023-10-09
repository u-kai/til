package pkg_test

import (
	"example/pkg"
	"fmt"
	"testing"
)

func TestBowling(t *testing.T) {
	t.Run("case all strike", func(t *testing.T) {
		player := pkg.NewPlayer("John")
		game := pkg.NewGame(player)

		for i := 0; i < 12; i++ {
			game = pkg.Play(game, func() pkg.HitPin {
				return pkg.HitPin(10)
			})
		}

		scores := pkg.AllScore(game)
		if scores[player.Name] != 300 {
			t.Errorf("Expected score to be 300, but got %d", scores[player.Name])
		}
	})
	t.Run("case double", func(t *testing.T) {
		player := pkg.NewPlayer("John")
		game := pkg.NewGame(player)

		for i := 0; i < 2; i++ {
			game = pkg.Play(game, func() pkg.HitPin {
				return pkg.HitPin(10)
			})
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(5)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(3)
		})

		fmt.Printf("game: %+v\n", game)
		scores := pkg.AllScore(game)

		if scores[player.Name] != 48 {
			t.Errorf("Expected score to be 48, but got %d", scores[player.Name])
		}

	})
	t.Run("case all spare", func(t *testing.T) {
		player := pkg.NewPlayer("John")
		game := pkg.NewGame(player)

		for i := 0; i < 21; i++ {
			game = pkg.Play(game, func() pkg.HitPin {
				return pkg.HitPin(5)
			})
		}

		scores := pkg.AllScore(game)
		if scores[player.Name] != 150 {
			t.Errorf("Expected score to be 150, but got %d", scores[player.Name])
		}
	})
	t.Run("player is only one", func(t *testing.T) {
		player := pkg.NewPlayer("John")
		game := pkg.NewGame(player)

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(3)
		})

		scores := pkg.AllScore(game)

		if scores[player.Name] != 3 {
			t.Errorf("Expected score to be 3, but got %d", scores[player.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})

		scores = pkg.AllScore(game)

		if scores[player.Name] != 7 {
			t.Errorf("Expected score to be 7, but got %d", scores[player.Name])
		}

	})
	t.Run("multi players", func(t *testing.T) {
		player1 := pkg.NewPlayer("John")
		player2 := pkg.NewPlayer("Jane")
		game := pkg.NewGame(player1, player2)

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(3)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})
		scores := pkg.AllScore(game)

		if scores[player1.Name] != 7 {
			t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 0 {
			t.Errorf("Expected score to be 0, but got %d", scores[player2.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(5)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})

		scores = pkg.AllScore(game)

		if scores[player1.Name] != 7 {
			t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 9 {
			t.Errorf("Expected score to be 9, but got %d", scores[player2.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(5)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})

		scores = pkg.AllScore(game)

		if scores[player1.Name] != 16 {
			t.Errorf("Expected score to be 16, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 9 {
			t.Errorf("Expected score to be 9, but got %d", scores[player2.Name])
		}
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(6)
		})
		scores = pkg.AllScore(game)

		if scores[player1.Name] != 16 {
			t.Errorf("Expected score to be 16, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 9 {
			t.Errorf("Expected score to be 9, but got %d", scores[player2.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(10)
		})
		scores = pkg.AllScore(game)

		if scores[player1.Name] != 16 {
			t.Errorf("Expected score to be 16, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 9 {
			t.Errorf("Expected score to be 9, but got %d", scores[player2.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(5)
		})
		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})
		scores = pkg.AllScore(game)

		if scores[player1.Name] != 16 {
			t.Errorf("Expected score to be 36, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 33 {
			t.Errorf("Expected score to be 33, but got %d", scores[player2.Name])
		}
		//if scores[player1.Name] != 7 {
		//	t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		//}
		//if scores[player2.Name] != 4 {
		//	t.Errorf("Expected score to be 4, but got %d", scores[player2.Name])
		//}

		//game = pkg.Play(game, func() pkg.HitPin {
		//	return pkg.HitPin(6)
		//})

		//scores = pkg.AllScore(game)

		//if scores[player1.Name] != 7 {
		//	t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		//}
		//if scores[player2.Name] != 0 {
		//	t.Errorf("Expected score to be 0, but got %d", scores[player2.Name])
		//}

		//game = pkg.Play(game, func() pkg.HitPin {
		//	return pkg.HitPin(9)
		//})

		//scores = pkg.AllScore(game)

		//if scores[player1.Name] != 16 {
		//	t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		//}
		//if scores[player2.Name] != 0 {
		//	t.Errorf("Expected score to be 0, but got %d", scores[player2.Name])
		//}

		//game = pkg.Play(game, func() pkg.HitPin {
		//	return pkg.HitPin(10)
		//})

		//scores = pkg.AllScore(game)

		//if scores[player1.Name] != 7 {
		//	t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		//}
		//fmt.Printf("game %v", game)
		//if scores[player2.Name] != 20 {
		//	t.Errorf("Expected score to be 20, but got %d", scores[player2.Name])
		//}

	})

}
