package main

import (
	"fmt"

	"github.com/abeltay/advent-of-code-2023/utilities"
)

/*
type line struct {
	first  int
	letter byte
	text   string
}
*/

func parseInput(filename string) [][]byte {
	in := utilities.ParseFile(filename)

	var arr [][]byte
	for _, row := range in {
		/*
			var l line
			_, err := fmt.Sscanf(row, "%d %c: %s", &l.first, &l.letter, &l.text)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			arr = append(arr, l)
		*/
		arr = append(arr, []byte(row))
	}
	return arr
}

func part1(filename string) int {
	in := parseInput(filename)
	for _, row := range in {
		for _, item := range row {
			fmt.Print(string(item))
		}
		fmt.Println()
	}
	// fmt.Println(in)
	var ans int
	return ans
}

func part2(filename string) int {
	return 0
}
