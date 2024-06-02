package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	context := &Context{}

	bubbleSort := &BubbleSort{}
	context.SetSorter(bubbleSort)
	fmt.Println("Bubble Sort:", context.ExecuteSort(arr))

	quickSort := &QuickSort{}
	context.SetSorter(quickSort)
	fmt.Println("Quick Sort:", context.ExecuteSort(arr))
}
