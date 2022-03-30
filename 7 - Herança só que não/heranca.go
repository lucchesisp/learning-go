package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	var p1 estudante = estudante{
		pessoa: pessoa{
			nome:  "João",
			idade: 30,
		},
		curso:     "Engenharia de Software",
		faculdade: "USP",
	}

	fmt.Println(p1)
}
