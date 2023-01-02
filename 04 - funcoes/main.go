package main

import (
	"fmt"
)

// declarar tipos e retorno
func somar(x, y int32) int32 {
	return x + y
}

// funcao com dois retornos
func somaEsubtrai(x, y int32) (int32, int32) {
	sum := x + y
	sub := x - y

	return sum, sub
}

// funcao com dois retornos (com eles nomeados)
func somaEsubtraiNomeada(x, y int32) (sum, sub int32) {
	sum = x + y
	sub = x - y

	return
}

// funcao com multiplos parametros
func somaVariatica(numeros ...int32) (sum, sub int32) {
	for _, numero := range numeros {
		sum += numero
		sub -= numero
	}

	fmt.Println(sum, sub)

	return
}

func main() {
	// soma := somar(5, 7)

	// atribuicao de dois retornos, por inferencia
	sum, sub := somaEsubtrai(10, 5)
	fmt.Println(sum, sub)

	somaVariatica(1, 2, 3, 4, 5)

	// funcao anonima
	func(text string) {
		fmt.Println("texto da funcao anonima:", text)
	}("text")
}
