package main

import "testing"

var testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "5" {
		t.Errorf("task1() returned %q, expected 5", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "12" {
		t.Errorf("task1() returned %q, expected 12", got)
	}
}

func BenchmarkTask2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		task2(puzzleInput)
	}
}
