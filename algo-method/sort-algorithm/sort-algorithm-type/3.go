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
func printArr(arr []string) {
	fmt.Println(strings.Join(arr, " "))
}

func readToInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func main() {
	// sc.Split(bufio.ScanLines)
	var n int
	fmt.Scan(&n)
	sc.Split(bufio.ScanWords)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readToInt()
	}

	for i := 1; i < n; i++ {
		pos := i
		for pos != 0 && arr[pos-1] > arr[pos] {
			arr[pos-1], arr[pos] = arr[pos], arr[pos-1]
			pos--
		}
		s := fmt.Sprintf("%v ", arr)
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		fmt.Println(s)
	}
}
