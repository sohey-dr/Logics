package main

import (
	"fmt"
	"search_tel_number/scraper"
	"strings"
)

func main() {
	address := "〒600-8118 京都府京都市下京区平居町５８番地 本池中 UNKNOWN"
	companyName := "Credo Ship."
	companyNameRemoveSpase := strings.Replace(companyName, " ", "", 1)
	url := "https://www.google.com/search?q=" + companyNameRemoveSpase

	fmt.Println(scraper.SearchTelNumber(url, address))
}
