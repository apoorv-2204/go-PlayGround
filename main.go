package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{
		16, 17, 18, 1, 1, 1, 2, 3, 4, 5, 6, 7, 1, 2,
	}

	fmt.Println(PivotIndexBinarySearch(arr))

}

// sorted rotated array
func PivotIndexBinarySearch(arr []int) int {
	start := 0
	end := len(arr) - 1
	var mid int

	for start <= end {
		mid = start + (end-start)/2
		// terminating condition
		if arr[mid] > arr[mid+1] || arr[mid-1] > arr[mid] {
			return mid
		} else if arr[start] > arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return -1
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	var mid int

	for start <= end {
		// infinitte loop case for one size
		// calculate mid
		mid = start + (end-start)/2
		if start == end {
			return start
		}
		// terminating condition
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			// reduce to RHS SS
			start = mid + 1
		} else {
			// reduce the left search space only
			end = mid
		}
	}

	return -1

}

// with binary search
// https://leetcode.com/problems/missing-number/
// func FirstMissingElement(arr []int) int {
// 	// index is always sorted,perhaps for index 0 val find anomaly with comparison b.w index and the value arr[index]
// 	// determine how to decide , left or the right search space
// 	// RHS searh space: if index == arr[index], start = mid+1
// 	// LHS search Space: index < arr[index], end = mid -1
// 	// terminating condition: index < arr[index] && arr[mid-1] ==mid-1

// }

// in montonically ascebnding array an pivot it that element that break the monotonocity not overall
// such array is obtained by sorting and then rotating it
// func BinarySearchPivotElement(arr []int) int {
// 	start := 0
// 	end := len(arr) - 1
// 	last_index := len(arr) - 1
// 	var mid int
// 	for start >= end {

// 		mid = start + (end-start)/2

// 		// terminating condition
// 		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
// 			return mid
// 		}

// 		// reduction of search space
// 		if arr[start] < arr[mid] {
// 			// reduction to only RHS SS
// 			start = mid + 1
// 		} else {
// 			// reduction to only LHS SS
// 			end = mid - 1
// 		}
// 	}

// 	return -1
// }

// first occurence of target element
func LowerBound(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	var index int = -1
	for start <= end {
		var mid int = start + (end-start)/2
		if arr[mid] == target {
			index = mid
			end = mid - 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return index

}

// last occurence in sorted arrry
func UpperBound(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	var index int = -1
	for start <= end {
		var mid int = start + (end-start)/2
		if arr[mid] == target {
			index = mid
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return index

}

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		}
	}
	return -1
}

func TakeInput() {
	var str1 string
	var size int
	_, err := fmt.Scanln(&size, &str1)
	if err != nil {
		log.Fatal(err)
	}
}

// binary as it splits searcg space in two
// func BinarySearch() {
// 	target: =2
// 	slice := []int{1, 2, 3, 4, 5}
// }
