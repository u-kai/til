package pkg_test

import (
	"example/pkg"
	"testing"
)

func TestCurry(t *testing.T) {
	words := []string{"ada", "haskell", "scala", "java", "rust"}
	wordsToWords := pkg.WordsWithScoreHitherThan()(1)(words)
	if wordsToWords[0] != "java" {
		t.Errorf("Expected java, got %s", wordsToWords[0])
	}
	if len(wordsToWords) != 1 {
		t.Errorf("Expected 1 word, got %d", len(wordsToWords))
	}
	words = []string{"football"}
	wordsToWords = pkg.WordsWithScoreHitherThan()(0)(words)
}
