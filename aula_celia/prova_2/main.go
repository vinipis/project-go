package main

import (
	"fmt"
)

func main() {
	var (
		num             int
		contVetor       = 0
		armazenaDivisor [10]int
	)

	for i := false; i != true; {
		fmt.Println("Digite um numero maior ou igual a 8 e menor ou igual a 48: ")
		//fmt.Scan(&num)
		num = 8
		if num >= 8 && num <= 48 {
			i = true
		}
	}

	for cont := 1; cont <= num; cont++ {
		if num%cont == 0 {
			armazenaDivisor[contVetor] = cont
			contVetor++
		}
	}

	var vetorTamanhoCerto []int

	for cont := 0; cont < contVetor; cont++ {
		fmt.Println(armazenaDivisor[cont])
		vetorTamanhoCerto[cont] = armazenaDivisor[cont]
	}
	fmt.Printf("Os divisores ocuparam um total de: %v espaços no vetor...", contVetor)

	valorDaSoma, qtdimpar := mediaImpar.Media(vetorTamanhoCerto)

}
