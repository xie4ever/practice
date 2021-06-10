package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func cal(ch chan int, wg *sync.WaitGroup, second int) {
	defer wg.Done()
	time.Sleep(time.Duration(second) * time.Second)
	ch <- 1
}

// TestCal ...
func TestCal(t *testing.T) {
	start := time.Now().Second()
	wg := sync.WaitGroup{}
	var chArrs []chan int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch := make(chan int, 1)
		chArrs = append(chArrs, ch)
		go cal(ch, &wg, i)
	}
	wg.Wait()
	var answer int
	for _, chArr := range chArrs {
		answer += <-chArr
	}
	fmt.Println(answer)
	fmt.Println(time.Now().Second() - start)
}
