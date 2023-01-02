package main

import "fmt"

func main() {
	// declaracoes normais
	var variavel1 string = "variavel tipo implicito"
	variavel2 := "variavel implicita"

	// Ambas sao do tipo string
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// declaracoes aninhadas
	var (
		variavel3 string = "variavel aninhada 1"
		variavel4 string = "variavel aninhada 2"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// constantes
	const constante1 string = "constante 1"
}
