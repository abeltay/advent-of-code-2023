package main

import (
	"bufio"
	"fmt"
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

func stringToIntArr(input string) []int {
	s := strings.Fields(input)
	var arr []int
	for k := 1; k < len(s); k++ {
		i, err := strconv.Atoi(s[k])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, i)
	}
	return arr
}

func parseInput(filename string) ([]int, []int) {
	input := parseFile(filename)

	time := stringToIntArr(input[0])
	dist := stringToIntArr(input[1])
	return time, dist
}

func findStart(time, dist int) int {
	for speed := 0; speed <= time; speed++ {
		timeLeft := time - speed
		distance := speed * timeLeft
		if distance > dist {
			return speed
		}
	}
	return 0
}

func findEnd(time, dist int) int {
	for speed := time; speed >= 0; speed-- {
		timeLeft := time - speed
		distance := speed * timeLeft
		if distance > dist {
			return speed
		}
	}
	return 0
}

func part1(filename string) int {
	time, dist := parseInput(filename)
	answer := 1
	for i := range time {
		start := findStart(time[i], dist[i])
		end := findEnd(time[i], dist[i])
		diff := 1 + end - start
		answer *= diff
	}
	return answer
}

func combineNum(input string) int {
	s := strings.Fields(input)
	var combine string
	for k := 1; k < len(s); k++ {
		combine += s[k]
	}
	i, err := strconv.Atoi(combine)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func part2(filename string) int {
	input := parseFile(filename)
	time := combineNum(input[0])
	dist := combineNum(input[1])

	start := findStart(time, dist)
	end := findEnd(time, dist)
	answer := 1 + end - start
	return answer
}
