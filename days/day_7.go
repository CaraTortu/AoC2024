package days

import (
	"log"
	"strconv"
	"strings"

	"AoC24/utils"

	"github.com/mowshon/iterium"
)

const (
	ADD = iota
	MULTIPLY
	APPEND
)

type Equation struct {
	Result  int
	Numbers []int
}

var permutations [][]int

func getEquations(filename string) []Equation {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	equations := make([]Equation, len(contents))
	for i, line := range contents {
		parts := strings.Split(line, ": ")
		result, _ := strconv.Atoi(parts[0])

		numbersStr := strings.Split(parts[1], " ")
		numbers := make([]int, len(numbersStr))
		for i, number := range numbersStr {
			numbers[i], err = strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
		}

		equations[i] = Equation{
			Result:  result,
			Numbers: numbers,
		}
	}

	return equations
}

func padding(n int) int {
	var p int = 1
	for p < n {
		p *= 10
	}
	return p
}

func (e *Equation) isPossible() bool {
	for _, perm := range permutations {
		result := e.Numbers[0]
		for i := 1; i < len(e.Numbers); i++ {

			number := e.Numbers[i]

			switch perm[i-1] {
			case ADD:
				result += number
				break
			case MULTIPLY:
				result *= number
				break
			case APPEND:
				result = result*padding(number) + number
			}
		}

		if result == e.Result {
			return true
		}
	}

	return false
}

func Day7A(filename string) int {
	equations := getEquations(filename)

	maxPermutations := 1
	for i := 0; i < len(equations); i++ {
		if maxPermutations < len(equations[i].Numbers) {
			maxPermutations = len(equations[i].Numbers)
		}
	}

	perm, _ := iterium.Product([]int{ADD, MULTIPLY}, maxPermutations-1).Slice()
	permutations = perm

	sum := 0
	for _, equation := range equations {
		if equation.isPossible() {
			sum += equation.Result
		}
	}

	return sum
}

func Day7B(filename string) int {
	equations := getEquations(filename)
	maxPermutations := 1

	for i := 0; i < len(equations); i++ {
		if maxPermutations < len(equations[i].Numbers) {
			maxPermutations = len(equations[i].Numbers)
		}
	}

	perm, _ := iterium.Product([]int{ADD, MULTIPLY, APPEND}, maxPermutations-1).Slice()
	permutations = perm

	sum := 0
	for _, equation := range equations {
		if equation.isPossible() {
			sum += equation.Result
		}
	}

	return sum
}
