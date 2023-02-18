package main

import (
	"fmt"
	"log"
)

func main() {
	arrBS := []int{
		1, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9, 10,
	}
	arr2 = []int{0, 1, 2, 3, 4, 5, 6, 8, 9, 10}

	fmt.Println(BinarySearch(arrBS, 11))
	fmt.Println(LowerBound(arrBS, 3))
	fmt.Println(UpperBound(arrBS, 4))

}

// with binary search
func FirstMissingElement(arr []int, target int) int {

}

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
