package csv

import (
	"encoding/csv"
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
