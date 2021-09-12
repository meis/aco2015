package main

import (
	"testing"

	"github.com/meis/aoc2015/input"
)

func Test_part1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputString()}, 252594},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputString()}, 3579328},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_lookAndSay(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 becomes 11", args{"1"}, "11"},
		{"11 becomes 21", args{"11"}, "21"},
		{"21 becomes 1211", args{"21"}, "1211"},
		{"1211 becomes 111221", args{"1211"}, "111221"},
		{"111221 becomes 312211", args{"111221"}, "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookAndSay(tt.args.input); got != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lookAndSayNTimes(t *testing.T) {
	type args struct {
		input string
		n     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"iteration 1 is 11", args{"1", 1}, "11"},
		{"iteration 2 is 21", args{"1", 2}, "21"},
		{"iteration 3 is 1211", args{"1", 3}, "1211"},
		{"iteration 4 is 111221", args{"1", 4}, "111221"},
		{"iteration 5 is 312211", args{"1", 5}, "312211"},
		{"iteration 6 is 13112221", args{"1", 6}, "13112221"},
		{"iteration 7 is 1113213211", args{"1", 7}, "1113213211"},
		{"iteration 8 is 21", args{"1", 8}, "31131211131221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookAndSayNTimes(tt.args.input, tt.args.n); got != tt.want {
				t.Errorf("lookAndSayNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
