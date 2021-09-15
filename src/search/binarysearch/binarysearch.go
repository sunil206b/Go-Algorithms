package main

import "fmt"

func binarySearch(key int, targetList []int) int {
	low := 0
	high := len(targetList) - 1

	if targetList[0] > key || targetList[high] < key {
		return -1
	}

	for low <= high {
		median := (low + high) / 2

		if targetList[median] == key {
			return median
		}

		if targetList[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
	fmt.Println(binarySearch(62, items))
	fmt.Println(binarySearch(1, items))
	fmt.Println(binarySearch(100, items))
	fmt.Println(binarySearch(110, items))
	fmt.Println(binarySearch(0, items))
	fmt.Println(binarySearch(50, items))
	fmt.Println(binarySearch(30, items))
}
