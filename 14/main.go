package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	template, methods := parse(input)

	pairs := process(template, methods, 10)

	max, min := minMax(pairs, template[len(template)-1])

	return fmt.Sprint(max - min)
}

func task2(input string) string {
	template, methods := parse(input)

	pairs := process(template, methods, 40)

	max, min := minMax(pairs, template[len(template)-1])

	return fmt.Sprint(max - min)
}

func parse(input string) (string, map[[2]byte]byte) {
	lines := strings.Split(input, "\n")
	methods := make(map[[2]byte]byte)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		methods[[2]byte{line[0], line[1]}] = line[6]
	}
	return lines[0], methods
}

func process(template string, methods map[[2]byte]byte, count int) map[[2]byte]uint {
	pairs := make(map[[2]byte]uint)
	for i := 0; i < len(template)-1; i++ {
		pairs[[2]byte{template[i], template[i+1]}]++
	}

	for i := 0; i < count; i++ {
		nextPairs := make(map[[2]byte]uint)

		for pair, count := range pairs {
			n := methods[pair]
			nextPairs[[2]byte{pair[0], n}] += count
			nextPairs[[2]byte{n, pair[1]}] += count
		}
		pairs = nextPairs
	}

	return pairs
}

func minMax(pairs map[[2]byte]uint, extra byte) (uint, uint) {
	counter := make(map[byte]uint)
	for pair, count := range pairs {
		counter[pair[0]] += count
	}
	counter[extra]++

	var max uint
	var min uint = math.MaxUint
	for _, v := range counter {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return max, min
}
