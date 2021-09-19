package main

import(
	"fmt"
	"get-github-contributions/scraper"
)

func main()  {
	fmt.Println(scraper.GetContributions("sohey-dr"))
}