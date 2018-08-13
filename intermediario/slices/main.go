package main

import (
	"fmt"
)

func main() {

	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "NY"
	cidades[1] = "Londres"
	cidades[2] = "tokio"
	cidades[3] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))

	capitais[1] = "Brasília"
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade [%d] = %s\n\r", indice, cidade)
	}
}
