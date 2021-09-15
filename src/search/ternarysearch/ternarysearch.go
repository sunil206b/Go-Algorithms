package main

import "fmt"

func ternarySearch(targetList []int, key int) int {
	low := 0
	high := len(targetList) - 1

	for low <= high {
		mid1 := low + ((high - low) / 3)
		mid2 := high - ((high - low) / 3)

		if targetList[mid1] == key {
			return mid1
		}
		if targetList[mid2] == key {
			return mid2
		}

		if key < targetList[mid1] {
			high = mid1 - 1
		} else if key > targetList[mid2] {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(ternarySearch(items, 5))
	fmt.Println(ternarySearch(items, 20))
	fmt.Println(ternarySearch(items, 0))
	fmt.Println(ternarySearch(items, 1))
	fmt.Println(ternarySearch(items, 10))
}
