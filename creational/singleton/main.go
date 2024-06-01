package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func() {
			defer wg.Done()
			s := GetInstance()
			s.Increase()
			fmt.Println(s.counter)
		}()
	}
	wg.Wait()
}
