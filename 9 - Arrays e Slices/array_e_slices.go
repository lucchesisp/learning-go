package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array e Slices")

	// Arrays
	fmt.Println("###### Arrays ######")
	var array1 [5]int // Array de 5 posições
	array1[0] = 23
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5} // Segunda forma de declarar um array
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Terceira forma de declarar um array, não torna flexível a quantidade de posições
	fmt.Println(array3)

	// Slices - São "arrays" de qualquer tamanho
	fmt.Println("###### Slices ######")

	slice := []int{1, 2, 3} // Primeira forma de declarar um slice
	fmt.Println(slice)

	slice = append(slice, 33) // Adiciona um elemento ao slice
	fmt.Println(slice)

	slice2 := array2[0:3] // Aponta(Referencia) para um pedaço do Array (Do 1 Inclusivo até o 3 Exclusivo) e guardando no slice
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("###### Arrays Internos ######")

	slice3 := make([]float32, 10, 11) // Array interno
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Capacidade do slice

	fmt.Println("Adicionando um elemento")
	slice3 = append(slice3, 1.1) // Adiciona um elemento ao slice

	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Capacidade do slice

	fmt.Println("Estourando a capacidade")
	slice3 = append(slice3, 1.1) // Adiciona um elemento ao slice

	fmt.Println(len(slice3)) // Tamanho do slice
	fmt.Println(cap(slice3)) // Capacidade do slice
}
