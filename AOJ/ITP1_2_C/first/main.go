package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var min, mid, max int
	if a == b && b == c {
		max = a
		mid = b
		min = a
		min = c
	} else if a == b {
		if a > c {
			max = a
			mid = b
			min = c
		} else {
			max = c
			mid = a
			min = b
		}
	} else if a == c {
		if a > b {
			max = a
			mid = c
			min = b
		} else {
			max = b
			mid = a
			min = c
		}
	} else if b == c {
		if a > b {
			max = a
			mid = b
			min = c
		} else {
			max = b
			mid = c
			min = a
		}
	} else if a > b && b > c {
		max = a
		mid = b
		min = c
	} else if a > c && c > b {
		max = a
		mid = c
		min = b
	} else if b > a && a > c {
		max = b
		mid = a
		min = c
	} else if b > c && c > a {
		max = b
		mid = c
		min = a
	} else if c > a && a > b {
		max = c
		mid = a
		min = b
	} else {
		max = c
		mid = b
		min = a
	}
	fmt.Println(min, mid, max)
}
