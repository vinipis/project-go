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

	for cont := 0; cont < contVetor; cont++ {
		fmt.Println(armazenaDivisor[cont])
	}
	fmt.Printf("Os divisores ocuparam um total de: %v espaços no vetor...", contVetor)

	valorDaSoma, qtdimpar := Media(armazenaDivisor)

	fmt.Printf("O valor da soma dos numero impares é: %v. E a quantidade de impares foi: %v ", valorDaSoma, qtdimpar)
}

//Media é um funçao para calcular os numeros inpares
func Media(v [10]int) (valorDaSoma int, qtdimpar int) {

	for cont := 0; cont < 10; {
		if v[cont]%2 != 0 {
			valorDaSoma = v[cont]
			valorDaSoma = valorDaSoma + valorDaSoma
			qtdimpar++
			cont++
		} else {
			cont++
		}
	}
	return
}
