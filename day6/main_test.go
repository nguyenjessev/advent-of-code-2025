package main

import (
	"bytes"
	"reflect"
	"slices"
	"testing"
)

func TestReadFile(t *testing.T) {
	contents := bytes.NewBufferString(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `)

	lines := readFile(contents)

	if len(lines) != 4 {
		t.Errorf("got %d lines, want 4", len(lines))
	}
}

func TestDataMatrix(t *testing.T) {
	contents := bytes.NewBufferString(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `)
	lines := readFile(contents)
	got := parseDataMatrix(lines)
	want := [][]string{
		{"123", "328", " 51", "64 "},
		{" 45", "64 ", "387", "23 "},
		{"  6", "98 ", "215", "314"},
		{"*", "+", "*", "+"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPerformMath(t *testing.T) {
	tests := []struct {
		name  string
		input [][]string
		want  int
	}{
		{
			name: "addition",
			input: [][]string{
				{"1"},
				{"2"},
				{"+"},
			},
			want: 3,
		},
		{
			name: "subtraction",
			input: [][]string{
				{"2"},
				{"1"},
				{"-"},
			},
			want: 1,
		},
		{
			name: "multiplication",
			input: [][]string{
				{"2"},
				{"3"},
				{"*"},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := performMath(tt.input, 0)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestConvertNumsToCephalopod(t *testing.T) {
	nums := []string{"   12", "23232", "31   ", " 123 "}
	want := []int{22, 133, 22, 311, 23}

	got := convertNumsToCephalopod(nums)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
