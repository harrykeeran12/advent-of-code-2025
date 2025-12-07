package day5

import (
	"testing"
)

func TestInput1(t *testing.T) {
	ans := part1("input.txt")
	if ans != 3 {
		t.Errorf("TestPart1 wrong, wanted 3, got %d", ans)
	}
}

func TestInput2(t *testing.T) {
	part1("input2.txt")
}
