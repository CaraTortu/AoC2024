package utils

import (
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	// Read the file
	byte_contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	contents := strings.TrimSuffix(string(byte_contents), "\n")

	return strings.Split(contents, "\n"), nil
}
