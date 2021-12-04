package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"line_bot_api_search_restaurants/service"

)

func LineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	bot, err := service.GetLineToken()
	if err != nil {
		// Do something when something bad happened.
		log.Fatal(err)
		return
	}

	//リクエスト取得
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
				SendRestoInfo(bot, event)
			}
		}
	}
}
