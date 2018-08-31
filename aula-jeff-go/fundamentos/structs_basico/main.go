package main

import (
	"fmt"
)

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A Casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu Apê", 760000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     22,
		valor: 55,
	}
	fmt.Printf("A Chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A Casa é: %+v\r\n", casa)
}
