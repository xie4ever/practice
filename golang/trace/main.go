package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		return
	}

	trace.Start(f)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "test"
	}()

	<-ch
}
