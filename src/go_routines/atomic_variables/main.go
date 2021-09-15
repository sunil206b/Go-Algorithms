package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func main() {
	go stingy()
	go spendy()
	time.Sleep(2500 * time.Millisecond)
	fmt.Println(money)
}

func stingy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, 10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy Done")
}