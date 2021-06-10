package close

import (
	"context"
	"fmt"
	"testing"
)

// TestClose ...
func TestClose(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
	}()
	value := <-c
	fmt.Println(value)
}

// TestClose2 ...
func TestClose2(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
		c <- 2
	}()
	value := <-c
	value = <-c
	fmt.Println(value)
}

// TestClose3 ...
func TestClose3(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
		close(c)
	}()
	value, ok := <-c
	fmt.Println(ok)
	fmt.Println(value)
	value, ok = <-c
	fmt.Println(ok)
	fmt.Println(value)
	value, ok = <-c
	fmt.Println(ok)
	fmt.Println(value)
}

// TestClose4 ...
func TestClose4(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
		close(c)
	}()
	fmt.Println(isChanClose(c))
	value, ok := <-c
	fmt.Println(ok)
	fmt.Println(value)
	fmt.Println(isChanClose(c))
}

// TestClose5 ...
func TestClose5(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
		close(c)
	}()
	_, received := <-c
	fmt.Println(received) // true
	value, ok := <-c
	fmt.Println(ok)    // false
	fmt.Println(value) // 0
}

// TestClose6 ...
func TestClose6(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		c <- 1
		close(c)
	}()
	select {
	case value, ok := <-c:
		fmt.Println(value) // 1
		fmt.Println(ok)    // true
	}
	value, ok := <-c
	fmt.Println(ok)    // false
	fmt.Println(value) // 0
}

// TestClose7 ...
func TestClose7(t *testing.T) {
	c := make(chan int, 1)
	close(c)
	select {
	case value, ok := <-c:
		fmt.Println(value) // 1
		fmt.Println(ok)    // true
	}
	value, ok := <-c
	fmt.Println(ok)    // false
	fmt.Println(value) // 0
}

// TestClose8 ...
func TestClose8(t *testing.T) {
	c := make(chan int, 1)
	close(c)
	fmt.Println(isChanClose(c))
}

// TestClose9 ...
func TestClose9(t *testing.T) {
	c := make(chan int, 1)
	fmt.Println(isChanClose(c))
}

// TestClose10 ...
func TestClose10(t *testing.T) {
	c := make(chan int, 1)
	select {
	case _, received := <-c:
		fmt.Println(received)
	default:
		fmt.Println("default")
	}
}

// TestClose11 ...
func TestClose11(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	close(c)
	select {
	case selected, received := <-c:
		fmt.Println(selected)
		fmt.Println(received)
	default:
		fmt.Println("default")
	}
}

// TestClose12 ...
func TestClose12(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	select {
	case _, received := <-c:
		fmt.Println(received) // true
	default:
		fmt.Println("default")
	}
}

// TestClose13 ...
func TestClose13(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(isChanClose(c)) // false
}

// TestClose14 ...
func TestClose14(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	c <- 2                      // block
	fmt.Println(isChanClose(c)) // false
}

// TestClose15 ...
func TestClose15(t *testing.T) {
	c := make(chan int, 5)
	c <- 1
	c <- 2
	for value := range c { // panic
		fmt.Println(value)
	}
}

// TestClose16 ...
func TestClose16(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan int, 1)
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("done")
		case _, ok := <-c:
			fmt.Println(ok) // true
		}
	}(ctx)
	cancel()
	c <- 1
}

// TestClose17 ...
func TestClose17(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	value, ok := <-c
	fmt.Println(value)
	fmt.Println(ok)

	close(c)

	value, ok = <-c
	fmt.Println(value)
	fmt.Println(ok)
}

func isChanClose(ch chan int) bool { // not well enough
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
