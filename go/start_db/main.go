package main
 
import (
	"start_db/scraper"
)

func main() {
	urls, _ := scraper.GetCatgoryUrls()
	for _, url := range urls {
		scraper.GetComUrlByList(url)	
	}
}