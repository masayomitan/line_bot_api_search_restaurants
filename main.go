package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"line_bot_api_search_restaurants/service"
	"line_bot_api_search_restaurants/controller"
)

var SECRET string
var ACCESS string

func main() {
	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", controller.LineHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "hello World!!!!"
	fmt.Fprintf(w, msg)
}