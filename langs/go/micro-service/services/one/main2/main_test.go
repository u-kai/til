package main2_test

import "testing"

// TestNewStrings is a test for NewStrings
func TestNewStrings(t *testing.T) {
	result := []string{}
	for _, v := range []string{"a", "b", "c"} {
		result = append(result, v)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3, got %d", len(result))
	}
}
