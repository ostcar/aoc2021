package main

import "testing"

const testInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "40" {
		t.Errorf("task1() == %s, expected 40", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "315" {
		t.Errorf("task2() == %s, expected 315", got)
	}
}
