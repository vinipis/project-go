package main

import "fmt"

func main() {
	var (
		nome, turma string
		t1          []string
		t2          []string
	)
	condicao := true

	fmt.Println("Digite 0 na turma para sair...")
	for condicao != false {
		fmt.Println("Digite um nome: ")
		fmt.Scan(&nome)

		fmt.Println("Digite uma turma: ")
		fmt.Scan(&turma)
		fmt.Println(" ")

		if turma == "t1" {

		}
		if turma == "t2" {
			for i := 0; i == i+1; i++ {
				t2[i] = nome
			}
		}
		if turma == "0" {
			condicao = false
		}
	}
	fmt.Println("A turma T1 tem: ", len(t1))
	fmt.Println("A turma T1 tem: ", len(t2))
}
