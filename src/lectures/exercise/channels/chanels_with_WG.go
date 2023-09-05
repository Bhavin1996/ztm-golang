//package channels
/*package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longCalculationAlt(jobID int) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", jobID, duration)
	return jobID * 30
}

func submitJobs(numJobs int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numJobs; i++ {
		result <- longCalculationAlt(i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var numJobs int
	fmt.Print("Enter the number of jobs: ")
	fmt.Scan(&numJobs)

	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start a goroutine to submit jobs
	wg.Add(1)
	go submitJobs(numJobs, results, &wg)

	// Wait for all jobs to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and calculate the total sum
	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println("Total sum is", sum)
}*/
