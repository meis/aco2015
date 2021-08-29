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
		{"Example", args{strings.Split("(())", "")}, 0},
		{"Example", args{strings.Split("()()", "")}, 0},
		{"Example", args{strings.Split("(((", "")}, 3},
		{"Example", args{strings.Split("(()(()(", "")}, 3},
		{"Example", args{strings.Split("))(((((", "")}, 3},
		{"Example", args{strings.Split("())", "")}, -1},
		{"Example", args{strings.Split("))(", "")}, -1},
		{"Example", args{strings.Split(")))", "")}, -3},
		{"Example", args{strings.Split(")())())", "")}, -3},

		{"Given Input", args{input.GetInputChars()}, 138},
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
		{"Example", args{strings.Split(")", "")}, 1},
		{"Example", args{strings.Split("()())", "")}, 5},
		{"Example", args{strings.Split("", "")}, 0},

		{"Given Input", args{input.GetInputChars()}, 1771},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.directions); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
