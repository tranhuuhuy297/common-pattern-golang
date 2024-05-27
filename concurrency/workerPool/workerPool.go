package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs chan int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Start job [%d]\n", j)
		time.Sleep(1 * time.Second)
		result <- j
		fmt.Printf("End job [%d]\n", j)
	}
}

func main() {
	var wg sync.WaitGroup

	numJobs := 10
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result: ", r)
	}
}
