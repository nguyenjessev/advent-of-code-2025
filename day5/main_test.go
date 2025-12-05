package main

import (
	"bytes"
	"testing"
)

func TestReadFile(t *testing.T) {
	contents := bytes.NewBufferString(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	freshRanges, availableIngredients := readFile(contents)

	if len(freshRanges) != 4 {
		t.Errorf("got %d want %d", len(freshRanges), 4)
	}

	if len(availableIngredients) != 6 {
		t.Errorf("got %d want %d", len(availableIngredients), 6)
	}
}

func TestParseStringToRange(t *testing.T) {
	tests := []struct {
		input string
		want  ingredientRange
	}{
		{
			input: "1-3",
			want:  ingredientRange{start: 1, end: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseStringToRange(tt.input)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestCountFreshIngredients(t *testing.T) {
	freshIngredientRanges := []ingredientRange{
		{start: 3, end: 5},
		{start: 10, end: 14},
		{start: 16, end: 20},
		{start: 12, end: 18},
	}
	availableIngredients := []int{1, 5, 8, 11, 17, 32}

	got := countFreshIngredients(freshIngredientRanges, availableIngredients)

	if got != 3 {
		t.Errorf("got %d want %d", got, 3)
	}
}

func TestCountAllValidIngredients(t *testing.T) {
	freshIngredientRanges := []ingredientRange{
		{start: 3, end: 5},
		{start: 10, end: 14},
		{start: 16, end: 20},
		{start: 12, end: 18},
	}

	got := countAllValidIngredients(freshIngredientRanges)

	if got != 14 {
		t.Errorf("got %d want %d", got, 14)
	}

	t.Run("overlapping ranges", func(t *testing.T) {
		freshIngredientRanges = []ingredientRange{
			{start: 1, end: 5},
			{start: 2, end: 7},
			{start: 0, end: 3},
		}

		got := countAllValidIngredients(freshIngredientRanges)
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("contained ranges", func(t *testing.T) {
		freshIngredientRanges = []ingredientRange{
			{start: 1, end: 5},
			{start: 2, end: 4},
		}

		got := countAllValidIngredients(freshIngredientRanges)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("whatever this is", func(t *testing.T) {
		freshIngredientRanges = []ingredientRange{
			{start: 1, end: 5},
			{start: 2, end: 5},
		}

		got := countAllValidIngredients(freshIngredientRanges)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("ranges that fully contain previous ranges", func(t *testing.T) {
		freshIngredientRanges = []ingredientRange{
			{start: 3, end: 5},
			{start: 2, end: 6},
		}

		got := countAllValidIngredients(freshIngredientRanges)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
