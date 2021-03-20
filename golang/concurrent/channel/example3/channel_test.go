package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Test ...
func Test(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(ch chan int) {
			time.Sleep(time.Second)
			ch <- i
		}(ch)
	}
	value := <-ch
	fmt.Println(value)
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test2 ...
func Test2(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(ch chan int, i int) {
			time.Sleep(time.Second)
			ch <- i
		}(ch, i)
	}
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test3 ...
func Test3(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(ch chan int, i int) {
			time.Sleep(time.Duration(i) * time.Second)
			ch <- i
		}(ch, i)
	}
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test4 ...
func Test4(t *testing.T) {
	ch := make(chan int)
	cch := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		go func(ch chan int, i int) {
			cch <- true
			ch <- i
		}(ch, i)
		<-cch
	}
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test5 ...
func Test5(t *testing.T) {
	ch := make(chan int)
	cch := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		go func(ch chan int) {
			cch <- true
			ch <- i
		}(ch)
		<-cch
	}
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test6 ...
func Test6(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go func(ch chan int) {
			ch <- i
		}(ch)
	}
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test7 ...
func Test7(t *testing.T) {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(ch chan int, i int) {
			wg.Done()
			ch <- i
		}(ch, i)
		wg.Add(1)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test8 ...
func Test8(t *testing.T) {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch chan int, i int) {
			ch <- i
			wg.Done()
		}(ch, i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

// Test9 ...
func Test9(t *testing.T) {
	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-chA; ok {
				fmt.Println("A")
				chB <- true
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-chB; ok {
				fmt.Println("B")
				chC <- true
			}
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-chC; ok {
				fmt.Println("C")
				chA <- true
			}
		}
	}()

	chA <- true
}

// Test10 ...
func Test10(t *testing.T) {
	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)

	go func() {
		for {
			if ok := <-chA; ok {
				if num == 0 {
					return
				}
				fmt.Println("A")
				wg.Done()
				num--
				chB <- true
			}
		}
	}()

	go func() {
		for {
			if ok := <-chB; ok {
				if num == 0 {
					return
				}
				fmt.Println("B")
				wg.Done()
				num--
				chC <- true
			}
		}
	}()

	go func() {
		for {
			if ok := <-chC; ok {
				if num == 0 {
					return
				}
				fmt.Println("C")
				wg.Done()
				num--
				chA <- true
			}
		}
	}()

	chA <- true
	wg.Wait()
}
