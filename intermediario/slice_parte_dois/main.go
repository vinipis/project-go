package main

import "fmt"

func main() {
	// var nums []int
	// fmt.Println(nums, len(nums), cap(nums))
	// nums = make([]int, 5)
	// fmt.Println(nums, len(nums), cap(nums))

	// capitais := []string{"Lisboa"}
	// fmt.Println(capitais, len(capitais), cap(capitais))

	// capitais = append(capitais, "Brasilia")
	// fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)
	cidades[0] = "NY"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "tokio"
	cidades[4] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))

	// capitais[1] = "Brasília"
	// fmt.Println(capitais, len(capitais), cap(capitais))

	// for indice, cidade := range cidades {
	// 	fmt.Printf("Cidade [%d] = %s\n\r", indice, cidade)
	// }

	//Primeiro itel começa com indice 0
	//Segundo item começa com o indice 1
	capitaisAsia := cidades[2:5]
	fmt.Println(capitaisAsia, len(capitaisAsia), cap(capitaisAsia))

	temp1 := cidades[:2]
	fmt.Println(temp1, len(temp1), cap(temp1))

	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2, len(temp2), cap(temp2))

	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)
}
