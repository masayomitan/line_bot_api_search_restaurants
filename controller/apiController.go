package controller

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/line/line-bot-sdk-go/linebot"
	"strconv"
	"line_bot_api_search_restaurants/domain"
	"line_bot_api_search_restaurants/service"
)


func SendRestoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := GetRestoInfo(lat, lng)

	_, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do()
	if err != nil {
		log.Print(err)
	}
}


func GetRestoInfo(lat string, lng string) string {
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

	var data domain.Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var info string
	for _, shop := range data.Results.Shop {
		info += shop.Name + "\n" + shop.Address + "\n\n"
	}
	return info
}