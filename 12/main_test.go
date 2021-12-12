package main

import (
	"fmt"
	"strings"
	"testing"
)

const testInput = `start-A
start-b
A-b
A-c
b-d
A-end
b-end
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "10" {
		t.Errorf("task1() == %s, expected 10", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "36" {
		t.Errorf("task2() == %s, expected 36", got)
	}
}

func TestInSliceTask2(t *testing.T) {
	for i, tt := range []struct {
		before string
		expect bool
	}{
		{"A,b,A,b,A,b", true},
		{"A,b,A,b", false},
		{"b,A,b,A,b", true},
		{"A,b,A,b,A,c,A,c", true},
	} {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			got := hasDoupleTask2(strings.Split(tt.before, ","))
			if got != tt.expect {
				t.Errorf("inSliceTask2(%v) == %t, expected %t", tt.before, got, tt.expect)
			}
		})
	}
}
