package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	b := newBoard(input)

	var count int
	//fmt.Printf("Before any steps:\n%v\n", b)
	for i := 0; i < 100; i++ {
		count += b.step()
		//fmt.Printf("After step %d:\n%v\n", i+1, b)
	}

	return fmt.Sprint(count)
}

func task2(input string) string {
	b := newBoard(input)

	var i int
	for {
		if b.step() == 100 {
			break
		}
		i++
	}

	return fmt.Sprint(i + 1)
}

type board struct {
	octopus []int
}

func newBoard(input string) *board {
	b := board{
		octopus: make([]int, 100),
	}
	for i, line := range strings.Split(input, "\n") {
		for j := range line {
			b.octopus[i*10+j] = int(line[j] - '0')
		}
	}
	return &b
}

func (b *board) step() int {
	var count int
	for i := 0; i < 100; i++ {
		b.octopus[i]++
	}

	for i := 0; i < 100; i++ {
		count += b.checkFlush(i, true)
	}
	return count
}

// checkFlush checks if a flush is necessary and returns the number of flushes
// it caused.
//
// checkFlush is called recursive. `orig` is true on the "outer" call and false
// on each recursive call.
func (b *board) checkFlush(idx int, orig bool) int {
	if b.octopus[idx] == 0 {
		return 0
	}

	if !orig {
		b.octopus[idx]++
	}

	var count int
	if b.octopus[idx] > 9 {
		count++
		b.octopus[idx] = 0

		for _, neighbor := range neighborIndex(idx) {
			count += b.checkFlush(neighbor, false)
		}
	}
	return count
}

// neighbors returns the indexes of all neighbors of a field.
func neighborIndex(index int) []int {
	var neighbor []int
	if top := index - 10; top >= 0 {
		neighbor = append(neighbor, top)
		if top%10 > 0 {
			neighbor = append(neighbor, top-1)
		}
		if top%10 < 9 {
			neighbor = append(neighbor, top+1)
		}
	}
	if index%10 > 0 {
		neighbor = append(neighbor, index-1)
	}
	if index%10 < 9 {
		neighbor = append(neighbor, index+1)
	}
	if bottom := index + 10; bottom < 100 {
		neighbor = append(neighbor, bottom)
		if bottom%10 > 0 {
			neighbor = append(neighbor, bottom-1)
		}
		if bottom%10 < 9 {
			neighbor = append(neighbor, bottom+1)
		}
	}
	return neighbor
}

func (b *board) String() string {
	buf := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Fprintf(buf, "%d", b.octopus[i*10+j])
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}
