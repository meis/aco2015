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
		{"Given Input", args{input.GetInputStrings()}, 1333},
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
		{"Given Input", args{input.GetInputStrings()}, 2046},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.strings); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_countCharactersOfCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{`"" is 2 characters of code`, args{`""`}, 2},
		{`"abc" is 5 characters of code`, args{`"abc"`}, 5},
		{`"aaa\"aaa" is 10 characters of code`, args{`"aaa\"aaa"`}, 10},
		{`"\x27" is 56haracters of code`, args{`"\x27"`}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharactersOfCode(tt.args.s); got != tt.want {
				t.Errorf("countCharactersOfCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCharactersOfMemory(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{`"" is 0 characters of code`, args{`""`}, 0},
		{`"abc" is 3 characters of code`, args{`"abc"`}, 3},
		{`"aaa\"aaa" is 7 characters of code`, args{`"aaa\"aaa"`}, 7},
		{`"\x27" is 1 characters of code`, args{`"\x27"`}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharactersOfMemory(tt.args.s); got != tt.want {
				t.Errorf("countCharactersOfMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{`"" encodes to "\"\""`, args{`""`}, `"\"\""`},
		{`"abc" encodes to "\"abc\""`, args{`"abc"`}, `"\"abc\""`},
		{`"aaa\"aaa" encodes to "\"aaa\\\"aaa\""`, args{`"aaa\"aaa"`}, `"\"aaa\\\"aaa\""`},
		{`"\x27" encodes to "\"\\x27\""`, args{`"\x27"`}, `"\"\\x27\""`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeString(tt.args.s); got != tt.want {
				t.Errorf("encodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
