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

// SetProxy プロキシサーバーのURLをhttpパッケージにセットする
func SetProxy(i int) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println(err)
	}

	proxyUser := os.Getenv("PROXY_USER")
	proxyPass := os.Getenv("PROXY_PASS")

	Url := `http://` + proxyUser + `:` + proxyPass + `@` + proxies[i]

	proxyUrl, _ := url.Parse(Url)
	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
}
