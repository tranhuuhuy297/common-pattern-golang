package main

import (
	"fmt"
	"time"
)

type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) []int {
	// Quick Sort implementation
	start := time.Now()
	pivot := arr[0]
	smaller := []int{}
	middle := []int{}
	bigger := []int{}

	for _, num := range arr {
		if num < pivot {
			smaller = append(smaller, num)
		} else if num == pivot {
			middle = append(middle, num)
		} else {
			bigger = append(bigger, num)
		}
	}

	result := []int{}
	result = append(result, smaller...)
	result = append(result, middle...)
	result = append(result, bigger...)

	fmt.Print("Execute time: ", time.Since(start), " | ")

	return result
}
