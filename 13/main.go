package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

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

	points := make(map[point]struct{})
	cutLines := false
	for _, line := range lines {
		if line == "" {
			cutLines = true
			continue
		}

		if cutLines {
			n, _ := strconv.Atoi(line[13:])
			if line[11] == 'x' {
				foldX(points, n)
			} else {
				foldY(points, n)
			}
			return fmt.Sprint(len(points))
		}

		var p point
		fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		points[p] = struct{}{}
	}

	return ""
}

func task2(input string) string {
	lines, err := aoc.ReadInput(input)
	if err != nil {
		return err.Error()
	}

	points := make(map[point]struct{})
	cutLines := false
	for _, line := range lines {
		if line == "" {
			cutLines = true
			continue
		}

		if cutLines {
			n, _ := strconv.Atoi(line[13:])
			if line[11] == 'x' {
				foldX(points, n)
			} else {
				foldY(points, n)
			}
			continue
		}

		var p point
		fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		points[p] = struct{}{}
	}

	return printPoints(points)
}

type point struct {
	x, y int
}

func foldY(points map[point]struct{}, line int) {
	// The idea is, that in go it is possible to edit a map while iterating on
	// it.
	for p := range points {
		if p.y < line {
			continue
		}
		points[point{
			y: 2*line - p.y,
			x: p.x,
		}] = struct{}{}
		delete(points, p)
	}
}

func foldX(points map[point]struct{}, line int) {
	for p := range points {
		if p.x < line {
			continue
		}
		points[point{
			x: 2*line - p.x,
			y: p.y,
		}] = struct{}{}
		delete(points, p)
	}
}

func printPoints(points map[point]struct{}) string {
	var maxX, maxY int
	for p := range points {
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
			if _, ok := points[point{x, y}]; ok {
				fmt.Fprintf(buf, "#")
				continue
			}
			fmt.Fprint(buf, " ")
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}
