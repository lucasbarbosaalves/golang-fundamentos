package main

import (
	"fmt"
	"sync"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // Quantidade de goroutines que serão executadas ao mesmo tempo

	// Função anônima que será executada em uma goroutine
	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // Indica que uma goroutine foi finalizada (decrementa o contador)
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()	
	}()

	waitGroup.Wait() // Aguarda todas as goroutines serem executadas, a contagem deve chegar a zero
	
	
}

func escrever(texto string) {

	for i:=0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}