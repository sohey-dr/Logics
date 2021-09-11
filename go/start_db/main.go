package main

import (
	"fmt"
	"strconv"
	"time"

	"start_db/csv"
	"start_db/scraper"
)

func main() {
	categoryUrls, _ := scraper.GetCatgoryUrls()

	companyUrls := [][]string{{}}

	for i, categoryUrl := range categoryUrls {
		time.Sleep(time.Second * 10)

		fmt.Println("再開します")
		companyUrlsInList, _ := scraper.GetComUrlByList(categoryUrl)

		companyUrls = append(companyUrls, companyUrlsInList...)
	}

	var records [][]string
	for i, companyUrl := range companyUrls {
		record := []string{strconv.Itoa(i), companyUrl, "2021-09-10 14:38:54", "2021-09-10 14:38:54"}
		records = append(records, record)
	}

	csv.Write(records)
}