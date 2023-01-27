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
	var n, x int
	for {
		fmt.Scan(&n, &x)
		if n == 0 && x == 0 {
			break
		}
		var cnt int
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				for k := j + 1; k < n; k++ {
					if i+1+j+1+k+1 == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
