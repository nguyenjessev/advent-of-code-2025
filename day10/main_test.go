package main

import (
	"os"
	"slices"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := readFile(file)
	want := 3

	if len(lines) != want {
		t.Errorf("got %d want %d", len(lines), want)
	}
}

func TestMachineComponents(t *testing.T) {
	contents := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"

	gotLights, gotButtons, gotJoltages := machineComponents(contents)
	wantLights := lights{false, true, true, false}
	wantButtons := []button{
		{3},
		{1, 3},
		{2},
		{2, 3},
		{0, 2},
		{0, 1},
	}
	wantJoltages := joltages{3, 5, 4, 7}

	assertLights(t, gotLights, wantLights)
	assertButtons(t, gotButtons, wantButtons)
	assertJoltages(t, gotJoltages, wantJoltages)
}

func TestSolveLights(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "one button",
			input: "[.##.] (1,2) {3,5,4,7}",
			want:  1,
		},
		{
			name:  "multiple buttons",
			input: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, b, _ := machineComponents(tt.input)
			got := solveLights(l, b)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestApplyLights(t *testing.T) {
	l := lights{false, false, false, false}
	b := button{0, 1, 2}
	want := lights{true, true, true, false}

	got, _ := b.applyLights(l, 0)

	assertLights(t, got, want)
}

func TestApplyJoltages(t *testing.T) {
	j := joltages{0, 0, 0}
	b := button{0, 2}
	want := joltages{1, 0, 1}

	got, _ := b.applyJoltages(j, 0)

	assertJoltages(t, got, want)
}

func assertButtons(t testing.TB, got, want []button) {
	t.Helper()

	if !slices.EqualFunc(got, want, func(a, b button) bool {
		return slices.Equal(a, b)
	}) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertJoltages(t testing.TB, got, want joltages) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertLights(t testing.TB, got, want lights) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
