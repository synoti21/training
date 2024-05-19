package main

import (
	"fmt"
	"sync"
)

func sayFirstHello(wg *sync.WaitGroup) {
	fmt.Println("Hello from function 1")
	wg.Done()
}

func saySecondHello(wg *sync.WaitGroup) {
	fmt.Println("Hello from function 2")
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayFirstHello(&wg)
		go saySecondHello(&wg)
	}

	wg.Wait()
}
