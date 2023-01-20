package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	var h = s / (60 * 60)
	var m = (s - (60 * 60 * h)) / 60
	var se = s - (60 * 60 * h) - (60 * m)
	fmt.Printf("%d:%d:%d\n", h, m, se)
}
