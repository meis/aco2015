package main

import (
	"strings"
	"testing"

	"github.com/meis/aoc2015/input"
)

func Test_part1(t *testing.T) {
	type args struct {
		directions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{strings.Split(">", "")}, 2},
		{"Example", args{strings.Split("^>v<", "")}, 4},
		{"Example", args{strings.Split("^v^v^v^v^v", "")}, 2},

		{"Given Input", args{input.GetInputChars()}, 2592},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.directions); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		directions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{strings.Split("^>", "")}, 3},
		{"Example", args{strings.Split("^>v<", "")}, 3},
		{"Example", args{strings.Split("^v^v^v^v^v", "")}, 11},

		{"Given Input", args{input.GetInputChars()}, 2360},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.directions); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
