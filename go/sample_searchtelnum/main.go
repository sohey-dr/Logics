package main

import "github.com/sohey-dr/searchtelnum"

func main() {
	telNum, err := searchtelnum.Run("株式会社ODYSSEY", "〒190-0182")
	if err != nil {
		panic(err)
	}
	println(telNum)
}