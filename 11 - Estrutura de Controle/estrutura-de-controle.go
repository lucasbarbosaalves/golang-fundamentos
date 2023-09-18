package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10

	fmt.Println("-------------- IF -----------------")

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init -> inicializa uma variavel que só pode ser usada no escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}

	fmt.Println("-------------- SWITCH -----------------")
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(8)) // retorna default
	
	fmt.Println("-------------- FOR -----------------")
	numero2 := 0
	for i := 0; i < 10; i++ {
		numero2 = numero2 + i
	}
	fmt.Println(numero2)


	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	nomes := [3]string{"João", "Davi", "Lucas"}

	// Iterando sobre um array -> range retorna o indice e o valor
	for i, nome := range nomes {
		fmt.Println(i, nome)

	}

	// Ocultando o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}


	
	

	func diaDaSemana(numero int) string {

		var diaDaSemana string
//		No Go não existe "break" implicito, ou seja, não é necessário colocar o break no final de cada case
	switch numero {
		case 1:
			diaDaSemana = "Domingo"
		case 2:
			diaDaSemana = "Segunda"
		case 3:
			diaDaSemana = "Terça"
		case 4:
			diaDaSemana = "Quarta"
		case 5:
			diaDaSemana = "Quinta"
		case 6:
			diaDaSemana = "Sexta"
		case 7:
			diaDaSemana = "Sabado"
		default:
			diaDaSemana = "Numero invalido"
		}
	return diaDaSemana
	}