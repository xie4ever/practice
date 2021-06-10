package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// TestConcurrent ...
func TestConcurrent(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	var mouseCount int64
	var catCount int64
	var dogCount int64

	mouseCount = 1
	catCount = 1
	dogCount = 1

	chMouse := make(chan struct{}, 1)
	chCat := make(chan struct{}, 1)
	chDog := make(chan struct{}, 1)

	chMouse <- struct{}{}

	go Mouse(chMouse, chCat, mouseCount, &wg)
	go Cat(chCat, chDog, catCount, &wg)
	go Dog(chDog, chMouse, dogCount, &wg)

	wg.Wait()
}

func Mouse(myCh, theirCh chan struct{}, count int64, wg *sync.WaitGroup) {
	for {
		if count > 100 {
			wg.Done()
			return
		}
		<-myCh
		fmt.Println(fmt.Sprintf("rice: %d", count))
		atomic.AddInt64(&count, 1)
		theirCh <- struct{}{}
	}
}

func Cat(myCh, theirCh chan struct{}, count int64, wg *sync.WaitGroup) {
	for {
		if count > 100 {
			wg.Done()
			return
		}
		<-myCh
		fmt.Println(fmt.Sprintf("fish: %d", count))
		atomic.AddInt64(&count, 1)
		theirCh <- struct{}{}
	}
}

func Dog(myCh, theirCh chan struct{}, count int64, wg *sync.WaitGroup) {
	for {
		if count > 100 {
			wg.Done()
			return
		}
		<-myCh
		fmt.Println(fmt.Sprintf("meat: %d", count))
		atomic.AddInt64(&count, 1)
		theirCh <- struct{}{}
	}
}
