package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var x int
	min, max := 1000000, -1000000
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		sum += x
	}
	fmt.Println(min, max, sum)
}
