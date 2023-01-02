package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var array1 [5]int8
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"posicao1", "posicao2", "posicao3", "posicao4", "posicao5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	array3[1] = 5
	// array3[4] = 5

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)

	fmt.Println(array3)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
}
