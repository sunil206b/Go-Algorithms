package main

import (
	"fmt"
)

func main() {
	// var cities []string
	// log.Println("cities is equal to nil: ", cities == nil)
	// log.Printf("Cities %#v\n", cities)
	// log.Println(len(cities))

	// // cities[0] = "London"
	// numbers := []int{2, 3, 4, 5}
	// log.Println(numbers)

	// nums := make([]int, 2)
	// log.Printf("%#v\n", nums)

	// for i, v := range numbers {
	// 	log.Printf("index: %v, value: %v", i, v)
	// }

	// var n []int
	// fmt.Println(n == nil)

	// m := []int{}
	// fmt.Println(m == nil)

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d \n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("After Slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d \n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("After extending the length")
	fmt.Printf("s = %v, len = %d, cap = %d \n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("After droping first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d \n", s, len(s), cap(s))
}