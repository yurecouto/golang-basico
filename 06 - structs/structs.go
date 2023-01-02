package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint
	endereco endereco
}

type endereco struct {
	rua  string
	casa int8
}

func main() {
	fmt.Println("Structs")

	var user usuario
	user.nome = "Yure"
	user.idade = 22
	fmt.Println("user =>", user)

	enderecoExemplo := endereco{rua: "rua 1", casa: 3}

	user2 := usuario{"Levi", 4, enderecoExemplo}
	fmt.Println("user =>", user2)

	user3 := usuario{nome: "Bira"}
	fmt.Println("user =>", user3)
}
