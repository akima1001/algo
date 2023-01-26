package main

import (
	"fmt"
)

func a(arr [13]bool, mark string) {
	for i := 0; i < 13; i++ {
		if arr[i] == false {
			fmt.Println(mark, i+1)
		}
	}
}

func main() {
	var n int
	var x string
	var y int
	var s [13]bool
	var h [13]bool
	var c [13]bool
	var d [13]bool
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		var key = y - 1
		if x == "S" {
			s[key] = true
		} else if x == "H" {
			h[key] = true
		} else if x == "C" {
			c[key] = true
		} else if x == "D" {
			d[key] = true
		}
	}
	a(s, "S")
	a(h, "H")
	a(c, "C")
	a(d, "D")
}
