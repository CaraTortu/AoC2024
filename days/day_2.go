package days

import (
	"log"
	"strconv"
	"strings"

	"AoC24/utils"
)

func day2GetReports(filename string) [][]int {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lists := make([][]int, len(contents))

	for i, line := range contents {
		numbers := strings.Split(line, " ")

		lists[i] = make([]int, len(numbers))

		for j, n := range numbers {
			lists[i][j], err = strconv.Atoi(strings.TrimSuffix(n, "\n"))
			if err != nil {
				log.Fatalf("Error converting string to int: %v", n)
			}
		}
	}

	return lists
}

func day2IsValidReport(report []int) bool {
	increasing := report[0] < report[1]
	valid := true
	for i := 1; i < len(report); i++ {
		// Check that it keeps increasing
		if increasing && report[i-1] > report[i] {
			valid = false
			break
		}

		// Check that it keeps decreasing
		if !increasing && report[i-1] < report[i] {
			valid = false
			break
		}

		// Check that the difference is between 1 and 3
		diff := abs_diff(report[i-1], report[i])
		if diff > 3 || diff < 1 {
			valid = false
			break
		}
	}

	return valid
}

func Day2A(filename string) int {
	reports := day2GetReports(filename)

	safe_reports := 0

	for _, report := range reports {
		if day2IsValidReport(report) {
			safe_reports++
		}
	}

	return safe_reports
}

func Day2B(filename string) int {
	reports := day2GetReports(filename)

	safe_reports := 0
	for _, report := range reports {
		if day2IsValidReport(report) {
			safe_reports++
			continue
		}

		// Try removing one element at a time
		// Naive approach, but it works
		for i := 0; i < len(report); i++ {
			new_report := make([]int, len(report)-1)
			copy(new_report, report[:i])
			copy(new_report[i:], report[i+1:])

			if day2IsValidReport(new_report) {
				safe_reports++
				break
			}
		}
	}

	return safe_reports
}
