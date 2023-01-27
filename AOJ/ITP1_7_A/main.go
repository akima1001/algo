package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

// sc.Split(bufio.ScanWords)
// sc.Split(bufio.ScanLines)

func main() {
	var m, f, r int
	for i := 0; i < 50; i++ {
		var rs string
		fmt.Scan(&m, &f, &r)
		if m == -1 && f == -1 && r == -1 {
			break
		} else if m == -1 || f == -1 {
			rs = "F"
		} else if m+f < 30 {
			rs = "F"
		} else if m+f >= 80 {
			rs = "A"
		} else if m+f >= 65 {
			rs = "B"
		} else if m+f >= 50 || r >= 50 {
			rs = "C"
		} else {
			rs = "D"
		}
		fmt.Println(rs)
	}
}
