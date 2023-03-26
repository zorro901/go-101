package main

import "fmt"

func main() {
	//sl := []int{100, 200}
	//sl2 := sl
	//sl2[0] = 1000
	//fmt.Println(sl) // 同じアドレスを指す、参照渡し
	//
	//var i int = 10
	//i2 := i
	//i2 = 100
	//fmt.Println(i, i2) // 基本型は参照渡しにならない

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

}
