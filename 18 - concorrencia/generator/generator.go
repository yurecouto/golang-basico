package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("texo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second / 2)
		}
	}()

	return canal
}
