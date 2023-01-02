package main

import "fmt"

func main() {
	// existem int8, int16, int32, int64...
	// sempre conferir o limite de cada um
	const inteiro int16 = 450
	fmt.Println(inteiro)

	// basicamente o mesmo para floats
	const flutuante float32 = 100.4353453

	// fora isso nada de muito novo
	const string1 string = "Sempre com aspas duplas"
	const boolean bool = false
}
