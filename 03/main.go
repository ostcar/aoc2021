package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"

	"github.com/ostcar/aoc2021/aoc"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(in string) string {
	input, err := aoc.ReadInput(in)
	if err != nil {
		return err.Error()
	}

	bits := len(input[0])
	lineCount := len(input)

	counters := make([]int, bits)
	for _, line := range input {
		for i, c := range line {
			if c == '1' {
				counters[i]++
			}
		}
	}

	var value int
	for _, v := range counters {
		value <<= 1
		if v > lineCount/2 {
			value++
		}
	}

	// The second value can be calculated from the first.
	maxNumber := (1 << bits) - 1
	result := value * (maxNumber - value)

	return fmt.Sprint(result)
}

func task2(in string) string {
	input, err := aoc.ReadInput(in)
	if err != nil {
		return err.Error()
	}

	sort.Strings(input)
	v1 := task2Filter(input, func(a, b int) bool { return a > b })

	v2 := task2Filter(input, func(a, b int) bool { return a <= b })

	return fmt.Sprint(v1 * v2)
}

// task2Filter applys the filter for the second task.
//
// For this filter to work, the input values have to be sorted.
//
// Look for the index of the first value '1'. If it is on the first half of the
// list, then there are more values with '1'. If it is on the second half, then
// there are more values with '0'.
//
// cmp is a function to check which half to use. It should be '>' for the first
// value and '<=' for the second value.
func task2Filter(input []string, cmp func(a, b int) bool) int {
	for i := 0; i < len(input[0]) && len(input) > 1; i++ {
		cut := sort.Search(len(input), func(n int) bool {
			return input[n][i] == '1'
		})

		if cmp(cut, len(input)/2) {
			input = input[:cut]
		} else {
			input = input[cut:]
		}
	}

	value, _ := strconv.ParseInt(input[0], 2, 64)
	return int(value)
}
