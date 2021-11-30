package service

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"

)

var SECRET string
var ACCESS string
var HOTPEPPER string

func GetLineToken() (*linebot.Client, error ) {
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

func GetHotpepperToken() (HOTPEPPER string){
	err := godotenv.Load((".env"))
	if err != nil {
		panic ("envファイルの読み込みに失敗しました。")
	}
  HOTPEPPER = os.Getenv("HOTPEPPER_TOKEN")
	return HOTPEPPER
}
