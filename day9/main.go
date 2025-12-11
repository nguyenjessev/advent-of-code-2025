package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := readFile(file)
	points := createPointsFromLines(lines)
	rectangles := createRectanglesFromPoints(points)

	slices.SortFunc(rectangles, func(a, b rectangle) int {
		return int(b.area() - a.area())
	})

	fmt.Println(rectangles[0].area())

	for _, r := range rectangles {
		valid := true

		for i, p1 := range points {
			j := i + 1
			if j >= len(points) {
				j = 0
			}
			p2 := points[j]

			if r.intersectedBy(p1, p2) {
				valid = false

				break
			}
		}

		if valid {
			fmt.Println(r.area())

			break
		}
	}
}

type point struct {
	x float64
	y float64
}

type rectangle struct {
	p1 point
	p2 point
}

func newPoint(x, y float64) point {
	return point{
		x: x,
		y: y,
	}
}

func newRectangle(p1, p2 point) rectangle {
	return rectangle{
		p1: p1,
		p2: p2,
	}
}

func (r *rectangle) area() float64 {
	width := math.Abs(r.p1.x-r.p2.x) + 1
	height := math.Abs(r.p1.y-r.p2.y) + 1

	return width * height
}

func (r *rectangle) intersectedBy(p1, p2 point) bool {
	if p1.x >= max(r.p1.x, r.p2.x) && p2.x >= max(r.p1.x, r.p2.x) {
		return false
	}

	if p1.x <= min(r.p1.x, r.p2.x) && p2.x <= min(r.p1.x, r.p2.x) {
		return false
	}

	if p1.y >= max(r.p1.y, r.p2.y) && p2.y >= max(r.p1.y, r.p2.y) {
		return false
	}

	if p1.y <= min(r.p1.y, r.p2.y) && p2.y <= min(r.p1.y, r.p2.y) {
		return false
	}

	return true
}

func createPointsFromLines(lines []string) []point {
	var points []point
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			panic(err)
		}

		point := newPoint(x, y)

		points = append(points, point)
	}

	return points
}

func createRectanglesFromPoints(points []point) []rectangle {
	var rectangles []rectangle
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			r := newRectangle(p1, p2)
			rectangles = append(rectangles, r)
		}
	}

	return rectangles
}

func readFile(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
