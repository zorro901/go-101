package main

import "fmt"

func main() {
	a := 1
	if a == 2 {
		fmt.Println("a is 2")
	} else if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 100 {
		fmt.Println("b is 100")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x)

}
