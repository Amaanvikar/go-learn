package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	mych := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receive only channel
	go func(mych <-chan int, wg *sync.WaitGroup) {
		fmt.Println("One")
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	// send only channel
	go func(mych chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Two")
		mych <- 5
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
