package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを解析する
	query := r.URL.Query()
	operator := query.Get("o")
	x := query.Get("x")
	y := query.Get("y")

	response := map[string]int{
		"result": calc(operator, x, y),
	}

	// Content-Typeヘッダーをapplication/jsonに設定
	w.Header().Set("Content-Type", "application/json")

	// マップをJSONにエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(response)
}

// operatorに応じて計算処理を行う関数
func calc(operator string, x string, y string) int {
	// string→intに変換
	xi, err := strconv.Atoi(x)
	if err != nil {
		// エラーの場合は0をリターン
		return 0
	}

	// string→intに変換
	yi, err := strconv.Atoi(y)
	if err != nil {
		// エラーの場合は0をリターン
		return 0
	}

	switch operator {
	case "add":
		return xi + yi
	case "sub":
		return xi - yi
	case "mul":
		return xi * yi
	case "div":
		if yi == 0 {
			return 0
		}
		return xi / yi
	default:
		return 0
	}
}
