package main

import (
	"fmt"
	"strings"

	"github.com/meis/aoc2015/input"
)

func main() {
	input := input.GetInputString()

	fmt.Printf("After 40 iterations, the length of the result is: %d\n", part1(input))
	fmt.Printf("After 50 iterations, the length of the result is: %d\n", part2(input))
}

func part1(input string) int {
	return len(lookAndSayNTimes(input, 40))
}

func part2(input string) int {
	return len(lookAndSayNTimes(input, 50))
}

func lookAndSayNTimes(input string, n int) string {
	output := input
	for i := 0; i < n; i++ {
		output = lookAndSay(output)
	}
	return output
}

func lookAndSay(input string) string {
	currChar := input[0]
	currCount := 1

	var acc strings.Builder

	for i := 1; i < len(input); i++ {
		c := input[i]
		if c == currChar {
			currCount++
		} else {
			fmt.Fprintf(&acc, "%d%s", currCount, string(currChar))
			currChar = c
			currCount = 1
		}
	}

	fmt.Fprintf(&acc, "%d%s", currCount, string(currChar))

	return acc.String()
}

// For Historical reasons, here is the original implementation, which is
// super slow because it concatenates strings:
//
// func lookAndSay(input string) string {
// 	currChar := input[0]
// 	currCount := 1

// 	acc := ""

// 	for i := 1; i < len(input); i++ {
// 		c := input[i]
// 		if c == currChar {
// 			currCount++
// 		} else {
// 			acc += strconv.Itoa(currCount) + string(currChar)
// 			currChar = c
// 			currCount = 1
// 		}
// 	}

// 	acc += strconv.Itoa(currCount) + string(currChar)

// 	return acc
// }
