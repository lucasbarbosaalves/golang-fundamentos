package main

import (
	"errors"
	"fmt"
)

func main() {
	// 	tipos de inteiros -> int8, int16, int32, int64
	number := 100
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// alias
	// int32 = rune
	var number3 rune = 123456
	fmt.Println(number3)

	// Byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	// tipos de ponto flutuante -> float32, float64
	var number5 float32 = 123.45
	fmt.Println(number5)

	// tipos de string
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto"
	fmt.Println(str2)

	char := 'B' // ele pega o valor da tabela ASCII -> não existe CHAR na linguagem GO
	fmt.Println(char)

	var texto string
	fmt.Println(texto) // Por padrão é inicializado com "" (vazio)

	// boolean
	var boolean bool
	fmt.Println(boolean) // Por padrão é inicializado com false

	// error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	var error2 error // Por padrão é inicializado com nil
	fmt.Println(error2)
}