package main

import "fmt"

func main() {
	fmt.Println("arrays internos")

	slice := make([]float32, 10, 20)
	fmt.Print(slice, len(slice), cap(slice))
}
