package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var i int = 1
	//fl64 := float64(i)
	//fmt.Printf("%T\n", i)
	//fmt.Printf("%T\n", fl64)
	//
	//i2 := int(fl64)
	//fmt.Printf("%T\n", i2)
	//
	//fl := 10.5
	//fmt.Println(int(fl)) // 10 切り捨てになる

	var s string = "hello"
	fmt.Printf("%T\n", s)

	//i, _ := strconv.Atoi(s)
	i, err := strconv.Atoi("s")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	var h string = "hello"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
