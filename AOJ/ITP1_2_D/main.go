package main

import "fmt"

func main() {
	var w, h, x, y, r int
	ans := "No"
	fmt.Scan(&w, &h, &x, &y, &r)
	if x < 0 || y < 0 {
	} else if x+r <= w && y+r <= h {
		ans = "Yes"
	}

	fmt.Println(ans)
}
