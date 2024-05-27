package main

import (
	"fmt"
)

// num -> *2 -> +1 -> ^2

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()

	return out
}

func addOne(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num + 1
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {
	in := make(chan int)

	double := double(in)
	addOne := addOne(double)
	square := square(addOne)

	go func() {
		for num := 0; num < 10; num++ {
			in <- num
		}
		close(in)
	}()

	for num := range square {
		fmt.Println(num)
	}
}
