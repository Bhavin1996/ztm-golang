package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	counter := 0
	var input int
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "gouritines remaining")
				counter -= 1
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("Waiting for duration", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("Running goroutines")
	wg.Wait()
	fmt.Println("Program exited")
}
