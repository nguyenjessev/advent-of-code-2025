package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ranges := getRangesFromFile(file)

	sum := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if isInvalidID(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}

func createRangeFromString(s string) Range {
	rangeElements := strings.Split(s, "-")

	start, err := strconv.Atoi(rangeElements[0])
	if err != nil {
		log.Fatal(err)
	}

	end, err := strconv.Atoi(rangeElements[1])
	if err != nil {
		log.Fatal(err)
	}

	return Range{
		start: start,
		end:   end,
	}
}

func getRangesFromFile(f io.Reader) []Range {
	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	stringContents := string(contents)

	stringContents = strings.Trim(stringContents, "\n")

	stringRanges := strings.Split(stringContents, ",")

	var ranges []Range
	for _, r := range stringRanges {
		newRange := createRangeFromString(r)

		ranges = append(ranges, newRange)
	}

	return ranges
}

func isInvalidID(id int) bool {
	stringID := strconv.Itoa(id)

	middleIndex := len(stringID) / 2

	for chunkSize := 1; chunkSize <= middleIndex; chunkSize += 1 {
		if (len(stringID) % chunkSize) != 0 {
			continue
		}

		var chunks []string

		for i := 0; i < len(stringID); i += chunkSize {
			chunks = append(chunks, stringID[i:i+chunkSize])
		}

		if allSliceElementsEqual(chunks) {
			return true
		}
	}

	return false
}

func allSliceElementsEqual(s []string) bool {
	for i := 1; i < len(s); i++ {
		if s[0] != s[i] {
			return false
		}
	}

	return true
}
