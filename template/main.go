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
	// sc.Split(bufio.ScanWords)
	// sc.Split(bufio.ScanLines)
	var i string
	fmt.Scan(&i)
}
