package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	matricula string
}

func main() {
	fmt.Println("heranca")

	yure := pessoa{
		nome:      "yure",
		sobrenome: "couto",
		idade:     22,
	}

	aluno := estudante{
		pessoa:    yure,
		curso:     "Engenharia de software",
		matricula: "6843843",
	}

	fmt.Println(aluno.idade)
}
