package pkg

import "strings"

type Option[T any] struct {
	value *T
}

type Result[T any] struct {
	Ok  *T
	Err error
}
type IO[T any] struct {
	Result T
}

type LargerThan func(int) bool

func LargerThanN(n int) LargerThan {
	return func(x int) bool {
		return x > n
	}
}

type WordScore func(string) int
type WordsToWords func([]string) []string
type ScoreToWordsToWords func(int) WordsToWords

func Score(word string) int {
	return len(strings.ReplaceAll(word, "a", ""))
}
func Bonus(word string) int {
	result := 0
	for _, char := range word {
		if char == 'c' {
			result += 5
		}
	}
	return result
}
func Penalty(word string) int {
	result := 0
	for _, char := range word {
		if char == 's' {
			result += 7
		}
	}
	return result
}

func WordsWithScoreHitherThan() ScoreToWordsToWords {
	return HighScoringWords(func(w string) int {
		return Score(w) + Bonus(w) - Penalty(w)
	})
}

func HighScoringWords(wordScore WordScore) ScoreToWordsToWords {
	return func(higherThan int) WordsToWords {
		return func(words []string) []string {
			return filter(words, func(word string) bool {
				return wordScore(word) > higherThan
			})
		}
	}
}

func filter(words []string, predicate func(string) bool) []string {
	filtered := []string{}
	for _, word := range words {
		if predicate(word) {
			filtered = append(filtered, word)
		}
	}
	return filtered
}
