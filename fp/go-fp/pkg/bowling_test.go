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
	//t.Run("multi players", func(t *testing.T) {
	//	player1 := pkg.NewPlayer("John")
	//	player2 := pkg.NewPlayer("Tom")
	//	game := pkg.NewGame(player1, player2)

	//

	//	johnThrow := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(3)))
	//	game = game.Throw(johnThrow)
	//	johnThrow2 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(3)))
	//	game = game.Throw(johnThrow2)

	//	if game.Scores()[player1.Name] != 6 {
	//		t.Errorf("Expected score to be 6, but got %d", game.Scores()[player1.Name])
	//	}
	//	if game.Scores()[player2.Name] != 0 {
	//		t.Errorf("Expected score to be 0, but got %d", game.Scores()[player2.Name])
	//	}

	//	tomThrow := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(4)))
	//	game = game.Throw(tomThrow)
	//	tomThrow2 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(4)))
	//	game = game.Throw(tomThrow2)

	//	if game.Scores()[player1.Name] != 6 {
	//		t.Errorf("Expected score to be 6, but got %d", game.Scores()[player1.Name])
	//	}
	//	if game.Scores()[player2.Name] != 8 {
	//		t.Errorf("Expected score to be 8, but got %d", game.Scores()[player2.Name])
	//	}
	//})

}
