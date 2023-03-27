package main

import (
	"fmt"
	. "fmt" // 短縮できるが非推奨
	f "fmt"
	"go-101/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//var b string = s // 返り値で宣言されているので重複宣言になる
	b = s
	{
		b := "BBBB" // ブロック内では再定義可能
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	Println(foo.Max)
	f.Println(foo.ReturnMin())

	fmt.Println(appName())

	fmt.Println(Do("hello"))

}
