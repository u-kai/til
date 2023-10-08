package pkg_test

import (
	"example/pkg"
	"fmt"
	"testing"
)

func fakeThrow(anyThrowResult pkg.ThrowResult) func() pkg.ThrowResult {
	return func() pkg.ThrowResult {
		return anyThrowResult
	}
}

func TestBowling(t *testing.T) {

	t.Run("Can calculate score", func(t *testing.T) {
		t.Run("add final round case not spare", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(3)))
			for i := 0; i < 18; i++ {
				rounds = throw(rounds)
			}
			fmt.Printf("rounds: %+v\n", rounds)
			score := rounds.Score()
			if score != 54 {
				t.Errorf("Expected score to be 54, but got %d", score)
			}
			finalFirst := throw(rounds)
			finalFirstScore := finalFirst.Score()
			if finalFirstScore != 57 {
				t.Errorf("Expected score to be 57, but got %d", finalFirstScore)
			}

			finalSecond := throw(finalFirst)
			finalSecondScore := finalSecond.Score()
			if finalSecondScore != 60 {
				t.Errorf("Expected score to be 60, but got %d", finalSecondScore)
			}
		})
		t.Run("all strike", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			strikeThrow := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(10)))
			for i := 0; i < 12; i++ {
				rounds = strikeThrow(rounds)
			}
			score := rounds.Score()
			if score != 300 {
				t.Errorf("Expected score to be 300, but got %d", score)
			}
		})
		t.Run("for a strike", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw10 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(10)))
			throwed10 := throw10(rounds)
			score := throwed10.Score()
			if score != 0 {
				t.Errorf("Expected score to be 10, but got %d", score)
			}

			throw4 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(4)))
			throwed4 := throw4(throwed10)
			score = throwed4.Score()
			if score != 0 {
				t.Errorf("Expected score to be 0, but got %d", score)
			}
			throw5 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(5)))
			throwed5 := throw5(throwed4)
			score = throwed5.Score()
			if score != 28 {
				t.Errorf("Expected score to be 23, but got %d", score)
			}

		})
		t.Run("for a spare", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw6 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(6)))
			throwed6 := throw6(rounds)
			throw4 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(4)))
			throwedSpareRounds := throw4(throwed6)
			score := throwedSpareRounds.Score()

			if score != 0 {
				t.Errorf("Expected score to be 0, but got %d", score)
			}

			throw7 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(7)))
			throwed7 := throw7(throwedSpareRounds)
			score = throwed7.Score()

			if score != 24 {
				t.Errorf("Expected score to be 24, but got %d", score)
			}

			throw2 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(2)))
			throwed2 := throw2(throwed7)

			if throwed2.Score() != 26 {
				t.Errorf("Expected score to be 26, but got %d", score)
			}
		})
		t.Run("for a simple throw", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(5)))
			throwed5Rounds := throw(rounds)
			score := throwed5Rounds.Score()
			if score != 5 {
				t.Errorf("Expected score to be 5, but got %d", score)
			}

			throw = pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(3)))
			throwed3Rounds := throw(throwed5Rounds)
			score = throwed3Rounds.Score()

			if score != 8 {
				t.Errorf("Expected score to be 8, but got %d", score)
			}
		})
	})
}
