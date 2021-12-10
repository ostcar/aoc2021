package main

import (
	_ "embed"
	"fmt"
	"sort"

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

	var count int
	for _, line := range lines {
		var s stack
		count += countLine(line, &s)
	}

	return fmt.Sprint(count)
}

func countLine(line string, s *stack) int {
	for i := 0; i < len(line); i++ {
		current := line[i]
		switch current {
		case '(', '[', '{', '<':
			s.push(current)
		case ')', ']', '}', '>':
			last := s.pop()
			if current != closing[last] {
				return point(current)
			}
		}
	}
	return 0
}

var closing = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

type stack struct {
	data []byte
}

func (s *stack) push(b byte) {
	s.data = append(s.data, b)
}

func (s *stack) pop() byte {
	if len(s.data) == 0 {
		return 0
	}

	b := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return b
}

func point(b byte) int {
	if b == ')' {
		return 3
	}
	if b == ']' {
		return 57
	}
	if b == '}' {
		return 1197
	}
	return 25137
}

func task2(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	var scores []int
	for _, line := range lines {
		var s stack
		if countLine(line, &s) != 0 {
			continue
		}

		var score int
		for {
			e := s.pop()
			if e == 0 {
				break
			}

			score *= 5
			score += task2Point(closing[e])
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return fmt.Sprint(scores[len(scores)/2])
}

func task2Point(b byte) int {
	if b == ')' {
		return 1
	}
	if b == ']' {
		return 2
	}
	if b == '}' {
		return 3
	}
	return 4
}
