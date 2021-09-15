package main

import "fmt"

func linearSearch(dataList []int, key int) int {
	for i, item := range dataList {
		if item == key {
			return i
		}
	}
	return -1
}

func main() {

	dataList := []int{95, 78, 46, 58, 86, 99, 251, 320}
	fmt.Println(linearSearch(dataList, 58))
	fmt.Println(linearSearch(dataList, 48))
}
