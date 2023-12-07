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

func parseInput(filename string) []int {
	in := parseFile(filename)

	type line struct {
		first  int
		second int
		letter byte
		text   string
	}

	var arr []int
	for _, v := range in {
		var l line
		_, err := fmt.Sscanf(v, "%d-%d %c: %s", &l.first, &l.second, &l.letter, &l.text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, l.first)
	}
	return arr
}

func part1(filename string) int {
	in := parseInput(filename)
	fmt.Println(in)
	var ans int
	return ans
}

func part2(filename string) int {
	var ans int
	return ans
}
