package domain

import (

	// "line_bot_api_search_restaurants/service"
)


// response APIレスポンス
type Response struct {
	Results Results `json:"results"`
}

// results APIレスポンスの内容
type Results struct {
	Shop []Shop `json:"shop"`
}

// shop レストラン一覧
type Shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}