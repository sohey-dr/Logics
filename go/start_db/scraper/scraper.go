package scraper

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetCategoryUrls カテゴリごとの企業一覧ページのリンクを返す
func GetCategoryUrls() ([]string, error) {
	res, err := http.Get("https://startup-db.com/tags")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var urls []string
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".tag-children-ul > li > a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		var url string = `https://startup-db.com` + href

		urls = append(urls, url)
	})

	return urls, nil
}

// GetComUrlByList 企業一覧ページから企業詳細ページのリンクを取得する
func GetComUrlByList(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var nextUrls []string
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".CompanyCard > a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		var nextUrl string = `https://startup-db.com` + href

		nextUrls = append(nextUrls, nextUrl)
	})

	return nextUrls, nil
}

// GetCompanyInfo 企業詳細ページから企業情報を取得
func GetCompanyInfo(url string) map[string]string {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	contents := make(map[string]string)
	doc.Find(".section-overview > dl > dd").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		switch i {
		case 0:
			contents["企業名"] = text
		case 1:
			contents["会社HP"] = text
		case 2:
			contents["設立年数"] = text
		case 3:
			contents["従業員数"] = text
		case 5:
			contents["住所"] = text
		}
	})

	doc.Find(".section-intro > p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		contents["事業内容"] = text
	})

	doc.Find(".funding-dl > dd > .SdbTextAmount > span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		contents["資金調達額"] = text
	})

	return contents
}

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
		if existsAddress && re.MatchString(text) {
			telNumber = text
		}
	})

	return telNumber
}
