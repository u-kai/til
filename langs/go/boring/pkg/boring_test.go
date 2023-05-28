package pkg_test

import (
	"boring/pkg"
	"testing"
)

//func TestIfGetASpareThenNextThrowScoreWillBeAdded(t *testing.T) {
//sut := pkg.NewBoring()
//sut.Slow(1)
//sut.Slow(9)
//sut.Slow(1)
//result := sut.Score()

//if result != 12 {
//t.Errorf("Expected 12, got %d", result)
//}
//}

func TestIfNotStrikeInOneFramePlayerWillThrowTwice(t *testing.T) {
	sut := pkg.NewFlame()
	sut.Slow(1)
	sut.Slow(2)
	result := sut.Score()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

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

func TestBoardFrameIsLimitedTo10Times(t *testing.T) {
	sut := pkg.NewBoring()
	for i := 0; i < 20; i++ {
		err := sut.Slow(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}
	err := sut.Slow(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
