package main

import "fmt"

func main() {
	fmt.Println("If Else")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual  15")
	}

	// If init (Mata a variavel após o fim do escopo)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}
}
