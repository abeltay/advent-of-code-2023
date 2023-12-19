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

type part struct {
	x int
	m int
	a int
	s int
}

func parseInput(filename string) (map[string]string, []part) {
	in := parseFile(filename)

	var rowPtr int
	workflow := make(map[string]string)
	for in[rowPtr] != "" {
		var i int
		for in[rowPtr][i] != '{' {
			i++
		}
		workflow[in[rowPtr][:i]] = in[rowPtr][i+1 : len(in[rowPtr])-1]
		rowPtr++
	}

	var arr []part
	for _, row := range in[rowPtr+1:] {
		var p part
		_, err := fmt.Sscanf(row, "{x=%d,m=%d,a=%d,s=%d}", &p.x, &p.m, &p.a, &p.s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, p)
	}
	return workflow, arr
}

type tester struct {
	xmas      byte
	isSmaller bool
	value     int
}

func (f tester) test(p part) bool {
	var value int
	switch f.xmas {
	case 'x':
		value = p.x
	case 'm':
		value = p.m
	case 'a':
		value = p.a
	default:
		value = p.s
	}
	if f.isSmaller {
		return value < f.value
	}
	return value > f.value
}

type step struct {
	name  string
	test  tester
	true  *step
	false *step
}

func buildStep(allSteps []step, cur *step, command string) {
	colon := strings.IndexByte(command, ':')
	num, err := strconv.Atoi(command[2:colon])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cur.test = tester{
		xmas:      command[0],
		isSmaller: command[1] == '<',
		value:     num,
	}
	comma := strings.IndexByte(command, ',')
	for i := range allSteps {
		if allSteps[i].name == command[colon+1:comma] {
			cur.true = &allSteps[i]
			break
		}
	}
	next := command[comma+1:]
	colon = strings.IndexByte(next, ':')
	if colon == -1 {
		for i := range allSteps {
			if allSteps[i].name == next {
				cur.false = &allSteps[i]
			}
		}
		return
	}
	cur.false = &step{}
	buildStep(allSteps, cur.false, command[comma+1:])
}

func createSteps(workflow map[string]string) *step {
	allSteps := []step{
		{"A", tester{}, nil, nil},
		{"R", tester{}, nil, nil},
	}
	for k := range workflow {
		allSteps = append(allSteps, step{k, tester{}, nil, nil})
	}
	for i := 2; i < len(allSteps); i++ {
		s := workflow[allSteps[i].name]
		buildStep(allSteps, &allSteps[i], s)
	}
	for i := range allSteps {
		if allSteps[i].name == "in" {
			return &allSteps[i]
		}
	}
	return &step{}
}

func stepNext(cur *step, p part) *step {
	if cur.test.test(p) {
		return cur.true
	}
	return cur.false
}

func part1(filename string) int {
	workflow, part := parseInput(filename)
	start := createSteps(workflow)
	var ans int
	for _, v := range part {
		current := start
		for {
			s := stepNext(current, v)
			if s.name == "R" {
				break
			}
			if s.name == "A" {
				ans += v.x + v.m + v.a + v.s
				break
			}
			current = s
		}
	}
	return ans
}

func handleLimits(isSmaller bool, min, max, value int) (int, int) {
	if isSmaller {
		if max > value {
			max = value
		}
	} else {
		if min < value {
			min = value
		}
	}
	return min, max
}

func findA(cur *step, path []tester) int {
	if cur.name == "R" {
		return 0
	}
	if cur.name == "A" {
		const start = 0
		const end = 4001
		min := part{x: start, m: start, a: start, s: start}
		max := part{x: end, m: end, a: end, s: end}
		for _, v := range path {
			switch v.xmas {
			case 'x':
				min.x, max.x = handleLimits(v.isSmaller, min.x, max.x, v.value)
			case 'm':
				min.m, max.m = handleLimits(v.isSmaller, min.m, max.m, v.value)
			case 'a':
				min.a, max.a = handleLimits(v.isSmaller, min.a, max.a, v.value)
			default:
				min.s, max.s = handleLimits(v.isSmaller, min.s, max.s, v.value)
			}
		}
		const modifier = -1
		return (max.x - min.x + modifier) * (max.m - min.m + modifier) * (max.a - min.a + modifier) * (max.s - min.s + modifier)
	}
	var total int
	newPath := append(path, cur.test)
	total += findA(cur.true, newPath)
	value := cur.test.value
	if cur.test.isSmaller {
		value--
	} else {
		value++
	}
	newPath = append(path, tester{
		xmas:      cur.test.xmas,
		isSmaller: !cur.test.isSmaller,
		value:     value,
	})
	total += findA(cur.false, newPath)
	return total
}

func part2(filename string) int {
	workflow, _ := parseInput(filename)
	start := createSteps(workflow)
	return findA(start, []tester{})
}
