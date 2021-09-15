package main

import "fmt"

func swap(xp *int, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}

type list []int

func (items list) bubbleSort() {
	n := len(items)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if items[j+1] < items[j] {
				swap(&items[j+1], &items[j])
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	unsortedItems := list{64, 34, 25, 25, 12, 22, 11, 90}
	fmt.Println(unsortedItems)
	unsortedItems.bubbleSort()
	fmt.Println(unsortedItems)
}
