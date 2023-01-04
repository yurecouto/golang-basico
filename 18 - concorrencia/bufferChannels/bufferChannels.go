package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Hello World"

	mensagem := <-canal
	fmt.Println(mensagem)
}
