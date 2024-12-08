package days

import (
	"testing"

	"github.com/mowshon/iterium"
)

func TestDay7A(t *testing.T) {
	filename := "../inputs/day7_test.txt"
	expected := 3749
	actual := Day7A(filename)
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay7B(t *testing.T) {
	filename := "../inputs/day7_test.txt"
	expected := 11387
	actual := Day7B(filename)
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestIsPossible(t *testing.T) {
	permutations, _ = iterium.Product([]int{ADD, MULTIPLY, APPEND}, 5).Slice()

	equation := Equation{
		Result:  10,
		Numbers: []int{1, 2, 3, 4},
	}
	expected := true
	actual := equation.isPossible()
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}

	equation = Equation{
		Result:  23,
		Numbers: []int{2, 3},
	}
	expected = true
	actual = equation.isPossible()
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}

	equation = Equation{
		Result:  7290,
		Numbers: []int{6, 8, 6, 15},
	}
	expected = true
	actual = equation.isPossible()
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
