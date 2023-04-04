package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}
func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello(&wg) // ゴルーチンの起動処理中にmainゴルーチンが終了するので待機をmainに追加

	wg.Add(1)
	// フォークジョイン、ゴルーチンは同期処理を保証しない
	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()
	//time.Sleep(2 * time.Second)

	wg.Wait() // Addカウンターが0になるまで待機
}
