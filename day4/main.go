package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	maxNeighbors     = 4
	occupiedSymbol   = '@'
	unoccupiedSymbol = '.'

	path = "input.txt"
)

func main() {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := readFile(file)

	removedRollCount := 0

	_, removedRollCount = removeAllRolls(lines, 0)

	fmt.Println(removedRollCount)
}

func occupiedNeighbors(lines []string, x, y int) int {
	occupiedCount := 0

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}

			if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0]) {
				continue
			}

			symbol := lines[i][j]

			if symbol == occupiedSymbol {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func readFile(f io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func removeAllRolls(lines []string, totalRemovedRolls int) ([]string, int) {
	newLines, rollsRemoved := removeMultipleRolls(lines)

	if rollsRemoved == 0 {
		return newLines, totalRemovedRolls
	}

	return removeAllRolls(newLines, totalRemovedRolls+rollsRemoved)
}

func removeMultipleRolls(lines []string) ([]string, int) {
	removedRollCount := 0

	for y, line := range lines {
		for x := range line {
			if lines[y][x] == occupiedSymbol && occupiedNeighbors(lines, x, y) < maxNeighbors {
				lines = removeRoll(lines, x, y)
				removedRollCount++
			}
		}
	}

	return lines, removedRollCount
}

func removeRoll(lines []string, x int, y int) []string {
	line := lines[y]

	runes := []rune(line)
	runes[x] = unoccupiedSymbol
	line = string(runes)

	lines[y] = line

	return lines
}
