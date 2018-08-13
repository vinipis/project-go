package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("/home/carlos/go/src/intermediario/ler_arquivos/cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro", err.Error())
		return
	}

	// scanner := bufio.NewScanner(arquivo)
	// for scanner.Scan() {
	// 	linha := scanner.Text()
	// 	fmt.Println("O conteudo da linha é: ", linha)
	// }

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo com leitor csv. Erro", err.Error())
		return
	}
	for indceLinha, linha := range conteudo {
		fmt.Printf("Linha [%d] é %s\r\n", indceLinha, linha)
		for indceItem, item := range linha {
			fmt.Printf("O Item [%d] é %s\r\n", indceItem, item)
		}
	}
	arquivo.Close()
}
