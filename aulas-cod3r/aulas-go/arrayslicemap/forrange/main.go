package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { //para pegar o valor dentro dos indices, declarar só uma variavel ele pega o indice
		fmt.Println(num)
	}
}
