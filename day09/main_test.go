package main

import (
	"testing"

	"github.com/meis/aoc2015/input"
)

func Test_part1(t *testing.T) {
	type args struct {
		distances []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputStrings()}, 207},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.distances); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		distances []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputStrings()}, 804},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.distances); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
