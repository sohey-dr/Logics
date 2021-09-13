package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
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

// ReadCompanyUrls companyUrls.csvからターゲットとなるurlを取得してスライスを返す
func ReadCompanyUrls() []string {
	file, err := os.Open("companyUrls.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)

	var urls []string
	// CSVの内容を1行ずつ読み取る
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, record[1])
	}

	return urls
}
