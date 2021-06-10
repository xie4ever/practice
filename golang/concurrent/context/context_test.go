package context

import (
	"context"
	"fmt"
	"testing"
)

// TestContext ...
func TestContext(t *testing.T) {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("done") // sometime done
		case _, ok := <-ch:
			fmt.Println(ok) // sometime true
		}
	}(ctx)
	cancel()
	ch <- 1
}

// TestContext2 ...
func TestContext2(t *testing.T) {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("done") // done
		case _, ok := <-ch:
			fmt.Println(ok) // true
		}
	}(ctx)
	cancel()
}

// TestContext3 ...
func TestContext3(t *testing.T) {
	ch := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx2 context.Context) {
		for {
			ch <- 1
			select {
			case <-ctx2.Done():
				fmt.Println("done")
				return
			case _, ok := <-ch:
				fmt.Println(ok)
			}
		}
	}(ctx)
	cancel()
}
