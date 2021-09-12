package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// SearchTelNumber Googleで「{企業名} 電話番号」で検索して出てきたものが同一の企業と判断できた場合にmapを返す
func SearchTelNumber(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	fmt.Println("探しましょう")
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
	})
}
