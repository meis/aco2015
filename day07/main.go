package main

import (
	"fmt"

	"github.com/meis/aoc2015/input"
)

func main() {
	strings := input.GetInputStrings()

	fmt.Printf("The signal from wire a is: %d\n", part1(strings))
	fmt.Printf("After rewiring, the signal from wire a is: %d\n", part2(strings))
}

func part1(strings []string) uint16 {
	board := NewBoard(strings)
	board.Run()

	return board.GetSignal("a")
}

func part2(strings []string) uint16 {
	aSignal := part1(strings)

	board := NewBoard(strings)
	board.SetConnection("b", fmt.Sprintf("%d", aSignal))
	board.Run()

	return board.GetSignal("a")
}
