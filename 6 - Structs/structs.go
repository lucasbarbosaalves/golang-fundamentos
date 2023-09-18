package main

import "fmt"

type user struct {
	name string
	idade uint8
	adress adress
}

type adress struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Archive Structs")
	
	var u user
	fmt.Println(u) // Por padrão é inicializado com { 0 }

	// Forma 1 de atribuir valores
	u.name = "Davi"
	u.idade = 21
	fmt.Println(u)

	// Forma 2 de atribuir valores (Mais utilizada)
	user2 := user{"Davi", 21, adress{"Rua dos bobos", 0}}
	fmt.Println(user2)

	// Forma 3 de atribuir valores sem um ou mais valores
	user3 := user{name: "Davi"}
	fmt.Println(user3)
}