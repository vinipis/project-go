package main

import (
	"fmt"
	"net/http"
	"project-go/avancado/conexao_bd_sql/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
