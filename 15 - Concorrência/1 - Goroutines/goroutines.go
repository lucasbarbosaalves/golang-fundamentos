package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
func main() {
	// Goroutines -> concorrência
	// Goroutines são funções ou métodos que são executados concorrentemente com a função main
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}