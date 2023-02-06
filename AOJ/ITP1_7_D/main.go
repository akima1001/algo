// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var sc = bufio.NewScanner(os.Stdin)

// func read() string {
// 	sc.Scan()
// 	return sc.Text()
// }

// // sc.Split(bufio.ScanLines)

// func main() {
// 	var n, m, l int
// 	fmt.Scan(&n, &m, &l)
// 	var nm [][]int
// 	sc.Split(bufio.ScanWords)
// 	for i := 0; i < n; i++ {
// 		var tmparr []int
// 		for j := 0; j < m; j++ {
// 			tmp, _ := strconv.Atoi(read())
// 			tmparr = append(tmparr, tmp)
// 		}
// 		nm = append(nm, tmparr)
// 	}
// 	var ml [][]int
// 	for i := 0; i < m; i++ {
// 		var tmparr []int
// 		for j := 0; j < l; j++ {
// 			tmp, _ := strconv.Atoi(read())
// 			tmparr = append(tmparr, tmp)
// 		}
// 		ml = append(ml, tmparr)
// 	}
// 	c := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		c[i] = make([]int, l)
// 		for j := 0; j < m; j++ {
// 			for k := 0; k < l; k++ {
// 				c[i][j] = nm[i][j] * ml[j][i]
// 				fmt.Println(foo)
// 			}
// 		}
// 		fmt.Printf("\n")
// 	}
// }
