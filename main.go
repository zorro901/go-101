package main

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		println("Hello, world!")
	})
}
