package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "testdata/test.txt"
	got := part1(testFile)
	want := 32000000
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
	testFile = "testdata/test2.txt"
	got = part1(testFile)
	want = 11687500
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	fmt.Println("Running input.txt")
	// actualFile := "testdata/input.txt"
	// fmt.Println("Answer for part 1:", part1(actualFile))
	// fmt.Println("Answer for part 2:", part2(actualFile))
}
