package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// メソッドのレシーバーはポインタ型が望ましい
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "John", Age: 20}
	user1.SetName2("Jane")
	user1.SayName()

	user2 := &User{Name: "Jane"}
	user2.SetName2("John")
	user2.SayName()
}
