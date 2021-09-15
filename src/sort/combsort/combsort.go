package main

import (
	"fmt"
	"math"
)

func swap(xp, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}

type list []int

func (items list) combSort() {
	n := len(items)
	gap := n
	sorted := false
	shrink := 1.3

	for sorted == false {
		gap = int(math.Floor(float64(gap) / shrink))
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		for i := 0; i < n-gap; i++ {
			if items[i] > items[i+gap] {
				swap(&items[i], &items[i+gap])
				sorted = false
			}
		}
	}
}

func main() {
	items := list{8, 4, 1, 56, 3, -44, 23, -6, 28, 0}
	fmt.Println(items)
	items.combSort()
	fmt.Println(items)
}
