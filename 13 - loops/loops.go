package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J")
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Gabriel", "Jessica", "Thawan"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuarios := map[string]string{
		"nome":  "Gabriel",
		"idade": "12",
	}

	for _, usuario := range usuarios {
		fmt.Println(usuario)
	}

	// Loop infinito
	// for {

	// }

	// Não dá pra usar em structs
}
