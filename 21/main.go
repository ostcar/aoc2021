package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Println(task1(puzzleInput))
	fmt.Println(task2(puzzleInput))
}

func task1(input string) string {
	//player := [2]*player{newPlayer(4), newPlayer(8)}

	g := game{
		player:   [2]*player{newPlayer(8), newPlayer(1)},
		dice:     dice{sides: 100},
		wonScore: 1000,
	}

	var turn int

	for g.next() {

	}

	return fmt.Sprint(g.player[(turn+1)%2].points * g.dice.rolled)
}

func task2(input string) string {
	// //player := [2]*player{newPlayer(8), newPlayer(1)}
	// player := [2]*player{newPlayer(4), newPlayer(8)}
	// var dice dice
	// var turn int

	// for {
	// 	player[turn].move(dice.roll(3))
	// 	if player[turn].points >= 1000 {
	// 		break
	// 	}
	// 	turn = (turn + 1) % 2
	// }

	// return fmt.Sprint(player[(turn+1)%2].points * dice.rolled)
	return ""
}

type game struct {
	dice     dice
	player   [2]*player
	turn     int
	wonScore int
}

// next plays the next turn. Returns false when the game is finished.
func (g *game) next() bool {
	g.player[g.turn].move(g.dice.roll(3))
	if g.player[g.turn].points >= g.wonScore {
		return false
	}
	g.turn = (g.turn + 1) % 2
	return true
}

// quantoDice is a map from all posible outcomes to the count how often it
// happens if the dice is rolled three times.
var quantomDice = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type dice struct {
	position int
	rolled   int
	sides    int
}

func (d *dice) roll(n int) int {
	d.rolled += n

	var out int
	for i := 0; i < n; i++ {
		d.position = ((d.position) % d.sides) + 1
		out += d.position
	}
	return out
}

type player struct {
	place  int
	points int
}

func newPlayer(start int) *player {
	return &player{
		place: start,
	}
}

func (p *player) move(n int) {
	p.place = ((p.place - 1 + n) % 10) + 1
	p.points += p.place
}
