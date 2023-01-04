package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("teste"), escrever("oi"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(chan1, chan2 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-chan1:
				outChannel <- mensagem
			case mensagem := <-chan2:
				outChannel <- mensagem
			}
		}
	}()

	return outChannel
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
