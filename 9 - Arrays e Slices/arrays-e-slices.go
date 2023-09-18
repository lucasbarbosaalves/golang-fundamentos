package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int	// array de 5 posições
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// array2[5] = "Posição 6" -> não é possível adicionar mais posições ao array

	array3 := [...]int{1, 2, 3, 4, 5} // tamanho baseado na quantidade de elementos
	fmt.Println(array3)

	slices := []int{1, 2, 3, 4, 5} // slices não tem tamanho fixo -> Muito Usado

	fmt.Println(reflect.TypeOf(slices)) // []int -> slice de inteiros
	fmt.Println(reflect.TypeOf(array3)) // [5]int -> array de 5 inteiros

	slices = append(slices, 18) // adiciona um novo elemento ao slice
	fmt.Println(slices)

	slice2 := array2[1:3] // pega do elemento 1 até o 3 (não inclui o 3)
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2) // slice é um ponteiro para o array -> altera o valor do array

	// Arrays Internos
	fmt.Println("----------")
	slice3 := make([]float32, 10, 15) // cria um slice de float32 com 10 posições e um array interno de 15 posições
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho do slice -> length(10)
	fmt.Println(cap(slice3)) // capacidade do slice -> capacity(15)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 11) // ultrapassa a capacidade do array interno
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // length(12)
}