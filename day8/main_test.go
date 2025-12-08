package main

import (
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestPoint(t *testing.T) {
	t.Run("create new point", func(t *testing.T) {
		got := newPoint(1, 2, 3)
		want := &point{
			x: 1,
			y: 2,
			z: 3,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get distance between points", func(t *testing.T) {
		p1 := newPoint(1, 2, 3)
		p2 := newPoint(4, 5, 6)
		got := p1.distanceTo(p2)
		want := math.Sqrt(27)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("find nearest neighbor", func(t *testing.T) {
		p := newPoint(1, 2, 3)
		want := newPoint(4, 5, 6)
		neighbors := []*point{
			newPoint(7, 8, 9),
			want,
			newPoint(10, 11, 12),
		}

		got := p.nearestNeighbor(neighbors)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestPair(t *testing.T) {
	t.Run("create pair", func(t *testing.T) {
		p1 := newPoint(1, 2, 3)
		p2 := newPoint(4, 5, 6)
		want := pair{
			p1:       p1,
			p2:       p2,
			distance: math.Sqrt(27),
		}

		got := newPair(p1, p2)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("find all pairs", func(t *testing.T) {
		points := []*point{
			newPoint(1, 2, 3),
			newPoint(4, 5, 6),
			newPoint(7, 8, 9),
		}
		want := []pair{
			newPair(points[0], points[1]),
			newPair(points[0], points[2]),
			newPair(points[1], points[2]),
		}

		got := findAllPairs(points)

		if len(got) != len(want) {
			t.Errorf("got length %v want %v", len(got), len(want))
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
