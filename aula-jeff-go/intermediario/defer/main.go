package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"intermediario/escrever_arquivos/model"
	"os"
	"strings"
)

func main() {
	arquivo, err := os.Open("/home/carlos/go/src/intermediario/ler_arquivos/cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo com leitor csv. Erro", err.Error())
		return
	}

	arquivoJSON, err := os.Create("/home/carlos/go/src/intermediario/escrever_arquivos/cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro", err.Error())
		return
	}
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Println(cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item", item, " ao criar o arquivo JSON, Erro: ", err.Error())
				return
			}

			escritor.WriteString("	" + string(cidadeJSON))
			if (indceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
}
