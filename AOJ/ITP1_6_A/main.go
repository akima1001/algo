package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Print(a[n-1])
	for j := n - 1 - 1; ; j-- {
		if j < 0 {
			break
		}
		fmt.Print(" ")
		fmt.Print(a[j])
	}
	fmt.Printf("\n")
}
