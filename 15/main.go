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

func parseInput(filename string) []string {
	in := parseFile(filename)

	return strings.Split(in[0], ",")
}

func hash(s string) int {
	var hash int
	for _, v := range s {
		hash += int(v)
		hash *= 17
		hash %= 256
	}
	return hash
}

func part1(filename string) int {
	in := parseInput(filename)
	var ans int
	for _, v := range in {
		h := hash(v)
		ans += h
	}
	return ans
}

type lens struct {
	tag    string
	remove bool
	focal  string
}

func convertToLens(s string) lens {
	for i := range s {
		switch s[i] {
		case '=':
			return lens{
				tag:   s[:i],
				focal: s[i+1:],
			}
		case '-':
			return lens{
				tag:    s[:i],
				remove: true,
			}
		}

	}
	return lens{}
}

func updateBox(box []lens, new lens) []lens {
	var i int
	if len(box) > 0 {
		for i < len(box) {
			if box[i].tag == new.tag {
				break
			}
			i++
		}
	}
	if new.remove {
		if len(box) == 0 {
			return []lens{}
		}
		arr := make([]lens, 0, len(box)-1)
		arr = append(arr, box[:i]...)
		if i+1 < len(box) {
			arr = append(arr, box[i+1:]...)
		}
		return arr
	}
	if i == len(box) {
		return append(box, new)
	}
	box[i].focal = new.focal
	return box
}

func part2(filename string) int {
	in := parseInput(filename)
	arr := make([]lens, len(in))
	for i, v := range in {
		arr[i] = convertToLens(v)
	}
	boxes := make([][]lens, 256)
	for _, lens := range arr {
		num := hash(lens.tag)
		boxes[num] = updateBox(boxes[num], lens)
	}
	var ans int
	for boxNum, box := range boxes {
		for lensNum, lens := range box {
			i, err := strconv.Atoi(lens.focal)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			num := (boxNum + 1) * (lensNum + 1) * i
			fmt.Println(num)
			ans += num
		}
	}
	return ans
}
