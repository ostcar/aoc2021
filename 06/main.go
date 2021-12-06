package main

import (
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
	return fishSim(input, 80)
}

func task2(input string) string {
	return fishSim(input, 256)
}

type day struct {
	adults int
	kids   int
}

// fishSim calculates how manys fishes there are after dayCount.
//
// It does not calculate each fish by itself, but puts all fishes, that are born
// on the same day in a group. All fish from the same day gets there children at
// the same day.
func fishSim(input string, dayCount int) string {
	fishes := strings.Split(strings.TrimSpace(input), ",")

	// Calculate days where a fish gets a child.
	days := make([]day, 7)
	for _, fish := range fishes {
		n, _ := strconv.Atoi(fish)
		days[n].adults++
	}

	for i := 0; i < dayCount; i++ {
		days[nextDay(i)].kids += days[i%7].adults
		days[i%7].adults += days[i%7].kids
		days[i%7].kids = 0
	}

	var count int
	for _, day := range days {
		count += day.adults + day.kids
	}

	return fmt.Sprint(count)
}

func nextDay(n int) int {
	return (n + 2) % 7
}
