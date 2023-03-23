package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Recovered in f", x)
		}
	}()
	panic("runtime error")
	fmt.Println("Hello, world!")
}
