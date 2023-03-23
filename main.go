package main

import "fmt"

func main() {
	//n := 1
	//switch n {
	//case 1, 2:
	//	println("1 or 2")
	//case 3, 4:
	//	println("3 or 4")
	//default:
	//	println("default")
	//}

	//switch n := 2; n {
	//case 1, 2:
	//	fmt.Println("1 or 2")
	//case 3, 4:
	//	fmt.Println("3 or 4")
	//default:
	//	fmt.Println("default")
	//}

	//n := 6
	//switch {
	//case n > 0 && n < 4:
	//	fmt.Println("n>0 && n<4")
	//case n > 3 && n < 7:
	//	fmt.Println("n>3 && n<7")
	//default:
	//	fmt.Println("default")
	//}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	//case n>3 && n<6:
	//	fmt.Println("n>3 && n<6")
	default:
		fmt.Println("default")
	}
}
