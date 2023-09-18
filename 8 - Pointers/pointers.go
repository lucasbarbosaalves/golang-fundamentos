package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int // valor zero de um ponteiro é nil

	variavel3 = 100
	ponteiro = &variavel3 // pegando o endereço de memória da variável -> 0xc00000a0c8

	fmt.Println(variavel3, *ponteiro) // *ponteiro -> desreferenciação
}