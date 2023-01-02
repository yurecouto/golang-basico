package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "veloso",
	}

	fmt.Println(usuario["nome"])

	usuario["nome"] = "yure"
	usuario["sobrenome"] = "couto"

	fmt.Println(usuario["nome"])
}
