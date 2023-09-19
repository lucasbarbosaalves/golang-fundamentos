package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func funcao3(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado")
	fmt.Println("Executando a função 3")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}




func main() {
	fmt.Println("Defer")

	// DEFER -> Adia a execução de uma função até que a função que a contém retorne.
	defer funcao1()
	funcao2()

	fmt.Println(funcao3(7, 8))
}