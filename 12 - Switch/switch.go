package main

import "fmt"

var diaDaSemana = func(dia int) string {

	var nomeDia string

	switch dia {
	case 1:
		nomeDia = "Segunda"
	case 2:
		nomeDia = "Terça"
		fallthrough // Joga o retorno pro caso de baixo
	default:
		nomeDia = "Número inválido"
	}

	return nomeDia
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(2)

	fmt.Println(dia)

}
