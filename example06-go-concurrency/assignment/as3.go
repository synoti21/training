package main

import (
	"fmt"
	"sync"
)

func send(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sending: %d\n", i)
		ch <- i
	}
	close(ch)
}

func receive(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}

func as3() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 5)

	wg.Add(2)
	go send(ch, &wg)
	go receive(ch, &wg)

	wg.Wait()
}
