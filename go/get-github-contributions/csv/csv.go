package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func GetGithubId() []map[string]string {
	file, err := os.Open("GitHubIdAndUser.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)

	var users []map[string]string
	// CSVの内容を1行ずつ読み取る
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var user = map[string]string{
			"id":       record[0],
			"name":     record[1],
			"githubId": record[2],
		}

		users = append(users, user)
	}

	return users
}

// WriteUserContributes 文字列のスライスからCSVに書き込む
func WriteUserContributes(records []string) {
	file, err := os.OpenFile("userContributes.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
