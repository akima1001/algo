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

// sc.Split(bufio.ScanLines)
func main() {
	var n int
	fmt.Scan(&n)
	sc.Split(bufio.ScanWords)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(read())
		arr[i] = num
	}
	for i := 0; i < n; i++ {
		isLoop := false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				isLoop = true
			}
		}
		if isLoop {
			for i, v := range arr {
				fmt.Print(v)
				if i < len(arr)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Printf("\n\n")
		} else {
			break
		}
	}
}
