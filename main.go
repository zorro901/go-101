package main

/*
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
*/

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("開始")
	time.Sleep(2 * time.Second)
	fmt.Println("終了")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)                                // チャネルを作成
	ctx := context.Background()                            // コンテキストを作成
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // 1秒間のタイムアウトを設定

	defer cancel()

	go longProcess(ctx, ch) // 処理に２秒かかる

	//cancel()

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("##########Error###########")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}

	}

	fmt.Println("ループ抜けた")
}
