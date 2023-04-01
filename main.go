package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func main() {
	h := md5.New()
	//io.WriteString(h, "ABCDE")
	if _, err := io.WriteString(h, "ABCDE"); err != nil {
		log.Fatal(err)
	}
	//b := []byte{12, 34, 55, 3}
	fmt.Println(h.Sum(nil))
	//fmt.Println(h.Sum(b))

	//fmt.Println(b)

	// %x で 16進数の文字列を生成する
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

}
