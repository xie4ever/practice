package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://mysql.taobao.org/monthly/2014/08/"
	body, err := download(url)
	if err != nil {
		log.Fatal(err)
	}

	contentList, err := getContentList(body)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range contentList {
		fmt.Println(content)
	}
}

func download(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("invalid status code")
	}

	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return dom, nil
}

func getContentList(dom *goquery.Document) ([]string, error) {
	var list []string
	dom.Find("a:contains(·)").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		list = append(list, fmt.Sprintf("%s%s %s", "http://mysql.taobao.org/", href, getTitle(selection.Text())))
	})

	return list, nil
}

func getTitle(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	rs := []rune(str)

	var (
		firstIdx     int
		lastIdx      int
		count        int
		lastPointIdx int
	)

	for idx, r := range rs {
		if firstIdx == 0 && r != ' ' {
			firstIdx = idx
		}

		if r != ' ' {
			lastIdx = idx
		}

		if r == '·' {
			count++
			if count == 2 {
				lastPointIdx = idx
			}
		}
	}

	prefix := string(rs[firstIdx : lastPointIdx+2])
	prefix = strings.Replace(prefix, " ", "", -1)
	suffix := string(rs[lastPointIdx+2 : lastIdx+1])

	return fmt.Sprintf("%s%s", prefix, suffix)
}
