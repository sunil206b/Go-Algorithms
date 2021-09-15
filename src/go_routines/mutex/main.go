package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money = 100
	lock = sync.Mutex{}
	monyDeposited = sync.NewCond(&lock)
)

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(money)
}

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		monyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			monyDeposited.Wait()
		}
		money -= 20
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy Done")
}