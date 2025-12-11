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

	lines := readFile(file)

	nodes := make(map[string]node)
	for _, line := range lines {
		n := newNode(line)
		nodes[n.name] = n
	}

	svrToDac := countPaths("svr", "dac", nodes)
	clear(memos)
	svrToFft := countPaths("svr", "fft", nodes)
	clear(memos)
	dacToFft := countPaths("dac", "fft", nodes)
	clear(memos)
	fftToDac := countPaths("fft", "dac", nodes)
	clear(memos)
	dacToOut := countPaths("dac", "out", nodes)
	clear(memos)
	fftToOut := countPaths("fft", "out", nodes)

	fmt.Println("svr to dac", svrToDac)
	fmt.Println("svr to fft", svrToFft)
	fmt.Println("dac to fft", dacToFft)
	fmt.Println("fft to dac", fftToDac)
	fmt.Println("dac to out", dacToOut)
	fmt.Println("fft to out", fftToOut)

	count := (svrToDac * dacToFft * fftToOut) + (svrToFft * fftToDac * dacToOut)
	fmt.Println("count", count)
}

type node struct {
	name        string
	connections []string
}

func newNode(line string) node {
	fields := strings.Fields(line)

	return node{
		name:        strings.TrimRight(fields[0], ":"),
		connections: fields[1:],
	}
}

var memos = make(map[string]int)

func countPaths(start string, end string, nodes map[string]node) int {
	if start == end {
		return 1
	}

	if count, ok := memos[start]; ok {
		return count
	}

	count := 0
	for _, connection := range nodes[start].connections {
		count += countPaths(connection, end, nodes)
	}

	memos[start] = count

	return count
}

func readFile(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
