package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x) // <nil>

	x = 42
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "hello"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}
