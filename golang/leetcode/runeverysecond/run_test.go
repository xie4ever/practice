package main

import (
	"fmt"
	"testing"
	"time"
)

// Test ...
func Test(t *testing.T) {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		var (
			i    int
			flag bool
		)

		for !flag {
			ticker := time.NewTicker(time.Second)
			select {
			case <-ticker.C:
				func(flag bool) {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()

					if i >= 10 {
						flag = true
						return
					}

					i++
					fmt.Println(i)
					proc()
				}(flag)
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
