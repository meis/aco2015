package main

import (
	"fmt"

	"github.com/meis/aoc2015/input"
)

func main() {
	directions := input.GetInputChars()

	fmt.Printf("Santa has visited %d houses\n", part1(directions))
	fmt.Printf("Santa and Robo-Santa had visited %d houses\n", part2(directions))
}

func part1(directions []string) int {
	santa := santa{}
	visited := make(map[position]bool)

	visited[santa.Position] = true

	for _, direction := range directions {
		visited[santa.move(direction)] = true
	}

	return len(visited)
}

func part2(directions []string) int {
	// Santa and Robo-Santa
	santas := [2]santa{}
	visited := make(map[position]bool)

	visited[santas[0].Position] = true

	for i, direction := range directions {
		s := &santas[i%2]
		visited[s.move(direction)] = true
	}

	return len(visited)
}
