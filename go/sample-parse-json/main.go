package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get the json file
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read the file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// parse the json file
	var result Request
	json.Unmarshal(byteValue, &result)

	// print the result
	fmt.Println(result)

	// get the value of the key
	fmt.Println(result.Id)

	// get the value of the key
	fmt.Println(result.Jsons)

	jsonStr, _ := json.Marshal(result.Jsons)

	var result2 Request2
	json.Unmarshal(jsonStr, &result2)

	fmt.Println(result2)

}

type Request struct {
	Id 	int    `json:"id"`
	Jsons map[string]string `json:"jsons"`
}

type Request2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
