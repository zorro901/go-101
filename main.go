package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	bs1, _ := io.ReadAll(f) // 巨大なデータには適さない
	fmt.Println(string(bs1))

	//bs2, _ := ioutil.ReadFile("test.txt")
	//fmt.Println(string(bs2))

	//if err := ioutil.WriteFile("test.txt", []byte("test"), 0666); err != nil {
	//	log.Fatalln(err)
	//}

	if err := os.WriteFile("test.txt", []byte("test1"), 0666); err != nil {
		log.Fatalln(err)
	}
}
