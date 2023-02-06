package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func read() int {
	sc.Scan()
	val, _ := strconv.Atoi((sc.Text()))
	return val
}

func main() {
	var n int
	sc.Split(bufio.ScanWords)
	n = read()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = read()
	}
	for i := 0; i < n; i++ {
		isLoop := false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isLoop = true
			}
		}
		if isLoop {
			/* 表示のループが原因のTLEだった */
			// for i, v := range arr {
			// 	fmt.Print(v)
			// 	if i < len(arr)-1 {
			// 		fmt.Print(" ")
			// 	}
			// }
			s := fmt.Sprintf("%d ", arr)
			s = strings.Replace(s, "[", "", -1)
			s = strings.Replace(s, "]", "", -1)
			fmt.Println(s)
		} else {
			break
		}
	}
}
