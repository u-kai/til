package pkg_test

import (
	"boring/pkg"
	"testing"
)

//func TestIfGetASpareThenNextThrowScoreWillBeAdded(t *testing.T) {
//sut := pkg.NewBoring()
//sut.Throw(1)
//sut.Throw(9)
//sut.Throw(1)
//result := sut.Score()

//if result != 12 {
//t.Errorf("Expected 12, got %d", result)
//}
//}

//func TestScoreCalculateRequireTheFramesThatAreTwoSpotsInFront(t *testing.T) {
//first := pkg.NewFrame()
//first.Throw(1)
//first.Throw(2)

//second := pkg.FrameFromTwoBefore(nil, first)
//second.Throw(3)
//second.Throw(4)

//third := pkg.FrameFromTwoBefore(first, second)
//third.Throw(3)
//third.Throw(4)

//result := third.Score()

// if result != 17 {
// t.Errorf("Expected 17, got %d", result)
// }
// }

func TestFirstBollStrikeAndSubsequentBallsAreNot(t *testing.T) {
	sut := pkg.NewBoring()
	sut.Throw(10)
	sut.Throw(1)
	sut.Throw(1)
	sut.Throw(2)
	sut.Throw(2)
	result := sut.Score()

	if result != 18 {
		t.Errorf("Expected 14, got %d", result)
	}
}

func TestNoSecondThrowInTheCaseOfAStrike(t *testing.T) {
	sut := pkg.NewFrame()
	sut.Throw(10)
	err := sut.Throw(1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestFrameCanDetermineIfItsAStrikeOrNot(t *testing.T) {
	sut := pkg.NewFrame()
	sut.Throw(10)

	if !sut.IsStrike() {
		t.Fatalf("Expected true, got false")
	}
	sut = pkg.NewFrame()
	sut.Throw(1)
	sut.Throw(9)

	if sut.IsStrike() {
		t.Fatalf("Expected false, got true")
	}
}
func TestFrameCanDetermineIfItsASpareOrNot(t *testing.T) {
	sut := pkg.NewFrame()
	sut.Throw(1)
	sut.Throw(2)

	if sut.IsSpare() {
		t.Fatalf("Expected false, got true")
	}
	sut = pkg.NewFrame()
	sut.Throw(1)
	sut.Throw(9)

	if !sut.IsSpare() {
		t.Fatalf("Expected true, got false")
	}
}
func TestFrameRecordsOnlyYourOwnScore(t *testing.T) {
	sut := pkg.NewFrame()
	sut.Throw(1)
	sut.Throw(2)
	result := sut.Score()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
func TestIfNotStrikeInOneFramePlayerWillThrowTwice(t *testing.T) {
	sut := pkg.NewFrame()
	sut.Throw(1)
	sut.Throw(2)
	result := sut.Score()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestScoreCanBeCalculated(t *testing.T) {
	sut := pkg.NewBoring()
	sut.Throw(1)
	sut.Throw(3)
	result := sut.Score()

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
	sut.Throw(4)
	sut.Throw(5)
	result = sut.Score()
	if result != 13 {
		t.Errorf("Expected 13, got %d", result)
	}
}

func TestBoardFrameIsLimitedTo10Times(t *testing.T) {
	sut := pkg.NewBoring()
	for i := 0; i < 20; i++ {
		err := sut.Throw(1)
		if err != nil {
			t.Errorf("Expected no error, got %v and i = %d", err, i)
		}
	}
	err := sut.Throw(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
