package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesFromFile(f)

	currentPosition := 50
	minPosition := 0
	maxPosition := 99
	zeroCount := 0

	for _, move := range lines {
		direction, distance := parseMove(move)

		zeroCount += countZeroCrossings(currentPosition, direction, distance, minPosition, maxPosition)

		currentPosition = getPositionFromMove(currentPosition, direction, distance, minPosition, maxPosition)

		if currentPosition == 0 {
			zeroCount += 1
		}
	}

	fmt.Println(zeroCount)
}

func countZeroCrossings(currentPosition int, direction string, distance, minPosition, maxPosition int) int {
	var distanceToZero int
	zeroCrossings := 0

	if direction == "R" {
		distanceToZero = maxPosition - currentPosition + 1
	} else {
		distanceToZero = currentPosition - minPosition
	}

	if distanceToZero != 0 && distance > distanceToZero {
		zeroCrossings += 1
	}

	remainingDistance := distance - distanceToZero
	rangeSize := maxPosition - minPosition + 1

	if remainingDistance/rangeSize > 0 {
		zeroCrossings += remainingDistance / rangeSize

		if remainingDistance%rangeSize == 0 {
			zeroCrossings -= 1
		}
	}

	return zeroCrossings
}

func getLinesFromFile(file io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getPositionFromMove(currentPosition int, direction string, distance int, minPosition, maxPosition int) int {
	normalizedDistance := normalizeDistance(distance, minPosition, maxPosition)

	if direction == "L" {
		normalizedDistance = -normalizedDistance
	}

	newPosition := currentPosition + normalizedDistance

	if newPosition > maxPosition {
		newPosition = minPosition + (newPosition - maxPosition - 1)
	}

	if newPosition < minPosition {
		newPosition = maxPosition - (minPosition - newPosition - 1)
	}

	return newPosition
}

func normalizeDistance(distance int, min, max int) int {
	rangeSize := max - min + 1

	return distance % rangeSize
}

func parseMove(move string) (direction string, distance int) {
	if strings.HasPrefix(move, "R") {
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Print("invalid move: could not convert movement to distance")
		}

		return "R", distance
	}

	if strings.HasPrefix(move, "L") {
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Print("invalid move: could not convert movement to distance")
		}

		return "L", distance
	}

	log.Print("invalid move: could not determine movement direction")

	return "", 0
}
