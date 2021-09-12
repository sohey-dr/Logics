package scraper

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SearchTelNumber Googleで「{企業名} 電話番号」で検索して出てきたものが同一の企業と判断できた場合にハイフンありの電話番号を返す
func SearchTelNumber(url, address string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var existsAddress = false
	var telNumber string

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		// NOTE: 郵便番号が一致する文字列がspanで見つかったらexistsAddressがtrueを返す
		// ex. address == "〒600-8118 京都府京都市下京区平居町５８番地 本池中 UNKNOWN"
		if !existsAddress && strings.HasPrefix(text, address[:11]) {
			existsAddress = true
		}

		re := regexp.MustCompile(`(\d{2,4})-(\d{2,4})-(\d{4})`)
		if re.MatchString(text) {
			telNumber = text
		}
	})

	return telNumber
}
