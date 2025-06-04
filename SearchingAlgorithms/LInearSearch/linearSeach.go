package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var target int

	fmt.Println("Enter the key to be searched: ")
	fmt.Scan(&target)

	index := linearSearch(arr, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in arry\n", target)
	}
}
