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

func main() {
	// soma := somar(5, 7)

	// atribuicao de dois retornos, por inferencia
	sum, sub := somaEsubtrai(10, 5)
	fmt.Println(sum, sub)
}
