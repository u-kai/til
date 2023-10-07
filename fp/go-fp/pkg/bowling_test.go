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
		t.Run("for a spare", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw6 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(6)))
			throwed6 := throw6(rounds)
			throw4 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(4)))
			throwedSpareRounds := throw4(throwed6)
			score := throwedSpareRounds.Score()

			if score != 10 {
				t.Errorf("Expected score to be 10, but got %d", score)
			}

			throw7 := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(7)))
			throwed7 := throw7(throwedSpareRounds)
			score = throwed7.Score()

			if score != 24 {
				t.Errorf("Expected score to be 24, but got %d", score)
			}
		})
		t.Run("for a simple throw", func(t *testing.T) {
			rounds := pkg.NewDefaultRounds()
			throw := pkg.ThrowByAnyLogic(fakeThrow(pkg.ThrowResult(5)))
			throwed5Rounds := throw(rounds)
			score := throwed5Rounds.Score()
			fmt.Printf("rounds = %+v\n", rounds)
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
