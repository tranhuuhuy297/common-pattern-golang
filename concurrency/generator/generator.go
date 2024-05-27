package main

import (
	"fmt"
	"sync"
)

func evenGenerator() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; ; i += 2 {
			c <- i
		}
	}()

	return c
}

func oddGenerator() <-chan int {
	c := make(chan int)

	// if this code doesn't use goroutine
	// so channel c doesn't have anyone get value from channel
	// then channel c will be block
	go func() {
		for i := 1; ; i += 2 {
			c <- i
		}
	}()

	return c
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for even := range evenGenerator() {
			fmt.Println("Even: ", even)
		}
	}()

	go func() {
		for odd := range oddGenerator() {
			fmt.Println("Odd: ", odd)
		}
	}()
	wg.Wait()
}
