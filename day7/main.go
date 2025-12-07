package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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
	beamMap := make(map[int]struct{})
	var beamLog []int
	startingIndex := strings.Index(lines[0], "S")
	beamMap[startingIndex] = struct{}{}
	beamLog = append(beamLog, startingIndex)
	splitCount := 0

	for _, line := range lines[1 : len(lines)-1] {
		var newBeamLog []int

		for i, char := range line {
			if char == '^' {
				for _, beam := range beamLog {
					if beam == i {
						newBeamLog = append(newBeamLog, i-1)
						newBeamLog = append(newBeamLog, i+1)
					}
				}

				if _, ok := beamMap[i]; ok {
					splitCount++
					beamMap[i-1] = struct{}{}
					beamMap[i+1] = struct{}{}
					delete(beamMap, i)
				}
			} else {
				for _, beam := range beamLog {
					if beam == i {
						newBeamLog = append(newBeamLog, beam)
					}
				}
			}
		}

		if strings.Contains(line, "^") {
			beamLog = slices.Clone(newBeamLog)

			fmt.Println(line)
		}
	}

	return splitCount, len(beamLog)
}

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
