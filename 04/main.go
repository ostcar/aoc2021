package main

import (
	"bytes"
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
	parts := strings.Split(input, "\n\n")

	var boards []bingo
	for _, part := range parts[1:] {
		boards = append(boards, *newBingo(part))
	}

	for _, step := range strings.Split(parts[0], ",") {
		n, _ := strconv.Atoi(step)

		for _, b := range boards {
			if b.call(byte(n)) {
				return fmt.Sprint(b.points() * n)
			}
		}
	}

	return "no boards have won :("
}

func task2(input string) string {
	parts := strings.Split(input, "\n\n")

	var boards []bingo
	for _, part := range parts[1:] {
		boards = append(boards, *newBingo(part))
	}

	for _, step := range strings.Split(parts[0], ",") {
		n, _ := strconv.Atoi(step)

		for i := 0; i < len(boards); i++ {
			if boards[i].call(byte(n)) {
				if len(boards) == 1 {
					return fmt.Sprint(boards[i].points() * n)
				}

				// Remove the winning board from the list of boards. Reset the
				// index by one to match the new order.
				boards = append(boards[:i], boards[i+1:]...)
				i--
			}
		}
	}

	return "no boards have won :("
}

type bingo struct {
	board []byte // use a byte slice instead of a int slice, so the bytes package can be used.
}

func newBingo(input string) *bingo {
	b := bingo{
		board: make([]byte, 25),
	}
	for i, field := range strings.Fields(input) {
		number, _ := strconv.Atoi(field)
		b.board[i] = byte(number)
	}
	return &b
}

// call removes the given number from the bord.
//
// returns true, if the bord is finished (one row or one collumn are done).
func (b *bingo) call(n byte) bool {
	idx := bytes.IndexByte(b.board, n)
	if idx == -1 {
		return false
	}

	b.board[idx] = 255

	return b.lineFinished(idx/5) || b.columnFinished(idx%5)
}

// lineFinished returns true, if all numbers of the line are finished.
//
// line has to be a number between 0 and 4.
func (b *bingo) lineFinished(line int) bool {
	return b.board[5*line] == 255 &&
		b.board[5*line+1] == 255 &&
		b.board[5*line+2] == 255 &&
		b.board[5*line+3] == 255 &&
		b.board[5*line+4] == 255
}

// columnFinished returns true, if all numbers of the column are finished.
//
// column has to be a number between 0 and 4.
func (b *bingo) columnFinished(column int) bool {
	return b.board[column] == 255 &&
		b.board[column+5] == 255 &&
		b.board[column+10] == 255 &&
		b.board[column+15] == 255 &&
		b.board[column+20] == 255
}

func (b *bingo) points() int {
	var count int
	for _, b := range b.board {
		if b != 255 {
			count += int(b)
		}
	}
	return count
}
