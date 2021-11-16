package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	url := "https://slack.com/api/conversations.list?types=public_channel,private_channel"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Bearer TOKEN")
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
        Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}