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
	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara linda", 280000}
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
