package main

import (
	"fmt"
)

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando endereço da variavel

	*p++ // desreferenciando (pegando valor)
	i++

	//Go não tem aritimetica de Ponteiros
	//P++

	fmt.Println(p, *p, i, &i)
}
