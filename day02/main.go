package main

import (
	"fmt"

	"github.com/meis/aoc2015/input"
)

func main() {
	lines := input.GetInputStrings()

	fmt.Printf("The elves should order %d sq feet if wrapping paper\n", part1(lines))
	fmt.Printf("The elves should order %d feet of ribbon\n", part2(lines))
}

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		total += NewBox(line).WrappingArea()
	}

	return total
}

func part2(lines []string) int {
	total := 0

	for _, line := range lines {
		total += NewBox(line).RibbonLenght()
	}

	return total
}
