package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) 
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// Pode gerar um erro em execução -> deadlock
	for {
		mensagem, aberto := <- canal // Recebendo dados do canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	// Outra maneira de fechar

	for mensagem := range canal {
		fmt.Println(mensagem)
	}


	fmt.Println("Fim do programa")

	// mensagem := <- canal // Recebendo dados do canal
	// fmt.Println(mensagem)

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Enviando dados para o canal
		time.Sleep(time.Second)
	}

	close(canal) // Fechando o canal
}