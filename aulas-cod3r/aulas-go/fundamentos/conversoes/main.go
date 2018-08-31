package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(int(x) / y)
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	//cuidado
	fmt.Println("Teste " + string(123))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // aceita 1 ou 2 para true or false
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Nao é verdadeiro")
	}
}
