package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/categories", categoriesHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
