package main

import (
	"fmt"
	"fundamentos/funcoes_avancada/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 5, 5)

	fmt.Printf("O resultado da Multiplicação foi: %d\n\r", resultado)

	resultado = matematica.Soma(3, 5)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 10, 2)
	fmt.Printf("O resultado da Divisão foi: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(30, 3)
	fmt.Printf("O resultado da Divisão foi: %d e o resto é: %d\r\n", resultado, resto)
}
