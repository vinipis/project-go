package main

import "fmt"

func main() {

	var v [6]int
	var cont int

	for cont = 0; cont <= 5; {
		fmt.Println("Digite um numero: ")
		fmt.Scan(&v[cont])
		cont++
	}

	fmt.Println(" ")

	for indice, numeros := range v {
		fmt.Printf("Indice [%d], valor digitado Na ordem foi: %v\n\r", indice, numeros)
	}
}
