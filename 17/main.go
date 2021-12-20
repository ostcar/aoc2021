package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	area := parse(input)

	minX, maxX := minMaxX(area)

	for y := 100; y > 0; y-- {
		for x := minX; x <= maxX; x++ {
			if success(point{x, y}, area) {
				return fmt.Sprint(y * (y + 1) / 2)
			}
		}
	}

	return ""
}

func task2(input string) string {
	area := parse(input)

	minX, maxX := minMaxX(area)
	maxX = area[1].x
	minY := area[1].y

	var count int
	for y := 100; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if success(point{x, y}, area) {
				count++
			}
		}
	}

	return fmt.Sprint(count)
}

func parse(input string) [2]point {
	var min, max point
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &min.x, &max.x, &max.y, &min.y)
	return [2]point{min, max}
}

func minMaxX(area [2]point) (minX, maxX int) {
	// Find a n where 1+2+3+...+n is smaller or equal area[0].x
	for n := 1; n*(n+1)/2 <= area[0].x; n++ {
		minX = n
	}

	// use a quoter from the area as a max value.
	maxX = area[1].x / 4
	return minX, maxX
}

func success(velocity point, area [2]point) bool {
	var probe point
	for {
		probe.x += velocity.x
		probe.y += velocity.y

		if probe.x > area[1].x || probe.y < area[1].y {
			return false
		}

		if probe.x >= area[0].x && probe.x <= area[1].x && probe.y <= area[0].y && probe.y >= area[1].y {
			return true
		}

		if velocity.x > 0 {
			velocity.x--
		}

		velocity.y--
	}
}

type point struct {
	x, y int
}
