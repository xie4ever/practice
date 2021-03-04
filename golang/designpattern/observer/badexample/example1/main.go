package main

import (
	"time"

	"golang/designpattern/observer/badexample/example1/article"
)

func main() {
	a := article.NewArticle("title", "content")
	_ = a.Add()
	time.Sleep(time.Second)
	_ = a.Modify()
	time.Sleep(time.Second)
	_ = a.Delete()
}
