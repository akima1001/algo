package main

import "fmt"

func main() {
	var x int
	i := 1
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		fmt.Printf("Case %v: %v\n", i, x)
		i++
	}
}
