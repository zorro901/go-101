package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) // [3]int 配列の数で型が変わる

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2) // [3]string

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3) // [3]int

	arr4 := [...]string{"a", "b"} // 要素の数を自動でカウントする
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4) // [2]string

	fmt.Println(arr1[0])
	fmt.Println(arr2)
	fmt.Println(arr2[2-1])

	arr2[2] = "c"
	fmt.Println(arr2)

	//var arr5 [4]int
	//arr5 = arr1

	fmt.Println(len(arr1))
}
