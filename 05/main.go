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
	values := parse(input)
	grid := make(map[point]int)

	for _, pointPair := range values {
		if pointPair[0].x != pointPair[1].x && pointPair[0].y != pointPair[1].y {
			// Skip points that are in the same line or same column.
			continue
		}

		v := vec(pointPair[0], pointPair[1])
		for p := pointPair[0]; p != pointPair[1]; {
			grid[p]++
			p.x += v.x
			p.y += v.y
		}
		grid[pointPair[1]]++
	}

	var counter int
	for _, v := range grid {
		if v > 1 {
			counter++
		}
	}

	return fmt.Sprint(counter)
}

func task2(input string) string {
	values := parse(input)
	grid := make(map[point]int)

	for _, pointPair := range values {
		v := vec(pointPair[0], pointPair[1])
		for p := pointPair[0]; p != pointPair[1]; {
			grid[p]++
			p.x += v.x
			p.y += v.y
		}
		grid[pointPair[1]]++

	}

	var counter int
	for _, v := range grid {
		if v > 1 {
			counter++
		}
	}

	return fmt.Sprint(counter)
}

func parse(input string) [][2]point {
	lines := strings.Split(input, "\n")

	out := make([][2]point, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		fmt.Sscanf(
			line,
			"%d,%d -> %d,%d",
			&out[i][0].x,
			&out[i][0].y,
			&out[i][1].x,
			&out[i][1].y,
		)
	}
	return out
}

// vec returns the direction from p1 to p2.
//
// The x and y values of the returned point are -1, 0 or 1.
func vec(p1, p2 point) point {
	return point{
		x: diff(p1.x, p2.x),
		y: diff(p1.y, p2.y),
	}
}

// diff returns if a or b is bigger
//
// -1 if smaler, 0 if the same and 1 if bigger
func diff(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return 1
	}
	return -1
}

type point struct {
	x, y int
}

// showGrid parses a grid to a string.
//
// Only for debugigng.
func showGrid(grid map[point]int) string {
	var maxX, maxY int
	for p := range grid {
		if maxX < p.x {
			maxX = p.x
		}
		if maxY < p.y {
			maxY = p.y
		}
	}

	buf := new(bytes.Buffer)
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if v := grid[point{x, y}]; v > 0 {
				fmt.Fprintf(buf, "%d", v)
				continue
			}
			fmt.Fprint(buf, ".")
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}
