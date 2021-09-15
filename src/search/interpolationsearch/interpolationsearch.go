package main

import "fmt"

func interpolationSearch(key int, targetList []int) int {
	low := 0
	high := len(targetList) - 1

	if targetList[low] > key || targetList[high] < key {
		return -1
	}
	for low <= high && key >= targetList[low] && key <= targetList[high] {
		if low == high {
			if targetList[low] == key {
				return low
			}
			return -1
		}

		pos := low + ((high-low)/(targetList[high]-targetList[low]))*(key-targetList[low])

		if targetList[pos] == key {
			return pos
		}

		if targetList[pos] < key {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(63, items))
	fmt.Println(interpolationSearch(1, items))
	fmt.Println(interpolationSearch(100, items))
	fmt.Println(interpolationSearch(120, items))
	fmt.Println(interpolationSearch(50, items))
}
