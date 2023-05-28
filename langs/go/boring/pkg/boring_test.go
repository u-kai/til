package pkg_test

import (
	"boring/pkg"
	"testing"
)

func TestScoreCanBeCalculated(t *testing.T) {

	sut := pkg.NewBoring()
	sut.Slow(1)
	sut.Slow(3)
	result := sut.Score()

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}

	sut.Slow(4)
	sut.Slow(5)

	result = sut.Score()

	if result != 13 {
		t.Errorf("Expected 13, got %d", result)
	}

}
