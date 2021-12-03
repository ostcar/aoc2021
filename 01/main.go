package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/ostcar/aoc2021/aoc"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(task1())
	fmt.Println(task2())
}

func task1() string {
	input, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	last := 0
	var count int
	for _, line := range input {
		i, _ := strconv.Atoi(line)
		if last != 0 && i > last {
			count++
		}
		last = i
	}
	return fmt.Sprint(count)
}

func task2() string {
	input, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	ints := make([]int, len(input))
	for i := range input {
		number, _ := strconv.Atoi(input[i])
		ints[i] = number
	}

	var count int
	var last int
	for i := 0; i < len(ints)-2; i++ {
		n := ints[i] + ints[i+1] + ints[i+2]
		if last != 0 && n > last {
			count++
		}
		last = n
	}
	return fmt.Sprint(count)
}
