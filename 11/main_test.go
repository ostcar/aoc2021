package main

import (
	"fmt"
	"sort"
	"testing"
)

const testInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "1656" {
		t.Errorf("task1() == %s, expected 1656", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "195" {
		t.Errorf("task2() == %s, expected 195", got)
	}
}

func TestNeighborIndex(t *testing.T) {
	for _, tt := range []struct {
		idx    int
		expect []int
	}{
		{0, []int{1, 10, 11}},
		{1, []int{0, 2, 10, 11, 12}},
		{2, []int{1, 3, 11, 12, 13}},
		{9, []int{8, 18, 19}},
		{11, []int{0, 1, 2, 10, 12, 20, 21, 22}},
	} {
		t.Run(fmt.Sprintf("%d", tt.idx), func(t *testing.T) {
			got := neighborIndex(tt.idx)
			sort.Ints(got)

			if !sliceEqual(got, tt.expect) {
				t.Errorf("neighborIndex(%d) == %v, expected %v", tt.idx, got, tt.expect)
			}
		})
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
