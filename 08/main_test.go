package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "testdata/input_test.txt"
	got := part1(testFile)
	want := 2
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
	testFile = "testdata/input2_test.txt"
	got = part1(testFile)
	want = 6
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
	testFile = "testdata/input3_test.txt"
	got2 := part2(testFile)
	want2 := 6
	if got2 != want2 {
		t.Fatalf("got %v want %v", got2, want2)
	}

	fmt.Println("Running input.txt")
	// actualFile := "testdata/input.txt"
	// fmt.Println("Answer for part 1:", part1(actualFile))
	// fmt.Println("Answer for part 2:", part2(actualFile))
}
