package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User // こちらの方が望ましい
//type Users struct {
//	Users []*User
//}

func main() {
	user1 := User{
		Name: "John",
		Age:  30,
	}
	user2 := User{
		Name: "Jane",
		Age:  20,
	}
	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	for _, user := range users {
		fmt.Println(*user)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)
	for _, user := range users2 {
		fmt.Println(*user)
	}

}
