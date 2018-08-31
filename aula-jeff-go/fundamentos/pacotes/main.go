package main

import (
	"fmt"
	"fundamentos/pacotes/operadora"
	"fundamentos/pacotes/prefixo"
)

var (
	NomeDoUsuario = "Vinicius"
)

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de Teste: %s\r\n", prefixo.TesteComPrefixo)
}
