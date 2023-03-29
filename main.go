package main

import (
	"log"
	"os"
)

func main() {
	// 標準出力させる
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	// 出力とともにプログラムを終了する
	//log.Fatal("Log\n")
	//log.Fatalln("Log2")
	//log.Fatalf("Log%d\n", 3)

	// 出力とともにプログラムを終了する
	//log.Panic("Log\n")
	//log.Panicln("Log2")
	//log.Panicf("Log%d\n", 3)

	// ファイルに出力する
	//f, err := os.Create("test.log")
	//if err != nil {
	//	return
	//}
	//log.SetOutput(f)
	//log.Print("Log\n")

	log.SetFlags(log.LstdFlags | log.Ltime | log.Lmicroseconds)
	log.Print("Log\n")
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetPrefix("[LOG] ")
	log.Print("Log\n")

	logger := log.New(os.Stdout, "[LOG] ", log.Ldate|log.Ltime|log.Lshortfile)
	_, err := os.Open("test1.log")
	if err != nil {
		logger.Fatalln(err)
	}
}
