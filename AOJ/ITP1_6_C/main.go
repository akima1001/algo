package main

import "fmt"

func aF(arr [3][10]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(" %d", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func bF() {
	for i := 0; i < 20; i++ {
		fmt.Print("#")
	}
	fmt.Printf("\n")
}

func main() {
	var n int
	fmt.Scan(&n)
	var b, f, r, v int
	var aT, bT, cT, dT [3][10]int
	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		if b == 1 {
			aT[f-1][r-1] += v
		} else if b == 2 {
			bT[f-1][r-1] += v
		} else if b == 3 {
			cT[f-1][r-1] += v
		} else if b == 4 {
			dT[f-1][r-1] += v
		}
	}
	aF(aT)
	bF()
	aF(bT)
	bF()
	aF(cT)
	bF()
	aF(dT)
}
