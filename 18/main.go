package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type row struct {
	direction string
	steps     int
	color     string
}

func parseInput(filename string) []row {
	in := parseFile(filename)

	var arr []row
	for _, str := range in {
		s := strings.Fields(str)
		i, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, row{
			direction: s[0],
			steps:     i,
			color:     s[2],
		})
	}
	return arr
}

func move(x, y int, row row) (int, int) {
	newX, newY := x, y
	switch row.direction {
	case "U":
		newY -= row.steps
	case "D":
		newY += row.steps
	case "L":
		newX -= row.steps
	case "R":
		newX += row.steps
	}
	return newX, newY
}

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func area(in []row) int {
	x, y := 0, 0
	vertices := []point{{x, y}}
	for _, row := range in {
		x, y = move(x, y, row)
		vertices = append(vertices, point{x, y})
	}
	sum := vertices[len(vertices)-1].x*vertices[0].y - vertices[0].x*vertices[len(vertices)-1].y
	var boundary int
	for i := range vertices {
		if i == 0 {
			continue
		}
		sum += vertices[i-1].x*vertices[i].y - vertices[i].x*vertices[i-1].y
		boundary += abs(vertices[i-1].x-vertices[i].x) + abs(vertices[i-1].y-vertices[i].y)
	}
	return abs(sum)/2 + 1 + boundary/2
}

func part1(filename string) int {
	in := parseInput(filename)
	return area(in)
}

func part2(filename string) int {
	in := parseInput(filename)
	newRows := make([]row, 0, len(in))
	for _, r := range in {
		var direction string
		switch r.color[len(r.color)-2] {
		case '0':
			direction = "R"
		case '1':
			direction = "D"
		case '2':
			direction = "L"
		default:
			direction = "U"
		}
		i, err := strconv.ParseInt(r.color[2:len(r.color)-2], 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		r := row{
			direction,
			int(i),
			r.color,
		}
		newRows = append(newRows, r)
	}
	return area(newRows)
}
