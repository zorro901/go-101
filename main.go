package main

import "fmt"

var i5 int = 5

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	var i int = 10
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	i4 := 450 // 暗黙的な定義 関数内でのみ使用可能
	i4 = 400
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

}
