package main

import "fmt"

func main() {
	var x, y int
	for {
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			break
		}
		if x < y {
			fmt.Printf("%v %v\n", x, y)
		} else {
			fmt.Printf("%v %v\n", y, x)
		}
	}
}
