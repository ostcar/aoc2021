package main

import "testing"

var testInput = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "15" {
		t.Errorf("task1() == %s, expected 15", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "1134" {
		t.Errorf("task2() == %s, expected 1134", got)
	}
}
