package main

import (
	"bufio"
	"fmt"
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
	defer file.Close()

	lines := readFile(file)

	dataMatrix := parseDataMatrix(lines)

	result := 0
	for i := range dataMatrix[0] {
		result += performMath(dataMatrix, i)
	}

	fmt.Println(result)
}

func convertNumsToCephalopod(nums []string) []int {
	maxDigits := len(nums[0])

	var result []int
	for i := maxDigits; i > 0; i-- {
		var newNum strings.Builder

		for _, num := range nums {
			digit := num[i-1]

			if digit == ' ' {
				continue
			}

			newNum.WriteString(string(digit))
		}

		newNumInt, err := strconv.Atoi(newNum.String())
		if err != nil {
			panic(err)
		}

		result = append(result, newNumInt)
	}

	return result
}

func parseDataMatrix(lines []string) [][]string {
	matrix := make([][]string, len(lines))
	opsRow := lines[len(lines)-1]

	start := 0
	var end int

	for i := 1; i < len(opsRow); i++ {
		if opsRow[i] != ' ' {
			end = i - 1

			for j, line := range lines[:len(lines)-1] {
				matrix[j] = append(matrix[j], line[start:end])
			}

			start = i
		}
	}

	for i, line := range lines[:len(lines)-1] {
		matrix[i] = append(matrix[i], line[start:])
	}

	matrix[len(matrix)-1] = strings.Fields(opsRow)

	return matrix
}

func performMath(input [][]string, index int) int {
	operation := input[len(input)-1][index]

	var operands []string
	for _, line := range input[:len(input)-1] {
		num := line[index]

		operands = append(operands, num)
	}

	cephalopodOperands := convertNumsToCephalopod(operands)

	result := cephalopodOperands[0]

	for _, num := range cephalopodOperands[1:] {
		result = performOperation(result, num, operation)
	}

	return result
}

func performOperation(a, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unknown operation")
	}
}

func readFile(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
