package main

import "fmt"

// binary search perfoms an iterative binary searchon a sorted array! the array must be sorted
// returns the index of the target if found, otherwise -1

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// to calculate the middle index to divice the search space
		mid := low + (high-low)/2

		// to calculate if target is found at mid, and return the index
		if arr[mid] == target {
			return mid
		}

		// then if the target is smaller, i will search the left half (move high pointer)
		if arr[mid] > target {
			high = mid - 1
		} else {
			// if the target is actually larger, then i will search the right half (move low pointer)
			low = mid + 1
		}
	}
	// if loop exits, then the target was found
	return -1
}

func main() {
	// in binary search algorithm the array must bse sorted! most important rule
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 22, 55}
	var target int

	// to get user input for target from the command line
	fmt.Print("Enter the target value to search: ")
	fmt.Scan(&target)

	// call binary search function
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array \n", target)
	}
}
