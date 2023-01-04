package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagemCh1 := <-canal1:
			fmt.Println(mensagemCh1)

		case mensagemCh2 := <-canal2:
			fmt.Println(mensagemCh2)

		}
	}
}
