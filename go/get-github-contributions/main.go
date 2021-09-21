package main

import (
	"fmt"
	"get-github-contributions/csv"
	"get-github-contributions/scraper"
	"log"
	"strconv"
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

		fmt.Println(user)
	}
}
