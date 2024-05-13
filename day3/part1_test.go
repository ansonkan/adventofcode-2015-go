package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	if Part1("input.txt") != 2081 {
		t.Fail()
	}
}
