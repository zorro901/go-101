package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: User{Name: "Alice", Age: 18},
		2: User{Name: "Bob", Age: 20},
	}
	fmt.Println(m)

	m2 := map[User]string{
		User{Name: "Alice", Age: 18}: "Tokyo",
		User{Name: "Bob", Age: 20}:   "London",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	m3[1] = User{Name: "Alice", Age: 18}
	fmt.Println(m3)

	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
