package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/meis/aoc2015/input"
)

func main() {
	secretKey := input.GetInputString()

	fmt.Printf("Answer using 5 zeroes is: %d \n", part1(secretKey))
	fmt.Printf("Answer using 6 zeroes is: %d \n", part2(secretKey))
}

func part1(secretKey string) int {
	return hashProducer(secretKey, 5)
}

func part2(secretKey string) int {
	return hashProducer(secretKey, 6)
}

func dataFrom(input io.Reader) string {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic("could not read input")
	}
	return string(bytes)
}

func hashProducer(secretKey string, zeroes int) int {
	i := 1
	prefix := strings.Repeat("0", zeroes)

	for !hashHasPrefix(secretKey, i, prefix) {
		i++
	}

	return i
}

func getHash(secretKey string, iteration int) string {
	data := []byte(secretKey + strconv.Itoa(iteration))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func hashHasPrefix(secretKey string, iteration int, prefix string) bool {
	hash := getHash(secretKey, iteration)
	return strings.HasPrefix(hash, prefix)
}
