package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "testdata/test.txt"
	got := part1(testFile, 6)
	want := 16
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	fmt.Println("Running input.txt")
	// actualFile := "testdata/input.txt"
	// fmt.Println("Answer for part 1:", part1(actualFile, 64))
	// fmt.Println("Answer for part 2:", part2(actualFile, 26501365))
}
