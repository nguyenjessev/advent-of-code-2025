package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGetRangesFromFile(t *testing.T) {
	contents := bytes.NewBufferString("20-39,40-67,2-7")

	ranges := getRangesFromFile(contents)

	if len(ranges) != 3 {
		t.Errorf("got %d ranges, want 3", len(ranges))
	}

	assertRangeEqual(t, ranges[0], Range{20, 39})
	assertRangeEqual(t, ranges[1], Range{40, 67})
	assertRangeEqual(t, ranges[2], Range{2, 7})
}

func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		{
			name: "valid id",
			id:   12345,
			want: false,
		},
		{
			name: "invalid id",
			id:   22,
			want: true,
		},
		{
			name: "large valid id",
			id:   222367848123812222,
			want: false,
		},
		{
			name: "large invalid id",
			id:   123456789123456789,
			want: true,
		},
		{
			name: "repeated more than twice",
			id:   123123123,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInvalidID(tt.id)

			if got != tt.want {
				t.Errorf("%d: got %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestAllSliceElementsEqual(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want bool
	}{
		{
			name: "all equal",
			s:    []string{"a", "a", "a"},
			want: true,
		},
		{
			name: "not all equal",
			s:    []string{"a", "a", "b"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allSliceElementsEqual(tt.s)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func assertRangeEqual(t *testing.T, got, want Range) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
