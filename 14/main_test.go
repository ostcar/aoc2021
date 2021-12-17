package main

import "testing"

const testInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "1588" {
		t.Errorf("task1() == %s, expected 1588", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "2188189693529" {
		t.Errorf("task2() == %s, expected 2188189693529", got)
	}
}
