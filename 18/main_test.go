package main

import (
	"fmt"
	"strings"
	"testing"
)

const testInput1 = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]
`

const testInput2 = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

func TestTask1(t *testing.T) {
	for i, tt := range []struct {
		testInput string
		expect    string
	}{
		{testInput1, "3488"},
		{testInput2, "4140"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := task1(tt.testInput); got != tt.expect {
				t.Errorf("task1() == %q, expected %q", got, tt.expect)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	if got := task2(testInput2); got != "3993" {
		t.Errorf("task2() == %q, expected 3993", got)
	}
}

func TestExplode(t *testing.T) {
	for _, tt := range []struct {
		before string
		after  string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	} {
		t.Run(tt.before, func(t *testing.T) {
			node, err := newNode(strings.NewReader(tt.before), nil)

			if err != nil {
				t.Fatalf("newNode(%q): %v", tt.before, err)
			}

			node.explode(4)

			if node.String() != tt.after {
				t.Errorf("explode(%q) == %q, expected %q", tt.before, node.String(), tt.after)
			}

			numbers := strings.ReplaceAll(tt.after, "[", "")
			numbers = strings.ReplaceAll(numbers, "]", "")
			numbers = strings.ReplaceAll(numbers, ",", " ")
			numbers = "[" + numbers + "]"

			if got := fmt.Sprint(node.linkedNumbers()); got != fmt.Sprint(numbers) {
				t.Errorf("linked Numbers are %q, expected %q", got, numbers)
			}

			if got := fmt.Sprint(node.linkedNumbersBackwards()); got != fmt.Sprint(numbers) {
				t.Errorf("linked Numbers Bakwards are %q, expected %q", got, numbers)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	// Build the nodes from the example. The parser does not support two
	// character numbers. So they have to be changed later.
	node1, _ := newNode(strings.NewReader("[[[[0,7],4],[0,[0,0]]],[1,1]]"), nil)
	node1.linkedHead().next.next.next.value = 15
	node1.linkedHead().next.next.next.next.next.value = 13
	node2, _ := newNode(strings.NewReader("[[[[0,7],4],[[7,8],[0,0]]],[1,1]]"), nil)
	node2.linkedHead().next.next.next.next.next.next.value = 13

	for _, tt := range []struct {
		before *node
		after  string
	}{
		{node1, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
		{node2, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
	} {
		t.Run(tt.before.String(), func(t *testing.T) {
			before := tt.before.String()
			tt.before.split()

			if got := tt.before.String(); got != tt.after {
				t.Errorf("split(%q) == %q, expected %q", before, got, tt.after)
			}

			numbers := strings.ReplaceAll(tt.after, "[", "")
			numbers = strings.ReplaceAll(numbers, "]", "")
			numbers = strings.ReplaceAll(numbers, ",", " ")
			numbers = "[" + numbers + "]"

			if got := fmt.Sprint(tt.before.linkedNumbers()); got != fmt.Sprint(numbers) {
				t.Errorf("linked Numbers are %q, expected %q", got, numbers)
			}

			if got := fmt.Sprint(tt.before.linkedNumbersBackwards()); got != fmt.Sprint(numbers) {
				t.Errorf("linked Numbers Bakwards are %q, expected %q", got, numbers)
			}
		})
	}
}
