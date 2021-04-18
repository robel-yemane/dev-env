package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("MinInt app running....")
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)

	min := MinInt(x, y)
	fmt.Println("Min of ", x, " and ", y, "is: ", min)
}
