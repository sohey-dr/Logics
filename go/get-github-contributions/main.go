package main

import (
	"fmt"
	"get-github-contributions/csv"
	"get-github-contributions/scraper"
	"log"
	"strconv"
	"time"
)

func main() {
	users := csv.GetGithubId()

	for _, user := range users {
		contributeNum, err := scraper.GetContributions(user["githubId"])
		if err != nil {
			log.Println(err)
			user["contributeNum"] = ""
		} else {
			user["contributeNum"] = strconv.FormatInt(contributeNum, 10)
		}

		time.Sleep(2 * time.Second)
		fmt.Println(user)
	}
}
