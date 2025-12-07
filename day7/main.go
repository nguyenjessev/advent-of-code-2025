package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := readLines(file)

	fmt.Println(countSplits(lines))
}

func countSplits(lines []string) (int, int) {
	beamMap := make(map[int]int, 0)
	startingIndex := strings.Index(lines[0], "S")
	beamMap[startingIndex] = 1
	splitCount := 0

	for _, line := range lines[1 : len(lines)-1] {
		if !strings.Contains(line, "^") {
			continue
		}

		fmt.Println(line)

		newBeamMap := make(map[int]int, 0)

		for i := range beamMap {
			if line[i] == '^' {
				splitCount++

				newBeamMap[i-1] += beamMap[i]
				newBeamMap[i+1] += beamMap[i]
			} else {
				newBeamMap[i] += beamMap[i]
			}
		}

		beamMap = newBeamMap

		fmt.Println(line)
	}

	universeCount := 0
	for _, value := range beamMap {
		universeCount += value
	}

	return splitCount, universeCount
}

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
