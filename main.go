package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello Goroutine %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("1st Goroutine Start")
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("1st Goroutine Done")
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("2st Goroutine Start")
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("2st Goroutine Done")
	//}()
	//
	//wg.Wait()

	var CPU int = runtime.NumCPU()

	wg.Add(CPU)
	for i := 0; i < CPU; i++ {
		go Hello(&wg, i)
	}

	wg.Wait()
}
