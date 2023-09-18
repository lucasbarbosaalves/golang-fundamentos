package main

import "fmt"

func main() {
	// Declaração de variáveis explícita
	var variable1 string = "Variable 1"
	fmt.Println(variable1)

	// Declaração de variáveis implícita (inferência de tipo)
	variable2 := "Variable 2"
	fmt.Println(variable2)

	// Declaração de variáveis múltiplas
	variable3, variable4 := "Variable 3", "Variable 4"
	fmt.Println(variable3, variable4)

	// Mudanndo valores das variaveis
	variable3, variable4 = variable4, variable3
	fmt.Println(variable3, variable4)
}