package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := readFile(file)

	points := createPointsFromLines(lines)

	pairs := findAllPairs(points)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	var circuits []*circuit

	for _, pair := range pairs {
		var circuitToBeDeleted *circuit

		if pair.p1.circuit != nil && pair.p2.circuit != nil {
			if pair.p1.circuit != pair.p2.circuit {

				circuitToBeDeleted = pair.p2.circuit
			}
		}

		pair.p1.connectTo(pair.p2)

		if len(pair.p1.circuit.points) == len(points) {
			fmt.Println("circuit complete")
			fmt.Println("last two points")

			fmt.Println(pair.p1)
			fmt.Println(pair.p2)

			break
		}

		if !slices.Contains(circuits, pair.p1.circuit) {
			circuits = append(circuits, pair.p1.circuit)
		}

		if circuitToBeDeleted != nil {
			circuits = slices.DeleteFunc(circuits, func(circuit *circuit) bool {
				return circuit == circuitToBeDeleted
			})
		}

	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].points) > len(circuits[j].points)
	})
}

type circuit struct {
	points []*point
}

type pair struct {
	distance float64
	p1       *point
	p2       *point
}

type point struct {
	circuit *circuit
	x       float64
	y       float64
	z       float64
}

func newPair(p1, p2 *point) pair {
	return pair{
		p1:       p1,
		p2:       p2,
		distance: p1.distanceTo(p2),
	}
}

func newPoint(x, y, z float64) *point {
	return &point{
		x: x,
		y: y,
		z: z,
	}
}

func (p *point) String() string {
	return fmt.Sprintf("%v, %v, %v", p.x, p.y, p.z)
}

func (p *point) connectTo(p2 *point) bool {
	if p.circuit == nil && p2.circuit == nil {
		newCircuit := &circuit{
			points: []*point{p, p2},
		}

		p.circuit = newCircuit
		p2.circuit = newCircuit

		return true
	}

	if p.circuit == nil {
		p.circuit = p2.circuit
		p.circuit.points = append(p.circuit.points, p)

		return true
	}

	if p2.circuit == nil {
		p2.circuit = p.circuit
		p2.circuit.points = append(p2.circuit.points, p2)

		return true
	}

	if p.circuit != p2.circuit {
		p.circuit.points = append(p.circuit.points, p2.circuit.points...)

		for _, point := range p2.circuit.points {
			point.circuit = p.circuit
		}

		return true
	}

	return false
}

func (p *point) distanceTo(p2 *point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2) + math.Pow(p.z-p2.z, 2))
}

func (p *point) nearestNeighbor(neighbors []*point) *point {
	nearest := neighbors[0]

	for _, neighbor := range neighbors {
		if p.distanceTo(neighbor) < p.distanceTo(nearest) {
			nearest = neighbor
		}
	}

	return nearest
}

func createPointsFromLines(lines []string) []*point {
	var points []*point

	for _, line := range lines {
		coordinateStrings := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(coordinateStrings[0], 64)
		y, _ := strconv.ParseFloat(coordinateStrings[1], 64)
		z, _ := strconv.ParseFloat(coordinateStrings[2], 64)
		points = append(points, newPoint(x, y, z))
	}

	return points
}

func findAllPairs(points []*point) []pair {
	var pairs []pair

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			pairs = append(pairs, newPair(p1, p2))
		}
	}

	return pairs
}

func readFile(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
