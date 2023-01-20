package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Println(strconv.FormatFloat(math.Pow(x, 3), 'f', -1, 64))
}
