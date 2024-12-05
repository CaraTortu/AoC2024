package days

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"AoC24/utils"
)

func abs_diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func getLists(filename string) ([]int, []int) {
	// Read the file
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Loop through the contents of the file
	listA := make([]int, len(contents))
	listB := make([]int, len(contents))

	for i, line := range contents {
		numbers := strings.Split(line, "   ")

		listA[i], err = strconv.Atoi(strings.TrimSuffix(numbers[0], "\n"))
		if err != nil {
			log.Fatalf("Error converting string to int: %v", numbers[0])
		}

		listB[i], err = strconv.Atoi(strings.TrimSuffix(numbers[1], "\n"))
		if err != nil {
			log.Fatalf("Error converting string to int: %v", numbers[1])
		}

	}

	return listA, listB
}

func Day1A(filename string) int {
	distance := 0

	listA, listB := getLists(filename)

	// Sort each array
	sort.Ints(listA)
	sort.Ints(listB)

	// Find the distance between the two arrays
	for i := 0; i < len(listA); i++ {
		distance += abs_diff(listA[i], listB[i])
	}

	return distance
}

func Day1B(filename string) int {
	distance := 0

	listA, listB := getLists(filename)

	// Get occurences for listB
	occurences := make(map[int]int)
	for _, num := range listB {

		// If the number is not in the map, add it
		if _, ok := occurences[num]; !ok {
			occurences[num] = 0
		}

		occurences[num]++
	}

	// For each item of listA, find the number of occurences in occurences
	// and add this value to the distance
	for _, num := range listA {
		occurence, ok := occurences[num]

		if !ok {
			continue
		}

		distance += occurence * num
	}

	return distance
}
