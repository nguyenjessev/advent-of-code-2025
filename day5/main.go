package main

import (
	"bufio"
	"io"
	"os"
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
	cleanedRanges := []ingredientRange{}

	for _, freshRange := range freshRanges {
		for _, existingRange := range cleanedRanges {
			if freshRange.start >= existingRange.start && freshRange.start <= existingRange.end {
				freshRange.start = existingRange.end + 1
			}

			if freshRange.end >= existingRange.start && freshRange.end <= existingRange.end {
				freshRange.end = existingRange.start - 1
			}
		}

		cleanedRanges = append(cleanedRanges, freshRange)
	}

	sum := 0

	for _, cleanedRange := range cleanedRanges {
		if cleanedRange.start > cleanedRange.end {
			continue
		}

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
