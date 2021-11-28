package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
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
	msg := "hellog World!!!!"
	fmt.Fprintf(w, msg)
}


func lineHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load((".env"))
    if err != nil {
			panic ("envファイルの読み込みに失敗しました。")
    }
		
	SECRET = os.Getenv("SECRET")
	ACCESS = os.Getenv("ACCESS")

	bot, err := linebot.New(
		SECRET,
		ACCESS,
	)
	if err != nil {
		log.Fatal(err)
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
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// メッセージがテキスト形式の場合
				case *linebot.TextMessage:
					replyMessage := message.Text
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}
				// 他にもスタンプや画像、位置情報など色々受信可能
			}
		}
}
