package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var variavel1 int = 100
	var variavel2 int = variavel1

	variavel1 = 300

	// a variavel 2 armazena uma copia do valor
	// que foi atribuido a ela, e caso a variavel1 mude, a
	//variavel2 ira manter o primeiro valor
	fmt.Println(variavel1, variavel2)

	var variavel3 int = 150
	var ponteiro *int = &variavel3

	variavel3 = 940

	// o ponteiro armazena o endereco de memoria, podendo assim
	// mostrat o valor desse endereco, mesmo caso haja alteracao
	fmt.Println(variavel3, ponteiro, *ponteiro)
}
