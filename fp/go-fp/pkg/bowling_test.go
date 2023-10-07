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
