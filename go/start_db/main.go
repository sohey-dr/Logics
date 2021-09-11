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
		scraper.GetComUrlByList(url)	
	}
}