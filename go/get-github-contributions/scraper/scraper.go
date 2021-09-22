package scraper

import (
	"errors"
	"log"
	"net/http"

	"get-github-contributions/strings"
	"github.com/PuerkitoBio/goquery"
)

// GetContributions contribute数をスクレイピングして返す。
func GetContributions(userName string) (int64, error) {
	url := "https://github.com/" + userName
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	// NOTE: -1はありえないので取得した時にスクレイピングできなかったことをわかるようにしている
	var contributeNum int64 = -1

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find(".position-relative > h2").Each(func(i int, s *goquery.Selection) {
		contributeNum = strings.FindNum(s.Text())
	})

	if contributeNum == -1 {
		return 0, errors.New("failed to retrieve value")
	}

	return contributeNum, nil
}
