package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	var div = a / b
	var pow = int(a) % int(b)
	fmt.Printf("%d ", int(div))
	fmt.Printf("%d ", pow)
	fmt.Println(strconv.FormatFloat(div, 'f', 5, 64))
}
