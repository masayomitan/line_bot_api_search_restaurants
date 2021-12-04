package main

import (
	"net/http"
	"context"
	"log"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/lambda"
	"line_bot_api_search_restaurants/handler"
)


func main() {
    lambda.Start(helloHandler)
}

func helloHandler(ctx context.Context) {
	// http.HandleFunc("/", handler.LineHandler)
	handler := http.HandlerFunc(handler.LineHandler)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
	lc, _ := lambdacontext.FromContext(ctx)
  log.Print(lc.Identity.CognitoIdentityPoolID)
}