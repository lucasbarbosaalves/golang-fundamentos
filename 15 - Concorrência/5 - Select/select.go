package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
		time.Sleep(time.Millisecond * 500) // 500 milissegundos
		canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
		time.Sleep(time.Second * 2) // 2 segundos
		canal2 <- "Canal 2"
		}
	}()

	for {
		// Select é como um switch, mas para canais
		select {
			case mensagemCanal1 := <- canal1:
				fmt.Println(mensagemCanal1)
			case mensagemCanal2 := <- canal2:
				fmt.Println(mensagemCanal2)
		}
	}
}