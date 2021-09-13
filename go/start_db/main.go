package main

import (
	"fmt"
	"github.com/sohey-dr/Logics/go/start_db/csv"
	"github.com/sohey-dr/Logics/go/start_db/proxy"
	"github.com/sohey-dr/Logics/go/start_db/scraper"
	"log"
	"strconv"
	"strings"
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

// CSVから企業名、住所を取得してそこから電話番号を探して出た場合は企業名と電話番号を返す
func searchTelNumber() {
	companyInfos := csv.ReadCompanyInfos()
	for _, companyInfo := range companyInfos {
		// サイト内で住所が空の場合は — となっている
		if companyInfo["住所"] == "—" {
			fmt.Printf("社名: %s なし \n", companyInfo["社名"])
			continue
		}

		time.Sleep(time.Second * 4)

		companyNameRemoveSpace := strings.Replace(companyInfo["社名"], " ", "", 1)
		url := "https://www.google.com/search?q=" + companyNameRemoveSpace
		telNum := scraper.SearchTelNumber(url, companyInfo["住所"])

		if telNum != "" {
			fmt.Printf("社名: %s, 電話番号: %s \n", companyInfo["社名"], telNum)
		} else {
			fmt.Printf("社名: %s なし \n", companyInfo["社名"])
		}
	}
}
