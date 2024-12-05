package days

import "testing"

func TestDay4A(t *testing.T) {
	filename := "../inputs/day4_test.txt"
	expected := 18
	actual := Day4A(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay4B(t *testing.T) {
	filename := "../inputs/day4_test.txt"
	expected := 9
	actual := Day4B(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
