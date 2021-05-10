package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// http://127.0.0.1:6060/debug/pprof/
// go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
// go tool pprof http://localhost:6060/debug/pprof/heap
// go tool pprof -http=:6001 profile

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}
