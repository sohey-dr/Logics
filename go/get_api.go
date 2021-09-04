package main

import (
	"fmt"
	"log"
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
	fmt.Fprintln(w, h[0])
}

// ヘッダから特定の要素を、カンマ区切りで取得する例
func headerComma(w http.ResponseWriter, r *http.Request) {
	//要素「Accept-Encoding」を取得
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/header", header)
	http.HandleFunc("/headerComma", headerComma)
	log.Fatal(http.ListenAndServe(":8080", nil))
}