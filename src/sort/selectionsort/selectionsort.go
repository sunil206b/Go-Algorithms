package main

import "fmt"

type list []int

func swap(xp *int, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}

func (items list) selectionSort() {
	n := len(items)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		swap(&items[i], &items[minIdx])
	}
}

func main() {
	unsortedList := list{64, 34, 25, 12, 22, 22, 11, 90, 45, 53, 100, 35}
	fmt.Println(unsortedList)
	unsortedList.selectionSort()
	fmt.Println(unsortedList)
}
