package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Map -> tipo da chave e valor devem ser o mesmo
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

}