package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type brick struct {
	cube [][3]int
}

type brickSlice []brick

func (x brickSlice) Len() int           { return len(x) }
func (x brickSlice) Less(i, j int) bool { return x[i].cube[0][2] < x[j].cube[0][2] }
func (x brickSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func parseInput(filename string) []brick {
	in := parseFile(filename)

	var arr []brick
	for _, row := range in {
		var x1, y1, z1, x2, y2, z2 int
		_, err := fmt.Sscanf(row, "%d,%d,%d~%d,%d,%d", &x1, &y1, &z1, &x2, &y2, &z2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bricks := brick{}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					bricks.cube = append(bricks.cube, [3]int{x, y, z})
				}
			}
		}
		arr = append(arr, bricks)
	}
	return arr
}

func dropOne(cur brick) brick {
	var nextPos brick
	for _, v := range cur.cube {
		nextPos.cube = append(nextPos.cube, [3]int{v[0], v[1], v[2] - 1})
	}
	return nextPos
}

func collisions(all []brick, cur int, newPos brick) []int {
	collision := make(map[int]bool)
	for i := range all {
		if i == cur {
			continue
		}
		for _, v := range all[i].cube {
			for _, v1 := range newPos.cube {
				if v[0] == v1[0] && v[1] == v1[1] && v[2] == v1[2] {
					collision[i] = true
				}
			}
		}
	}
	var arr []int
	for k, ok := range collision {
		if ok {
			arr = append(arr, k)
		}
	}
	return arr
}

func lower(all []brick, cur int, newPos brick) {
	const zFloor = 1
	if all[cur].cube[0][2] == zFloor {
		return
	}
	if i := collisions(all, cur, newPos); len(i) > 0 {
		return
	}
	all[cur] = newPos
	lower(all, cur, dropOne(all[cur]))
}

func gravity(all []brick) {
	for cur := range all {
		lower(all, cur, all[cur])
	}
}

func checkSupports(all []brick) [][]int {
	support := make([][]int, len(all))
	for i := len(all) - 1; i >= 0; i-- {
		support[i] = collisions(all, i, dropOne(all[i]))
	}
	return support
}

func canDisintegrate(all []brick, supports [][]int, cur int) bool {
	for _, supp := range supports {
		for _, v := range supp {
			if v == cur && len(supp) == 1 {
				return false
			}
		}
	}
	return true
}

func part1(filename string) int {
	in := parseInput(filename)
	sort.Sort(brickSlice(in))
	gravity(in)
	supports := checkSupports(in)
	var ans int
	for i := range in {
		if canDisintegrate(in, supports, i) {
			ans++
		}
	}
	return ans
}

func isSubset(super map[int]bool, sub []int) bool {
	if len(sub) == 0 {
		return false
	}
	for _, v := range sub {
		if !super[v] {
			return false
		}
	}
	return true
}

func chain(supports [][]int, cur int) int {
	drops := make(map[int]bool)
	drops[cur] = true
	i := 0
	for i < len(supports) {
		if !drops[i] && isSubset(drops, supports[i]) {
			drops[i] = true
			i = 1
		} else {
			i++
		}
	}
	return len(drops) - 1
}

func part2(filename string) int {
	in := parseInput(filename)
	sort.Sort(brickSlice(in))
	gravity(in)
	supports := checkSupports(in)
	var ans int
	// Can add a cache here cache:= make(map[int]map[int]bool)
	for i := range in {
		ans += chain(supports, i)
	}
	return ans
}
