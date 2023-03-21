package main

import "fmt"

func main() {
	var fl64 float64 = 3.14
	fmt.Println(fl64)

	fl := 3.14 // 自動だとfloat64型になる
	fmt.Printf("%T\n", fl)

	var fl32 float32 = 3.14 // あまり使わない
	fmt.Printf("%T\n", fl32)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
