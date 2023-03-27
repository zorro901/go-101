package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "James Bond"
}

func main() {
	t := T{
		User: User{
			Name: "John",
			Age:  20,
		},
	}
	fmt.Println(t.Name) // T構造体で省略されている場合は直接参照できる
	t.SetName()
	fmt.Println(t.Name)
}
