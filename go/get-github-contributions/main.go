package main

import (
	"fmt"
	"get-github-contributions/csv"
	"get-github-contributions/scraper"
)

func main() {
	users := csv.GetGithubId()
	for _, user := range users {
		contributeNum := scraper.GetContributions(user["githubId"])
		fmt.Println(contributeNum)
	}
}
