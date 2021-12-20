package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	lines := strings.Split(input, "\n")

	var result *node
	for _, line := range lines {
		if line == "" {
			continue
		}

		n, err := newNode(strings.NewReader(line), nil)
		if err != nil {
			return fmt.Sprintln(err)
		}

		if result == nil {
			result = n
			continue
		}
		result = add(result, n)
		reduce(result)
	}

	return fmt.Sprint(result.magnitude())
}

func task2(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	var max int
	for a := 0; a < len(lines); a++ {
		for b := 0; b < len(lines); b++ {
			if a == b {
				continue
			}

			node1, _ := newNode(strings.NewReader(lines[a]), nil)
			node2, _ := newNode(strings.NewReader(lines[b]), nil)

			got := add(node1, node2)
			reduce(got)
			if m := got.magnitude(); m > max {
				max = m
			}
		}
	}

	return fmt.Sprint(max)
}

// Node is a tree. All leafes are also a linked list.
type node struct {
	left, right *node
	pre, next   *node
	value       int
}

func (n *node) isValue() bool {
	return n.left == nil
}

func (n *node) rightLeafNode() *node {
	if n.isValue() {
		return n
	}
	return n.right.rightLeafNode()
}

func (n *node) magnitude() int {
	if n.isValue() {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *node) String() string {
	if n.isValue() {
		return fmt.Sprint(n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func newNode(r io.Reader, pre *node) (*node, error) {
	c, err := nextChar(r)
	if err != nil {
		return nil, fmt.Errorf("reading first char: %w", err)
	}

	switch {
	case c >= '0' && c <= '9':
		n := &node{pre: pre, value: int(c - '0')}
		if pre != nil {
			pre.next = n
		}
		return n, nil

	case c == '[':
		left, err := newNode(r, pre)
		if err != nil {
			return nil, fmt.Errorf("left: %w", err)
		}

		if err := expectNext(r, ','); err != nil {
			return nil, err
		}

		right, err := newNode(r, left.rightLeafNode())
		if err != nil {
			return nil, fmt.Errorf("right: %w", err)
		}

		if err := expectNext(r, ']'); err != nil {
			return nil, err
		}

		return &node{left: left, right: right}, nil

	default:
		return nil, fmt.Errorf("invalid chat: %v", c)
	}
}

func nextChar(r io.Reader) (byte, error) {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		return 0, fmt.Errorf("reading next byte: %w", err)
	}
	return buf[0], nil
}

func expectNext(r io.Reader, expect byte) error {
	got, err := nextChar(r)
	if err != nil {
		return fmt.Errorf("reading next: %w", err)
	}
	if got != expect {
		return fmt.Errorf("got %q, expected %q", got, expect)
	}
	return nil

}

func add(a, b *node) *node {
	n := &node{left: a, right: b}
	a.linkedTail().next = b.linkedHead()
	b.linkedHead().pre = a.linkedTail()
	return n
}

func reduce(n *node) {
	splited := true
	for splited {
		exploded := true
		for exploded {
			_, exploded = n.explode(4)
		}
		splited = n.split()
	}
}

// actionNested checks if a node is nested four times and if so, explode it.
//
// Returns true if an explosion happen
func (n *node) explode(level int) (*node, bool) {
	if n.isValue() {
		return nil, false
	}

	if level > 0 {
		if newNode, ok := n.left.explode(level - 1); ok {
			if newNode != nil {
				n.left = newNode
			}
			return nil, true
		}

		if newNode, ok := n.right.explode(level - 1); ok {
			if newNode != nil {
				n.right = newNode
			}
			return nil, true
		}
		return nil, false
	}

	leftV := n.left.value
	rightV := n.right.value
	newNode := &node{value: 0}

	if n.left.pre != nil {
		n.left.pre.value += leftV
		newNode.pre = n.left.pre
		n.left.pre.next = newNode
	}

	if n.right.next != nil {
		n.right.next.value += rightV

		newNode.next = n.right.next
		n.right.next.pre = newNode
	}
	return newNode, true
}

func (n *node) split() bool {
	if !n.isValue() {
		if ok := n.left.split(); ok {
			return true
		}

		if ok := n.right.split(); ok {
			return true
		}
		return false
	}

	if n.value < 10 {
		return false
	}

	n.left = &node{value: n.value / 2}
	n.right = &node{value: n.value - n.left.value}
	n.left.next = n.right
	n.right.pre = n.left
	n.value = 0

	if n.pre != nil {
		n.pre.next = n.left
		n.left.pre = n.pre
		n.pre = nil
	}

	if n.next != nil {
		n.next.pre = n.right
		n.right.next = n.next
		n.next = nil
	}
	return true
}

func (n *node) linkedHead() *node {
	if n.isValue() {
		return n
	}
	return n.left.linkedHead()
}

func (n *node) linkedNumbers() []int {
	var numbers []int
	for head := n.linkedHead(); head != nil; head = head.next {
		numbers = append(numbers, head.value)
	}
	return numbers
}

func (n *node) linkedTail() *node {
	if n.isValue() {
		return n
	}
	return n.right.linkedTail()
}

func (n *node) linkedNumbersBackwards() []int {
	var numbers []int
	for tail := n.linkedTail(); tail != nil; tail = tail.pre {
		numbers = append(numbers, tail.value)
	}

	var reversed []int
	for i := len(numbers); i > 0; i-- {
		reversed = append(reversed, numbers[i-1])

	}
	return reversed
}
