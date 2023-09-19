package main

import "fmt"

// init é chamado antes do main
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}