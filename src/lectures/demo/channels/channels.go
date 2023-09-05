package main

import (
	"fmt"
	"time"
)

const (
	DoExit = iota
	ExitOk
)

type ControlMsg int

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func cal(jobs <-chan Job, result chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit the goroutine")
				control <- ExitOk
				return
			default:
				panic("Unhandled operations")
			}
		case job := <-jobs:
			result <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	jobs := make(chan Job, 30)
	results := make(chan Result, 50)
	controlMsg := make(chan ControlMsg)
	go cal(jobs, results, controlMsg)
	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timed out")
			controlMsg <- DoExit
			<-controlMsg
			fmt.Println("Prohram exit")
			return
		}
	}

}
