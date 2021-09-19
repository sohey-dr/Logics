package scraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"get-github-contributions/strings"
)

// GetContributions contribute数をスクレイピングして返す
func GetContributions(userName string) int64 {
	url := "https://github.com/" + userName
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var contributeNum int64

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find(".position-relative > h2").Each(func(i int, s *goquery.Selection) {
		contributeNum = strings.FindNum(s.Text())
	})

	return contributeNum
}
