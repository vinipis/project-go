package main

import (
	"fmt"
	"net/http"
	"project-go/avancado/conexao_bd_sql/manipulador"
	"project-go/avancado/conexao_bd_sql/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")

}

func main() {

	err := repo.AbreConexaoComBdSQL()
	if err != nil {
		fmt.Println("erro ao abrir o banco de dados: ", err.Error())
		return

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
