package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"


	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
	"line_bot_api_search_restaurants/service"
	"line_bot_api_search_restaurants/domain"

)


func LineHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	log.Print(request.Headers)
	log.Print(request.Body)
	fmt.Println()

	bot, err := service.GetLineToken()

	if err != nil {
		// Do something when something bad happened.
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			// サーバー側のエラー
			Body:       fmt.Sprintf(`{"message:":"%s"}`+"\n", http.StatusText(http.StatusInternalServerError)),
			StatusCode: http.StatusInternalServerError,
			}, nil
	}

	if !validateSignature(os.Getenv("SECRET_TOKEN"), request.Headers["X-Line-Signature"], []byte(request.Body)) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", linebot.ErrInvalidSignature.Error()),
		}, nil
	}

	var api domain.Api
	
	if err := json.Unmarshal([]byte(request.Body), &api); err != nil {
		log.Print(err)
		// クライアントのエラーを返す
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", http.StatusText(http.StatusBadRequest)),
		}, nil
	}

	for _, event := range api.Events {
fmt.Println(event)
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
	
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}


func validateSignature(channelSecret string, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	hash := hmac.New(sha256.New, []byte(channelSecret))
	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
