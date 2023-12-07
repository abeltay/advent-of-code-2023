package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func convertHand(s string) []int {
	letters := map[string]int{
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	var arr []int
	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			i = letters[string(v)]
		}
		arr = append(arr, i)
	}
	return arr
}

type hand struct {
	cards []int
	bid   int
}

func parseInput(filename string) []hand {
	in := parseFile(filename)

	var arr []hand
	for _, v := range in {
		s := strings.Split(v, " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, hand{
			convertHand(s[0]),
			bid,
		})
	}
	return arr
}

func maxRepeats(cards []int) (int, int) {
	count := make([]int, 15)
	for _, v := range cards {
		count[v]++
	}
	sort.Ints(count)
	max, max2 := count[len(count)-1], count[len(count)-2]
	return max, max2
}

func less(card1, card2 []int) bool {
	for i := range card1 {
		if card1[i] == card2[i] {
			continue
		}
		return card1[i] < card2[i]
	}
	return false
}

type handSlice []hand

func (x handSlice) Len() int { return len(x) }
func (x handSlice) Less(i, j int) bool {
	first, first2 := maxRepeats(x[i].cards)
	second, second2 := maxRepeats(x[j].cards)
	if first != second {
		return first < second
	}
	switch first {
	case 3:
		if first2 == 2 && second2 < 2 {
			return false
		}
		if second2 == 2 && first2 < 2 {
			return true
		}
	case 2:
		if first2 == 2 && second2 < 2 {
			return false
		}
		if second2 == 2 && first2 < 2 {
			return true
		}
	}
	return less(x[i].cards, x[j].cards)
}
func (x handSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func part1(filename string) int {
	in := parseInput(filename)
	sort.Sort(handSlice(in))
	var ans int
	for k, v := range in {
		rank := (1 + k) * v.bid
		ans += rank
	}
	return ans
}

func convertHand2(s string) []int {
	letters := map[string]int{
		"T": 10,
		"J": 0,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	var arr []int
	for _, v := range s {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			i = letters[string(v)]
		}
		arr = append(arr, i)
	}
	return arr
}

func parseInput2(filename string) []hand {
	in := parseFile(filename)

	var arr []hand
	for _, v := range in {
		s := strings.Split(v, " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		arr = append(arr, hand{
			convertHand2(s[0]),
			bid,
		})
	}
	return arr
}

func maxRepeats2(cards []int) (int, int) {
	count := make([]int, 15)
	for _, v := range cards {
		count[v]++
	}
	sort.Ints(count[1:])
	max, max2 := count[len(count)-1]+count[0], count[len(count)-2]
	return max, max2
}

type handSlice2 []hand

func (x handSlice2) Len() int { return len(x) }
func (x handSlice2) Less(i, j int) bool {
	first, first2 := maxRepeats2(x[i].cards)
	second, second2 := maxRepeats2(x[j].cards)
	if first != second {
		return first < second
	}
	switch first {
	case 3:
		if first2 == 2 && second2 < 2 {
			return false
		}
		if second2 == 2 && first2 < 2 {
			return true
		}
	case 2:
		if first2 == 2 && second2 < 2 {
			return false
		}
		if second2 == 2 && first2 < 2 {
			return true
		}
	}
	return less(x[i].cards, x[j].cards)
}
func (x handSlice2) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func part2(filename string) int {
	in := parseInput2(filename)
	sort.Sort(handSlice2(in))
	var ans int
	for k, v := range in {
		rank := (1 + k) * v.bid
		ans += rank
	}
	return ans
}
