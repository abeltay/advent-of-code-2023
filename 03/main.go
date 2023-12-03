package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open %q: %q\n", filename, err)
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
	input := parseFile(filename)

	var arr2d [][]string
	for _, v := range input {
		var arr []string
		for _, c := range v {
			arr = append(arr, string(c))
		}
		arr2d = append(arr2d, arr)
	}
	return arr2d
}

func isNum(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func hasAdjacentSymbol(schema [][]string, x, y int) bool {
	for y1 := y - 1; y1 <= y+1; y1++ {
		if y1 < 0 || y1 >= len(schema) {
			continue
		}
		for x1 := x - 1; x1 <= x+1; x1++ {
			if x1 < 0 || x1 >= len(schema[y1]) {
				continue
			}
			c := schema[y1][x1][0]
			if !(isNum(c)) && c != '.' {
				return true
			}
		}
	}
	return false
}

func convertToNum(sarr []string) int {
	var s string
	for _, v := range sarr {
		s += v
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	return num
}

func part1(filename string) int {
	input := parseInput(filename)
	var answer int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			var valid bool
			if isNum(input[y][x][0]) {
				x1 := x
				for ; x1 < len(input[y]); x1++ {
					if !isNum(input[y][x1][0]) {
						break
					}
					if hasAdjacentSymbol(input, x1, y) {
						valid = true
					}
				}
				if valid {
					sarr := input[y][x:x1]
					answer += convertToNum(sarr)
				}
				if x1-1 > x {
					x = x1 - 1
				}
			}
		}
	}
	return answer
}

type coord struct {
	x, y int
}

func findAsterix(input [][]string) []coord {
	var arr []coord
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == "*" {
				arr = append(arr, coord{x, y})
			}
		}
	}
	return arr
}

func findParts(schema [][]string, asterix coord) []int {
	var parts []int
	for y1 := asterix.y - 1; y1 <= asterix.y+1; y1++ {
		if y1 < 0 || y1 >= len(schema) {
			continue
		}
		for x1 := asterix.x - 1; x1 <= asterix.x+1; x1++ {
			if x1 < 0 || x1 >= len(schema[y1]) {
				continue
			}
			if isNum(schema[y1][x1][0]) {
				start := x1 - 1
				for ; start >= 0; start-- {
					if !isNum(schema[y1][start][0]) {
						break
					}
				}
				start++
				end := x1 + 1
				for ; end < len(schema[y1]); end++ {
					if !isNum(schema[y1][end][0]) {
						break
					}
				}
				x1 = end
				parts = append(parts, convertToNum(schema[y1][start:end]))
			}
		}
	}
	return parts
}

func part2(filename string) int {
	input := parseInput(filename)
	coord := findAsterix(input)
	var answer int
	for _, v := range coord {
		p := findParts(input, v)
		if len(p) == 2 {
			answer += p[0] * p[1]
		}
	}
	return answer
}
