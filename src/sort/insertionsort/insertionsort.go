package main

import "fmt"

type list []int

func (items list) insertionSort() {
	n := len(items)
	for i := 1; i < n; i++ {
		key := items[i]
		j := i - 1
		for j >= 0 && items[j] > key {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = key
	}
}

func main() {
	items := list{8, 4, 1, 56, 3, -44, 23, -6, 28, 0}
	fmt.Println(items)
	items.insertionSort()
	fmt.Println(items)
}
