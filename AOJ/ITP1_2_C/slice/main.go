package main

import "fmt"

func bubbleSort(numbers []int, size int) {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	arr := []int{a, b, c}

	bubbleSort(arr, 3)
	fmt.Println(arr[0], arr[1], arr[2])
}
