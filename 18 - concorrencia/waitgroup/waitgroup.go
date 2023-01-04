package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Primeira chamada")
		waitGroup.Done()
	}()

	go func() {
		escrever("SEGUNDA chamada")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second / 10)
	}
}
