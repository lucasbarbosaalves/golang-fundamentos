package main

import "fmt"

func main() {

	// Função anônima
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}