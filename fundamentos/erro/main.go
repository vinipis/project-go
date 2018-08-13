package main

import (
	"encoding/json"
	"fmt"
	"fundamentos/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(100); err != nil {
		fmt.Println("[main] houve um erro na atribuição de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique sua avaliação...")
		}
		return
	}
	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da Casa é: %d\r\n", casa.GetValor())

	objJson, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON:", string(objJson))
}
