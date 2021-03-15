package main

import (
	"fmt"
	"time"
)

func main() {
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
				func() {
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
				}()
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
