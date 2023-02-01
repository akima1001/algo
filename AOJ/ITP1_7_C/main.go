package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Split(bufio.ScanLines)
	var r, c int
	var input [][]int
	fmt.Scan(&r, &c)
	var all int
	for i := 0; i < r; i++ {
		var xarr []int
		for j := 0; j < c; j++ {
			x, _ := strconv.Atoi(read())
			xarr = append(xarr, x)
		}
		input = append(input, xarr)
	}
	for i := 0; i < r; i++ {
		xsum := 0
		for j := 0; j < c; j++ {
			x := input[i][j]
			fmt.Print(x)
			fmt.Print(" ")
			xsum += input[i][j]
		}
		fmt.Print(xsum)
		fmt.Printf("\n")
	}
	for i := 0; i < c; i++ {
		ysum := 0
		for j := 0; j < r; j++ {
			y := input[j][i]
			ysum += y
		}
		fmt.Print(ysum)
		fmt.Print(" ")
		all += ysum
	}
	fmt.Println(all)
}
