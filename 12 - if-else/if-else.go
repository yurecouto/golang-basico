package main

import "fmt"

func main() {
	numero := 2

	if numero > 5 {
		fmt.Println("maior q 5")
	} else if numero == 5 {
		fmt.Println("igual a 5")
	} else {
		fmt.Println("menor q 5")
	}
}
