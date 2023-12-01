package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func parseInput(filename string) []int {
	input, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}

	type line struct {
		first  int
		second int
		letter byte
		text   string
	}

	var arr []int
	for _, v := range input {
		var l line
		_, err := fmt.Sscanf(v, "%d-%d %c: %s", &l.first, &l.second, &l.letter, &l.text)
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, l.first)
	}
	return arr
}

func part1(filename string) int {
	input := parseInput(filename)
	fmt.Println(input)
	var answer int
	return answer
}

func part2(filename string) int {
	var answer int
	return answer
}
