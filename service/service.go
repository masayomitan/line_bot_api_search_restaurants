package service

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"

)

var SECRET string
var ACCESS string

func GetEnvData() (*linebot.Client, error ) {
	err := godotenv.Load((".env"))
	if err != nil {
		panic ("envファイルの読み込みに失敗しました。")
	}
	SECRET = os.Getenv("SECRET_TOKEN")
	ACCESS = os.Getenv("ACCESS_TOKEN")

	bot, err := linebot.New(
		SECRET,
		ACCESS,
	)
	return bot, err
}