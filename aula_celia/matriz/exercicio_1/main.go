package main

import "fmt"

func main() {
	var matriz [2][3]int

	for coluna := 0; coluna < 2; coluna++ {
		fmt.Println("Digite trê numeros: ")
		for linha := 0; linha < 3; linha++ {
			fmt.Scan(&matriz[coluna][linha])
		}
	}

	for coluna := 0; coluna < 2; coluna++ {
		for linha := 0; linha < 3; linha++ {
			fmt.Print(matriz[coluna][linha])
		}
		fmt.Println(" ")
	}
}
