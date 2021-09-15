package main

import (
	"fmt"
	"strings"
)

func main() {
	var com complex128
	fmt.Println(com)
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
	sl := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}
  fmt.Println(strings.Join(sl, ""))
}