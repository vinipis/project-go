package main

import (
	"fmt"
	"fundamentos/funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 5, 5)
	fmt.Printf("O resultado da Multiplicação foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 3, 3)
	fmt.Printf("O resultado da Soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisao, 3, 3)
	fmt.Printf("O resultado da Divisão foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Subtracao, 3, 3)
	fmt.Printf("O resultado da Subtração foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 6)
	fmt.Printf("O resultado da Divisão foi: %d e o Resto foi: %d\r\n", resultado, resto)
}
