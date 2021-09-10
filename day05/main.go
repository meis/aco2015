package main

import (
	"fmt"
	"os/exec"

	"github.com/meis/aoc2015/input"
)

func main() {
	strings := input.GetInputStrings()

	fmt.Printf("There are %d nice strings\n", part1(strings))
	fmt.Printf("There are %d nice strings (new definition)\n", part2(strings))
}

func part1(strings []string) int {
	niceStrings := 0

	for _, s := range strings {
		if isNiceString(s) {
			niceStrings++
		}
	}
	return niceStrings
}

func part2(strings []string) int {
	niceStrings := 0

	for _, s := range strings {
		if isNiceStringNow(s) {
			niceStrings++
		}
	}
	return niceStrings
}

func isNiceString(s string) bool {
	atLeastThreeVowels := false
	oneLetterTwiceInARow := false
	containsForbiddenStrings := false

	vowels := 0
	prevChar := '\n'
	for _, currentChar := range s {
		if prevChar != '\n' {
			if prevChar == currentChar {
				oneLetterTwiceInARow = true
			}
			switch string(prevChar) + string(currentChar) {
			case "ab", "cd", "pq", "xy":
				containsForbiddenStrings = true
			}
		}

		switch currentChar {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		prevChar = currentChar
	}

	atLeastThreeVowels = vowels >= 3

	return !containsForbiddenStrings && atLeastThreeVowels && oneLetterTwiceInARow
}

func isNiceStringNow(s string) bool {
	// That's what I want to do, but Go does not support backreferences in regexes:
	//
	// condition1, _ := regexp.MatchString("([[:alpha:]][[:alpha:]])[[:alpha:]]*\1", s)
	// condition2, _ := regexp.MatchString("([[:alpha:]])[[:alpha:]]\1", s)

	// return condition1 && condition2

	// So here we are:

	cmd := fmt.Sprintf("echo -n %s | grep -P '([[:alpha:]][[:alpha:]])[[:alpha:]]*\\1' | grep -P '([[:alpha:]])[[:alpha:]]\\1'", s)
	stdout, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	return string(stdout) == s+"\n"
}
