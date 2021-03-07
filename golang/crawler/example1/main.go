package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackdanger/collectlinks"
)

func main() {
	url := "http://mysql.taobao.org/monthly/2014/08/"
	if err := download(url); err != nil {
		log.Fatal(err)
	}
}

func download(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// 函数结束后关闭相关链接
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		fmt.Println(link)
	}
	return nil
}
