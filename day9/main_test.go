package main

import "testing"

func TestPoint(t *testing.T) {
	t.Run("create new point", func(t *testing.T) {
		got := newPoint(1, 2)
		want := point{x: 1, y: 2}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestRectangle(t *testing.T) {
	t.Run("create new rectangle", func(t *testing.T) {
		p1 := newPoint(1, 2)
		p2 := newPoint(3, 4)
		got := newRectangle(p1, p2)
		want := rectangle{p1: p1, p2: p2}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get area", func(t *testing.T) {
		p1 := newPoint(1, 2)
		p2 := newPoint(3, 4)
		r := newRectangle(p1, p2)
		got := r.area()
		want := 9.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("intersections", func(t *testing.T) {
		p1 := newPoint(10, 10)
		p2 := newPoint(20, 20)
		r := newRectangle(p1, p2)
		tests := []struct {
			name string
			p1   point
			p2   point
			want bool
		}{
			{
				name: "no intersection",
				p1:   newPoint(0, 0),
				p2:   newPoint(0, 1),
				want: false,
			},
			{
				name: "vertical intersection with one end inside",
				p1:   newPoint(15, 0),
				p2:   newPoint(15, 15),
				want: true,
			},
			{
				name: "horizontal intersection with one end inside",
				p1:   newPoint(0, 15),
				p2:   newPoint(15, 15),
				want: true,
			},
			{
				name: "vertical intersection all the way through",
				p1:   newPoint(15, 0),
				p2:   newPoint(15, 25),
				want: true,
			},
			{
				name: "horizontal intersection all the way through",
				p1:   newPoint(0, 15),
				p2:   newPoint(25, 15),
				want: true,
			},
			{
				name: "partial shared edge",
				p1:   newPoint(10, 15),
				p2:   newPoint(10, 20),
				want: false,
			},
			{
				name: "complete shared edge",
				p1:   newPoint(10, 10),
				p2:   newPoint(10, 20),
				want: false,
			},
			{
				name: "extended shared edge",
				p1:   newPoint(10, 5),
				p2:   newPoint(10, 25),
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := r.intersectedBy(tt.p1, tt.p2)

				if got != tt.want {
					t.Errorf("got %v want %v", got, tt.want)
				}
			})
		}
	})
}
