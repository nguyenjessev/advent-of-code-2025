package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

	sumLights := 0
	sumJoltages := 0
	for _, line := range lines {
		l, b, j := machineComponents(line)

		sumLights += solveLights(l, b)
		sumJoltages += solveJoltages(j, b)
	}

	fmt.Println(sumLights)
	fmt.Println(sumJoltages)
}

type button []int

type joltages []int

type lights []bool

func newButton(s string) button {
	numbers := strings.Split(s[1:len(s)-1], ",")

	var result []int
	for _, numStr := range numbers {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		result = append(result, numInt)
	}

	return result
}

func newJoltages(s string) joltages {
	numbers := strings.Split(s[1:len(s)-1], ",")

	var result []int
	for _, numStr := range numbers {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		result = append(result, numInt)
	}

	return result
}

func newLights(s string) lights {
	var bits []bool

	for _, c := range s[1 : len(s)-1] {
		switch c {
		case '#':
			bits = append(bits, true)
		case '.':
			bits = append(bits, false)
		}
	}

	return bits
}

func (b *button) applyJoltages(j joltages, n int) (joltages, int) {
	result := slices.Clone(j)

	for _, index := range *b {
		result[index]++
	}

	return result, n + 1
}

func (b *button) applyLights(l lights, n int) (lights, int) {
	result := slices.Clone(l)

	for _, index := range *b {
		result[index] = !result[index]
	}

	return result, n + 1
}

func (j *joltages) String() string {
	var output strings.Builder

	for _, num := range *j {
		output.WriteString(fmt.Sprintf("%d", num))
	}

	return fmt.Sprintf("{%s}", output.String())
}

func (l *lights) String() string {
	var output strings.Builder

	for _, bit := range *l {
		if bit {
			output.WriteString("#")
		} else {
			output.WriteString(".")
		}
	}

	return fmt.Sprintf("[%s]", output.String())
}

func (j *joltages) greaterThan(other joltages) bool {
	for i, val := range *j {
		if val > other[i] {
			return true
		}
	}

	return false
}

func machineComponents(s string) (lights, []button, joltages) {
	fields := strings.Fields(s)

	resultLights := newLights(fields[0])

	var resultButtons []button
	for _, field := range fields[1 : len(fields)-1] {
		resultButtons = append(resultButtons, newButton(field))
	}

	resultJoltages := newJoltages(fields[len(fields)-1])

	return resultLights, resultButtons, resultJoltages
}

func readFile(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

type joltageState struct {
	vals  [10]int
	round int
}

func solveJoltages(target joltages, buttons []button) int {
	fmt.Println("solving for", target)

	var tgt [10]int
	copy(tgt[:], target)

	equal := func(a [10]int) bool {
		for i := range len(tgt) {
			if a[i] != tgt[i] {
				return false
			}
		}

		return true
	}

	greater := func(a [10]int) bool {
		for i := range len(tgt) {
			if a[i] > tgt[i] {
				return true
			}
		}

		return false
	}

	var queue []joltageState
	seenJoltages := make(map[[10]int]struct{})

	for _, button := range buttons {
		result := joltageState{round: 0}

		for _, val := range button {
			result.vals[val]++
		}

		queue = append(queue, result)

		seenJoltages[result.vals] = struct{}{}
	}

	for head := 0; head < len(queue); head++ {
		current := queue[head]

		for _, button := range buttons {
			nextResult := joltageState{
				vals:  current.vals,
				round: current.round + 1,
			}

			for _, val := range button {
				nextResult.vals[val]++
			}

			if equal(nextResult.vals) {
				return nextResult.round
			}

			if _, ok := seenJoltages[nextResult.vals]; ok {
				continue
			}

			if greater(nextResult.vals) {
				continue
			}

			queue = append(queue, nextResult)

			seenJoltages[nextResult.vals] = struct{}{}
		}
	}

	return -1
}

func solveLights(target lights, buttons []button) int {
	currentLights := make(lights, len(target))

	if slices.Equal(currentLights, target) {
		return 0
	}

	var queue []struct {
		result lights
		round  int
	}
	seenLights := make(map[string]struct{})

	for _, button := range buttons {
		result, round := button.applyLights(currentLights, 0)

		queue = append(queue, struct {
			result lights
			round  int
		}{result, round})

		seenLights[result.String()] = struct{}{}
	}

	for len(queue) > 0 {
		result := queue[0]
		queue = queue[1:]

		if slices.Equal(result.result, target) {
			return result.round
		}

		for _, button := range buttons {
			nextResult, nextRound := button.applyLights(result.result, result.round)

			if slices.Equal(nextResult, result.result) {
				continue
			}

			if _, ok := seenLights[nextResult.String()]; ok {
				continue
			}

			queue = append(queue, struct {
				result lights
				round  int
			}{nextResult, nextRound})

			seenLights[nextResult.String()] = struct{}{}
		}
	}

	return -1
}
