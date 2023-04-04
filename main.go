package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC() // 使われていないメモリを解放する

	var s runtime.MemStats

	runtime.ReadMemStats(&s) // メモリの状態を取得

	return s.Sys
}

func main() {
	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch // チャネルの受信を永遠に待つ
	}

	const numGoroutines = 1000000

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	// 8.878kb
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000) // 小数点3桁まで表示

}
