package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	freshRanges, availableIngredients := readFile(file)

	println(countFreshIngredients(freshRanges, availableIngredients))

	println(countAllValidIngredients(freshRanges))
}

type ingredientRange struct {
	start int
	end   int
}

func countAllValidIngredients(freshRanges []ingredientRange) int {
	sort.Slice(freshRanges, func(i, j int) bool {
		if freshRanges[i].start == freshRanges[j].start {
			return freshRanges[i].end < freshRanges[j].end
		}
		return freshRanges[i].start < freshRanges[j].start
	})

	merged := []ingredientRange{freshRanges[0]}
	for _, current := range freshRanges[1:] {
		last := &merged[len(merged)-1]

		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}

			continue
		}

		merged = append(merged, current)
	}

	sum := 0
	for _, cleanedRange := range merged {
		sum += cleanedRange.end - cleanedRange.start + 1
	}

	return sum
}

func countFreshIngredients(freshRanges []ingredientRange, available []int) int {
	freshCount := 0

	for _, ingredient := range available {
		for _, freshRange := range freshRanges {

			if ingredient >= freshRange.start && ingredient <= freshRange.end {
				freshCount++

				break
			}
		}
	}

	return freshCount
}

func parseStringToRange(s string) ingredientRange {
	rangeElements := strings.Split(s, "-")

	start, err := strconv.Atoi(rangeElements[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(rangeElements[1])
	if err != nil {
		panic(err)
	}

	return ingredientRange{
		start: start,
		end:   end,
	}
}

func readFile(r io.Reader) ([]ingredientRange, []int) {
	scanner := bufio.NewScanner(r)
	readingIngredients := false

	var freshRanges []ingredientRange
	var availableIngredients []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingIngredients = true

			continue
		}

		if readingIngredients {
			ingredient, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			availableIngredients = append(availableIngredients, ingredient)
		} else {
			ingredientRange := parseStringToRange(line)
			freshRanges = append(freshRanges, ingredientRange)
		}
	}

	return freshRanges, availableIngredients
}
