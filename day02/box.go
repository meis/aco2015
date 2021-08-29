package main

import (
	"math"
	"strconv"
	"strings"
)

type Box struct {
	Length int
	Width  int
	Height int
}

func NewBox(s string) Box {
	dimensions := strings.Split(s, "x")

	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic("not a number")
	}

	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic("not a number")
	}

	height, err := strconv.Atoi(dimensions[2])
	if err != nil {
		panic("not a number")
	}

	return Box{
		Length: length,
		Width:  width,
		Height: height,
	}
}

func (b Box) WrappingArea() int {
	smaller := math.MaxInt64
	total := 0

	sides := []int{
		b.Length * b.Width,
		b.Width * b.Height,
		b.Height * b.Length,
	}

	for _, side := range sides {
		if side < smaller {
			smaller = side
		}

		total += 2 * side
	}

	return total + smaller
}

func (b Box) RibbonLenght() int {
	wrapping := 0
	bow := b.Length * b.Width * b.Height

	biggerSide := 0
	for _, side := range []int{b.Length, b.Width, b.Height} {
		if side > biggerSide {
			biggerSide = side
		}

		wrapping += side * 2
	}

	wrapping -= biggerSide * 2

	return wrapping + bow
}
