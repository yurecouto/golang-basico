package main

import (
	"fmt"
	"time"
)

func main() {
	// canal vai receber apenas dados do tipo string
	canal := make(chan string)
	go escrever("Hello World!", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second / 10)
	}

	close(canal)
}
