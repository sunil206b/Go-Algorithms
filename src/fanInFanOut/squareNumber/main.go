package main

import (
	"log"
	"sync"
)

func main() {
	nums := []int{58, 39, 44, 32, 56, 84, 45, 78, 25}
	inputCh := generatePipeline(nums)
	dbCh1 := double(inputCh)
	dbCh2 := double(inputCh)
	dbCh3 := double(inputCh)
	result := fanIn(dbCh1, dbCh2, dbCh3)
	for r := range result {
		log.Println(r)
	}
}

func generatePipeline(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func double(numCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range numCh {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func fanIn(numChs ...<-chan int) <-chan int {
	result := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(numChs))
	for _, numCh := range numChs {
		go func(numCh <-chan int) {
			for num := range numCh {
				result <- num
			}
			wg.Done()
		}(numCh)	
	}

	go func(){
		wg.Wait()
		close(result)
	}()
	return result
}