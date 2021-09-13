package main

import (
	"fmt"
	"github.com/sohey-dr/Logics/go/search_tel_number/scraper"
	"strings"
)

func main() {
	address := "〒600-8118 京都府京都市下京区平居町５８番地 本池中 UNKNOWN"
	companyName := "Credo Ship."
	companyNameRemoveSpace := strings.Replace(companyName, " ", "", 1)
	url := "https://www.google.com/search?q=" + companyNameRemoveSpace

	fmt.Println(scraper.SearchTelNumber(url, address))
}
