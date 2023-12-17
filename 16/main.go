package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type direction int

const (
	north direction = iota
	south
	east
	west
)

type traveller struct {
	x    int
	y    int
	face direction
}

func follow(layout [][]byte, visited map[string]bool, travel traveller) []traveller {
	if travel.x < 0 || travel.x >= len(layout[0]) || travel.y < 0 || travel.y >= len(layout) {
		return []traveller{}
	}
	s := fmt.Sprintf("%d,%d,%d", travel.x, travel.y, travel.face)
	if visited[s] {
		return []traveller{}
	}
	visited[s] = true
	var newTraveller []traveller
	switch layout[travel.y][travel.x] {
	case '/':
		switch travel.face {
		case north:
			travel.face = east
			travel.x++
		case south:
			travel.face = west
			travel.x--
		case east:
			travel.face = north
			travel.y--
		case west:
			travel.face = south
			travel.y++
		}
	case '\\':
		switch travel.face {
		case north:
			travel.face = west
			travel.x--
		case south:
			travel.face = east
			travel.x++
		case east:
			travel.face = south
			travel.y++
		case west:
			travel.face = north
			travel.y--
		}
	case '|':
		switch travel.face {
		case north:
			travel.y--
		case south:
			travel.y++
		case east:
			fallthrough
		case west:
			newTraveller = append(newTraveller, traveller{x: travel.x, y: travel.y - 1, face: north})
			travel.face = south
			travel.y++
		}
	case '-':
		switch travel.face {
		case north:
			fallthrough
		case south:
			newTraveller = append(newTraveller, traveller{x: travel.x - 1, y: travel.y, face: west})
			travel.face = east
			travel.x++
		case east:
			travel.x++
		case west:
			travel.x--
		}
	default:
		switch travel.face {
		case north:
			travel.y--
		case south:
			travel.y++
		case east:
			travel.x++
		case west:
			travel.x--
		}
	}
	newTraveller = append(newTraveller, travel)
	return newTraveller
}

func energise(layout [][]byte, travellerStack []traveller) int {
	visited := make(map[string]bool)
	for len(travellerStack) > 0 {
		t := travellerStack[len(travellerStack)-1]
		travellerStack = travellerStack[:len(travellerStack)-1]
		newT := follow(layout, visited, t)
		travellerStack = append(travellerStack, newT...)
	}
	compressed := make(map[string]bool)
	for key, found := range visited {
		if found {
			compressed[key[:len(key)-2]] = true
		}
	}
	var ans int
	for _, found := range compressed {
		if found {
			ans++
		}
	}
	return ans
}

func part1(filename string) int {
	in := parseInput(filename)
	travellerStack := []traveller{
		{x: 0, y: 0, face: east},
	}
	return energise(in, travellerStack)
}

func part2(filename string) int {
	in := parseInput(filename)
	var arr []int
	for y := range in {
		if y == 0 || y == len(in)-1 {
			continue
		}
		arr = append(arr, energise(in, []traveller{
			{x: 0, y: y, face: east},
		}))
		arr = append(arr, energise(in, []traveller{
			{x: len(in[y]) - 1, y: y, face: west},
		}))
	}
	for x := range in[0] {
		if x == 0 || x == len(in[0])-1 {
			continue
		}
		arr = append(arr, energise(in, []traveller{
			{x: x, y: 0, face: south},
		}))
		arr = append(arr, energise(in, []traveller{
			{x: x, y: len(in) - 1, face: north},
		}))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr[0]
}
