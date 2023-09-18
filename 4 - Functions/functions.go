package main

import "fmt"

func main() {

	fmt.Println(somar(10, 20))

	// Função anônima
	var f = func (txt string) string {
		return txt
	} 
	fmt.Println(f("Texto da função f"))


	// Função com mais de um retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Ocultando um dos retornos
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)
	
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função com mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}