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
	outputCompanyInfo()
}

// 企業詳細ページ取得までの一連の処理
func outputCompanyUrlsCSV() {
	categoryUrls, _ := scraper.GetCategoryUrls()

	var companyUrls []string
	for _, categoryUrl := range categoryUrls {
		time.Sleep(time.Second * 7)

		fmt.Println("再開します")
		companyUrlsInList, _ := scraper.GetComUrlByList(categoryUrl)

		companyUrls = append(companyUrls, companyUrlsInList...)
	}

	var records [][]string
	for i, companyUrl := range companyUrls {
		record := []string{strconv.Itoa(i), companyUrl, "2021-09-10 14:38:54", "2021-09-10 14:38:54"}
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
		}
		csv.WriteCompanyInfos(companyInfoSlice)
		log.Println(companyUrl)
	}
}

func searchTelNumber() {
	//	TODO:プライベートパッケージを読み取って動かす

}
