package main

import (
	"fmt"
	"net/http"
)

// ヘッダ全体を取得する例
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

// ヘッダから特定の要素を取得する例
func header(w http.ResponseWriter, r *http.Request) {
	//要素「Accept-Encoding」を取得
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}

// ヘッダから特定の要素を、カンマ区切りで取得する例
func headerComma(w http.ResponseWriter, r *http.Request) {
	//要素「Accept-Encoding」を取得
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func main() {
	//サーバを作り、各ハンドラ関数をハンドルする
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/header", header)
	http.HandleFunc("/headerComma", headerComma)
	fmt.Println("Start Server")
	server.ListenAndServe()
}