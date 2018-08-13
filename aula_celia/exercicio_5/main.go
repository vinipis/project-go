package main

import "fmt"

func main() {
	var (
		v    [5]string
		nome string
		cont int
	)
	achou := false

	for cont = 0; cont <= 4; {
		fmt.Println("Digite um nome: ")
		fmt.Scan(&v[cont])
		cont++
	}

	fmt.Println(" ")

	fmt.Println("Digite um nome para consulta: ")
	fmt.Scan(&nome)

	for cont = 0; cont <= 4 && !achou; {
		if v[cont] == nome {
			fmt.Printf("O nome %v esta no vetor. no Indice: [%d]", nome, cont)
			achou = true
		}
		cont++
	}
	if !achou {
		fmt.Println("Nome não existe no vetor...")
	}
}
