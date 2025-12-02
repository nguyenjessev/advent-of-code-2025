package main

import (
	"bytes"
	"testing"
)

func TestReadFile(t *testing.T) {
	contents := bytes.NewBufferString(`R1
R2
R3
L1
L2
L3`)

	lines := getLinesFromFile(contents)

	if len(lines) != 6 {
		t.Fatalf("got %d lines, want 6", len(lines))
	}
}

func TestMove(t *testing.T) {
	currentPosition := 50
	tests := []struct {
		name        string
		move        string
		minPosition int
		maxPosition int
		want        int
	}{
		{
			name:        "move right",
			move:        "R1",
			minPosition: 0,
			maxPosition: 99,
			want:        51,
		},
		{
			name:        "move left",
			move:        "L1",
			minPosition: 0,
			maxPosition: 99,
			want:        49,
		},
		{
			name:        "move right past limit",
			move:        "R50",
			minPosition: 0,
			maxPosition: 99,
			want:        0,
		},
		{
			name:        "move right far past limit",
			move:        "R60",
			minPosition: 0,
			maxPosition: 99,
			want:        10,
		},
		{
			name:        "move left past limit",
			move:        "L51",
			minPosition: 0,
			maxPosition: 99,
			want:        99,
		},
		{
			name:        "move left far past limit",
			move:        "L61",
			minPosition: 0,
			maxPosition: 99,
			want:        89,
		},
		{
			name:        "invalid move",
			move:        "X1",
			minPosition: 0,
			maxPosition: 99,
			want:        50,
		},
		{
			name:        "move that loops multiple times to the right",
			move:        "R201",
			minPosition: 0,
			maxPosition: 99,
			want:        51,
		},
		{
			name:        "move that loops multiple times to the left",
			move:        "L201",
			minPosition: 0,
			maxPosition: 99,
			want:        49,
		},
		{
			name:        "move right with different min and max positions",
			move:        "R10",
			minPosition: 46,
			maxPosition: 53,
			want:        52,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, distance := parseMove(tt.move)

			got := getPositionFromMove(currentPosition, direction, distance, tt.minPosition, tt.maxPosition)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNormalizeDistance(t *testing.T) {
	maxDistance := 99
	minDistance := 0
	tests := []struct {
		name     string
		distance int
		want     int
	}{
		{
			name:     "within range",
			distance: 50,
			want:     50,
		},
		{
			name:     "within range",
			distance: 60,
			want:     60,
		},
		{
			name:     "just outside range",
			distance: 100,
			want:     0,
		},
		{
			name:     "far outside range",
			distance: 110,
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeDistance(tt.distance, minDistance, maxDistance)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCountZeroCrossings(t *testing.T) {
	tests := []struct {
		name            string
		currentPosition int
		direction       string
		distance        int
		want            int
	}{
		{
			name:            "move right from close to max",
			currentPosition: 9,
			direction:       "R",
			distance:        2,
			want:            1,
		},
		{
			name:            "move left from close to min",
			currentPosition: 1,
			direction:       "L",
			distance:        2,
			want:            1,
		},
		{
			name:            "landing directly on zero",
			currentPosition: 9,
			direction:       "R",
			distance:        1,
			want:            0,
		},
		{
			name:            "multiple crossings from close to max",
			currentPosition: 9,
			direction:       "R",
			distance:        32,
			want:            4,
		},
		{
			name:            "multiple crossings from close to min",
			currentPosition: 1,
			direction:       "L",
			distance:        32,
			want:            4,
		},
		{
			name:            "multiple crossings from middle of range",
			currentPosition: 5,
			direction:       "R",
			distance:        25,
			want:            2,
		},
		{
			name:            "no crossings",
			currentPosition: 1,
			direction:       "R",
			distance:        5,
			want:            0,
		},
		{
			name:            "crossing right but then landing on zero",
			currentPosition: 9,
			direction:       "R",
			distance:        11,
			want:            1,
		},
		{
			name:            "crossing left but then landing on zero",
			currentPosition: 1,
			direction:       "L",
			distance:        11,
			want:            1,
		},
		{
			name:            "crossing multiple times but then landing on zero",
			currentPosition: 9,
			direction:       "R",
			distance:        21,
			want:            2,
		},
		{
			name:            "crossing left multiple times but then landing on zero",
			currentPosition: 2,
			direction:       "L",
			distance:        42,
			want:            4,
		},
		{
			name:            "starting at zero and ending at zero",
			currentPosition: 0,
			direction:       "L",
			distance:        10,
			want:            0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countZeroCrossings(tt.currentPosition, tt.direction, tt.distance, 0, 9)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
