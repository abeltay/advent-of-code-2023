package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testFile := "testdata/test.txt"
	got := part1(testFile)
	want := 54
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	fmt.Println("Running input.txt")
	actualFile := "testdata/input.txt"
	fmt.Println("Answer for part 1:", part1(actualFile))
}
