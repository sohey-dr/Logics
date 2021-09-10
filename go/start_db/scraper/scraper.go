package scraper

import (
	"log"
	"net/http"
 
	"github.com/PuerkitoBio/goquery"
)

// カテゴリごとの企業一覧ページのリンクを返す
func GetCatgoryUrls() ([]string, error) {
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

// 企業一覧ページから企業詳細ページのリンクを取得する
func GetComUrlByList(url string) ([]string, error){
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