package test

import (
	"context"
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

func cal2(ch chan int, second int) {
	time.Sleep(time.Duration(second) * time.Second)
	ch <- 1
}

// TestCal2 ...
func TestCal2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		cancel()
	}()

	ch := make(chan int, 1)
	go cal2(ch, 10)

	select {
	case i, _ := <-ch:
		fmt.Println(i)
	case <-ctx.Done():
		fmt.Println("ctx break")
	}
}

func cal3(ctx context.Context, second int) {
	time.Sleep(time.Duration(second) * time.Second) // 这里一直阻塞了
	select {
	case <-ctx.Done():
		fmt.Println("ctx break")
	default:
		fmt.Println(second)
	}
}

// TestCal3 ...
func TestCal3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go cal3(ctx, i)
	}

	time.Sleep(time.Duration(5) * time.Second)
	cancel()

	time.Sleep(time.Duration(5) * time.Second)
}

// TestCal4 ...
func TestCal4(t *testing.T) {
	go func() {
		go func() {
			time.Sleep(time.Duration(5) * time.Second)
			fmt.Println("son quit")
		}()

		time.Sleep(time.Duration(3) * time.Second)
		fmt.Println("father quit")
	}()

	time.Sleep(time.Duration(10) * time.Second)
}
