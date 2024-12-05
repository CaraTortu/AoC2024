package days

import "testing"

func TestDay3A(t *testing.T) {
	filename := "../inputs/day3A_test.txt"
	want := 161
	got := Day3A(filename)
	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestDay3B(t *testing.T) {
	filename := "../inputs/day3B_test.txt"
	want := 48
	got := Day3B(filename)
	if got != want {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
