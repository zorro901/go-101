package main

import "fmt"

func main() {
	var ch1 chan int

	//var ch2 <-chan int // ch2 is a receive-only channel
	//var ch3 chan<- int // ch3 is a send-only channel

	ch1 = make(chan int)
	ch2 := make(chan int)
	fmt.Println(cap(ch1), cap(ch2)) // 容量を調べる

	ch3 := make(chan int, 5)
	fmt.Println("ch3 cap is", cap(ch3))

	ch3 <- 1
	fmt.Println("ch3 len is", len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("ch3 len is", len(ch3))

	i := <-ch3
	fmt.Println(i) // 1
	fmt.Println("ch3 len is", len(ch3))

	i2 := <-ch3
	fmt.Println(i2) // 2
	fmt.Println("ch3 len is", len(ch3))

	fmt.Println(<-ch3) // 3
	fmt.Println("ch3 len is", len(ch3))

}
