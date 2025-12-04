package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestReadFile(t *testing.T) {
	contents := bytes.NewBufferString(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

	lines := readFile(contents)
	if len(lines) != 10 {
		t.Errorf("got %d lines want 10", len(lines))
	}
}

func TestOccupiedNeighbors(t *testing.T) {
	contents := bytes.NewBufferString(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)
	lines := readFile(contents)
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "middle",
			x:    5,
			y:    5,
			want: 6,
		},
		{
			name: "top edge",
			x:    5,
			y:    0,
			want: 3,
		},
		{
			name: "bottom edge",
			x:    5,
			y:    9,
			want: 5,
		},
		{
			name: "left edge",
			x:    0,
			y:    5,
			want: 4,
		},
		{
			name: "right edge",
			x:    9,
			y:    5,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := occupiedNeighbors(lines, tt.x, tt.y)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestRemoveRoll(t *testing.T) {
	contents := bytes.NewBufferString(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)
	lines := readFile(contents)
	want := bytes.NewBufferString(`...@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)
	wantLines := readFile(want)

	lines = removeRoll(lines, 2, 0)

	if !slices.Equal(lines, wantLines) {
		t.Errorf("got %v want %v", lines, wantLines)
	}
}
