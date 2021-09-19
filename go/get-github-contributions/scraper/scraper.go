package scraper

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// GetContributions 
func GetContributions(userName string) string {
	url := "https://github.com/" + userName
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var contributeNum string

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// NOTE: 要素が変わることがあるため広いspanタグで指定している
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		contributeNum := s.Text()
		
		contributions, _ = strconv.Atoi(contributeNum)
	})

	return contributeNum
}
