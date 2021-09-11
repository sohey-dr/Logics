package main

import (
	"fmt"
	"start_db/csv"
	"start_db/scraper"
	"strconv"
	"time"
)

func main() {

}

// 企業詳細ページ取得までの一連の処理
func outputCompanyUrlsCSV() {
	categoryUrls, _ := scraper.GetCategoryUrls()

	var companyUrls []string
	for _, categoryUrl := range categoryUrls {
		time.Sleep(time.Second * 7)

		fmt.Println("再開します")
		companyUrlsInList, _ := scraper.GetComUrlByList(categoryUrl)

		companyUrls = append(companyUrls, companyUrlsInList...)
	}

	var records [][]string
	for i, companyUrl := range companyUrls {
		record := []string{strconv.Itoa(i), companyUrl, "2021-09-10 14:38:54", "2021-09-10 14:38:54"}
		records = append(records, record)
	}

	csv.WriteCompanyUrls(records)
}
