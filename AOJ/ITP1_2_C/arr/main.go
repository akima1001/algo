package main

import "fmt"

func bubbleSort(numbers [3]int, size int) {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	fmt.Println(numbers[0], numbers[1], numbers[2])
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	arr := [...]int{a, b, c}

	bubbleSort(arr, 3)
}
