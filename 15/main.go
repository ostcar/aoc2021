package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	nodes, width, height := parse(input)

	return fmt.Sprint(dijkstra(nodes, width, height))
}

func task2(input string) string {
	nodes, width, height := parse(input)

	nodes, width, height = nodesTransform(nodes, width, height)

	return fmt.Sprint(dijkstra(nodes, width, height))
}

func parse(input string) (nodes []int, width, height int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	width = len(lines[0])
	height = len(lines)
	nodes = make([]int, width*height)
	for i, line := range lines {
		for j := range line {
			nodes[i*width+j] = int(line[j] - '0')
		}
	}
	return nodes, width, height
}

func dijkstra(nodes []int, width, height int) int {
	// unvisitedRiskLVL is a list of node-indexes where the risk level was
	// caludlated (at least once) but where not visited.
	//
	// A slice with a capacety for all nodes is created so no copy is needed.
	unvisitedRiskLVL := make([]int, 0, len(nodes))

	// riskLVL is the risk level for each node. The slice is initialized with 0.
	// Normaly dijksta would initialize it with math.MaxInt but since there is
	// no unvisited node with a risk of 0, 0 can also be used as default.
	riskLVL := make([]int, len(nodes))

	// Set the first risk level to 0 and append it as only value to the list of
	// unvisited risk levels.
	riskLVL[0] = 0
	unvisitedRiskLVL = append(unvisitedRiskLVL, 0)

	// Go though all unvisited nodes that have a risk level. Always take the
	// value with the lowest level (this has the lowest risk).
	for current := unvisitedRiskLVL[0]; current < width*height-1; current = unvisitedRiskLVL[0] {
		// Remove the current element from the list of unvisited nodes. We visit
		// it now.
		unvisitedRiskLVL = unvisitedRiskLVL[1:]

		// Calculate the risk level of the neibors of the current node.
		for _, neibor := range neibors(width, height, current) {
			newRisk := riskLVL[current] + nodes[neibor]
			oldRisk := riskLVL[neibor]

			if oldRisk == 0 {
				// Calculating the rist level for this node for the first time.
				// Add it to the list of unvisited Risk level.
				unvisitedRiskLVL = append(unvisitedRiskLVL, neibor)
				riskLVL[neibor] = newRisk
				continue
			}

			// The rist level was calculated before. Use the smaller value.
			riskLVL[neibor] = min(oldRisk, newRisk)
		}

		// Reorder the list of unvisited nodes so the node with the lowest
		// risk level is first.
		sort.Slice(unvisitedRiskLVL, func(i, j int) bool {
			return riskLVL[unvisitedRiskLVL[i]] < riskLVL[unvisitedRiskLVL[j]]
		})
	}
	return riskLVL[width*height-1]
}

func nodesTransform(smallNodes []int, smallWidth, smallHeight int) (nodes []int, bigWidth, bigHeight int) {
	bigWidth = smallWidth * 5
	bigHeight = smallHeight * 5
	bigMap := make([]int, bigWidth*bigHeight)

	for i := 0; i < len(bigMap); i++ {
		bigX := i % bigWidth
		bigY := i / bigWidth

		smallX := bigX % smallWidth
		smallY := bigY % smallHeight

		fieldX := bigX / smallWidth
		fieldY := bigY / smallHeight

		smallIndex := smallX + smallY*smallWidth
		bigMap[i] = ((smallNodes[smallIndex] + fieldX + fieldY - 1) % 9) + 1
	}

	return bigMap, bigWidth, bigHeight
}

func neibors(width, height, index int) []int {
	var neighbor []int
	if top := index - width; top >= 0 {
		neighbor = append(neighbor, top)
	}
	if index%height > 0 {
		neighbor = append(neighbor, index-1)
	}
	if index%height < height-1 {
		neighbor = append(neighbor, index+1)
	}
	if bottom := index + width; bottom < height*width {
		neighbor = append(neighbor, bottom)
	}
	return neighbor
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printNodes(nodes []int, width, height int) string {
	buf := new(bytes.Buffer)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Fprint(buf, nodes[i*height+j])
		}
		fmt.Fprintln(buf)
	}
	fmt.Fprintln(buf)
	return buf.String()
}
