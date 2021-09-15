package main

import (
	"fmt"
	"math"
)

func jumpSearch(targetList []int, key int) int {
	length := len(targetList)
	step := math.Floor(math.Sqrt(float64(length)))

	if targetList[0] > key || targetList[length-1] < key {
		return -1
	}
	prev := 0
	for targetList[int(math.Min(step, float64(length))-1)] < key {
		prev = int(step)
		step += math.Floor(math.Sqrt(float64(length)))

		if prev >= length {
			return -1
		}
	}

	for targetList[prev] < key {
		prev++

		if prev == int(math.Min(step, float64(length))) {
			return -1
		}
	}

	if targetList[prev] == key {
		return prev
	}
	return -1
}

func main() {
	items := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	fmt.Println(jumpSearch(items, 55))
	fmt.Println(jumpSearch(items, 0))
	fmt.Println(jumpSearch(items, 610))
	fmt.Println(jumpSearch(items, 9))
	fmt.Println(jumpSearch(items, 620))
	fmt.Println(jumpSearch(items, -1))
}
