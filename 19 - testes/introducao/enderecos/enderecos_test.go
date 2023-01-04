package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua misteriosa", "Rua"},
		{"Avenida maldita", "Avenida"},
		{"Estrada do terror", "Estrada"},
		{"Rodovia do medo", "Rodovia"},
		{"Praca do medo", "Tipo invalido"},
		{"Rodovia macabra", "Rodovia"},
		{"Avenida sinistra", "Avenida"},
		{"Rua das trevas", "Rua"},
		{"Estrada sangrenta", "Estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido diferente do esperado, recebido %s => esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
