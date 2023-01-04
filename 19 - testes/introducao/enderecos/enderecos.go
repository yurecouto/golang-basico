package enderecos

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	primeiraPalavra := strings.Split(endereco, " ")[0]
	tipoEValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			tipoEValido = true
		}
	}

	if tipoEValido {
		return primeiraPalavra
	}

	return "Tipo invalido"
}
