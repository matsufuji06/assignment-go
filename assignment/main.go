package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server!")

	// ルートとハンドラ関数を定義
	http.HandleFunc("/api/categories", categoriesHandler)
	http.HandleFunc("/api/calculator", calculatorHandler)

	// 8000番ポートでサーバを開始
	http.ListenAndServe(":8000", nil)
}
