package scraper

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	"get-github-contributions/strings"
)

// GetContributions contribute数をスクレイピングして返す
func GetContributions(userName string) int {
	url := "https://github.com/" + userName
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	// NOTE: -1はありえないので取得した時にスクレイピングできなかったことをわかるようにしている
	contributeNum := -1

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find(".position-relative > h2").Each(func(i int, s *goquery.Selection) {
		contributeNum, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Println("scraping err")
		}
	})

	return contributeNum
}
