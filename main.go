package main

import "fmt"

type User struct {
	Name string
	Age  int
	//X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	user2.Name = "Jane"
	fmt.Println(user2)

	user3 := User{
		Name: "Bob",
	}
	fmt.Println(user3)

	user4 := User{
		"user4", 40,
	}
	fmt.Println(user4)

	user5 := User{
		Name: "Bob",
	}
	fmt.Println(user5)

	user6 := new(User) // ポインタ型で宣言できる
	fmt.Println(user6)

	user7 := &User{} // ポインタ型でこちらの方が書かれやすい
	fmt.Println(user7)

	UpdateUser(user1)
	UpdateUser2(user7)

	fmt.Println(user1)
	fmt.Println(user7)
}
