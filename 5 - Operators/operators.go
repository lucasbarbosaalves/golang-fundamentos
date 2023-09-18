package main

import "fmt"

func main() {

	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "Texto"
	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println("---------------------")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true) // Invertendo o valor

	// Operadores unários
	fmt.Println("---------------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /=10
	numero %=3

	fmt.Println(numero)

	// Operadores ternários (Somente if e else)
	// texto :=  numero > 5 ? "Maior que 5" : "Menor que 5"  -> Não existe na linguagem GO


}