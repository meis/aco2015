package main

import (
	"reflect"
	"testing"
)

func Test_getPositionRange(t *testing.T) {
	type args struct {
		startX int
		startY int
		endX   int
		endY   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0,0 through 999,999", args{0, 0, 999, 999}, 1_000_000},
		{"0,0 through 999,0", args{0, 0, 999, 0}, 1_000},
		{"499,499 through 500,500", args{499, 499, 500, 500}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPositionRange(tt.args.startX, tt.args.startY, tt.args.endX, tt.args.endY); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("getPositionRange() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
