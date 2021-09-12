package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/meis/aoc2015/input"
)

func main() {
	strings := input.GetInputStrings()

	fmt.Printf("Number of characters of code minus number of characters in memory: %d\n", part1(strings))
	fmt.Printf("Number of characters for newly encoded strings minuts the originals: %d\n", part2(strings))
}

func part1(strings []string) int {
	var codeChars, memoryChars int

	for _, s := range strings {
		codeChars += countCharactersOfCode(s)
		memoryChars += countCharactersOfMemory(s)
	}

	return codeChars - memoryChars
}

func part2(strings []string) int {
	var originalChars, encodedChars int

	for _, s := range strings {
		originalChars += countCharactersOfCode(s)
		encodedChars += countCharactersOfCode(encodeString(s))
	}

	return encodedChars - originalChars
}
func countCharactersOfCode(s string) int {
	return utf8.RuneCountInString(s)
}

func countCharactersOfMemory(s string) int {
	unquoted, _ := strconv.Unquote(s)
	return len(unquoted)
}

func encodeString(s string) string {
	return fmt.Sprintf("%q", s)
}
