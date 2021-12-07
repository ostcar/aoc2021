package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	values := strings.Split(strings.TrimSpace(input), ",")

	numbers := make([]int, len(values))
	for i := range values {
		numbers[i], _ = strconv.Atoi(values[i])
	}

	var last = int(^uint(0) >> 1) // heihest int
	for i := 0; ; i++ {
		f := fuelTask1(numbers, i)
		if f > last {
			// F will get smaller and smaller until it gets bigger. Stop the
			// loop if the numbers get bigger again.
			break
		}
		last = f
	}

	return fmt.Sprint(last)
}

func task2(input string) string {
	values := strings.Split(strings.TrimSpace(input), ",")

	numbers := make([]int, len(values))
	for i := range values {
		numbers[i], _ = strconv.Atoi(values[i])
	}

	var last = int(^uint(0) >> 1) // heihest int
	for i := 0; ; i++ {
		f := fuelTask2(numbers, i)
		if f > last {
			// F will get smaller and smaller until it gets bigger. Stop the
			// loop if the numbers get bigger again.
			break
		}
		last = f
	}

	return fmt.Sprint(last)
}

func fuelTask1(numbers []int, to int) int {
	var count int
	for _, number := range numbers {
		count += abs(number - to)
	}
	return count
}

func fuelTask2(numbers []int, to int) int {
	var count int
	for _, number := range numbers {
		count += fuelStepCost(abs(number - to))
	}
	return count
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

// fuelStepCost calculates 1+2+3+n
func fuelStepCost(steps int) int {
	return steps * (steps + 1) / 2
}
