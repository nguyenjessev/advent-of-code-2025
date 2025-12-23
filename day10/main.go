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

type frac struct {
	num int64
	den int64
}

// newFrac normalizes rationals so we can do exact math during row reduction.
func newFrac(num, den int64) frac {
	if den < 0 {
		num = -num
		den = -den
	}
	if num == 0 {
		return frac{0, 1}
	}
	g := gcd(abs(num), den)
	return frac{num / g, den / g}
}

func (f frac) isZero() bool {
	return f.num == 0
}

func (f frac) add(g frac) frac {
	return newFrac(f.num*g.den+g.num*f.den, f.den*g.den)
}

func (f frac) sub(g frac) frac {
	return newFrac(f.num*g.den-g.num*f.den, f.den*g.den)
}

func (f frac) mul(g frac) frac {
	return newFrac(f.num*g.num, f.den*g.den)
}

func (f frac) div(g frac) frac {
	return newFrac(f.num*g.den, f.den*g.num)
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b
}

// rref performs Gauss-Jordan elimination and returns pivot columns plus consistency.
func rref(matrix [][]frac, rhs []frac) ([][]frac, []frac, []int, bool) {
	if len(matrix) == 0 {
		return matrix, rhs, nil, true
	}

	m := len(matrix)
	n := len(matrix[0])
	pivotCols := make([]int, 0, m)

	row := 0
	for col := 0; col < n && row < m; col++ {
		// Find a pivot row with a non-zero entry in this column.
		pivot := -1
		for r := row; r < m; r++ {
			if !matrix[r][col].isZero() {
				pivot = r
				break
			}
		}
		if pivot == -1 {
			continue
		}

		if pivot != row {
			matrix[row], matrix[pivot] = matrix[pivot], matrix[row]
			rhs[row], rhs[pivot] = rhs[pivot], rhs[row]
		}

		// Normalize the pivot row and eliminate this column everywhere else.
		pivotVal := matrix[row][col]
		for c := 0; c < n; c++ {
			matrix[row][c] = matrix[row][c].div(pivotVal)
		}
		rhs[row] = rhs[row].div(pivotVal)

		for r := 0; r < m; r++ {
			if r == row {
				continue
			}
			if matrix[r][col].isZero() {
				continue
			}
			factor := matrix[r][col]
			for c := 0; c < n; c++ {
				matrix[r][c] = matrix[r][c].sub(factor.mul(matrix[row][c]))
			}
			rhs[r] = rhs[r].sub(factor.mul(rhs[row]))
		}

		pivotCols = append(pivotCols, col)
		row++
	}

	// Any zero row with a non-zero RHS means the system is inconsistent.
	for r := 0; r < m; r++ {
		allZero := true
		for c := 0; c < n; c++ {
			if !matrix[r][c].isZero() {
				allZero = false
				break
			}
		}
		if allZero && !rhs[r].isZero() {
			return matrix, rhs, pivotCols, false
		}
	}

	return matrix, rhs, pivotCols, true
}

func solveJoltages(target joltages, buttons []button) int {
	if len(target) == 0 {
		return 0
	}

	// Precompute a tight upper bound for each button's presses.
	type buttonInfo struct {
		btn      button
		maxPress int
	}

	infos := make([]buttonInfo, len(buttons))
	for i, b := range buttons {
		maxPress := int(^uint(0) >> 1)
		for _, idx := range b {
			if target[idx] < maxPress {
				maxPress = target[idx]
			}
		}
		if maxPress == int(^uint(0)>>1) {
			maxPress = 0
		}
		infos[i] = buttonInfo{btn: b, maxPress: maxPress}
	}

	slices.SortFunc(infos, func(a, b buttonInfo) int {
		if a.maxPress != b.maxPress {
			return b.maxPress - a.maxPress
		}
		return len(b.btn) - len(a.btn)
	})

	// Build the button-position matrix (rows = positions, cols = buttons).
	m := len(target)
	k := len(infos)
	matrix := make([][]frac, m)
	zero := newFrac(0, 1)
	for i := range matrix {
		matrix[i] = make([]frac, k)
		for j := range matrix[i] {
			matrix[i][j] = zero
		}
	}
	for col, info := range infos {
		for _, idx := range info.btn {
			matrix[idx][col] = newFrac(1, 1)
		}
	}

	rhs := make([]frac, m)
	for i, v := range target {
		rhs[i] = newFrac(int64(v), 1)
	}

	// Reduce to RREF so we can separate forced vs free variables.
	rrefMatrix, rrefRHS, pivotCols, ok := rref(matrix, rhs)
	if !ok {
		return -1
	}

	pivotSet := make(map[int]struct{}, len(pivotCols))
	for _, col := range pivotCols {
		pivotSet[col] = struct{}{}
	}
	var freeCols []int
	for col := 0; col < k; col++ {
		if _, ok := pivotSet[col]; !ok {
			freeCols = append(freeCols, col)
		}
	}

	// Convert each pivot row into integer coefficients by clearing denominators.
	type rowEq struct {
		pivotCol int
		denom    int64
		coeffs   []int64
		constNum int64
	}

	rowEqs := make([]rowEq, len(pivotCols))
	for r := 0; r < len(pivotCols); r++ {
		den := int64(1)
		den = lcm(den, rrefRHS[r].den)

		for _, col := range freeCols {
			den = lcm(den, rrefMatrix[r][col].den)
		}

		coeffs := make([]int64, len(freeCols))
		for i, col := range freeCols {
			coeffs[i] = rrefMatrix[r][col].num * (den / rrefMatrix[r][col].den)
		}

		rowEqs[r] = rowEq{
			pivotCol: pivotCols[r],
			denom:    den,
			coeffs:   coeffs,
			constNum: rrefRHS[r].num * (den / rrefRHS[r].den),
		}
	}

	maxPresses := make([]int64, k)
	for i, info := range infos {
		maxPresses[i] = int64(info.maxPress)
	}

	best := int64(^uint64(0) >> 1)
	freeVals := make([]int64, len(freeCols))

	// Enumerate free variables within bounds, compute forced values, keep min sum.
	var dfs func(idx int, sum int64)
	dfs = func(idx int, sum int64) {
		if sum >= best {
			return
		}
		if idx == len(freeCols) {
			total := sum
			for _, eq := range rowEqs {
				num := eq.constNum
				for i, coeff := range eq.coeffs {
					num -= coeff * freeVals[i]
				}
				// Skip if the forced value is not an integer.
				if num%eq.denom != 0 {
					return
				}
				val := num / eq.denom
				if val < 0 {
					return
				}
				if val > maxPresses[eq.pivotCol] {
					return
				}
				total += val
				if total >= best {
					return
				}
			}
			if total < best {
				best = total
			}
			return
		}

		col := freeCols[idx]
		maxVal := maxPresses[col]
		for v := int64(0); v <= maxVal; v++ {
			freeVals[idx] = v
			dfs(idx+1, sum+v)
		}
	}

	dfs(0, 0)
	if best == int64(^uint64(0)>>1) {
		return -1
	}
	return int(best)
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
