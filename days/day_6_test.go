package days

import (
	"testing"
)

func TestDay6A(t *testing.T) {
	actual := Day6A("../inputs/day6_test.txt")
	expected := 41
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func TestDay6B(t *testing.T) {
	actual := Day6B("../inputs/day6_test.txt")
	expected := 6
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func TestLoop(t *testing.T) {
	board := day6GetBoard("../inputs/day6_test_loop.txt")
	actual := board.isLoop()
	expected := true
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}

	board = day6GetBoard("../inputs/day6_test.txt")
	actual = board.isLoop()
	expected = false
	if actual != expected {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
}

func TestClone(t *testing.T) {
	board := day6GetBoard("../inputs/day6_test.txt")
	clone := board.clone()
	clone.Player.X = 1
	clone.Player.Y = 1

	clone.setCell(0, 0, 1)
	clone.setCell(1, 1, 1)

	if clone.Player.X == board.Player.X {
		t.Errorf("clone.Player.X == board.Player.X")
	}

	if clone.Player.Y == board.Player.Y {
		t.Errorf("clone.Player.Y == board.Player.Y")
	}

	if clone.getCell(0, 0) == board.getCell(0, 0) {
		t.Errorf("clone.Board[0][0] == board.Board[0][0]")
	}

	if clone.getCell(1, 1) == board.getCell(1, 1) {
		t.Errorf("clone.Board[1][1] == board.Board[1][1]")
	}
}
