package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var memoryAccess sync.Mutex
	var data int // クリティカルセクション

	wg.Add(1)
	go func() {
		defer wg.Done()
		memoryAccess.Lock() // Lockのコストは小さくない
		data++
		memoryAccess.Unlock()
	}()

	wg.Wait()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(data)
	}
	memoryAccess.Unlock()

}
