package test

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

// TestSelect ...
func TestSelect(t *testing.T) {
	c := make(chan int, 1)

	go func(c chan int) {
		c <- 1
		c <- 2
	}(c)

	for {
		select {
		case ch := <-c:
			fmt.Println(ch)
			if ch == 2 {
				return
			}
		}
	}
}

func asyncCall(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	return c
}

func asyncCall2(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	// gc or some other reason cost some time
	time.Sleep(200 * time.Millisecond)
	return c
}

// TestSelect2 ...
func TestSelect2(t *testing.T) {
	select {
	case resp := <-asyncCall(50):
		println(resp)
	case resp := <-asyncCall(200):
		println(resp)
	case resp := <-asyncCall2(3000):
		println(resp)
	}
}

// TestSelect3 ...
func TestSelect3(t *testing.T) {
	select {
	case resp := <-asyncCall(50):
		println(resp)
	case resp := <-asyncCall(10):
		println(resp)
	case <-time.After(100 * time.Millisecond):
		fmt.Fprint(os.Stderr, "time out")
	}
}

// TestSelect4 ...
func TestSelect4(t *testing.T) {
	c := make(chan int, 1)
	go func(chan int) {
		time.Sleep(200 * time.Millisecond)
		c <- 1
	}(c)
	select {
	case <-c:
		return
	case <-time.After(100 * time.Millisecond):
		fmt.Fprint(os.Stderr, "time out")
	}
}

func sender(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
func receiver(ch chan int, wg *sync.WaitGroup) {
	for {
		if val, ok := <-ch; ok {
			fmt.Println(fmt.Sprintf("%d,%s", val, "revevier"))
		} else {
			fmt.Println("quit recevier")
			break
		}
	}
	wg.Done()
}
func receiver2(ch chan int, wg *sync.WaitGroup) {
	for {
		if val, ok := <-ch; ok {
			fmt.Println(fmt.Sprintf("%d,%s", val, "revevier2"))
		} else {
			fmt.Println("quit recevier2")
			break
		}
	}
	wg.Done()
}

// TestSelect5 ...
func TestSelect5(t *testing.T) {
	ch := make(chan int, 0)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sender(ch, wg)
	wg.Add(1)
	go receiver(ch, wg)
	wg.Add(1)
	go receiver2(ch, wg)
	wg.Wait()
}

// TestSelect6 ...
func TestSelect6(t *testing.T) {
	c := make(chan func(), 1)

	go func(chan func()) {
		f := <-c
		f()
	}(c)

	select {
	case c <- func() {
		println("fuck")
	}:
	default:
		fmt.Println("do nothing")
	}
}
