package days

import "testing"

func TestDay1A(t *testing.T) {
	filename := "../inputs/day1_test.txt"
	expected := 11
	actual := Day1A(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay1B(t *testing.T) {
	filename := "../inputs/day1_test.txt"
	expected := 31
	actual := Day1B(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
