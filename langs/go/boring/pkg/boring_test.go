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

//func TestScoreCalculateRequireTheFlamesThatAreTwoSpotsInFront(t *testing.T) {
//first := pkg.NewFlame()
//first.Slow(1)
//first.Slow(2)

//second := pkg.FlameFromTwoBefore(nil, first)
//second.Slow(3)
//second.Slow(4)

//third := pkg.FlameFromTwoBefore(first, second)
//third.Slow(3)
//third.Slow(4)

//result := third.Score()

// if result != 17 {
// t.Errorf("Expected 17, got %d", result)
// }
// }

func TestNoSecondThrowInTheCaseOfAStrike(t *testing.T) {
	sut := pkg.NewFlame()
	sut.Slow(10)
	err := sut.Slow(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestFrameCanDetermineIfItsAStrikeOrNot(t *testing.T) {
	sut := pkg.NewFlame()
	sut.Slow(10)

	if !sut.IsStrike() {
		t.Fatalf("Expected true, got false")
	}
	sut = pkg.NewFlame()
	sut.Slow(1)
	sut.Slow(9)

	if sut.IsStrike() {
		t.Fatalf("Expected false, got true")
	}
}
func TestFrameCanDetermineIfItsASpareOrNot(t *testing.T) {
	sut := pkg.NewFlame()
	sut.Slow(1)
	sut.Slow(2)

	if sut.IsSpare() {
		t.Fatalf("Expected false, got true")
	}
	sut = pkg.NewFlame()
	sut.Slow(1)
	sut.Slow(9)

	if !sut.IsSpare() {
		t.Fatalf("Expected true, got false")
	}
}
func TestFrameRecordsOnlyYourOwnScore(t *testing.T) {
	sut := pkg.NewFlame()
	sut.Slow(1)
	sut.Slow(2)
	result := sut.Score()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
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
