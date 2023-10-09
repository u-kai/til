package pkg_test

import (
	"example/pkg"
	"testing"
)

func TestBowling(t *testing.T) {
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

		scores := pkg.AllScore(game)

		if scores[player1.Name] != 3 {
			t.Errorf("Expected score to be 3, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 0 {
			t.Errorf("Expected score to be 0, but got %d", scores[player2.Name])
		}

		game = pkg.Play(game, func() pkg.HitPin {
			return pkg.HitPin(4)
		})

		scores = pkg.AllScore(game)

		if scores[player1.Name] != 7 {
			t.Errorf("Expected score to be 7, but got %d", scores[player1.Name])
		}
		if scores[player2.Name] != 0 {
			t.Errorf("Expected score to be 0, but got %d", scores[player2.Name])
		}
	})

}
