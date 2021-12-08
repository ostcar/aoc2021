package main

import (
	_ "embed"
	"fmt"
	"strings"

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
		pipeIdx := strings.Index(line, "|")
		for _, element := range strings.Split(line[pipeIdx+1:], " ") {
			switch len(element) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	return fmt.Sprint(count)
}

func task2(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	var count int
	for _, line := range lines {
		pipeIdx := strings.Index(line, "|")
		elements := strings.Split(line[:pipeIdx]+line[pipeIdx+2:], " ")
		code := findCode(elements)
		count += elementToDigit(elements[len(elements)-4:], code)
	}

	return fmt.Sprint(count)
}

// validLetters are 7-bit combinations, that represent valid seven-segment
// displays.
//
// The top segment is the first bit, the segmet top left the second bit and so
// on.
//
// The 10 numbers in this map are the ints that are the result of the 10 letters
// that a sevent-segment can show.
var validLetters = map[uint8]bool{
	119: true, // 0: 1110111
	18:  true, // 1: 0010010
	93:  true, // 2: 1011101
	91:  true, // 3: 1011011
	58:  true, // 4: 0111010
	107: true, // 5: 1101011
	111: true, // 6: 1101111
	82:  true, // 7: 1010010
	127: true, // 8: 1111111
	123: true, // 9: 1111011
}

// letterToInt contains the same ten numbers then validLetters but directs to
// the number shon by the sevent-segment display.
var letterToInt = map[uint8]int{
	119: 0,
	18:  1,
	93:  2,
	91:  3,
	58:  4,
	107: 5,
	111: 6,
	82:  7,
	127: 8,
	123: 9,
}

// isValid taks a list of strings from the puzzle, like 'dbefa' and a docoding
// code and returns if this code fits.
func isValid(elements []string, code [7]uint8) bool {
	for _, e := range elements {
		if !validLetters[elementToInt(e, code)] {
			return false
		}
	}
	return true
}

// findCode tries every possible code to solve the puzzle. Returns the code that
// fits.
func findCode(elements []string) [7]uint8 {
	code := [7]uint8{0, 1, 2, 3, 4, 5, 6}
	for perm := [7]uint8{}; perm[0] < uint8(len(perm)); nextPerm(&perm) {
		permutedCode := getPerm(code, perm)
		if isValid(elements, permutedCode) {
			return permutedCode
		}
	}
	return [7]uint8{}
}

// elementToInt translates a puzzle-string like 'dbefa' to a bitcode that can be
// used by validLetters and letterToInt.
//
// code is used as a mapping from the letter 'a' to the bit possition.
func elementToInt(input string, code [7]uint8) uint8 {
	var n uint8
	for _, l := range input {
		n |= (1 << code[(l-'a')])
	}
	return n
}

// elementToDigit expects a slice with 4 elements and returns the int that would
// be shown on the sevent-elment-display.
func elementToDigit(input []string, code [7]uint8) int {
	return letterToInt[elementToInt(input[0], code)]*1000 +
		letterToInt[elementToInt(input[1], code)]*100 +
		letterToInt[elementToInt(input[2], code)]*10 +
		letterToInt[elementToInt(input[3], code)]
}

// getPerm returns a permutation of the given code.
//
// getPerm and nextPerm are from https://stackoverflow.com/a/30230552
func getPerm(code, perm [7]uint8) [7]uint8 {
	for i, v := range perm {
		code[i], code[i+int(v)] = code[i+int(v)], code[i]
	}
	return code
}

func nextPerm(p *[7]uint8) {
	for i := uint8(6); i >= 0; i-- {
		if i == 0 || p[i] < 6-i {
			p[i]++
			return
		}
		p[i] = 0
	}
}
