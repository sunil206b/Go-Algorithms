package main

import "fmt"

func main() {
	arr := []int{1, 5, 3, 4, 1, 6, 6, 6, 8, 7, 13, 5}
	result := uniqueArray(arr)
	fmt.Println(result)
}

func uniqueArray(arr []int) []int {
	var result []int
	uniqueMap := make(map[int]bool)

	for _, a := range arr {
		if uniqueMap[a] != true {
			result = append(result, a)
			uniqueMap[a] = true
		}
	}
	return result
}