package main

import "fmt"

func main() {	
	generica("String")
	generica(1)
	generica(true)
	generica(1.9)

	// Mapa com tipos genéricos -> não recomendado (Sempre usar tipos específicos)
	mapa := map[interface{}]interface{}{
		1: "String",
		float32(100): true,
	}

	fmt.Println(mapa)
}

// interface{} é um tipo genérico
func generica(interf interface{}){
	fmt.Println(interf)
	
}