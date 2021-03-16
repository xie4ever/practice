package test

import (
	"fmt"
	"testing"
	"time"
)

// Run ...
func Run(taskId, sleepTime, timeout int, ch chan string) {
	chRun := make(chan string)
	go run(taskId, sleepTime, chRun)
	select {
	case re := <-chRun:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", taskId)
		ch <- re
	}
}
func run(taskId, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskId, sleepTime)
}

// Test ...
func Test(t *testing.T) {
	input := []int{3, 2, 1}
	timeout := 2
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")

	for i, sleepTime := range input {
		chs[i] = make(chan string)
		go Run(i, sleepTime, timeout, chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}
