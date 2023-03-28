package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x   bool
	)

	// オプション
	flag.IntVar(&max, "n", 32, "最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション") // 入力がなければfalse

	// コマンドラインをパース
	flag.Parse()

	fmt.Println(max, msg, x)
}
