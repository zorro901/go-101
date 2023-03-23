package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
}

func main() {
	TestDefer()

	//defer func() {
	//	fmt.Println("Defer 1")
	//	fmt.Println("Defer 2")
	//	fmt.Println("Defer 3")
	//}()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("test")
}
