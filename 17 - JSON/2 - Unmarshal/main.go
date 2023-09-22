package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade uint   `json:"idade"`
}

func main() {
	// json para struct
	userJson := `{"nome":"João","email":"joao@email", "idade": 20}`

	var u user
	// Recebe um slice de bytes
	if erro := json.Unmarshal([]byte(userJson), &u); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(u)
		
	// json para map
	userJson2 := `{"nome":"João","email":"joao@email"}`

	u2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(userJson2), &u2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(u2)
}
