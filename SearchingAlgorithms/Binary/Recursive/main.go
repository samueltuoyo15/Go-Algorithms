package main

import (
	"fmt"
	"log"
)

func recursiveBinarySearch(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if arr[mid] == target {
		// if the target is found then i return it
		return mid
	} else if target < arr[mid] {
		// if the target is lesser than the mid then i will search left
		return recursiveBinarySearch(arr, target, low, mid-1)
	} else {
		// else then i search right
		return recursiveBinarySearch(arr, target, mid+1, high)
	}
}

func main() {
	arr := []int{20, 21, 22, 23, 24, 25, 26, 66, 88}
	var target int

	log.Print("Enter target value to search: ")
	_, err := fmt.Scan(&target)
	if err != nil {
		log.Fatal("Failed to read input:", err)
	}

	index := recursiveBinarySearch(arr, target, 0, len(arr)-1)
	if index != -1 {
		log.Printf("target %d found at index %d\n", target, index)
	} else {
		log.Printf("target %d not found in the provided array of integers\n", target)
	}
}
