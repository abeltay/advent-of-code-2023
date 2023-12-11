package main

import (
	"bufio"
	"fmt"
	"math"
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
		var bArr []string
		for _, b := range v {
			bArr = append(bArr, string(b))
		}
		arr = append(arr, bArr)
	}
	return arr
}

type coord struct {
	x, y int
}

type traveller struct {
	from coord
	to   coord
	cost int
}

func fromNorth(prev, at coord) bool {
	return prev.x == at.x && prev.y == at.y-1
}
func toNorth(at coord) coord {
	return coord{at.x, at.y - 1}
}

func fromSouth(prev, at coord) bool {
	return prev.x == at.x && prev.y == at.y+1
}
func toSouth(at coord) coord {
	return coord{at.x, at.y + 1}
}

func fromEast(prev, at coord) bool {
	return prev.x == at.x+1 && prev.y == at.y
}
func toEast(at coord) coord {
	return coord{at.x + 1, at.y}
}

func fromWest(prev, at coord) bool {
	return prev.x == at.x-1 && prev.y == at.y
}
func toWest(at coord) coord {
	return coord{at.x - 1, at.y}
}

func followPipe(pipes [][]string, prev, at coord) (coord, bool) {
	var next coord
	ok := true
	switch pipes[at.y][at.x] {
	case "|":
		switch {
		case fromNorth(prev, at):
			next = toSouth(at)
		case fromSouth(prev, at):
			next = toNorth(at)
		default:
			ok = false
		}
	case "-":
		switch {
		case fromWest(prev, at):
			next = toEast(at)
		case fromEast(prev, at):
			next = toWest(at)
		default:
			ok = false
		}
	case "L":
		switch {
		case fromNorth(prev, at):
			next = toEast(at)
		case fromEast(prev, at):
			next = toNorth(at)
		default:
			ok = false
		}
	case "J":
		switch {
		case fromNorth(prev, at):
			next = toWest(at)
		case fromWest(prev, at):
			next = toNorth(at)
		default:
			ok = false
		}
	case "7":
		switch {
		case fromWest(prev, at):
			next = toSouth(at)
		case fromSouth(prev, at):
			next = toWest(at)
		default:
			ok = false
		}
	case "F":
		switch {
		case fromEast(prev, at):
			next = toSouth(at)
		case fromSouth(prev, at):
			next = toEast(at)
		default:
			ok = false
		}
	default:
		ok = false
	}
	return next, ok
}

func travel(pipes [][]string, path [][]string, t traveller) (traveller, bool) {
	if t.to.x < 0 || t.to.x >= len(pipes[0]) {
		return traveller{}, false
	}
	if t.to.y < 0 || t.to.y >= len(pipes) {
		return traveller{}, false
	}
	path[t.to.y][t.to.x] = pipes[t.to.y][t.to.x]
	if pipes[t.to.y][t.to.x] == "S" {
		return t, true
	}
	next, ok := followPipe(pipes, t.from, t.to)
	if !ok {
		return traveller{}, false
	}
	return travel(pipes, path, traveller{t.to, next, t.cost + 1})
}

func createDistance(pipes [][]string) ([][]string, int) {
	var startX, startY int
	for y := range pipes {
		for x := range pipes[y] {
			if pipes[y][x] == "S" {
				startY = y
				startX = x
				break
			}
		}
	}
	path := make([][]string, len(pipes))
	for i := range pipes {
		arr := make([]string, len(pipes[0]))
		for j := range arr {
			arr[j] = " "
		}
		path[i] = arr
	}
	travellers := []traveller{
		{coord{startX, startY}, coord{startX - 1, startY}, 0},
		{coord{startX, startY}, coord{startX + 1, startY}, 0},
		{coord{startX, startY}, coord{startX, startY - 1}, 0},
		{coord{startX, startY}, coord{startX, startY + 1}, 0},
	}
	for _, t := range travellers {
		end, ok := travel(pipes, path, t)
		if ok {
			if path[startY][startX+1] == "7" {
				path[startY][startX] = "F"
			} else {
				path[startY][startX] = "|"
			}
			return path, int(math.Round(float64(end.cost) / 2))
		}
	}
	return nil, 0
}

func part1(filename string) int {
	in := parseInput(filename)
	_, ans := createDistance(in)
	return ans
}

func part2(filename string) int {
	in := parseInput(filename)
	path, _ := createDistance(in)
	for y := range path {
		var inner bool
		var encountered int
		for x := range path[y] {
			switch path[y][x] {
			case "|":
				inner = !inner
			case "L":
				encountered = 1
			case "J":
				if encountered != 1 {
					inner = !inner
				}
				encountered = 0
			case "F":
				encountered = -1
			case "7":
				if encountered == 1 {
					inner = !inner
				}
				encountered = 0
			case "S":
				inner = !inner
			case "-":
			default:
				if inner {
					path[y][x] = "X"
				} else {
					path[y][x] = "0"
				}
			}
		}
	}
	var ans int
	for y := range path {
		for x := range path[y] {
			if path[y][x] == "X" {
				ans++
			}
		}
	}
	return ans
}
