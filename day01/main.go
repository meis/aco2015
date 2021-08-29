package main

import (
	"fmt"

	"github.com/meis/aoc2015/input"
)

func main() {
	instructions := input.GetInputChars()

	fmt.Printf("The instructions take Santa to floor: %d\n", part1(instructions))
	fmt.Printf("The position of the character that causes Santa to first enter the basement is: %d\n", part2(instructions))
}

func part1(directions []string) int {
	i := 0

	for _, c := range directions {
		switch c {
		case "(":
			i++
		case ")":
			i--
		}
	}
	return i
}

func part2(directions []string) int {
	i := 0

	for pos, c := range directions {
		switch c {
		case "(":
			i++
		case ")":
			i--
		}

		if i == -1 {
			return pos + 1
		}
	}

	return 0
}
