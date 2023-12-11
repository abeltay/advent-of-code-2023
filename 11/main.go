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

	var image [][]string
	for _, v := range in {
		var arr []string
		for _, v1 := range v {
			arr = append(arr, string(v1))
		}
		image = append(image, arr)
	}
	return image
}

func emptyColumns(image [][]string) []int {
	var arr []int
	for x := 0; x < len(image[0]); x++ {
		var hasGalaxy bool
		for y := range image {
			if image[y][x] == "#" {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			arr = append(arr, x)
		}
	}
	return arr
}

func emptyRows(image [][]string) []int {
	var arr []int
	for y := range image {
		var hasGalaxy bool
		for x := range image[y] {
			if image[y][x] == "#" {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			arr = append(arr, y)
		}
	}
	return arr
}

type coord struct {
	x, y int
}

func galaxies(image [][]string) []coord {
	var c []coord
	for y := range image {
		for x := range image[y] {
			if image[y][x] == "#" {
				c = append(c, coord{x, y})
			}
		}
	}
	return c
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func distance(image [][]string, emptyCol, emptyRow []int, expand int, g1, g2 coord) int {
	var times int
	for _, v := range emptyCol {
		if g1.x < v && v < g2.x || g2.x < v && v < g1.x {
			times++
		}
	}
	x := abs(g1.x-g2.x) + times*(expand-1)
	times = 0
	for _, v := range emptyRow {
		if g1.y < v && v < g2.y || g2.y < v && v < g1.y {
			times++
		}
	}
	y := abs(g1.y-g2.y) + times*(expand-1)
	return x + y
}

func totalDistance(in [][]string, expand int) int {
	emptyCol := emptyColumns(in)
	emptyRow := emptyRows(in)
	g := galaxies(in)
	var ans int
	for i := range g {
		for j := i + 1; j < len(g); j++ {
			ans += distance(in, emptyCol, emptyRow, expand, g[i], g[j])
		}
	}
	return ans
}

func part1(filename string) int {
	in := parseInput(filename)
	return totalDistance(in, 2)
}

func part2(filename string, expand int) int {
	in := parseInput(filename)
	return totalDistance(in, expand)
}
