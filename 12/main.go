package main

import (
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
	start := parseCaves(input)
	ways := start.startWalk(hasDoupleTask1)
	return fmt.Sprint(ways)
}

func task2(input string) string {
	start := parseCaves(input)
	ways := start.startWalk(hasDoupleTask2)
	return fmt.Sprint(ways)
}

// parseCaves parses the input and returns the start cave.
//
// The start cave is linked to its neighbor cave and each other cave to its
// neighbor caves. No cave is linked to the start cave again.
func parseCaves(input string) *cave {
	lines := strings.Split(input, "\n")

	caves := make(map[string]*cave)
	for _, line := range lines {
		if line == "" {
			continue
		}

		idx := strings.Index(line, "-")
		caveName1 := line[:idx]
		caveName2 := line[idx+1:]

		cave1 := caves[caveName1]
		if cave1 == nil {
			cave1 = new(cave)
			cave1.name = caveName1
			caves[caveName1] = cave1
		}

		cave2 := caves[caveName2]
		if cave2 == nil {
			cave2 = new(cave)
			cave2.name = caveName2
			caves[caveName2] = cave2
		}

		if caveName2 != "start" {
			cave1.links = append(cave1.links, cave2)
		}

		if caveName1 != "start" {
			cave2.links = append(cave2.links, cave1)
		}
	}

	return caves["start"]
}

type cave struct {
	name  string
	links []*cave
}

// startWalk has to be called from the start cave and returns the number of
// possible walks though the caves.
func (c cave) startWalk(inSlice func([]string) bool) int {
	var before []string
	var count int
	for _, link := range c.links {
		count += link.walk(before, inSlice)
	}
	return count
}

// walk is like startWalk but also takes a list of cave-names where that where
// visited before.
func (c cave) walk(before []string, hasDouple func([]string) bool) int {
	if c.name == "end" {
		return 1
	}

	before = append(before, c.name)
	defer func() { before = before[:len(before)-1] }()

	if hasDouple(before) {
		return 0
	}

	var count int
	for _, link := range c.links {
		count += link.walk(before, hasDouple)
	}

	return count
}

func (c cave) String() string {
	return c.name
}

// hasDoupleTask1 returns true, if there is a small cave more then once in the
// given list.
func hasDoupleTask1(slice []string) bool {
	set := make(map[string]struct{})
	for _, v := range slice {
		if !isLow(v) {
			continue
		}

		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

// hasDoupleTask2 returns true, if there are more then one small cave more then
// once in the given list.
func hasDoupleTask2(slice []string) bool {
	set := make(map[string]struct{})
	hasDouple := false
	for _, v := range slice {
		if !isLow(v) {
			continue
		}

		if _, ok := set[v]; ok {
			if hasDouple {
				return true
			}
			hasDouple = true
		}
		set[v] = struct{}{}
	}
	return false
}

// isLow returns true, if the given string starts with a lower letter (not
// unicode save).
func isLow(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}
