package days

import (
	"testing"
)

func TestDay5A(t *testing.T) {
	expected := 143
	result := Day5A("../inputs/day5_test.txt")
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay5B(t *testing.T) {
	expected := 123
	result := Day5B("../inputs/day5_test.txt")
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
