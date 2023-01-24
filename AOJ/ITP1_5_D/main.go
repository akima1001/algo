package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 1
	for {
		x := i
		if x%3 == 0 {
			fmt.Printf(" %d", i)
		} else {
			for {
				if x%10 == 3 {
					fmt.Printf(" %d", i)
					break
				}
				x /= 10
				if x <= 0 {
					break
				}
			}
		}
		i++
		if i > n {
			break
		}
	}
	fmt.Printf("\n")
}
