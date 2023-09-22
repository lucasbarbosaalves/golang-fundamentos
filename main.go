package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade uint   `json:"idade"`
}

func main() {
	u := user{"Lucas", "lucas@email.com", 28}
	fmt.Println(u)

	userEmJson, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userEmJson)
		
	}
