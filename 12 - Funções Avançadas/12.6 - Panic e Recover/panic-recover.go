package main

import "fmt"

 func recuperarExecucao() {
 	if r := recover(); r != nil {
 		fmt.Println("Execução recuperada com sucesso!")
	} }

func alunoEstaAprovado(n1, n2 float64) bool {

	defer recuperarExecucao() // defer -> adia a execução de uma função até que a função que a contém retorne
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	} else {
		panic("A MÉDIA É EXATAMENTE 6!!!") // panic() -> interrompe a execução do programa
	}
}

func main() {

	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")

}