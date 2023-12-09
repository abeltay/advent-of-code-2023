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

func parseInput(filename string) [][]int {
	in := parseFile(filename)

	var arr [][]int
	for _, v := range in {
		s := strings.Fields(v)
		var numArr []int
		for _, num := range s {
			i, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			numArr = append(numArr, i)
		}
		arr = append(arr, numArr)
	}
	return arr
}

func next(in []int, backwards bool) int {
	var foundNonZero bool
	for _, v := range in {
		if v != 0 {
			foundNonZero = true
		}
	}
	if !foundNonZero {
		return 0
	}
	out := make([]int, len(in)-1)
	for i := range out {
		out[i] = in[i+1] - in[i]
	}
	if !backwards {
		below := next(out, backwards)
		return in[len(in)-1] + below
	} else {
		below := next(out, backwards)
		return in[0] - below
	}
}

func part1(filename string) int {
	in := parseInput(filename)
	var ans int
	for _, v := range in {
		n := next(v, false)
		ans += n
	}
	return ans
}

func part2(filename string) int {
	in := parseInput(filename)
	var ans int
	for _, v := range in {
		n := next(v, true)
		ans += n
	}
	return ans
}
