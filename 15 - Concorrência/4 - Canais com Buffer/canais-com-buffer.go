package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Canal com buffer (tamanho 2)

	canal <- "Olá Mundo!" // Enviando dados para o canal

	mensagem := <-canal // Recebendo dados do canal
	fmt.Println(mensagem)

}