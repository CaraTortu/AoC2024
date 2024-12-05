package days

import (
	"log"
	"regexp"
	"strconv"

	"AoC24/utils"
)

func Day3A(filename string) int {
	content, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	total := 0

	mulRe := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	for _, line := range content {
		found := mulRe.FindAllStringSubmatch(line, -1)

		for _, expr := range found {
			a, _ := strconv.Atoi(expr[1])
			b, _ := strconv.Atoi(expr[2])
			total += a * b
		}
	}

	return total
}

func Day3B(filename string) int {
	content, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	total := 0
	canAdd := true
	mulRe := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)

	for _, line := range content {
		found := mulRe.FindAllStringSubmatch(line, -1)

		for _, expr := range found {
			if expr[0] == "do()" {
				canAdd = true
				continue
			} else if expr[0] == "don't()" {
				canAdd = false
				continue
			}

			if !canAdd {
				continue
			}

			a, _ := strconv.Atoi(expr[1])
			b, _ := strconv.Atoi(expr[2])
			total += a * b
		}
	}

	return total
}
