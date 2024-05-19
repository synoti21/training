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

func as4_fix() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string, firstTask)
	ch2 := make(chan string, secondTask)

	wg.Add(2)
	go first(ch1, &wg)
	go second(ch2, &wg)

	wg.Wait()

	for {
		select {
		case msg1, ok1 := <-ch1:
			if ok1 {
				fmt.Printf("Received from channel 1: %s\n", msg1)
			} else {
				ch1 = nil
			}
		case msg2, ok2 := <-ch2:
			if ok2 {
				fmt.Printf("Received from channel 2: %s\n", msg2)
			} else {
				ch2 = nil
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
