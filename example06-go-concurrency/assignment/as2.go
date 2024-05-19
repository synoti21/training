package main

import (
	"fmt"
	"sync"
)

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}
func as2() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
