package main

import "fmt"

func main() {
	var a, b int
	var op string
	var ans int
	for {
		fmt.Scan(&a, &op, &b)
		if op == "?" {
			break
		}
		if op == "+" {
			ans = a + b
		} else if op == "-" {
			ans = a - b
		} else if op == "*" {
			ans = a * b
		} else if op == "/" {
			ans = a / b
		}
		fmt.Println(ans)
	}
}
