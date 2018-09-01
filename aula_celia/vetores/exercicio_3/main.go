package main

import (
	"fmt"
)

func main() {
	var (
		num        int
		continuada int
		semestral  int
		media      int
	)
	for {
		fmt.Println("Digite um numero: ")
		fmt.Scan(&num)

		if num >= 3 && num <= 10 {
			fmt.Println("Número no range")
			break
		}
	}
	//	for {
	fmt.Println("Digite a nota continuada: ")
	fmt.Scan(&continuada)

	fmt.Println("Digite a nota semestral: ")
	fmt.Scan(&semestral)

	calculaMedia(continuada, semestral)

	if  media = calculaMedia() {
		fmt.Println("A media foi: ", media)
		return
	}
	
}

func calculaMedia(continuada int, semestral int) (media int) {
	media = 3 + 3
	return

}
