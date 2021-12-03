package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

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

	var forward, depth int
	for _, line := range input {
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			forward += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return fmt.Sprint(forward * depth)
}

func task2() string {
	input, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	var forward, depth, aim int
	for _, line := range input {
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			forward += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return fmt.Sprint(forward * depth)
}
