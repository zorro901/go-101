package main

import "fmt"

const Pi = 3.14

const (
	URL      = "https://www.google.com"
	SiteName = "Google"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 9223372036854775807 + 1 // overflow
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}
