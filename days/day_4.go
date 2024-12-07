package days

import (
	"log"
	"regexp"
	"strings"

	"AoC24/utils"
)

func day4FindSubstrings(s, sub string) int {
	count := 0
	subLen := len(sub)
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		if sLen-i < subLen {
			break
		}
		if s[i:i+subLen] == sub {
			count++
		}
	}
	return count
}

func day4GetWindow(contents []string, i, j, x, y int) string {
	window := ""
	for k := 0; k < x; k++ {
		for l := 0; l < y; l++ {
			window += string(contents[i+k][j+l])
		}
	}
	return window
}

func Day4A(filename string) int {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	count := 0

	// Check for horizontal matches
	horizontal := strings.Join(contents, " ")
	count += day4FindSubstrings(horizontal, "XMAS")
	count += day4FindSubstrings(horizontal, "SAMX")

	// Check for vertical matches
	vertical := ""
	for i := 0; i < len(contents[0]); i++ {
		for j := 0; j < len(contents); j++ {
			vertical += string(contents[j][i])
		}
		vertical += " "
	}

	count += day4FindSubstrings(vertical, "XMAS")
	count += day4FindSubstrings(vertical, "SAMX")

	// Check for diagonal matches
	// Diagonal from top left to bottom right
	diagonal := ""
	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents); j++ {
			for k := 0; k < len(contents[j]); k++ {
				if j == k && j+i < len(contents) && k+i < len(contents[j]) {
					diagonal += string(contents[j+i][k])
				}
			}
		}
		diagonal += " "
	}

	for i := 1; i < len(contents); i++ {
		for j := 0; j < len(contents); j++ {
			for k := 0; k < len(contents[j]); k++ {
				if j == k && j+i < len(contents) && k+i < len(contents[j]) {
					diagonal += string(contents[k][j+i])
				}
			}
		}
		diagonal += " "
	}

	count += day4FindSubstrings(diagonal, "XMAS")
	count += day4FindSubstrings(diagonal, "SAMX")

	// Diagonal from top right to bottom left
	diagonal = ""
	for c := 0; c < len(contents); c++ {
		for k := 0; k < len(contents[c]); k++ {
			if c-k >= 0 {
				diagonal += string(contents[k][c-k])
			}
		}
		diagonal += " "
	}

	for c := 1; c < len(contents); c++ {
		for k := 0; k < len(contents[c]); k++ {
			if c+k < len(contents) {
				diagonal += string(contents[c+k][len(contents[c])-1-k])
			}
		}
		diagonal += " "
	}

	count += day4FindSubstrings(diagonal, "XMAS")
	count += day4FindSubstrings(diagonal, "SAMX")

	return count
}

func Day4B(filename string) int {
	contents, err := utils.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	reg := regexp.MustCompile(`M.S.A.M.S|M.M.A.S.S|S.S.A.M.M|S.M.A.S.M`)
	count := 0

	for i := 0; i < len(contents)-2; i++ {
		for j := 0; j < len(contents[i])-2; j++ {
			window := day4GetWindow(contents, i, j, 3, 3) // 3x3 window
			count += len(reg.FindAllString(window, -1))
		}
	}

	return count
}
