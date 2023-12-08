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

type node struct {
	left  string
	right string
}

func parseInput(filename string) (string, map[string]node) {
	in := parseFile(filename)

	nodes := make(map[string]node)
	for i := 2; i < len(in); i++ {
		var name, left, right string
		_, err := fmt.Sscanf(in[i], "%s = %s %s", &name, &left, &right)
		if err != nil {
			fmt.Println("scan error:", err)
			os.Exit(1)
		}
		nodes[name] = node{left[1 : len(left)-1], right[:len(right)-1]}
	}
	return in[0], nodes
}

func steps(start, end, direction string, dPtr int, nodes map[string]node) (int, string, int) {
	var steps int
	cur := start
	for {
		for dPtr < len(direction) {
			if direction[dPtr] == 'L' {
				cur = nodes[cur].left
			} else {
				cur = nodes[cur].right
			}
			steps++
			dPtr++
			if cur[len(cur)-len(end):] == end {
				return steps, cur, dPtr
			}
		}
		dPtr = 0
	}
}

func part1(filename string) int {
	direction, nodes := parseInput(filename)
	step, _, _ := steps("AAA", "ZZZ", direction, 0, nodes)
	return step
}

func factorise(start, direction string, nodes map[string]node) (int, int) {
	offset, item, ptr := steps(start, "Z", direction, 0, nodes)
	loopLength, _, _ := steps(item, item, direction, ptr, nodes)
	return offset, loopLength
}

func lcm(steps []int, factor []int) int {
	for {
		var notEquals bool
		for _, v := range steps {
			if v != steps[0] {
				notEquals = true
				break
			}
		}
		if !notEquals {
			return steps[0]
		}
		var highest int
		for i := range steps {
			if steps[i] > highest {
				highest = steps[i]
			}
		}
		for i := range steps {
			for steps[i] < highest {
				steps[i] += factor[i]
			}
		}
	}
}

func part2(filename string) int {
	direction, nodes := parseInput(filename)
	var cur []string
	for k := range nodes {
		if k[len(k)-1] == 'A' {
			cur = append(cur, k)
		}
	}
	offset := make([]int, len(cur))
	loopLength := make([]int, len(cur))
	for k, v := range cur {
		offset[k], loopLength[k] = factorise(v, direction, nodes)
	}
	return lcm(offset, loopLength)
}
