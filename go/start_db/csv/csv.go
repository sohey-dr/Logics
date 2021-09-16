package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/sohey-dr/Logics/go/start_db/proxy"
	"github.com/sohey-dr/Logics/go/start_db/scraper"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// WriteCompanyUrls 二次元配列から企業詳細ページのリンクをCSVに書き込む
func WriteCompanyUrls(records [][]string) {
	file, err := os.Create("companyUrls.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// WriteCompanyInfos 二次元配列から企業詳細ページのリンクをCSVに書き込む
func WriteCompanyInfos(records []string) {
	file, err := os.OpenFile("CompanyInfos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// WriteCompanyInfoAndTelNum 二次元配列から企業詳細ページのリンクをCSVに書き込む
func WriteCompanyInfoAndTelNum(record []string) {
	file, err := os.OpenFile("CompanyInfoAndTelNum.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write(record)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// ReadCompanyUrls companyUrls.csvからターゲットとなるurlを取得してマップの配列を返す
func ReadCompanyUrls() []map[string]string {
	file, err := os.Open("companyUrls.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)

	var urlAndCategoryNames []map[string]string
	// CSVの内容を1行ずつ読み取る
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var urlAndCategoryName = map[string]string{"url": record[1], "categoryName": record[2]}
		urlAndCategoryNames = append(urlAndCategoryNames, urlAndCategoryName)
	}

	return urlAndCategoryNames
}

// ReadCompanyInfos duplicateDeletedCompanyInfos.csvからターゲットとなる社名と住所を取得してmapを返す
func ReadCompanyInfos() {
	file, err := os.Open("duplicateDeletedCompanyInfos.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)

	i := 0
	// CSVの内容を1行ずつ読み取る
	for {
		companyInfo, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// サイト内で住所が空の場合は — となっている
		if companyInfo[4] == "—" {
			fmt.Printf("社名: %s なし \n", companyInfo[0])
			companyInfo[7] = ""
			WriteCompanyInfoAndTelNum(companyInfo)
			continue
		}

		if len(companyInfo[7]) != 0 {
			fmt.Printf("社名: %s, 電話番号: %s \n", companyInfo[0], companyInfo[7])
			WriteCompanyInfoAndTelNum(companyInfo)
			continue
		}

		time.Sleep(time.Second * 2)
		proxy.SetProxy(i % 2)

		companyNameRemoveSpace := strings.Replace(companyInfo[0], " ", "", 1)
		url := "https://www.google.com/search?q=" + companyNameRemoveSpace
		telNum := scraper.SearchTelNumber(url, companyInfo[4])

		if telNum != "" {
			fmt.Printf("社名: %s, 電話番号: %s \n", companyInfo[0], telNum)
			companyInfo[7] = telNum
			WriteCompanyInfoAndTelNum(companyInfo)
		} else {
			fmt.Printf("社名: %s なし \n", companyInfo[0])
			companyInfo[7] = ""
			WriteCompanyInfoAndTelNum(companyInfo)
		}
		i++

	}

}
