package main

import "fmt"

func main() {
	var v [10]int
	var cont int
	var soma int
	var media int

	for cont = 0; cont <= 9; {
		fmt.Println("Digite um numero: ")
		fmt.Scan(&v[cont])
		soma = soma + v[cont]

		cont++
	}
	media = soma / 10

	fmt.Printf("A media de todos os numeros é: %v", media)
	fmt.Println(" ")

	for cont = 0; cont <= 9; {
		if v[cont] > media {
			fmt.Printf("O numero %v é maior que a media %v e esta no indice [%v]\r\n", v[cont], media, cont)
		}
		cont++
	}
}
