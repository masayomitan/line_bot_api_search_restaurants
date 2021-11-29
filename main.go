package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/line/line-bot-sdk-go/linebot"
	"strconv"
	"encoding/json"

	"line_bot_api_search_restaurants/service"
)

var SECRET string
var ACCESS string

func main() {
	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "hello World!!!!"
	fmt.Fprintf(w, msg)
}


func lineHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := service.GetLineToken()
	if err != nil {
		log.Fatal(err)
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				sendRestoInfo(bot, event)
			}
		}
	}
}

func sendRestoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := getRestoInfo(lat, lng)

	_, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do()
	if err != nil {
		log.Print(err)
	}
}
// response APIレスポンス
type response struct {
	Results results `json:"results"`
}

// results APIレスポンスの内容
type results struct {
	Shop []shop `json:"shop"`
}

// shop レストラン一覧
type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func getRestoInfo(lat string, lng string) string {


	apikey :=  service.GetHotpepperToken()
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s",
		apikey, lat, lng)

	// リクエストしてボディを取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var info string
	for _, shop := range data.Results.Shop {
		info += shop.Name + "\n" + shop.Address + "\n\n"
	}
	return info
}