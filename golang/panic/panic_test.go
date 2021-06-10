package panic

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"
)

// TestPanic ...
func TestPanic(t *testing.T) {
	go func() {
		panic("fuck")
	}()
	defer func() { // useless
		if p := recover(); p != nil {
			info := debug.Stack()
			fmt.Println(string(info), p)
		}
	}()

	time.Sleep(time.Minute)
}

// TestPanic2 ...
func TestPanic2(t *testing.T) {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				info := debug.Stack()
				fmt.Println(string(info), p)
			}
		}()
		panic("fuck")
	}()

	time.Sleep(time.Minute)
}

func inner() {
	panic("inner func fuck")
}

// TestPanic3 ...
func TestPanic3(t *testing.T) {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				info := debug.Stack()
				fmt.Println(string(info), p)
			}
		}()
		inner()
	}()

	time.Sleep(time.Minute)
}
