package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func main() {
	u1 := NewUser("Alice", 20)
	fmt.Println(u1)
	fmt.Println(*u1)
}
