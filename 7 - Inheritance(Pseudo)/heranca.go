package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person // herança -> recebe todos os atributos e métodos da struct person
	course  string
	college string
}

func main() {
	fmt.Println("Herança")

	p1 := person{"João", 20}
	fmt.Println(p1)

	s1 := student{p1, "Engenharia", "USP"}
	fmt.Println(s1)
	fmt.Println(s1.name) // acessando o atributo name da struct person através da struct student
}