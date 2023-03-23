package main

import "fmt"

func main() {
	//i := 0
	//for {
	//	i++
	//	if i == 3 {
	//		break
	//	}
	//	println("Hello, world!")
	//}

	//point := 0
	//for point < 10 {
	//	println(point)
	//	point++
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		continue
	//	}
	//	if i == 6 {
	//		break
	//	}
	//	println(i)
	//}

	//arr := [3]int{1, 2, 3}
	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}

	//arr := [3]int{1, 2, 3}
	//for i, v := range arr {
	//	println(i, v)
	//}

	//sl := []string{"a", "b", "c"}
	//for i, v := range sl {
	//	println(i, v)
	//}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
