package main

import (
	"reflect"
	"testing"
)

func TestNewBox(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Box
	}{
		{"Example", args{"29x13x26"}, Box{29, 13, 26}},
		{"Example", args{"11x11x14"}, Box{11, 11, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBox(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_WrappingArea(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Example", fields{2, 3, 4}, 58},
		{"Example", fields{1, 1, 10}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := b.WrappingArea(); got != tt.want {
				t.Errorf("Box.WrappingArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_RibbonLenght(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Example", fields{2, 3, 4}, 34},
		{"Example", fields{1, 1, 10}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := b.RibbonLenght(); got != tt.want {
				t.Errorf("Box.RibbonLenght() = %v, want %v", got, tt.want)
			}
		})
	}
}
