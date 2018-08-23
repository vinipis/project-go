package main

import "fmt"

func main() {
	var matriz [3][3]int

	for coluna := 0; coluna < 3; coluna++ {
		fmt.Println("Digite trê numeros: ")
		for linha := 0; linha < 3; linha++ {
			fmt.Scan(&matriz[coluna][linha])
		}
	}
}
