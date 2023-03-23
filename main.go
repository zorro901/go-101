package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("Hello from ReturnFunc()")
	}
}

func main() {
	f := ReturnFunc()
	f()
}
