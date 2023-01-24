package main

import "fmt"

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < w; i++ {
			fmt.Print("#")
		}
		fmt.Println()
		for i := 0; i < h-2; i++ {
			fmt.Print("#")
			for j := 0; j < w-2; j++ {
				fmt.Print(".")
			}
			fmt.Println("#")
		}
		for i := 0; i < w; i++ {
			fmt.Print("#")
		}
		fmt.Printf("\n\n")
	}
}
