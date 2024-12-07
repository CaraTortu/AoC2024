package days

import (
	"log"
	"strconv"
	"strings"

	"AoC24/utils"
)

func day5CheckReport(report []int, rules map[int][]int) bool {
	for i := 0; i < len(report)-1; i++ {
		page := report[i]
		pageToCompare := report[i+1]

		pageRules, ok := rules[page]
		if !ok {
			return false
		}

		found := false
		for _, rule := range pageRules {
			if pageToCompare == rule {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func day5GetRulesAndReports(filename string) (map[int][]int, [][]int) {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Get rules
	rules := make(map[int][]int)
	printIndex := -1
	for i, line := range contents {
		if line == "" {
			printIndex = i + 1
			break
		}

		rule := strings.Split(line, "|")
		ruleNum, _ := strconv.Atoi(rule[0])
		ruleVal, _ := strconv.Atoi(rule[1])

		rules[ruleNum] = append(rules[ruleNum], ruleVal)
	}

	// Get reports
	reports := make([][]int, len(contents[printIndex:]))
	for i, report := range contents[printIndex:] {
		numbersStr := strings.Split(report, ",")
		numbers := make([]int, len(numbersStr))

		for i, num := range numbersStr {
			num, _ := strconv.Atoi(num)
			numbers[i] = num
		}

		reports[i] = numbers
	}

	return rules, reports
}

func sortReport(report []int, rules map[int][]int) {
	for {
		if day5CheckReport(report, rules) {
			return
		}

		for i := 0; i < len(report)-1; i++ {
			page := report[i]
			pageToCompare := report[i+1]

			pageRules, ok := rules[page]
			if !ok {
				report[i], report[i+1] = report[i+1], report[i]
				continue
			}

			found := false
			for _, rule := range pageRules {
				if pageToCompare == rule {
					found = true
					break
				}
			}

			if !found {
				report[i], report[i+1] = report[i+1], report[i]
			}
		}
	}
}

func Day5A(filename string) int {
	rules, reports := day5GetRulesAndReports(filename)
	validReports := 0

	for _, report := range reports {
		if day5CheckReport(report, rules) {
			validReports += report[len(report)/2]
		}
	}

	return validReports
}

func Day5B(filename string) int {
	rules, reports := day5GetRulesAndReports(filename)
	validReports := 0

	for _, report := range reports {

		if day5CheckReport(report, rules) {
			continue
		}

		sortReport(report, rules)
		validReports += report[len(report)/2]
	}

	return validReports
}
