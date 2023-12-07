package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

type gardenMap struct {
	destination int
	source      int
	length      int
}

func strToIntArr(s string) []int {
	s1 := strings.Split(s, " ")
	var seeds []int
	for _, v := range s1 {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		seeds = append(seeds, n)
	}
	return seeds
}

func convertToGardenMap(input []string, line int) ([]gardenMap, int) {
	var category []gardenMap
	for line < len(input) && input[line] != "" {
		arr := strToIntArr(input[line])
		category = append(category, gardenMap{
			destination: arr[0],
			source:      arr[1],
			length:      arr[2],
		})
		line++
	}
	return category, line
}

func parseInput(filename string) ([]int, [][]gardenMap) {
	input := parseFile(filename)

	s := strings.Split(input[0], ": ")
	seeds := strToIntArr(s[1])
	line := 3
	var specs [][]gardenMap
	for line < len(input) && input[line] != "" {
		var category []gardenMap
		category, line = convertToGardenMap(input, line)
		specs = append(specs, category)
		line += 2
	}
	return seeds, specs
}

func part1(filename string) int {
	seed, specs := parseInput(filename)
	answer := math.MaxInt64
	for _, v := range seed {
		mapping := v
		for _, v1 := range specs {
			for _, v2 := range v1 {
				if v2.source <= mapping && mapping <= v2.source+v2.length {
					mapping = mapping - v2.source + v2.destination
					break
				}
			}
		}
		if mapping < answer {
			answer = mapping
		}
	}
	return answer
}

type point struct {
	point  int
	length int
}

func translate(seed point, garden []gardenMap) []point {
	var newPoints []point
	i := seed.point
	for i < seed.point+seed.length {
		var p point
		var found bool
		for _, v := range garden {
			if v.source <= i && i < v.source+v.length {
				end := min(seed.point+seed.length, v.source+v.length)
				p = point{i + v.destination - v.source, end - i}
				i = end
				found = true
				break
			}
		}
		if !found {
			p = seed
			i = seed.point + seed.length
		}
		newPoints = append(newPoints, p)
	}
	return newPoints
}

func lowestLocationNumber(seeds []point, specs [][]gardenMap) int {
	points := seeds
	for _, garden := range specs {
		var newPoints []point
		for _, seed := range points {
			p := translate(seed, garden)
			newPoints = append(newPoints, p...)
		}
		points = newPoints
	}
	answer := math.MaxInt64
	for _, v := range points {
		if v.point < answer {
			answer = v.point
		}
	}
	return answer
}

func part2(filename string) int {
	seed, specs := parseInput(filename)
	var seeds []point
	for i := 0; i < len(seed); i += 2 {
		seeds = append(seeds, point{
			point:  seed[i],
			length: seed[i+1],
		})
	}
	answer := math.MaxInt64
	for _, v := range seeds {
		num := lowestLocationNumber([]point{v}, specs)
		if num < answer {
			answer = num
		}
	}
	return answer
}
