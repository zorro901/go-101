package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(256) // global scopeのため非推奨
	r := rand.New(rand.NewSource(256))
	// 乱数生成
	fmt.Println(r.Float64())
	fmt.Println(r.Float64())
	fmt.Println(r.Float64())

	// 現在の時刻をシードに使った疑似乱数の生成
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(100)) // 0〜99の範囲の乱数を生成
	fmt.Println(r.Int())
	fmt.Println(r.Int31())
	fmt.Println(r.Uint32())

}
