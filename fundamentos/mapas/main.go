package main

import (
	"fmt"
	"fundamentos/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "casa Azul"
	casa.X = 28
	casa.Y = 25
	casa.SetValor(60000)

	cache["casa Azul"] = casa

	fmt.Println("Há uma casa azul no cache?")
	imovel, achou := cache["casa Azul"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "apto vermelho"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["casa Azul"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}
}
