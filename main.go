package main

import "fmt"

func main() {
	var i int = 100
	var i2 int64 = 200
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))
}
