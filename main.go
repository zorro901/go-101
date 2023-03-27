package main

import "fmt"

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (p *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", p.Number, p.Model)
}

func main() {
	vs := []Stringify{
		&Person{Name: "Tom", Age: 18},
		&Car{Number: "123", Model: "BMW"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
