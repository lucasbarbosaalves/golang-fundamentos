package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return // retorno limpo
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}