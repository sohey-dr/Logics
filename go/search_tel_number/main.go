package main

import "search_tel_number/scraper"

func main() {
	address := "〒600-8118 京都府京都市下京区平居町５８番地 本池中 UNKNOWN"

	scraper.SearchTelNumber("https://www.google.com/search?q=株式会社CredoShip.", address)
}
