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

func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan string) //all senders are blocked because of unbuffered channel
	ch2 := make(chan string) //which caused deadlock between all goroutines
	//senders are blocked if a channel is full

	wg.Add(2)
	go first(ch1, &wg)
	go second(ch2, &wg)

	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from channel 1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from channel 2: %s\n", msg2)
		}
	}
}
