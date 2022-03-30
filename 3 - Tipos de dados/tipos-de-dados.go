package main

import (
	"errors"
	"fmt"
)

func main() {

	// Números inteiros
	var numero int64 = 100000
	fmt.Println(numero)

	var numero2 uint64 = 10000
	fmt.Println(numero2)

	// alias
	// INT 32 = RUNE
	var numero3 rune = 1000
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 100
	fmt.Println(numero4)
	// Fim números inteiros

	// Números reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.45
	fmt.Println(numeroReal2)
	// Fim números reais

	// Strings
	numeroReal3 := 123.40 // Inferecia do tipo float com base na arquitetura do sistema
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	char := '%'
	fmt.Println(char)

	var teste int32 = 'B'
	fmt.Println(teste)
	// Fim Strings

	var texto string
	fmt.Println(texto)

	var boleano bool
	fmt.Println(boleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
