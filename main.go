package main

import "fmt"

func anything(a interface{}) {
	//fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Println("unknown", v)
	}
}

func main() {
	anything(42)
	anything("hello")

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt)
	fmt.Println(i + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Println("x is int", i)
	} else if s, isString := x.(string); isString {
		fmt.Println("x is string", s)
	} else {
		fmt.Println("I don't know what x is")
	}

	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Println("I don't know what x is")
	}

	switch v := x.(type) {
	case int:
		fmt.Println("x is int", v)
	case string:
		fmt.Println("x is string", v)
	default:
		fmt.Println("I don't know what x is")
	}

}
