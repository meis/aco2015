package main

import (
	"fmt"

	"github.com/meis/aoc2015/input"
)

func main() {
	strings := input.GetInputStrings()

	fmt.Printf("After following the instructions, there are %d lights lit\n", part1(strings))
	fmt.Printf("The total brightness of all lights combined is %d\n", part2(strings))
}

func part1(strings []string) int {
	lights := make(map[position]bool, 1_000_000)

	for _, s := range strings {
		in := parseInstruction(s)
		for _, p := range getPositionRange(in.startX, in.startY, in.endX, in.endY) {
			switch in.op {
			case TURN_ON:
				lights[p] = true
			case TURN_OFF:
				delete(lights, p)
			case TOGGLE:
				_, exists := lights[p]
				if exists {
					delete(lights, p)
				} else {
					lights[p] = true
				}
			default:
				panic("should not happen")
			}
		}
	}

	return len(lights)
}

func part2(strings []string) int {
	lights := make(map[position]int, 1_000_000)
	totalBrightness := 0

	for _, s := range strings {
		in := parseInstruction(s)
		for _, p := range getPositionRange(in.startX, in.startY, in.endX, in.endY) {
			currentValue := lights[p]
			var nextValue int
			switch in.op {
			case TURN_ON:
				nextValue = currentValue + 1
			case TURN_OFF:
				nextValue = currentValue - 1
				if nextValue < 0 {
					nextValue = 0
				}
			case TOGGLE:
				nextValue = currentValue + 2
			default:
				panic("should not happen")
			}
			totalBrightness -= currentValue
			totalBrightness += nextValue
			lights[p] = nextValue
		}
	}

	return totalBrightness
}
