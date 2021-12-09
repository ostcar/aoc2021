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

	width := len(lines[0])
	height := len(lines)
	board := make([]int, 0, width*height)
	for _, line := range lines {
		for _, v := range line {
			board = append(board, int(v-'0'))
		}
	}

	var count int
	for i, v := range board {
		if lowestNeighbor(board, i, width, height) {
			count += 1 + v
		}
	}

	return fmt.Sprint(count)
}

func lowestNeighbor(board []int, index int, width, height int) bool {
	value := board[index]
	if top := index - width; top >= 0 {
		if topValue := board[top]; topValue <= value {
			return false
		}
	}
	if index%width > 0 {
		if leftValue := board[index-1]; leftValue <= value {
			return false
		}
	}
	if index%width < width-1 {
		if rightValue := board[index+1]; rightValue <= value {
			return false
		}
	}
	if bottom := index + width; bottom < len(board) {
		if bottomValue := board[bottom]; bottomValue <= value {
			return false
		}
	}

	return true
}

func task2(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	width := len(lines[0])
	height := len(lines)
	board := make([]int, 0, width*height)
	for _, line := range lines {
		for _, v := range line {
			board = append(board, int(v-'0'))
		}
	}

	basins := findBasins(board, width, height)

	// Get the heighest 3 basins.
	sort.Ints(basins)
	v := basins[len(basins)-3] * basins[len(basins)-2] * basins[len(basins)-1]

	return fmt.Sprint(v)
}

// fincBasis returns the size of every basin on the board.
func findBasins(board []int, width, height int) []int {
	var basins []int
	for i, value := range board {
		if value == 9 {
			continue
		}

		basins = append(basins, walkBasin(board, width, height, i))
	}

	return basins
}

// walkBasin walks on the board and returns how big the basin is. It changes the
// board by replacing all places of the basin with a 9
func walkBasin(board []int, width, height int, start int) int {
	if board[start] == 9 {
		return 0
	}

	count := 1
	board[start] = 9
	for _, neighbor := range neighborIndex(board, width, height, start) {
		count += walkBasin(board, width, height, neighbor)
	}
	return count
}

// neighborIndex returns the value of all valid neighbors.
//
// Does not return indexes that are outside of the board and only returns
// indexes for values, that are not 9 (are in the same basin).
func neighborIndex(board []int, width, height int, index int) []int {
	var neighbor []int
	if top := index - width; top >= 0 {
		if topValue := board[top]; topValue != 9 {
			neighbor = append(neighbor, top)
		}
	}
	if index%width > 0 {
		if leftValue := board[index-1]; leftValue != 9 {
			neighbor = append(neighbor, index-1)
		}
	}
	if index%width < width-1 {
		if rightValue := board[index+1]; rightValue != 9 {
			neighbor = append(neighbor, index+1)
		}
	}
	if bottom := index + width; bottom < len(board) {
		if bottomValue := board[bottom]; bottomValue != 9 {
			neighbor = append(neighbor, bottom)
		}
	}
	return neighbor
}
