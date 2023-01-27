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

func main() {
	var i string
	fmt.Scan(&i)
	// sc.Split(bufio.ScanWords)
	// sc.Split(bufio.ScanLines)
}
