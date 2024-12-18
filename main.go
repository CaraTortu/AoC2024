package main

import (
	"fmt"
	"time"

	"AoC24/days"
)

func main() {
	/*f, perr := os.Create("cpu.pprof")
	if perr != nil {
		log.Fatal(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	*/
	day_funcs := []func(string) int{
		days.Day1A, days.Day1B,
		days.Day2A, days.Day2B,
		days.Day3A, days.Day3B,
		days.Day4A, days.Day4B,
		days.Day5A, days.Day5B,
		days.Day6A, days.Day6B,
		days.Day7A, days.Day7B,
	}

	for day, solver := range day_funcs {
		actual_day := day/2 + 1
		letter := ""

		if day%2 == 0 {
			letter = "A"
		} else {
			letter = "B"
		}

		start := time.Now()

		result := solver(fmt.Sprintf("./inputs/day%d.txt", actual_day))
		elapsed := time.Since(start)
		fmt.Printf("[i] Day %d%s took %v: %d\n", actual_day, letter, elapsed, result)
	}
}
