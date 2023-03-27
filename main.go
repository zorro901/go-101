package main

import "fmt"

func Double(i int) {
	fmt.Println(i * 2)
}

func DoubleV2(i *int) {
	*i *= 2
}

func DoubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n) // 100
	fmt.Println(&n)
	Double(n)       // 上のnとは別の値を関数内に持つ
	fmt.Println(n)  // 100
	fmt.Println(&n) // ポインタは変わらない

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p) // アドレスが指す実態の値を表示

	*p = 300
	fmt.Println(n)
	n = 200
	fmt.Println(*p)

	// 参照渡しでメモリの値を直接書き換える
	DoubleV2(&n)
	fmt.Println(n)

	// 参照型は元々、参照渡しになる
	var sl []int = []int{1, 2, 3}
	DoubleV3(sl)
	fmt.Println(sl)
}
