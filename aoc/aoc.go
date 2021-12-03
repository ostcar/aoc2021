package aoc

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadInput reads the input and return each line.
func ReadInput(input string) ([]string, error) {
	var out []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning content: %w", err)
	}
	return out, nil
}
