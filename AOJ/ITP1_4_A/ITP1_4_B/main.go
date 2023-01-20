package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var r float64
	fmt.Scan(&r)
	ans1 := strconv.FormatFloat(r*r*math.Pi, 'f', 6, 64)
	ans2 := strconv.FormatFloat(r*2*math.Pi, 'f', 6, 64)
	fmt.Printf("%v %v\n", ans1, ans2)
}
