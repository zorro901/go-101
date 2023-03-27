package main

import (
	"fmt"
	"time"
)

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " END")
}

func main() {
	ch1 := make(chan int, 2)

	//ch1 <- 1 // close後、送信は出来ないが受信はできる
	//close(ch1)
	////fmt.Println(<-ch1)
	//
	//i, ok := <-ch1
	//fmt.Println(i, ok) // closeしてデータがない場合はfalse

	go receiver("1.goroutin", ch1)
	go receiver("2.goroutin", ch1)
	go receiver("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)

}
