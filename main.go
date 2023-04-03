package main

import "fmt"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type MyInt int

func main() {
	fmt.Println(Max(1, 2))
	fmt.Println(Max(1.1, 1.3))

	var x, y MyInt = 1, 2
	fmt.Println(Max[MyInt](x, y))
}
