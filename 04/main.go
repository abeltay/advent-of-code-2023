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

type cards struct {
	winning []int
	numbers []int
}

func convertNumbers(s string) []int {
	s1 := strings.Fields(s)
	var arr []int
	for _, v := range s1 {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, i)
	}
	return arr
}

func parseInput(filename string) []cards {
	input := parseFile(filename)

	var arr []cards
	for _, v := range input {
		s := strings.Split(v, ": ")
		s1 := strings.Split(s[1], " | ")
		arr = append(arr, cards{
			winning: convertNumbers(s1[0]),
			numbers: convertNumbers(s1[1]),
		})
	}
	return arr
}

func winningNumbers(win, num []int) int {
	var count int
	for _, v := range num {
		for _, v1 := range win {
			if v == v1 {
				count++
			}
		}
	}
	return count
}

func part1(filename string) int {
	input := parseInput(filename)
	var answer int
	for _, v := range input {
		win := winningNumbers(v.winning, v.numbers)
		if win > 0 {
			points := math.Pow(2, float64(win-1))
			answer += int(points)
		}
	}
	return answer
}

func part2(filename string) int {
	input := parseInput(filename)
	wins := make([][]int, len(input))
	for i := range input {
		win := winningNumbers(input[i].winning, input[i].numbers)
		var arr []int
		for j := i + 1; j <= i+win; j++ {
			arr = append(arr, j)
		}
		wins[i] = arr
	}
	fmt.Println(wins)
	card := make([]int, len(input))
	for i := len(wins) - 1; i >= 0; i-- {
		card[i]++
		for _, v := range wins[i] {
			card[i] += card[v]
		}
	}
	var answer int
	for _, v := range card {
		answer += v
	}
	return answer
}
