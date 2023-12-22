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

func parseInput(filename string) [][]byte {
	in := parseFile(filename)

	var arr [][]byte
	for _, row := range in {
		arr = append(arr, []byte(row))
	}
	return arr
}

type point struct {
	x     int
	y     int
	steps int
}

func explore(in [][]byte, cost [][]int, p point) []point {
	if p.x < 0 || p.x >= len(in[0]) || p.y < 0 || p.y >= len(in) || in[p.y][p.x] == '#' || (cost[p.y][p.x] != 0 && cost[p.y][p.x] <= p.steps) {
		return []point{}
	}
	cost[p.y][p.x] = p.steps
	newPoints := []point{
		{p.x - 1, p.y, p.steps + 1},
		{p.x + 1, p.y, p.steps + 1},
		{p.x, p.y - 1, p.steps + 1},
		{p.x, p.y + 1, p.steps + 1},
	}
	return newPoints
}

func calcCost(in [][]byte, start point, steps int) [][]int {
	cost := make([][]int, len(in))
	for i := range in {
		arr := make([]int, len(in[0]))
		cost[i] = arr
	}
	points := []point{start}
	for len(points) > 0 {
		var newPoints []point
		for _, p := range points {
			if p.steps > steps {
				break
			}
			np := explore(in, cost, p)
			newPoints = append(newPoints, np...)
		}
		points = newPoints
	}
	return cost
}

func part1(filename string, steps int) int {
	in := parseInput(filename)
	var start point
	for y := range in {
		for x := range in[y] {
			if in[y][x] == 'S' {
				start = point{x: x, y: y}
				break
			}
		}
	}
	cost := calcCost(in, start, steps)
	var ans int
	for _, row := range cost {
		for _, item := range row {
			if item != 0 && item%2 == 0 {
				ans++
			}
			// fmt.Print(item, " ")
		}
		// fmt.Println()
	}
	return ans
}

// Part 2 solution from: https://github.com/omotto/AdventOfCode2023/blob/main/src/day21/main.go
