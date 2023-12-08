package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "testdata/test.txt"
	got := part1(testFile)
	want := 0
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
	// got2 := part2(testFile)
	// want2 := 0
	// if got2 != want2 {
	// 	t.Fatalf("got %v want %v", got2, want2)
	// }

	fmt.Println("Running input.txt")
	actualFile := "testdata/input.txt"
	fmt.Println("Answer for part 1:", part1(actualFile))
	// fmt.Println("Answer for part 2:", part2(actualFile))
}
