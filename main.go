package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//os.Exit(1)
	//fmt.Println("Hello, world!") // 実行されない

	//defer func() {
	//	fmt.Println("defer") // 実行されない
	//}()
	//os.Exit(0)

	//_, err := os.Open("test.txt")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1])
	//
	//fmt.Printf("length=%d\n", len(os.Args))
	//for i, arg := range os.Args {
	//	fmt.Printf("arg[%d]=%s\n", i, arg)
	//}

	f, err := os.Open("test.txt")
	if err != nil {
		//log.Fatalln(err)
		f, err = os.Create("test.txt")
		f.Write([]byte("Hello, world!"))
		f.WriteAt([]byte("Golang"), 6) // 6文字目から書き込み
		f.Seek(0, io.SeekEnd)          // 書き込み位置を末尾に移動
		f.WriteString("You're welcome!")
	}
	fmt.Println("ファイルを閉じます")
	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
	bs2 := make([]byte, 128)
	nm, err := f.ReadAt(bs2, 10)
	fmt.Println(nm)
	fmt.Println(string(bs2))

	f, err = os.OpenFile("test.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	bs = make([]byte, 128)
	n, err = f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
