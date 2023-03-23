package main

import "fmt"

func main() {
	//Loop:
	//	for {
	//		for {
	//			for {
	//				fmt.Println("start")
	//				break Loop
	//			}
	//		}
	//	}

Loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

}
