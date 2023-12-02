package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var arr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr, nil
}

type colors struct {
	red   int
	green int
	blue  int
}

type game struct {
	num    int
	record []colors
}

func parseInput(filename string) []game {
	input, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}

	var arr []game
	for _, v := range input {
		line := strings.Split(v, ": ")
		var num int
		_, err := fmt.Sscanf(line[0], "Game %d", &num)
		if err != nil {
			log.Fatalln(err)
		}
		parts := strings.Split(line[1], "; ")
		var record []colors
		for _, part := range parts {
			var c colors
			colors := strings.Split(part, ", ")
			for _, v1 := range colors {
				v1s := strings.Split(v1, " ")
				i, err := strconv.Atoi(v1s[0])
				if err != nil {
					return nil
				}
				switch v1s[1] {
				case "red":
					c.red = i
				case "blue":
					c.blue = i
				default:
					c.green = i
				}
			}
			record = append(record, c)
		}
		arr = append(arr, game{
			num:    num,
			record: record,
		})
	}
	return arr
}

func isExceeded(actual []colors, max colors) bool {
	for _, v := range actual {
		if v.red > max.red {
			return true
		}
		if v.green > max.green {
			return true
		}
		if v.blue > max.blue {
			return true
		}
	}
	return false
}

func part1(filename string) int {
	input := parseInput(filename)
	total := colors{
		red:   12,
		green: 13,
		blue:  14,
	}
	var answer int
	for _, v := range input {
		if !isExceeded(v.record, total) {
			answer += v.num
		}
	}
	return answer
}

func minimum(actual []colors) colors {
	var max colors
	for _, v := range actual {
		if v.red > max.red {
			max.red = v.red
		}
		if v.green > max.green {
			max.green = v.green
		}
		if v.blue > max.blue {
			max.blue = v.blue
		}
	}
	return max
}

func part2(filename string) int {
	input := parseInput(filename)
	var answer int
	for _, v := range input {
		c := minimum(v.record)
		answer += c.red * c.green * c.blue
	}
	return answer
}
