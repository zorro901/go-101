package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
	go sub()
	for {
		fmt.Println("Main loop")
		time.Sleep(200 * time.Millisecond)
	}
}
