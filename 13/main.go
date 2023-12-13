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

func parseInput(filename string) [][][]string {
	in := parseFile(filename)

	var out [][][]string
	var arr [][]string
	for _, v := range in {
		if v == "" {
			out = append(out, arr)
			arr = make([][]string, 0)
			continue
		}
		var row []string
		for _, n := range v {
			row = append(row, string(n))
		}
		arr = append(arr, row)
	}
	out = append(out, arr)
	return out
}

func rowMirrored(maps [][]string, margin int) int {
	for y := range maps {
		if y == 0 {
			continue
		}
		var notMirrored bool
		smudge := margin
		for offset := 0; offset+y < len(maps) && y-offset-1 >= 0; offset++ {
			for x := range maps[y] {
				if maps[y+offset][x] != maps[y-offset-1][x] {
					if smudge > 0 {
						smudge--
					} else {
						notMirrored = true
						break
					}
				}
			}
			if notMirrored {
				break
			}
		}
		if !notMirrored && smudge == 0 {
			return y
		}
	}
	return 0
}

func columnMirrored(maps [][]string, margin int) int {
	for x := range maps[0] {
		if x == 0 {
			continue
		}
		var notMirrored bool
		smudge := margin
		for offset := 0; x+offset < len(maps[0]) && x-offset-1 >= 0; offset++ {
			for y := range maps {
				if maps[y][x+offset] != maps[y][x-offset-1] {
					if smudge > 0 {
						smudge--
					} else {
						notMirrored = true
						break
					}
				}
			}
			if notMirrored {
				break
			}
		}
		if !notMirrored && smudge == 0 {
			return x
		}
	}
	return 0
}

func sumMap(maps [][]string, margin int) int {
	if row := rowMirrored(maps, margin); row != 0 {
		return row * 100
	}
	if col := columnMirrored(maps, margin); col != 0 {
		return col
	}
	return 0
}

func part1(filename string) int {
	in := parseInput(filename)
	var ans int
	for _, v := range in {
		s := sumMap(v, 0)
		ans += s
	}
	return ans
}

func part2(filename string) int {
	in := parseInput(filename)
	var ans int
	for _, v := range in {
		s := sumMap(v, 1)
		ans += s
	}
	return ans
}
