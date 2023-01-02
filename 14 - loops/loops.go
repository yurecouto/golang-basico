package main

import "fmt"

func main() {
	i := 0

	for i <= 10 {
		// fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		// fmt.Println(j)
	}

	nomes := [5]string{"Yure", "Maria", "Alpha", "Kiara", "Bira"}

	for k, v := range nomes {
		fmt.Println(k, "=>", v)
	}
}
