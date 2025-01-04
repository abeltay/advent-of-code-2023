package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abeltay/advent-of-code-2023/utilities"
)

func parseInput(filename string) ([]string, [][]int) {
	in := utilities.ParseFile(filename)

	groups := make([]string, 0, len(in))
	nums := make([][]int, 0, len(in))
	for _, v := range in {
		split := strings.Split(v, " ")
		n := strings.Split(split[1], ",")
		num := make([]int, 0, len(n))
		for _, v := range n {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			num = append(num, i)
		}
		groups = append(groups, split[0])
		nums = append(nums, num)
	}
	return groups, nums
}

func compressGroup(group string) string {
	var i int
	for i < len(group) && group[i] == '.' {
		i++
	}
	var s string
	for i < len(group) {
		if group[i] != '.' {
			if i-1 > 0 && group[i-1] == '.' {
				s += "."
			}
			s += string(group[i])
		}
		i++
	}
	return s
}

func valid(group string, springs []int) bool {
	length := springs[0]
	var i int
	for i < len(group) && i < length {
		if group[i] == '.' {
			return false
		}
		i++
	}
	if i != length {
		return false
	}
	if i < len(group) && group[i] == '#' {
		return false
	}
	if len(springs) == 1 {
		for j := i; j < len(group); j++ {
			if group[j] == '#' {
				return false
			}
		}
	}
	return true
}

func key(group string, num []int) string {
	s := group
	for _, v := range num {
		s += "," + strconv.Itoa(v)
	}
	return s
}

func combinations(group string, springs []int, cache map[string]int) int {
	if a, ok := cache[key(group, springs)]; ok {
		return a
	}
	length := springs[0]
	var total int
	for ptr := 0; ptr < len(group); ptr++ {
		if ptr > 0 && group[ptr-1] == '#' {
			break
		}
		if valid(group[ptr:], springs) {
			// fmt.Println("valid, 0 index", group[ptr:], springs)
			if len(springs) == 1 {
				total++
			} else {
				if len(group) >= ptr+length+1 {
					c := combinations(compressGroup(group[ptr+length+1:]), springs[1:], cache)
					total += c
				}
			}
		}
	}
	cache[key(group, springs)] = total
	return total
}

func part1(filename string) int {
	groups, springs := parseInput(filename)
	cache := make(map[string]int)
	var ans int
	for i := range groups {
		c := compressGroup(groups[i])
		ans += combinations(c, springs[i], cache)
	}
	return ans
}

func part2(filename string) int {
	groups, springs := parseInput(filename)
	cache := make(map[string]int)
	var ans int
	for i := range groups {
		newGroup := groups[i]
		for j := 1; j < 5; j++ {
			newGroup += "?" + groups[i]
		}
		newSprings := make([]int, 0, len(springs[i])*5)
		for j := 0; j < 5; j++ {
			newSprings = append(newSprings, springs[i]...)
		}
		ans += combinations(compressGroup(newGroup), newSprings, cache)
	}
	return ans
}
