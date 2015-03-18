package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	searchFor := 5
	fmt.Printf("Index for the number %d: %d \n", searchFor, binarySearch(numbers, searchFor))
}

func binarySearch(numbers []int, key int) int {
	return binarySearchAux(numbers, key, 0, len(numbers)-1)
}

func binarySearchAux(numbers []int, key int, min int, max int) int {
	if max < min {
		return -1
	} else {
		mid := (min + max) / 2

		if numbers[mid] > key {
			return binarySearchAux(numbers, key, min, mid-1)
		} else if numbers[mid] < key {
			return binarySearchAux(numbers, key, mid+1, max)
		} else {
			return mid
		}
	}
}
