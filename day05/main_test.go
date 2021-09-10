package main

import (
	"testing"

	"github.com/meis/aoc2015/input"
)

func Test_part1(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputStrings()}, 236},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.strings); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{input.GetInputStrings()}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.strings); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNiceString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ugknbfddgicrmopn is nice", args{"ugknbfddgicrmopn"}, true},
		{"aaa              is nice", args{"aaa"}, true},
		{"jchzalrnumimnmhp is not nice", args{"jchzalrnumimnmhp"}, false},
		{"haegwjzuvuyypxyu is not nice", args{"haegwjzuvuyypxyu"}, false},
		{"dvszwmarrgswjxmb is not nice", args{"dvszwmarrgswjxmb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNiceString(tt.args.s); got != tt.want {
				t.Errorf("isNiceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNiceStringNow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"qjhvhtzxzqqjkmpb is nice", args{"qjhvhtzxzqqjkmpb"}, true},
		{"xxyxx            is nice", args{"xxyxx"}, true},
		{"uurcxstgmygtbstg is not nice", args{"uurcxstgmygtbstg"}, false},
		{"ieodomkazucvgmuy is not nice", args{"ieodomkazucvgmuy"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNiceStringNow(tt.args.s); got != tt.want {
				t.Errorf("isNiceStringNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
