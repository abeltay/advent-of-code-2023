package main

import (
	"math/rand"
	"sort"
	"strings"

	"github.com/abeltay/advent-of-code-2023/utilities"
)

func parseInput(filename string) [][]bool {
	in := utilities.ParseFile(filename)

	all := make(map[string]bool)
	for _, row := range in {
		s := strings.Split(row, ": ")
		all[s[0]] = true
		for _, v := range strings.Split(s[1], " ") {
			all[v] = true
		}
	}
	names := make([]string, 0, len(all))
	for k := range all {
		names = append(names, k)
	}
	sort.Strings(names)
	connections := make([][]bool, len(names))
	for i := range names {
		connections[i] = make([]bool, len(names))
	}
	for _, row := range in {
		s := strings.Split(row, ": ")
		from := sort.SearchStrings(names, s[0])
		for _, v := range strings.Split(s[1], " ") {
			to := sort.SearchStrings(names, v)
			connections[from][to] = true
			connections[to][from] = true
		}
	}
	return connections
}

type node struct {
	num    int
	parent *node
}

func bfs(connections [][]bool, breaks []connect, from, to int) *node {
	visited := make([]bool, len(connections))
	visited[from] = true
	queue := []node{{num: from}}
	for len(queue) > 0 {
		last := queue[0]
		queue = queue[1:]
		if last.num == to {
			return &last
		}
		for i := range connections[last.num] {
			if !connections[last.num][i] || visited[i] {
				continue
			}
			var broken bool
			for _, v := range breaks {
				if v.from == last.num && v.to == i {
					broken = true
				}
			}
			if broken {
				continue
			}
			visited[i] = true
			queue = append(queue, node{num: i, parent: &last})
		}
	}
	return nil
}

type connect struct {
	from int
	to   int
}

func connectedAfter3Breaks(connections [][]bool, from, to int) []connect {
	var breaks []connect
	for i := 0; i < 3; i++ {
		path := bfs(connections, breaks, from, to)
		from := path.num
		next := path.parent
		for next != nil {
			breaks = append(breaks, connect{from, next.num})
			breaks = append(breaks, connect{next.num, from})
			from = next.num
			next = next.parent
		}
	}
	path := bfs(connections, breaks, from, to)
	if path != nil {
		return nil
	}
	return breaks
}

func findSize(connections [][]bool, breaks []connect, from int) int {
	visited := make([]bool, len(connections))
	visited[from] = true
	queue := []node{{num: from}}
	for len(queue) > 0 {
		last := queue[0]
		queue = queue[1:]
		for i := range connections[last.num] {
			if !connections[last.num][i] || visited[i] {
				continue
			}
			var broken bool
			for _, v := range breaks {
				if v.from == last.num && v.to == i {
					broken = true
					break
				}
			}
			if broken {
				continue
			}
			visited[i] = true
			queue = append(queue, node{num: i, parent: &last})
		}
	}
	var count int
	for _, v := range visited {
		if v {
			count++
		}
	}
	return count
}

func part1(filename string) int {
	connected := parseInput(filename)
	sameGroup := []int{0}
	source := 0
	r := rand.New(rand.NewSource(99))
	var breaks []connect
	for {
		n := int(r.Int31n(int32(len(connected))))
		var exists bool
		for _, v := range sameGroup {
			if v == n {
				exists = true
				break
			}
		}
		if exists {
			continue
		}
		breaks = connectedAfter3Breaks(connected, source, n)
		if breaks != nil {
			break
		}
		sameGroup = append(sameGroup, n)
	}
	size := findSize(connected, breaks, 0)
	return size * (len(connected) - size)
}
