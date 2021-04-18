package main

import (
	"fmt"
	"testing"
)

func TestMinIntBasic(t *testing.T) {
	fmt.Println("Running Basic Test...")
	min := MinInt(6, -5)
	if min != -5 {
		t.Errorf("MinInt(5, -5) returned %d; expecting -5", min)
	}
}

func TestMinIntTableDriven(tdt *testing.T) {
	var tests = []struct {
		x, y int
		want int
	}{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{0, -1, -1},
		{-2, -5, -5},
		{-5, -2, -5},
	}
	fmt.Println("Running Table driven tests")
	for _, t := range tests {
		testname := fmt.Sprintf("TestMinInt(): %d,%d", t.x, t.y)
		tdt.Run(testname, func(tdt *testing.T) {
			min := MinInt(t.x, t.y)
			if min != t.want {
				tdt.Errorf("MinInt(%d, %d)returned %d, expecting %d", t.x, t.y, min, t.want)
			}
		})
	}
}
