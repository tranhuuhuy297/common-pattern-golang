package main

import (
	"fmt"
	"time"
)

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
	// Bubble Sort implementation
	start := time.Now()
	l := len(arr)
	for i := l - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Print("Execute time: ", time.Since(start), " | ")

	return arr
}
