package main

import "testing"

const testInput = `16,1,2,0,4,2,7,1,2,14`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "37" {
		t.Errorf("task1() returned %s, expected 37", got)
	}
}
