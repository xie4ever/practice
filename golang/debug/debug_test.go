package debug

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
)

// Go ...
func Go(c context.Context, f func(c context.Context)) {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				info := debug.Stack()
				fmt.Println(string(info), p)
			}
		}()
		f(c)
	}()
}

// TestDebug ...
func TestDebug(t *testing.T) {
	f := func(c context.Context) {
		panic("test panic")
	}
	Go(context.TODO(), f)
}

// TestDebug2 ...
func TestDebug2(t *testing.T) {
	f := func(c context.Context) {
		panic("test panic")
	}
	go f(context.TODO())
}
