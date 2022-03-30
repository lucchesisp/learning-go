package main

import "fmt"

func main() {

	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	var (
		variavel3 string = "la la la"
		variavel4 string = "la la la"
	)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
