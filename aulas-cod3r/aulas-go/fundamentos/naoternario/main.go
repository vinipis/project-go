package main

import (
	"fmt"
)

//Não tem operador ternário
func obterResultado(nota float64) string {
	// return nota >= 6 ? "aprovado" : "reprovado" não da para fazer
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado" //não usar else em caso de return unicos
}

func main() {
	fmt.Println(obterResultado(6.2))
}
