package main
 
import (
	"time"
	"start_db/scraper"
)

func main() {
	urls, _ := scraper.GetCatgoryUrls()
	for _, url := range urls {
		time.Sleep(time.Second * 10)
		scraper.GetComUrlByList(url)	
	}
}