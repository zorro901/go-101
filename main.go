package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	// 2の平方根
	fmt.Println(math.Sqrt2)

	// float32の最大値 float64版もある
	fmt.Println(math.MaxFloat32)
	// float32で表現可能な0ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	// int64
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	// 累乗
	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(2, 2))

	// 平方根、立方根
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Cbrt(8))

	// 最大値、最小値
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	// 数値の正負を問わず、小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Floor(-3.14))

	// 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Ceil(-3.14))

}
