package main

import (

	"github.com/aws/aws-lambda-go/lambda"
	"line_bot_api_search_restaurants/handler"
)

func main() {
	lambda.Start(handler.LineHandler)
}
