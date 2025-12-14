package main

import (
	"encoding/json"
	"net/http"

	lorem "github.com/drhodes/golorem"
)

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンス用のマップを作成
	response := map[string]string{
		"message": lorem.Paragraph(1, 3),
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}
