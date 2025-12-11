package main

import (
	"os"
	"testing"
)

func TestFindPathToEnd(t *testing.T) {
	file, err := os.Open("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	lines := readFile(file)

	nodes := make(map[string]node)
	for _, line := range lines {
		n := newNode(line)
		nodes[n.name] = n
	}

	if len(nodes) != 10 {
		t.Errorf("got %d nodes want 10", len(nodes))
	}

	t.Run("path from end to end", func(t *testing.T) {
		got := countPaths("out", "out", nodes)

		if got != 1 {
			t.Errorf("got %d want 1", got)
		}
	})

	t.Run("multiple paths", func(t *testing.T) {
		got := countPaths("ccc", "out", nodes)

		if got != 3 {
			t.Errorf("got %d want 3", got)
		}
	})
}
