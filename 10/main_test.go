package main

import "testing"

var testInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "26397" {
		t.Errorf("task1() == %s, expected 26397", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "288957" {
		t.Errorf("task2() == %s, expected 288957", got)
	}
}
