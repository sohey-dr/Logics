package proxy

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

var proxies = []string{"13.114.181.4:80", "35.76.99.107:80"}

// SetProxyUrl プロキシサーバーのURLをhttpパッケージにセットする
func SetProxyUrl(i int) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println(err)
	}

	proxyUser := os.Getenv("PROXY_USER")
	proxyPass := os.Getenv("PROXY_PASS")

	Url := `http://` + proxyUser + `:` + proxyPass + `@` + proxies[i]
	log.Println(Url)

	proxyUrl, _ := url.Parse(Url)
	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
}
