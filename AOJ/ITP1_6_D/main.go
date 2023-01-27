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
	var n, m int
	fmt.Scan(&n, &m)
	sc.Split(bufio.ScanWords)
	var a [][]int
	var b []int
	for i := 0; i < n; i++ {
		var tmp_arr []int
		for j := 0; j < m; j++ {
			tmp, _ := strconv.Atoi(read())
			tmp_arr = append(tmp_arr, tmp)
		}
		a = append(a, tmp_arr)
	}
	for i := 0; i < m; i++ {
		tmp_b, _ := strconv.Atoi(read())
		b = append(b, tmp_b)
	}
	for i := 0; i < n; i++ {
		var ans int
		for j := 0; j < m; j++ {
			ans += a[i][j] * b[j]
		}
		fmt.Println(ans)
	}
}
