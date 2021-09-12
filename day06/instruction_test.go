package main

import (
	"reflect"
	"testing"
)

func Test_parseInstruction(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{"turn on 583,543 through 846,710", args{"turn on 583,543 through 846,710"}, instruction{TURN_ON, 583, 543, 846, 710}},
		{"turn off 367,664 through 595,872", args{"turn off 367,664 through 595,872"}, instruction{TURN_OFF, 367, 664, 595, 872}},
		{"toggle 39,799 through 981,872", args{"toggle 39,799 through 981,872"}, instruction{TOGGLE, 39, 799, 981, 872}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInstruction(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
