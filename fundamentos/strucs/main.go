package main

import (
	"fmt"
)

//Imovel é um cara que é exportado e sempre precisa de comentario algo exportado em Go
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apto := Imovel{17, 56, "Meu apê", 760000}
	fmt.Printf("O apto é: %+v\r\n", apto)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		X:     22,
		valor: 22,
	}
	fmt.Printf("A Cacara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.valor = 45000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)
}
