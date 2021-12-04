package domain

import (


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
	Photo   photo  `json:"photo"`
	URLS    urls   `json:"urls"`
}

// photo 写真URL一覧
type photo struct {
	Mobile mobile `json:"mobile"`
}

// mobile モバイル用の写真URL
type mobile struct {
	L string `json:"l"`
}

// urls URL一覧
type urls struct {
	PC string `json:"pc"`
}
