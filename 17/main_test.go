package main

import (
	"fmt"
	"testing"
)

const testInput = `target area: x=20..30, y=-10..-5
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "45" {
		t.Errorf("task1() == %q, expected 45", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "112" {
		t.Errorf("task2() == %q, expected 112", got)
	}
}

func TestSuccess(t *testing.T) {
	area := parse("target area: x=20..30, y=-10..-5")
	for _, tt := range []struct {
		area     [2]point
		velocity point
		expect   bool
	}{
		{area, point{7, 2}, true},
		{area, point{6, 3}, true},
		{area, point{3, 3}, false},
		{area, point{9, 0}, true},
		{area, point{17, -4}, false},
	} {
		t.Run(fmt.Sprint(tt.area), func(t *testing.T) {
			got := success(tt.velocity, tt.area)
			if got != tt.expect {
				t.Errorf("success(%v, %v) == %v, expected %v", tt.velocity, tt.area, got, tt.expect)
			}
		})
	}
}
