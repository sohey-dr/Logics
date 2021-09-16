package main

import (
	"fmt"
	"github.com/sohey-dr/Logics/go/start_db/csv"
	"github.com/sohey-dr/Logics/go/start_db/proxy"
	"github.com/sohey-dr/Logics/go/start_db/scraper"
	"log"
	"strconv"
	"time"
)

func main() {
	searchTelNumber()
}

// 企業詳細ページ取得までの一連の処理
func outputCompanyUrlsCSV() {
	categoryUrls, _ := scraper.GetCategoryUrls()

	var companyUrls []map[string]string
	for i, categoryUrl := range categoryUrls {
		time.Sleep(time.Second * 4)

		//NOTE: プロキシサーバーが2つの場合で対応しているため2の余りを渡している
		proxy.SetProxy(i % 2)

		urls, _ := scraper.GetComUrlByList(categoryUrl["url"])
		if urls == nil {
			urls, _ = scraper.GetComUrlByList(categoryUrl["url"])
			fmt.Println(categoryUrl["categoryName"])
		}
		for _, url := range urls {
			var companyUrlInList = map[string]string{"url": url, "categoryName": categoryUrl["categoryName"]}

			companyUrls = append(companyUrls, companyUrlInList)
		}
	}

	var records [][]string
	for i, companyUrl := range companyUrls {
		record := []string{strconv.Itoa(i), companyUrl["url"], companyUrl["categoryName"], "2021-09-10 14:38:54", "2021-09-10 14:38:54"}
		records = append(records, record)
	}

	csv.WriteCompanyUrls(records)
}

// CSVから企業詳細ページのリンクを取得しそのページから情報を取得
func outputCompanyInfo() {
	companyUrls := csv.ReadCompanyUrls()
	for i, companyUrl := range companyUrls {
		time.Sleep(time.Second * 4)

		//NOTE: プロキシサーバーが2つの場合で対応しているため2の余りを渡している
		proxy.SetProxy(i % 2)

		companyInfoMap := scraper.GetCompanyInfo(companyUrl)

		// NOTE: 順番を揃えるためにスライスに書き換えている
		companyInfoSlice := []string{
			companyInfoMap["企業名"],
			companyInfoMap["設立年数"],
			companyInfoMap["資金調達額"],
			companyInfoMap["事業内容"],
			companyInfoMap["住所"],
			companyInfoMap["従業員数"],
			companyInfoMap["会社HP"],
			companyInfoMap["代表名"],
			companyInfoMap["最新のプレスリリース日"],
			companyInfoMap["最新の資金調達日"],
			companyUrlsAndName["categoryName"],
		}
		csv.WriteCompanyInfos(companyInfoSlice)
		log.Println(companyUrl)
	}
}

// CSVから企業名、住所を取得してそこから電話番号を探して出た場合は企業名と電話番号を返す
func searchTelNumber() {
	csv.ReadCompanyInfos()
}
