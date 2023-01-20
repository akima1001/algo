package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var ans string
	if a < b {
		ans = "a < b"
	} else if a > b {
		ans = "a > b"
	} else {
		ans = "a == b"
	}
	fmt.Println(ans)
}
