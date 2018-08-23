package main

import "fmt"

func main() {
	var (
		v1         [10]int
		num1       int
		numInterno int
	)

	for cont := 0; cont < len(v1); cont++ {
		fmt.Println("Digite um numero: ")
		fmt.Scan(&v1[cont])
	}

	fmt.Println(" ")

	fmt.Println("Digite um numero para procurar dentro do vetor: ")
	fmt.Scan(&num1)

	for cont := 0; cont < len(v1); cont++ {

		switch {
		case num1 == v1[cont]:
			numInterno++
		case numInterno == 0 && cont == len(v1):
			fmt.Printf("O numero %v não ocorre nenhuma vez...", num1)
		}
	}
	fmt.Printf("O número %v aparece %v vezes\r\n", num1, numInterno)
}
