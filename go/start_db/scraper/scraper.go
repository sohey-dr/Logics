package scraper

import (
	"log"
	"net/http"

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

	return contents
}
