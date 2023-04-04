package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//say := "Hello"
	//
	//wg.Add(1)
	//
	//go func() {
	//	defer wg.Done()
	//	say = "Good bye"
	//}()
	//
	//wg.Wait()
	//
	//fmt.Println(say) // クロージャーの編集が行われる

	tasks := []string{"A", "B", "C"}
	for _, task := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			fmt.Println(task)
		}(task)
	}
	wg.Wait()

}
