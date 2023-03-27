package main

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1) // チャネルでfor文をするときはcloseするのが基本
	for i := range ch1 {
		println(i)
	}
}
