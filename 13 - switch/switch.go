package main

import "fmt"

func main() {
	fmt.Println("Switch")

	var teste string = "blablabla"

	switch teste {
	case "teste":
		fmt.Println("igual teste")
	case "abacate":
		fmt.Println("igual abacate")
	case "malabare":
		fmt.Println("igual malabare")
	default:
		fmt.Println("palavra nenhuma")
	}
}
