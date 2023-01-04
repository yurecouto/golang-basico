package main

import (
	"fmt"
	"teste/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua dos bobos")
	fmt.Println(tipoEndereco)
}
