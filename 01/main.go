package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
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

func number(input string) int {
	var num string
	for _, v := range input {
		if v > '0' && v <= '9' {
			num += string(v)
			break
		}
	}
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if c > '0' && c <= '9' {
			num += string(c)
			break
		}
	}
	i, _ := strconv.Atoi(num)
	return i
}

func part1(filename string) int {
	input := parseFile(filename)
	var answer int
	for _, v := range input {
		answer += number(v)
	}
	return answer
}

type wordNum struct {
	word string
	num  string
}

var wordNumList = []wordNum{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func findWord(input string) string {
	for _, v := range wordNumList {
		if len(v.word) > len(input) {
			continue
		}
		if v.word == input[:len(v.word)] {
			return v.num
		}
	}
	return ""
}

func numberFromWords(input string) int {
	var number string
	for i := range input {
		c := input[i]
		if c >= '0' && c <= '9' {
			number += string(c)
			break
		}
		if s := findWord(input[i:]); s != "" {
			number += s
			break
		}
	}
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if c >= '0' && c <= '9' {
			number += string(c)
			break
		}
		if s := findWord(input[i:]); s != "" {
			number += s
			break
		}
	}
	i, _ := strconv.Atoi(number)
	return i
}

func part2(filename string) int {
	input := parseFile(filename)
	var answer int
	for _, v := range input {
		answer += numberFromWords(v)
	}
	return answer
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 142 {
		fmt.Println("Wrong answer, got:", answer)
	}
	if answer := part2("input2_test.txt"); answer != 281 {
		fmt.Println("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
