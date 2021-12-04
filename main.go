package main

import (
	"fmt"
	"log"
	"net/http"

	"line_bot_api_search_restaurants/controller"
)

var SECRET string
var ACCESS string

func main() {
	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", controller.LineHandler)

	fmt.Println("http://localhost:8080")
	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "top page"
	fmt.Fprintf(w, msg)
}