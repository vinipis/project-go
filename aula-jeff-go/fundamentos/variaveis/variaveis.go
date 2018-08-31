package main

import (
	"fmt"
)

var (
	Endereco, Telefone string  //endereço, telefone = ""
	qntd, estoque      int     // qntd = 0
	comprou            bool    // comprou = false
	valor              float64 //valor = 0.00
	palavras           rune
)

func main() {
	teste := "valor de teste"
	fmt.Printf("endereço: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", qntd)
	fmt.Printf("Comprou: %v\r\n", comprou)
	fmt.Printf("Palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
