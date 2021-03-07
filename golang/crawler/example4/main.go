package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	statusCodeError = errors.New("invalid status code")
)

func main() {
	fileName := "README.md"
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		if err := os.Remove(fileName); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("|address|content|\n")
	file.WriteString("|----|----|\n")

	var monthList []string
	monthMap := make(map[string][]string, 8)
	for i := 2014; i <= 2021; i++ {
		month := strconv.Itoa(i)
		monthList = append(monthList, month)
		monthMap[month] = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	}

	for _, month := range monthList {
		idList, ok := monthMap[month]
		if !ok {
			continue
		}
		for _, id := range idList {
			url := fmt.Sprintf("http://mysql.taobao.org/monthly/%s/%s/", month, id)
			body, err := download(url)
			if errors.Is(err, statusCodeError) {
				continue
			}
			if err != nil {
				log.Fatal(err)
			}

			contentList, err := getContentList(body)
			if err != nil {
				log.Fatal(err)
			}

			for _, content := range contentList {
				text := fmt.Sprintf("%s\n", content)
				fmt.Print(text)
				file.WriteString(text)
			}

			time.Sleep(2 * time.Second)
		}
	}
}

func download(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, statusCodeError
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
		list = append(list, fmt.Sprintf("|%s%s|%s|", "http://mysql.taobao.org/", href, getTitle(selection.Text())))
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
			lastPointIdx = idx
		}
	}

	prefix := string(rs[firstIdx : lastPointIdx+2])
	prefix = strings.Replace(prefix, " ", "", -1)
	suffix := string(rs[lastPointIdx+2 : lastIdx+1])

	return fmt.Sprintf("%s%s", prefix, suffix)
}
