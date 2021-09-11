package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func Write(records [][]string) {
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
