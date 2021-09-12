package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SearchTelNumber Googleで「{企業名} 電話番号」で検索して出てきたものが同一の企業と判断できた場合にmapを返す
func SearchTelNumber(url, address string) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("探しましょう")
	existsAddress := false
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		// NOTE: 郵便番号が一致する文字列がspanで見つかったらexistsAddressがtrueを返す
		// ex. address == "〒600-8118 京都府京都市下京区平居町５８番地 本池中 UNKNOWN"
		if strings.HasPrefix(text, address[:11]) {
			existsAddress = true
		}
	})
}
