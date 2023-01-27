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
	var i string
	fmt.Scan(&i)
}
