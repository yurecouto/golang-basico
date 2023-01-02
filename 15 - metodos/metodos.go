package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do metodo")
}

func main() {
	usuario1 := usuario{"Yure", 22}

	usuario1.salvar()
}
