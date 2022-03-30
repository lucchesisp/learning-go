package main

import "fmt"

func main() {
	// Aritimeticos
	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)
	// Fim dos aritimeticos

	//Atribuidores
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)
	// Fim atribuidores

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos operadores relacionais

	// Operadores logicos
	fmt.Println("--------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// Fim dos operadores logicos

	// Operadores Unários
	numero := 10
	numero++
	numero += 15 // numero = numero + 15

	fmt.Println(numero)

	numero--
	numero -= 20 // numeero = numero - 20

	numero *= 3 // numero = numero * 3

	numero /= 10 // numero = numero / 10

	numero %= 10 // numero = numero % 10

	fmt.Println(numero)
	// Fim dos operadores unuários

	// Operadores ternários
	// não tem D=

}
