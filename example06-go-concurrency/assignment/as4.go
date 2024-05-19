package main

import (
	"fmt"
	"sync"
)

const firstTask = 50
const secondTask = 60

func first(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= firstTask; i++ {
		ch <- fmt.Sprintf("data%d", i)
	}
	close(ch)
}

func second(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= secondTask; i++ {
		ch <- fmt.Sprintf("data%d", i)
	}
	close(ch)
}

func as4() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(2)
	go first(ch1, &wg)
	go second(ch2, &wg)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from channel 1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from channel 2: %s\n", msg2)
		}
	}
	wg.Wait()
}
