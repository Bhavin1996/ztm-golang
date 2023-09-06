package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Hits struct {
	count int
	sync.Mutex
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func serve(hits *Hits, wg *sync.WaitGroup, counter int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Println("Served iteratio number", counter)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	hitcounter := Hits{}
	for i := 0; i < 2000; i++ {
		iteration := i
		wg.Add(1)
		go serve(&hitcounter, &wg, iteration)
	}
	fmt.Println("waiting for goroutines")
	wg.Wait()
	hitcounter.Lock()
	totalCounter := hitcounter.count
	hitcounter.Unlock()
	fmt.Println("Total counter is", totalCounter)
}
