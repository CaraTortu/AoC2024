package days

import "testing"

func TestDay2A(t *testing.T) {
	filename := "../inputs/day2_test.txt"
	expected := 2
	actual := Day2A(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay2B(t *testing.T) {
	filename := "../inputs/day2_test.txt"
	expected := 4
	actual := Day2B(filename)

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
