package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func readToInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func main() {
	// sc.Split(bufio.ScanLines)
	// input
	sc.Split(bufio.ScanWords)
	n := readToInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readToInt()
	}

	// sort
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		s := fmt.Sprintf("%d ", arr)
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		fmt.Println(s)
	}
}
