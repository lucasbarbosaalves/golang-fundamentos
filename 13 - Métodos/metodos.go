package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar() \nNome:", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Lucas", 20}
	usuario1.salvar()

	usuarios2 := usuario{"Mariana", 17}
	usuarios2.salvar()

	usuarios3 := usuario{"João", 50}
	usuarios3.salvar()

	maiorDeIdade := usuarios3.maiorIdade()
	fmt.Println("É maior de idade?",maiorDeIdade, "\nIdade:", usuarios3.idade)

	usuarios3.fazerAniversario()
	fmt.Println("Nova idade:", usuarios3.idade)


}