package main

import (
	"fmt"
	"testing"
)

var testInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`

func TestTask1(t *testing.T) {
	if got := task1(testInput); got != "26" {
		t.Errorf("task1() returned %q, expected 26", got)
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput); got != "61229" {
		t.Errorf("task2() returned %q, expected 61229", got)
	}
}

func TestElementToInt(t *testing.T) {
	for _, tt := range []struct {
		letters string
		mapper  [7]byte
		value   string
	}{
		{"abc", [7]byte{0, 1, 2, 3, 4, 5, 6}, "111"},
		{"b", [7]byte{0, 1, 2, 3, 4, 5, 6}, "10"},
		{"gdcbef", [7]byte{0, 1, 2, 3, 4, 5, 6}, "1111110"},
		{"abc", [7]byte{1, 2, 3, 4, 5, 6, 0}, "1110"},
	} {
		t.Run(tt.letters, func(t *testing.T) {
			got := fmt.Sprintf("%b", elementToInt(tt.letters, tt.mapper))
			if got != tt.value {
				t.Errorf("elementToInt(%s) == %s, expected %s", tt.letters, got, tt.value)
			}
		})
	}
}
