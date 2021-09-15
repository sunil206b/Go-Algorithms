package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 10, 10, 100, 2, 10, 11, 2, 11, 2}
	result := duplicateNumbers(arr)
	fmt.Println(result)
}

func duplicateNumbers(arr []int) []int {
	var result []int
	dupMap := make(map[int]int)
	for _, a := range arr {
		if count, ok := dupMap[a]; ok {
			dupMap[a] = count + 1
		} else {
			dupMap[a] = 1
		}
	}
	
	for key, val := range dupMap {
		if val > 1 {
			result = append(result, key)
		}
	}
	return result
}