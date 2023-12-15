package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var arr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr
}

func parseInput(filename string) [][]string {
	in := parseFile(filename)

	var arr [][]string
	for _, v := range in {
		var row []string
		for _, v1 := range v {
			row = append(row, string(v1))
		}
		arr = append(arr, row)
	}
	return arr
}

func shiftLeft(in []string) []string {
	var lastEmpty int
	for i := range in {
		if in[i] == "." {
			lastEmpty = i
			break
		}
	}
	for i := lastEmpty; i < len(in); i++ {
		switch in[i] {
		case "O":
			if lastEmpty >= len(in) {
				break
			}
			in[lastEmpty], in[i] = in[i], in[lastEmpty]
			lastEmpty++
		case "#":
			lastEmpty = i + 1
		}
	}
	return in
}

func tiltNorth(in [][]string) {
	for col := range in[0] {
		arr := make([]string, len(in))
		for k, v := range in {
			arr[k] = v[col]
		}
		shiftLeft(arr)
		for k := range in {
			in[k][col] = arr[k]
		}
	}
}

func tiltSouth(in [][]string) {
	for col := range in[0] {
		arr := make([]string, 0, len(in))
		for y := len(in) - 1; y >= 0; y-- {
			arr = append(arr, in[y][col])
		}
		shiftLeft(arr)
		for i := range arr {
			in[len(in)-1-i][col] = arr[i]
		}
	}
}

func tiltWest(in [][]string) {
	for y := range in {
		shiftLeft(in[y])
	}
}

func tiltEast(in [][]string) {
	for y := range in {
		arr := make([]string, 0, len(in[y]))
		for x := len(in[y]) - 1; x >= 0; x-- {
			arr = append(arr, in[y][x])
		}
		shiftLeft(arr)
		for i := range arr {
			in[y][len(in)-1-i] = arr[i]
		}
	}
}

func calculateLoad(in [][]string) int {
	var ans int
	for i, row := range in {
		var count int
		for _, v := range row {
			if v == "O" {
				count++
			}
		}
		ans += (len(in) - i) * count
	}
	return ans
}

func part1(filename string) int {
	in := parseInput(filename)
	tiltNorth(in)
	return calculateLoad(in)
}

func printMap(in [][]string) string {
	var s string
	for _, row := range in {
		for _, v := range row {
			s += v
		}
		s += " "
	}
	return s
}

func spinCycle(in [][]string) {
	tiltNorth(in)
	tiltWest(in)
	tiltSouth(in)
	tiltEast(in)
}

func part2(filename string) int {
	in := parseInput(filename)
	cache := make(map[string]int)
	var num1, num2 int
	for i := 0; i < 1000000000; i++ {
		spinCycle(in)
		m := printMap(in)
		if cache[m] != 0 {
			num1 = cache[m]
			num2 = i
			break
		}
		cache[m] = i
	}
	i := 1000000000 - num1
	remainder := i % (num2 - num1)
	for i := 1; i < remainder; i++ {
		spinCycle(in)
	}
	return calculateLoad(in)
}
