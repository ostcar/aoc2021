package main

import (
	_ "embed"
	"fmt"

	"github.com/ostcar/aoc2021/aoc"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	_ = lines

	return ""
}

func task2(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	_ = lines

	return ""
}
